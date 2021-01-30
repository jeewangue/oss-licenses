// Package shell provides a method to simply exec shell commands and get the results.
package shell

import (
	"bytes"
	"os/exec"
)

// Command represents a shell command to run.
type Command struct {
	ShellToUse string
	Command    string
}

// Run executes shell commands and returns error, stdout, stderr
func (c *Command) Run() (stdout string, stderr string, err error) {
	var stdoutBytes bytes.Buffer
	var stderrBytes bytes.Buffer
	cmd := exec.Command(c.ShellToUse, "-c", c.Command)
	cmd.Stdout = &stdoutBytes
	cmd.Stderr = &stderrBytes

	err = cmd.Run()
	stdout = stdoutBytes.String()
	stderr = stderrBytes.String()
	return stdout, stderr, err
}
