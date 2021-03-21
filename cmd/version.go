package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime/debug"
)

var Version string

func init() {
	if Version == "" {
		i, ok := debug.ReadBuildInfo()
		if !ok {
			return
		}
		Version = i.Main.Version
	}
}

func NewCmdVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show the version",
		Args:  cobra.NoArgs,
		Run: func(_ *cobra.Command, _ []string) {
			if Version == "" {
				fmt.Println("Could not determine build information")
			} else {
				fmt.Println(Version)
			}
		},
	}
}
