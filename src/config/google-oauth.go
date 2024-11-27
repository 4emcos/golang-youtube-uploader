package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOAuthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8099/oauth/callback",
	ClientID:     "xxxxxxxxxxxxxxxxxxxxxxxx",
	ClientSecret: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/youtube",
		"https://www.googleapis.com/auth/youtube.channel-memberships.creator",
		"https://www.googleapis.com/auth/youtube.force-ssl",
		"https://www.googleapis.com/auth/youtube.readonly",
		"https://www.googleapis.com/auth/youtube.upload",
		"https://www.googleapis.com/auth/youtubepartner",
		"https://www.googleapis.com/auth/youtubepartner-channel-audit",
	},
	Endpoint: google.Endpoint,
}

var LocalToken *oauth2.Token
