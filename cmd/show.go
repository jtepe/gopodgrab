package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/jtepe/gopodgrab/pod"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:     "show",
	Example: "gopodgrab show FooPodcast",
	Short:   "Short summary of a managed podcast",
	Long: `Display a short summary of the specified managed podcast.
Shows all properties of the podcast.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := pod.Get(args[0])
		if err != nil {
			return err
		}

		show(p)

		return nil
	},
}

func show(p *pod.Podcast) {
	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', tabwriter.AlignRight)
	fmt.Fprintf(tw, "Name\t%s\n", p.Name)
	fmt.Fprintf(tw, "Episodes directory\t%s\n", p.LocalStore)
	tw.Flush()
}
