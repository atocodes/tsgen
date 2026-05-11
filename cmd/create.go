/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/atocodes/tsgen/internal/generator"
)

var installPackages bool

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a typescript starter code project with basic pacakges preconfigured with nodemon",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide project name")
			return
		}

		projectName := args[0]

		fmt.Println("Creating:", projectName, "...")

		err := generator.CreateProject(projectName, installPackages)

		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
			return
		}
	},
}

func init() {
	createCmd.Flags().BoolVarP(
		&installPackages,
		"install-packages",
		"p",
		false,
		"Install packages after generate",
	)
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
