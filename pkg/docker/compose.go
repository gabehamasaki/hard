package docker

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/clebsonsh/hard/internal/config"
)

func RunComposeCommand(command string, args ...string) error {
	cmd := getDockerComposeCmd()
	cmd.Args = append(cmd.Args, command)
	cmd.Args = append(cmd.Args, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func RunInContainer(container, command string, args ...string) error {
	cmd := getDockerComposeCmd()
	cmd.Args = append(cmd.Args, "exec", container, command)
	cmd.Args = append(cmd.Args, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func RunInProject(projectCmd string, args []string) error {
	project := determineProject(args)
	if project == "" {
		return fmt.Errorf("couldn't determine project")
	}

	cmd := getDockerComposeCmd()
	cmd.Args = append(cmd.Args, "exec", "-it", "--workdir", "/var/www/"+project, "app", projectCmd)
	cmd.Args = append(cmd.Args, args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func getDockerComposeCmd() *exec.Cmd {
	dockerComposePath := filepath.Join(config.HardPath, "docker-compose.yml")
	if _, err := exec.LookPath("docker"); err != nil {
		fmt.Println("ERROR: Docker is not installed or not in PATH")
		os.Exit(1)
	}

	if _, err := exec.LookPath("docker-compose"); err == nil {
		return exec.Command("docker-compose", "-f", dockerComposePath)
	}

	return exec.Command("docker", "compose", "-f", dockerComposePath)
}

func determineProject(args []string) string {
	if len(args) > 0 {
		if _, err := os.Stat(filepath.Join(config.WwwPath, args[len(args)-1])); err == nil {
			return args[len(args)-1]
		}
	}

	dir, _ := os.Getwd()
	project := filepath.Base(dir)
	if _, err := os.Stat(filepath.Join(config.WwwPath, project)); err != nil {
		fmt.Printf("ERROR: Couldn't find project '%s' directory\n", project)
		return ""
	}
	return project
}
