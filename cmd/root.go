package cmd

import (
	"github.com/gasperio/gasper/pkg/gasper"
	"github.com/spf13/cobra"
)

const (
	use   = "gasper"
	short = "Gasper"
)

var Root = New(use, short, []gasper.Option{})

func New(use, short string, options []gasper.Option) *cobra.Command {
	root := &cobra.Command{
		Use:   use,
		Short: short,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	commands := []*cobra.Command{
		NewCmdLogin(),
		NewCmdLogout(),
		NewCmdPull(&options),
		NewCmdPush(&options),
		NewCmdVersion(),
	}

	root.AddCommand(commands...)

	return root
}
