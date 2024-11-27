package commands

import (
	"fmt"
)

func PrintHelp() {
	fmt.Println("Available Commands:")
	fmt.Println("  `auth`   - Authenticate with Google and save your OAuth token.")
	fmt.Println("  `upload`  - Upload a video to YouTube.")
	fmt.Println("        Arguments:\n" +
		"            -t <title> - Title of the video.\n" +
		"            -d <description> - Description of the video.\n" +
		"            -p <path> - Path to the video file for upload.\n" +
		"            -pn <playlistName> - Name of the playlist to add the video to.\n" +
		"            -ps <privacyStatus> - Privacy setting of the video.\n" +
		"               \tPossible values:\n " +
		"                  public\n " +
		"                  unlisted\n " +
		"                  private")
	fmt.Println("  `mass-upload`  - Upload multiple videos to YouTube.")
	fmt.Println("  `exit`    - Exit the tool.")
}
