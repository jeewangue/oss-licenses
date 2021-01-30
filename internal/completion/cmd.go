package completion

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// EnvName is the name of target conda environment
var EnvName string

// Command is a subcommand for conda package
var Command *cli.Command = &cli.Command{
	Name:            "completion",
	Usage:           "generate shell completions",
	HideHelpCommand: true,
	Subcommands: []*cli.Command{
		{
			Name:        "bash",
			Description: "generate bash completion",
			Action: func(c *cli.Context) error {
				fmt.Println(BashCompletion)
				return nil
			},
		},
		{
			Name:        "zsh",
			Description: "generate zsh completion",
			Action: func(c *cli.Context) error {
				fmt.Println(ZshCompletion)
				return nil
			},
		},
	},
}
