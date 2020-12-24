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
	newEps := make(map[*pod.Podcast][]*pod.Episode)

	for _, p := range pods {
		eps, err := p.NewEpisodes()
		if err != nil {
			return err
		}

		newEps[p] = eps
	}

	if len(newEps) == 0 {
		fmt.Println("No new episodes. Nothing to do.")
		return nil
	}

	for p, eps := range newEps {
		fmt.Printf("%s:\n------------------\n", p.Name)
		for _, e := range eps {
			fmt.Println(e.Title)
		}
	}

	var numEps int
	for _, eps := range newEps {
		numEps += len(eps)
	}

	if waitApproval(numEps) {
		for p, eps := range newEps {
			if err := p.DownloadEpisodes(eps); err != nil {
				return err
			}
		}
	}

	return nil
}

// waitApproval blocks until the user confirms the download of numEps podcast
// episodes with a "y" or "yes" input. Every other input (or error) is
// interpreted as disapproval.
func waitApproval(numEps int) bool {
	fmt.Printf("\nDo you want to download %d episodes? (yes/no) ", numEps)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() == "y" || s.Text() == "yes" {
			return true
		}

		break
	}

	return false
}
