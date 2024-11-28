# YouTube Uploader

## Overview
The **YouTube Uploader** is a simple command-line application for uploading videos to YouTube. It supports both single video uploads and batch uploads, with options for setting video metadata such as title, description, playlist, and privacy settings.

---

## Features
- **Authentication**: Authenticates with Google using OAuth and saves the token for future use.
- **Single Video Upload**: Uploads a single video with customizable metadata.
- **Mass Upload**: Uploads multiple videos from a predefined list or directory(**under construction**).
- **Help Command**: Lists all available commands and their usage.

---

## Running
1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/uploader-youtube-video.git
   ```
2. Navigate to the project directory:
   ```bash
   cd uploader-youtube-video
   ```
3. Build the application:
   ```bash
   go run cmd\main.go
   ```
   
You will see the following prompt:

```
************************************
********* YouTube Uploader *********
************************************
 Type 'help' to view available commands.
>
```
---

## Usage

### Available Commands

| Command       | Description                                               |
|---------------|-----------------------------------------------------------|
| `help`        | Displays a list of available commands.                    |
| `auth`        | Authenticates with Google and saves the OAuth token.      |
| `upload`      | Uploads a single video to YouTube with custom metadata.   |
| `mass-upload` | Uploads multiple videos at once.(**under construction**). |
| `exit`        | Exits the application.                                    |

---

### Upload Command

Use the `upload` command with the following parameters:

```bash
upload -t <title> -d <description> -p <path> -pn <playlistName> -ps <privacyStatus>
```

| Parameter       | Description                                          |
|-----------------|------------------------------------------------------|
| `-t <title>`    | Title of the video.                                  |
| `-d <description>` | Description of the video.                          |
| `-p <path>`     | Path to the video file for upload.                   |
| `-pn <playlistName>` | Name of the playlist to add the video to.         |
| `-ps <privacyStatus>` | Privacy setting of the video (`public`, `private`, `unlisted`). |

#### Example:
```bash
upload -t "Video Title" -d "Description about the video" -p "./videos/video.mp4" -pn "Vlogs" -ps "public"
```

---

## Technologies Used

### 1. [Golang (Go)](https://golang.org/)
- The core programming language used to build the application.

### 2. [go-redis](https://github.com/redis/go-redis)
- A Redis client for Go, used for caching and session management.
- **_Disclaimer_** : Having a Redis server running is completely optional. It is included to cache OAuth tokens and avoid re-authenticating with Google repeatedly. If Redis is not configured, the tool will prompt for authentication each time.

### 3. [oauth2](https://github.com/golang/oauth2)
- Provides OAuth 2.0 support for Go applications.
- Handles authentication with Google, enabling secure access to YouTube APIs.

### 4. [Google APIs (YouTube Data API)](https://github.com/googleapis/google-api-go-client)
- A client library for interacting with Google services, including YouTube.


