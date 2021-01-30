package conda

import (
	"fmt"
	"path"
	"strings"

	"github.com/jeewangue/oss-licenses/internal/shell"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// GetCondaPkgLicenses retrieves packages info from the conda environment's `conda-meta` and make it human-readable.
func GetCondaPkgLicenses(c *cli.Context) error {

	infoCommand := &shell.Command{
		ShellToUse: "bash",
		Command:    "conda info --json | jq -r '.conda_prefix'",
	}
	condaPrefix, _, _ := infoCommand.Run()
	condaPrefix = strings.TrimSpace(condaPrefix)
	condaEnvPath := condaPrefix
	if EnvName != "base" {
		condaEnvPath = path.Join(condaPrefix, "envs", EnvName)
	}
	log.WithFields(log.Fields{
		"conda_prefix": condaPrefix,
		"env_name":     EnvName,
		"env_path":     condaEnvPath,
	}).Infof("conda environment")

	log.Info("generating csv")
	command := fmt.Sprintf("cat %s/conda-meta/*.json | jq --slurp -r 'map({ name: .name, version: .version, license: .license, license_family: .license_family, channel: .channel }) | (.[0] | to_entries | map(.key)), .[] | [.[]] | @csv'", condaEnvPath)
	licCommand := &shell.Command{
		ShellToUse: "bash",
		Command:    command,
	}
	stdout, _, _ := licCommand.Run()
	pkgsCount := strings.Count(strings.TrimSpace(stdout), "\n")
	log.Info(fmt.Sprintf("%d packages exported", pkgsCount))
	fmt.Print(stdout)
	return nil
}
