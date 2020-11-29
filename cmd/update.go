package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jtepe/gopodgrab/pod"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [<podcast>|all] [<podcast>...]",
	Short: "updates the specifed podcast",
	Long: `Updates the specified podcast's episodes, downloading all
episodes that are not yet present in the local storage.

The special name "all" updates all managed podcasts.`,

	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pods := make([]*pod.Podcast, 0, len(args))

		for _, arg := range args {
			if arg == pod.ReservedPodName {
				all, err := pod.List()
				if err != nil {
					return err
				}

				return updatePods(all)
			}

			p, err := pod.Get(arg)
			if err != nil {
				return err
			}

			pods = append(pods, p)
		}

		return updatePods(pods)
	},
}

func updatePods(pods []*pod.Podcast) error {
	for _, p := range pods {
		eps, err := p.NewEpisodes()
		if err != nil {
			return err
		}

		if len(eps) == 0 {
			return nil
		}

		fmt.Printf("%s:\n------------------\n", p.Name)
		for _, e := range eps {
			fmt.Println(e.Title)
		}

		if waitApproval() {
			if err := p.DownloadEpisodes(eps); err != nil {
				return err
			}
		}
	}

	return nil
}

// waitApproval blocks until the user confirms the progression with
// a "y" or "yes" input. Every other input (or error) is interpreted
// as disapproval.
func waitApproval() bool {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() == "y" || s.Text() == "yes" {
			return true
		}
	}

	return false
}
