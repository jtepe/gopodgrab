package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gopodgrab",
	Short: "gopodgrab downloads your podcasts by feed URL",
	Long: `By providing a podcast feed URL gopodgrab manages your favorite podcasts.
It lets you download, update, and search your list of podcasts and episodes.`,
}

func init() {
	rootCmd.AddCommand(addCmd, listCmd, showCmd, versionCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
