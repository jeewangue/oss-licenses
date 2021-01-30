package conda

import (
	"encoding/json"
	"fmt"
	"path"
	"strings"

	"github.com/itchyny/gojq"
	"github.com/jeewangue/oss-licenses/internal/shell"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// GetCondaPkgLicenses retrieves packages info from the conda environment's `conda-meta` and make it human-readable.
func GetCondaPkgLicenses(c *cli.Context) error {

	infoCommand := &shell.Command{
		ShellToUse: "bash",
		Command:    "conda info --json",
	}
	infoOutStr, _, _ := infoCommand.Run()
	infoOutStr = strings.TrimSpace(infoOutStr)

	infoOutJSON := make(map[string]interface{})
	_ = json.Unmarshal([]byte(infoOutStr), &infoOutJSON)

	query, _ := gojq.Parse(".conda_prefix")
	v, _ := query.Run(infoOutJSON).Next()

	infoOutStr = v.(string)

	condaEnvPath := infoOutStr
	if EnvName != "base" {
		condaEnvPath = path.Join(infoOutStr, "envs", EnvName)
	}
	log.WithFields(log.Fields{
		"conda_prefix": infoOutStr,
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
