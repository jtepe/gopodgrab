package cmd

import (
	"log"

	"github.com/jtepe/gopodgrab/pod"
	"github.com/spf13/cobra"
)

const (
	flagFeedURL = "feed-url"
	flagName    = "name"
	flagStorage = "storage"
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

		name := cmd.Flag(flagName).Value.String()
		feedURL := cmd.Flag(flagFeedURL).Value.String()
		storage := cmd.Flag(flagStorage).Value.String()

		return add(name, feedURL, storage)
	},
}

func add(name, feedURL, storage string) error {
	podcast, err := pod.New(name, feedURL, storage)
	if err != nil {
		return err
	}

	log.Printf("podcast %s added under %s", podcast.Name, podcast.LocalStore)

	return nil
}

func init() {
	addCmd.Flags().StringP("feed-url", "u", "", "URL of the podcast feed")
	addCmd.Flags().StringP("name", "n", "", "Name under which the podcast should be managed")
	addCmd.Flags().StringP("storage", "s", "", "Path to directory (absolute) where to store episodes")
	_ = addCmd.MarkFlagRequired("feed-url")
	_ = addCmd.MarkFlagRequired("name")
	_ = addCmd.MarkFlagRequired("storage")
}
