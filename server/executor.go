package server

import (
	"os/exec"
)

// executeScript executes the bash script and returns the output
func executeScript(script string) ([]byte, error) {
	cmd := exec.Command("/bin/bash", "-c", script)
	output, err := cmd.CombinedOutput()
	return output, err
}
