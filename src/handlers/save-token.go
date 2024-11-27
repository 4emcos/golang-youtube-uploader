package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"uploader-youtube-video/src/config"
)

func SaveToken(token *oauth2.Token) {
	tokenJSON, err := json.Marshal(token)
	if err != nil {
		fmt.Errorf("error marshalling token: %v", err)
	}
	config.RedisClient.Del(context.Background(), "google-token")
	err = config.RedisClient.Set(context.Background(), "google-token", tokenJSON, 0).Err()
	if err != nil {
		config.LocalToken = token
		fmt.Errorf("error saving token to Redis: %v", err)
	}
}
