package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"uploader-youtube-video/src/commands"
	"uploader-youtube-video/src/handlers"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("************************************")
	fmt.Println("********* YouTube Uploader *********")
	fmt.Println("************************************")
	fmt.Println("Type 'help' to view available commands.")

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading input: %v", err)
		}

		command := strings.TrimSpace(input)

		if strings.HasPrefix(command, "upload") {
			handleUploadCommand(command)
			continue
		}

		switch command {
		case "help":
			commands.PrintHelp()
		case "auth":
			commands.Auth()
		case "mass-upload":
			commands.MassUploadVideo()
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Printf("Unknown command: %s. Type 'help' for a list of commands.\n", command)
		}
	}
}

func handleUploadCommand(command string) {
	args := handlers.ParseUploadArgs(command)
	if args == nil {
		return
	}
	commands.UploadVideo(args["-t"], args["-d"], args["-p"], args["-pn"], args["-ps"])
}
