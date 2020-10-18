package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a podcast to your library",
	Long: `Adds a podcast to your list of managed podcasts.
Providing a name, feed URL, and storage location for episodes.
Unless a podcast by that name is already managed, this initializes the storage location and
downloads the newest feed.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmd.Flags().Parse(args); err != nil {
			return err
		}

		log.Println("feed-URL", cmd.Flag("feed-url").Value)
		log.Println("name", cmd.Flag("name").Value)
		log.Println("storage", cmd.Flag("storage").Value)

		return nil
	},
}

func init() {
	addCmd.Flags().StringP("feed-url", "u", "", "URL of the podcast feed")
	addCmd.Flags().StringP("name", "n", "", "Name under which the podcast should be managed")
	addCmd.Flags().StringP("storage", "s", "", "Path to directory (absolute) where to store episodes")
	_ = addCmd.MarkFlagRequired("feed-url")
	_ = addCmd.MarkFlagRequired("name")
	_ = addCmd.MarkFlagRequired("storage")
}
