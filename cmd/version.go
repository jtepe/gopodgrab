package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "development"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Long:  "Show the full version information from gopodgrab.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gopodgrab", Version)
	},
}
