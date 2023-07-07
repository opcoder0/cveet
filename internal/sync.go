package cve

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

const (
	cveSyncDir = ".cveet"
	cveProject = "https://github.com/CVEProject/cvelistV5.git"
)

// Sync the CVEProject to local directory
func Sync(cmd *cobra.Command, args []string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("unable to determine home directory: %v\n", err)
		os.Exit(1)
	}
	syncPath := filepath.Join(homeDir, cveSyncDir)
	fileInfo, err := os.Stat(syncPath)
	if fileInfo != nil && !fileInfo.IsDir() {
		fmt.Printf("found %s file; expected directory\n", syncPath)
		os.Exit(1)
	}
	if err != nil {
		if err = initRepo(syncPath); err != nil {
			fmt.Printf("init: %v\n", err)
			os.Exit(1)
		}
	} else {
		if err = updateRepo(syncPath); err != nil {
			fmt.Printf("update: %v\n", err)
			os.Exit(1)
		}
	}
}

func initRepo(syncPath string) error {
	err := os.Mkdir(syncPath, 0700)
	if err != nil && !os.IsExist(err) {
		fmt.Println("Unable to create sync directory", syncPath)
		return err
	}
	_, err = git.PlainClone(syncPath, false, &git.CloneOptions{
		URL:      cveProject,
		Progress: os.Stdout,
	})
	return err
}

func updateRepo(syncPath string) error {
	repo, err := git.PlainOpen(syncPath)
	if err != nil {
		return err
	}
	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}
	err = worktree.Pull(&git.PullOptions{
		RemoteName: "origin",
		RemoteURL:  cveProject,
		Progress:   os.Stdout,
	})
	return err
}
