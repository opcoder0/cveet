package cmds

import (
	"github.com/opcoder0/cveet/pkg/cve"
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
		Short: "Sync the local CVE database from CVEProject",
		Run:   cve.Sync,
	}

	searchCmd := &cobra.Command{
		Use:   "search",
		Short: "search CVE database",
	}

	rootCmd.AddCommand(syncCmd)
	rootCmd.AddCommand(searchCmd)

	return rootCmd
}
