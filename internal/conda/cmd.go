package conda

import "github.com/urfave/cli/v2"

// EnvName is the name of target conda environment
var EnvName string

// Command is a subcommand for conda package
var Command *cli.Command = &cli.Command{
	Name:            "conda",
	Usage:           "export licenses of installed python packages in a conda environment.",
	Flags:           []cli.Flag{&cli.StringFlag{Name: "env", Value: "base", Usage: "name of the conda environment", Destination: &EnvName}},
	Action:          GetCondaPkgLicenses,
	HideHelpCommand: true,
}
