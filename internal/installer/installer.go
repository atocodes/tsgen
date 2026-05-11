package installer

import (
	"os"
	"os/exec"
)

func InstallDependencies(projectPath string) error {
	cmd := exec.Command("npm", "install")

	// Important
	// Run Command inside generated project
	cmd.Dir = projectPath

	// Show terminal output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
