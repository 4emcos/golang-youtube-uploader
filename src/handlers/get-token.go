package handlers

import (
	"context"
	"encoding/json"
	"golang.org/x/oauth2"
	"uploader-youtube-video/src/config"
)

func GetToken() *oauth2.Token {
	tokenJSON, err := config.RedisClient.Get(context.Background(), "google-token").Result()
	if err != nil {
		if config.LocalToken != nil {
			return config.LocalToken
		}
		return nil
	}

	var token oauth2.Token
	err = json.Unmarshal([]byte(tokenJSON), &token)
	if err != nil {
		return nil
	}

	return &token
}
