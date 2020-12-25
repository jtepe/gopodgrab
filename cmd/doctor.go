package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/jtepe/gopodgrab/pod"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Managed podcast maintenance",
	Long:  `Checks all managed podcasts for broken storage or missing feed files, suggesting actions where possible.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		pods, err := pod.List()
		if err != nil {
			return err
		}

		return checkStorage(pods)
	},
}

func checkStorage(pods []*pod.Podcast) error {
	for _, p := range pods {
		stat, err := os.Stat(p.LocalStore)

		if errors.Is(err, os.ErrNotExist) {
			msg := fmt.Sprintf("%s: storage dir %s does not exist. Create and fetch feed?",
				p.Name, p.LocalStore)

			if waitApproval(msg) {
				fmt.Printf("... create directory %s and download feed from %s\n", p.LocalStore, p.FeedURL)

				if err := p.RefreshFeed(); err != nil {
					reportError(p, "storage "+p.LocalStore, err)
				}
			}

			continue
		}

		if err != nil {
			reportError(p, "storage "+p.LocalStore, err)
			continue
		}

		if !stat.IsDir() {
			reportError(p, p.LocalStore+" is not a directory.", err)
			continue
		}

		if err := checkFeed(p); err != nil {
			reportError(p, p.LocalStore+" feed file "+p.FeedFile(), err)
			continue
		}
	}

	return nil
}

// checkFeed checks the existence of the feed zip archive inside pods storage directory.
func checkFeed(pod *pod.Podcast) error {
	feedFile := pod.FeedFile()

	stat, err := os.Stat(pod.FeedFile())
	if errors.Is(err, os.ErrNotExist) {
		msg := fmt.Sprintf("Feed file %s does not exist. Download?", feedFile)

		if waitApproval(msg) {
			if err := pod.RefreshFeed(); err != nil {
				return err
			}
		}

		return nil
	}

	if err != nil {
		return err
	}

	if stat.IsDir() {
		return errors.New("is a directory")
	}

	return nil
}

func reportError(p *pod.Podcast, msg string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %s: %v\n", p.Name, msg, err)
}
