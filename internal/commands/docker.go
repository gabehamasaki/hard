package commands

import (
	"fmt"
	"os/exec"

	"github.com/clebsonsh/hard/internal/config"
	"github.com/clebsonsh/hard/pkg/docker"
	"github.com/spf13/cobra"
)

var (
	upCmd = &cobra.Command{
		Use:                "up",
		Short:              "Start the environment",
		DisableFlagParsing: true,
		Run: func(cmd *cobra.Command, args []string) {
			docker.RunComposeCommand("up", args...)
		},
	}

	downCmd = &cobra.Command{
		Use:   "down",
		Short: "Stop the environment",
		Run: func(cmd *cobra.Command, args []string) {
			docker.RunComposeCommand("down")
		},
	}

	restartCmd = &cobra.Command{
		Use:   "restart",
		Short: "Restart the environment",
		Run: func(cmd *cobra.Command, args []string) {
			docker.RunComposeCommand("restart")
		},
	}

	psCmd = &cobra.Command{
		Use:   "ps",
		Short: "Display the status of all containers",
		Run: func(cmd *cobra.Command, args []string) {
			docker.RunComposeCommand("ps")
		},
	}

	buildCmd = &cobra.Command{
		Use:   "build",
		Short: "Build the environment",
		Run: func(cmd *cobra.Command, args []string) {
			docker.RunComposeCommand("build")
		},
	}

	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update the Laravel Hard CLI",
		Long:  "Update the Laravel Hard CLI and rebuild all containers",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Updating Laravel Hard CLI...")

			// Pull latest changes
			gitCmd := exec.Command("git", "-C", config.HardPath, "pull")
			if err := gitCmd.Run(); err != nil {
				fmt.Printf("Error pulling latest changes: %v\n", err)
				return
			}

			// Drop container images
			fmt.Println("Drop container images...")
			if err := docker.RunComposeCommand("down", "--rmi", "all", "--volumes"); err != nil {
				fmt.Printf("Error dropping containers: %v\n", err)
				return
			}

			// Rebuild container images
			fmt.Println("Rebuild container images...")
			if err := docker.RunComposeCommand("build", "--no-cache"); err != nil {
				fmt.Printf("Error rebuilding containers: %v\n", err)
				return
			}

			// Start the environment
			fmt.Println("Starting the environment...")
			if err := docker.RunComposeCommand("up", "-d"); err != nil {
				fmt.Printf("Error starting containers: %v\n", err)
				return
			}

			fmt.Println("Laravel Hard CLI updated successfully!")
		},
	}
)
