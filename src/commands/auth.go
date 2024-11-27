package commands

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/oauth2"
	"net/http"
	"time"
	"uploader-youtube-video/src/config"
	"uploader-youtube-video/src/handlers"
)

func Auth() {
	if handlers.GetToken() != nil {
		fmt.Println("already has token")
		return
	}

	state := generateStateOauth()

	authURL := config.GoogleOAuthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	fmt.Println("Visit this URL to authenticate:", authURL)

	server := &http.Server{Addr: ":8099"}
	codeChan := make(chan string)

	http.HandleFunc("/oauth/callback", func(w http.ResponseWriter, r *http.Request) {
		queryState := r.URL.Query().Get("state")
		if queryState != state {
			http.Error(w, "Invalid state parameter", http.StatusBadRequest)
			fmt.Println("Invalid state parameter received.")
			return
		}

		code := r.URL.Query().Get("code")
		if code == "" {
			http.Error(w, "No code received", http.StatusBadRequest)
			fmt.Println("No code parameter received.")
			return
		}
		fmt.Fprintln(w, "Authentication successful! You can close this window.")
		codeChan <- code
	})

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Errorf("error starting server: %v", err)
		}
	}()

	fmt.Println("Waiting for the callback...")
	code := <-codeChan

	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Errorf("error shutting down server: %v", err)
	}

	token, err := exchangeToken(code)
	if err != nil {
		fmt.Errorf("error exchanging token: %v", err)
	}
	fmt.Println("Token successfully retrieved and saved!")
	handlers.SaveToken(token)
}

func exchangeToken(code string) (*oauth2.Token, error) {
	return config.GoogleOAuthConfig.Exchange(context.Background(), code)
}

func generateStateOauth() string {
	b := make([]byte, 16)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}
