package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/clebsonsh/hard/internal/commands"
	"github.com/clebsonsh/hard/internal/config"
	"github.com/clebsonsh/hard/internal/utils"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	// Only initialize config if we're not running the install command
	if len(os.Args) <= 1 || (len(os.Args) > 1 && os.Args[1] != "install") {
		homeDir, _ := os.UserHomeDir()
		hardPath := filepath.Join(homeDir, ".hard")
		envPath := filepath.Join(hardPath, ".env")

		if _, err := os.Stat(envPath); os.IsNotExist(err) {
			return fmt.Errorf("%s", utils.Red("Hard is not installed. Please run 'hard install' first"))
		}

		if err := config.InitializeConfig(envPath); err != nil {
			return fmt.Errorf("error initializing config: %v", err)
		}
	}

	commands.InitializeCommands()
	return commands.RootCmd.Execute()
}
