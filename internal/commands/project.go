package commands

import (
	"os"

	"github.com/clebsonsh/hard/internal/config"
	"github.com/clebsonsh/hard/internal/utils"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:                "hard",
	Short:              "PHP Hard CLI",
	Long:               `PHP Hard CLI is a tool for managing Laravel projects with Docker.`,
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		utils.DisplayHelp()
	},
}

func InitializeCommands() {
	// Add commands
	RootCmd.AddCommand(installCmd)
	RootCmd.AddCommand(upCmd)
	RootCmd.AddCommand(downCmd)
	RootCmd.AddCommand(restartCmd)
	RootCmd.AddCommand(psCmd)
	RootCmd.AddCommand(buildCmd)
	RootCmd.AddCommand(bashCmd)
	RootCmd.AddCommand(laravelCmd)
	RootCmd.AddCommand(phpCmd)
	RootCmd.AddCommand(composerCmd)
	RootCmd.AddCommand(nodeCmd)
	RootCmd.AddCommand(npmCmd)
	RootCmd.AddCommand(yarnCmd)
	RootCmd.AddCommand(updateCmd)

	// Add project commands
	projects := projectsHandles()
	RootCmd.AddCommand(projects...)
}

func projectsHandles() []*cobra.Command {
	projects := []string{}
	files, err := os.ReadDir(config.WwwPath)
	if err != nil {
		return []*cobra.Command{}
	}

	for _, file := range files {
		if file.IsDir() {
			projects = append(projects, file.Name())
		}
	}

	commands := []*cobra.Command{}
	for _, project := range projects {
		projectName := project // Create a new variable to avoid closure issues
		projectCmd := &cobra.Command{
			Use:                projectName,
			Short:              "Run commands in " + projectName,
			DisableFlagParsing: true,
			Run: func(cmd *cobra.Command, args []string) {
				if len(args) > 0 {
					// Check if the first argument is a valid command from rootCmd commands list
					for _, c := range RootCmd.Commands() {
						if c.Use == args[0] {
							args = append(args, projectName)
							c.Run(cmd, args[1:])
							return
						}
					}
				}
				utils.DisplayHelp()
			},
		}
		commands = append(commands, projectCmd)
	}

	return commands
}
