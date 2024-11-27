package handlers

import (
	"fmt"
	"strings"
)

func ParseUploadArgs(command string) map[string]string {
	cleanedInput := strings.ReplaceAll(command, `"`, "")
	parts := strings.Split(cleanedInput, " ")
	if len(parts) < 2 {
		fmt.Println("Invalid args, type `help`")
		return nil
	}

	args := make(map[string]string)
	for i := 1; i < len(parts)-1; i += 2 {
		args[parts[i]] = parts[i+1]
	}

	required := []string{"-t", "-p", "-ps"}
	privacyRequired := []string{"public", "unlisted", "private"}

	for _, key := range required {
		if _, exists := args[key]; !exists {
			fmt.Printf("Missing required argument: %s\n", key)
			return nil
		}
	}

	psValue := args["-ps"]
	valid := false
	for _, validValue := range privacyRequired {
		if psValue == validValue {
			valid = true
			break
		}
	}

	if !valid {
		fmt.Printf("Invalid argument for -ps: %s\n", psValue)
		fmt.Println("Possible values:\n" +
			"     public\n" +
			"     unlisted\n" +
			"     private")
		return nil
	}

	if _, exists := args["-d"]; !exists {
		args["-d"] = ""
	}
	if _, exists := args["-pn"]; !exists {
		args["-pn"] = ""
	}

	return args
}
