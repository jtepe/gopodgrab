package cmd

import (
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update <podcast-name>",
	Short: "updates the specifed podcast",
	Long: `Updates the specified podcast's episodes, downloading all
episodes that are not yet present in the local storage.

The special name "all" updates all managed podcasts.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
