package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var (
	HardPath string
	WwwPath  string
)

func InitializeConfig(envPath string) error {
	// Load .env file
	if err := loadEnv(envPath); err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	homeDir, _ := os.UserHomeDir()
	HardPath = filepath.Join(homeDir, ".hard")
	if HardPath == "" {
		return fmt.Errorf("ERROR: HARD_PATH not set in .env file")
	}

	WwwPath = os.Getenv("WWW_PATH")
	if WwwPath == "" {
		return fmt.Errorf("ERROR: WWW_PATH not set in .env file")
	}

	return nil
}

func loadEnv(filename string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return fmt.Errorf("ERROR: Couldn't find '.env' file at %s", filename)
	}
	return godotenv.Load(filename)
}
