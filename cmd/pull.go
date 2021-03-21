package cmd

import (
	"github.com/gasperio/gasper/pkg/gasper"
	"github.com/spf13/cobra"
	"log"
)

func NewCmdPull(options *[]gasper.Option) *cobra.Command {
	return &cobra.Command{
		Use:   "pull [IMAGE] [PATH]",
		Short: "Pull an image from a registry",
		Args:  cobra.ExactArgs(2),
		Run: func(_ *cobra.Command, args []string) {
			img, dst := args[0], args[1]

			if err := gasper.Pull(img, dst, *options...); err != nil {
				log.Fatal(err)
			}
		},
	}
}
