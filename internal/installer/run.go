package installer

import (
	"os"
	"os/exec"
)

func RunApp(projectPath string) error {

	cmd := exec.Command("npm", "run", "dev")

	cmd.Dir = projectPath

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
