package commands

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"log"
	"os"
	"uploader-youtube-video/src/config"
	"uploader-youtube-video/src/handlers"
)

func UploadVideo(title string, description string, videoPath string, playListName string, privacyStatus string) {
	tags := []string{"youTube", "uploader"}

	file, err := os.Open(videoPath)
	if err != nil {
		fmt.Println("error opening video file: %w", err)
	}
	defer file.Close()

	fmt.Printf("Uploading video: \n"+
		"\nWith title: %s"+
		"\nWith description: %s"+
		"\nIn playlist name: %s\n", title, description, playListName)

	service, err := youtube.NewService(context.Background(), option.WithTokenSource(config.GoogleOAuthConfig.TokenSource(context.Background(), handlers.GetToken())))

	video := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       title,
			Description: description,
			Tags:        tags,
			CategoryId:  "22",
		},
		Status: &youtube.VideoStatus{
			PrivacyStatus: privacyStatus,
		},
	}

	call := service.Videos.Insert([]string{"snippet", "status"}, video)
	call.Media(file)

	uploadedVideo, err := call.Do()
	if err != nil {
		fmt.Println("error uploading video: %w", err)
	}

	videoId := uploadedVideo.Id
	fmt.Printf("video uploaded successfully: https://www.youtube.com/watch?v=%s\n", videoId)

	if playListName != "" {
		createPlaylist(playListName, service, uploadedVideo)
	}
}

func createPlaylist(playListName string, service *youtube.Service, uploadedVideo *youtube.Video) {

	var playlistId string
	playlistCall := service.Playlists.List([]string{"id", "snippet"}).Mine(true)
	playlistResponse, err := playlistCall.Do()

	if err != nil {
		log.Fatalf("error fetching playlists: %v", err)
	}

	for _, item := range playlistResponse.Items {
		if item.Snippet.Title == playListName {
			fmt.Printf("playlists: %s", item.Id)

			playlistId = item.Id
			break
		}
	}

	if playlistId == "" {
		playlist := &youtube.Playlist{
			Snippet: &youtube.PlaylistSnippet{
				Title: playListName,
			},
			Status: &youtube.PlaylistStatus{
				PrivacyStatus: "private",
			},
		}

		playlistInsertCall := service.Playlists.Insert([]string{"snippet", "status"}, playlist)
		newPlaylist, err := playlistInsertCall.Do()
		if err != nil {
			log.Fatalf("error creating playlist: %v", err)
		}
		playlistId = newPlaylist.Id
		fmt.Printf("Playlist created successfully: %s\n", playListName)
	}

	playlistItem := &youtube.PlaylistItem{
		Snippet: &youtube.PlaylistItemSnippet{
			PlaylistId: playlistId,
			ResourceId: &youtube.ResourceId{
				Kind:    "youtube#video",
				VideoId: uploadedVideo.Id,
			},
		},
	}

	playlistItemCall := service.PlaylistItems.Insert([]string{"snippet"}, playlistItem)
	_, err = playlistItemCall.Do()
	if err != nil {
		fmt.Errorf("error adding video to playlist: %v", err)
	}

	fmt.Printf("Video added to playlist: %s\n", playListName)

}
