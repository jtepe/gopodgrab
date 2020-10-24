package cmd

import (
	"fmt"
	"sort"

	"github.com/jtepe/gopodgrab/pod"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List managed podcasts",
	Long: `List all podcasts currently managed by gopodgrab sorted by name.
These are the ones stored in the configuration file. The tool does
not actually go look and see whether there are any episodes available`,
	RunE: func(cmd *cobra.Command, args []string) error {
		pods, err := pod.ListPodcasts()
		if err != nil {
			return err
		}

		sort.Slice(pods, func(i, j int) bool {
			return pods[i].Name < pods[j].Name
		})

		printPods(pods)

		return nil
	},
}

// printPods the list of podcasts to stdout.
func printPods(pods []*pod.Podcast) {
	for _, p := range pods {
		fmt.Println(p.Name)
	}
}
