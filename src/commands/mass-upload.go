package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func MassUploadVideo() {
	err := filepath.Walk("/static", func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() {
			parentDir := filepath.Base(filepath.Dir(path))
			fileName := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
			UploadVideo(path, fileName, "", parentDir, "unlisted")
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", "/static", err)
	}
}
