package cmd

import (
	"github.com/gasperio/gasper/pkg/gasper"
	"github.com/spf13/cobra"
	"log"
)

func NewCmdPush(options *[]gasper.Option) *cobra.Command {
	return &cobra.Command{
		Use:   "push [IMAGE] [PATH]",
		Short: "Push an image to a registry",
		Args:  cobra.ExactArgs(2),
		Run: func(_ *cobra.Command, args []string) {
			img, src := args[0], args[1]

			if err := gasper.Push(img, src, *options...); err != nil {
				log.Fatal(err)
			}
		},
	}
}
