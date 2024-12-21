package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/clebsonsh/hard/internal/utils"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install Hard CLI",
	Long:  "Install Hard CLI and set up the necessary environment",
	Run:   runInstall,
}

func runInstall(cmd *cobra.Command, args []string) {
	fmt.Println(utils.Green("Installing Hard..."))

	// Check for required tools
	checkRequiredTools()

	// Set up paths
	homeDir, _ := os.UserHomeDir()
	hardPath := filepath.Join(homeDir, ".hard")
	wwwPath, _ := cmd.Flags().GetString("www-path")
	if wwwPath == "" {
		wwwPath = filepath.Join(homeDir, "hard")
	}

	// Clone or update repository
	setupRepository(hardPath)

	// Set up .env file
	setupEnvFile(hardPath, wwwPath)

	// Create WWW_PATH
	if err := os.MkdirAll(wwwPath, os.ModePerm); err != nil {
		fmt.Printf("Error creating WWW_PATH: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(utils.Green("Hard installed successfully!"))
}

func checkRequiredTools() {
	tools := []string{"git", "docker"}
	for _, tool := range tools {
		if _, err := exec.LookPath(tool); err != nil {
			fmt.Printf("%s is not installed. Please install %s and try again.\n", tool, tool)
			os.Exit(1)
		}
	}

	// Check for docker-compose
	_, err1 := exec.LookPath("docker-compose")
	err2 := exec.Command("docker", "compose").Run()
	if err1 != nil && err2 != nil {
		fmt.Println("Docker Compose is not installed. Please install docker-compose and try again.")
		os.Exit(1)
	}
}

func setupRepository(hardPath string) {
	if _, err := os.Stat(hardPath); os.IsNotExist(err) {
		cmd := exec.Command("git", "clone", "https://github.com/clebsonsh/hard.git", hardPath)
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error cloning repository: %v\n", err)
			os.Exit(1)
		}
	} else {
		cmd := exec.Command("git", "-C", hardPath, "pull")
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error pulling repository: %v\n", err)
			os.Exit(1)
		}
	}
}

func setupEnvFile(hardPath string, wwwPath string) {
	envPath := filepath.Join(hardPath, ".env")
	envExamplePath := filepath.Join(hardPath, ".env.example")
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		input, err := os.ReadFile(envExamplePath)
		if err != nil {
			fmt.Printf("Error reading .env.example: %v\n", err)
			os.Exit(1)
		}

		output := strings.ReplaceAll(string(input), "USER=hard", fmt.Sprintf("USER=%s", os.Getenv("USER")))
		output = strings.ReplaceAll(output, "USER_ID=1001", fmt.Sprintf("USER_ID=%d", os.Getuid()))
		output = strings.ReplaceAll(output, "WWW_PATH=\"/home/${USER}/hard\"", fmt.Sprintf("WWW_PATH=\"%s\"", wwwPath))

		if err := os.WriteFile(envPath, []byte(output), 0644); err != nil {
			fmt.Printf("Error writing .env file: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Println("If you want to change the default values, please edit the .env file in", hardPath)
}

func init() {
	installCmd.Flags().StringP("www-path", "p", "", "Custom path for WWW directory")
}
