package cmd

import (
	"github.com/gasperio/gasper/pkg/gasper"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/spf13/cobra"
	"log"
)

func NewCmdLogout() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logout [SERVER]",
		Short: "Log out from a registry",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			reg, err := name.NewRegistry(args[0])
			if err != nil {
				log.Fatal(err)
			}

			if err := gasper.Logout(reg.Name()); err != nil {
				log.Fatal(err)
			}
		},
	}

	return cmd
}
