package cve

import (
	"log"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

const (
	cveSyncDir = ".cveet"
)

// Sync the CVEProject to local directory
func Sync(cmd *cobra.Command, args []string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("Unable to determine user's home directory. Using /tmp/%s for sync", cveSyncDir)
		homeDir = "/tmp"
	}
	syncPath := filepath.Join(homeDir, cveSyncDir)
	err = os.Mkdir(syncPath, 0700)
	if err != nil && !os.IsExist(err) {
		log.Fatalf("Unable to create sync directory %s/%s", homeDir, cveSyncDir)
	}
	_, err = git.PlainClone(syncPath, false, &git.CloneOptions{
		URL:      "https://github.com/CVEProject/cvelistV5.git",
		Depth:    1,
		Progress: os.Stdout,
	})
	if err != nil {
		log.Println("Sync:", err)
	}
}
