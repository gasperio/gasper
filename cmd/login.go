package cmd

import (
	"github.com/gasperio/gasper/pkg/gasper"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/spf13/cobra"
	"log"
)

func NewCmdLogin() *cobra.Command {
	var username, password string
	var passwordStdin bool

	cmd := &cobra.Command{
		Use:   "login [SERVER]",
		Short: "Log in to a registry",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			reg, err := name.NewRegistry(args[0])
			if err != nil {
				log.Fatal(err)
			}

			if err := gasper.Login(reg.Name(), username, password, passwordStdin); err != nil {
				log.Fatal(err)
			}
		},
	}

	flags := cmd.Flags()

	flags.StringVarP(&username, "username", "u", "", "username")
	flags.StringVarP(&password, "password", "p", "", "password")
	flags.BoolVarP(&passwordStdin, "password-stdin", "", false, "take the password from stdin")

	return cmd
}
