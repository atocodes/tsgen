package generator

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/atocodes/tsgen/internal/installer"
	"github.com/atocodes/tsgen/internal/template"
)

func CreateProject(name string, installPackages bool) error {
	err := os.MkdirAll(name, os.ModePerm)

	if err != nil {
		return err
	}

	// Create src folder
	err = os.MkdirAll(filepath.Join(name, "src"), os.ModePerm)

	if err != nil {
		return err
	}

	//Copy package.json
	err = copyTemplate(
		"base/package.json",
		filepath.Join(name, "package.json"),
		name,
	)

	if err != nil {
		return err
	}

	// Copy nodemon.json
	err = copyTemplate(
		"base/nodemon.json",
		filepath.Join(name, "nodemon.json"),
		name,
	)

	if err != nil {
		return err
	}

	// Copy tsconfig.json
	err = copyTemplate(
		"base/tsconfig.json",
		filepath.Join(name, "tsconfig.json"),
		name,
	)

	if err != nil {
		return err
	}

	// Copy index.json
	err = copyTemplate(
		"base/src/index.ts",
		filepath.Join(name, "src/index.ts"),
		name,
	)

	if err != nil {
		return err
	}

	// Copy gitignore
	err = copyTemplate(
		"base/.gitignore",
		filepath.Join(name, ".gitignore"),
		name,
	)

	if err != nil {
		return err
	}

	if installPackages {
		println("Installing Packages...")

		err = installer.InstallDependencies(name)

		if err != nil {
			return err
		}
	}

	println("Project created successfully")

	// Undecided feature
	// err = installer.RunApp(name)

	// if err != nil {
	// 	return err
	// }

	return nil
}

func copyTemplate(src string, dest string, projectName string) error {
	// Read template file
	content, err := template.Files.ReadFile(src)

	if err != nil {
		return err
	}

	// Replace placeholders
	finalContent := strings.ReplaceAll(
		string(content),
		"{{PROJECT_NAME}}",
		projectName,
	)

	// Write new file
	err = os.WriteFile(dest, []byte(finalContent), os.ModePerm)

	if err != nil {
		return err
	}

	return nil
}
