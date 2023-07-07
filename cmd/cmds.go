package cmd

import (
	cve "github.com/opcoder0/cveet/internal"
	"github.com/spf13/cobra"
)

var (
	instance string
)

// New returns a new typesense shell command tree
func New() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "cveet",
		Short: "CVE Exploration Tool is an offline CVE search tool",
	}

	syncCmd := &cobra.Command{
		Use:   "sync",
		Short: "Sync with the remote CVEProject repository",
		Run:   cve.Sync,
	}

	searchCmd := &cobra.Command{
		Use:   "search",
		Short: "Search CVE database",
	}

	rootCmd.AddCommand(syncCmd)
	rootCmd.AddCommand(searchCmd)

	return rootCmd
}
