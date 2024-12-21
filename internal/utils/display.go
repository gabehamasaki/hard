package utils

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

var (
	Yellow = color.New(color.FgYellow).SprintFunc()
	Green  = color.New(color.FgGreen).SprintFunc()
	Red    = color.New(color.FgRed).SprintFunc()
)

func DisplayHelp() {
	fmt.Println("Laravel Hard")
	fmt.Println()
	fmt.Printf("%s\n", Yellow("docker-compose Commands:"))
	fmt.Printf("  %s       Build the environment\n", Green("hard build"))
	fmt.Printf("  %s        Start the environment\n", Green("hard up"))
	fmt.Printf("  %s     Start the environment in the background\n", Green("hard up -d"))
	fmt.Printf("  %s      Stop the environment\n", Green("hard down"))
	fmt.Printf("  %s   Restart the environment\n", Green("hard restart"))
	fmt.Printf("  %s        Display the status of all containers\n", Green("hard ps"))
	fmt.Printf("  %s      Start a shell session within the app container\n", Green("hard bash"))
	fmt.Println()
	fmt.Printf("%s\n", Yellow("Laravel Commands:"))
	fmt.Printf("  %s Run the Laravel command\n", Green("hard laravel ..."))
	fmt.Printf("  %s\n", Green("hard laravel new awesome-project"))
	fmt.Println()
	fmt.Printf("%s\n", Yellow("PHP Commands:"))
	fmt.Printf("  %s Run a snippet of PHP code\n", Green("hard php ..."))
	fmt.Printf("  %s\n", Green("hard awesome-project php artisan migrate"))
	fmt.Println()
	fmt.Printf("%s\n", Yellow("Composer Commands:"))
	fmt.Printf("  %s Run a Composer command\n", Green("hard composer ..."))
	fmt.Printf("  %s\n", Green("hard awesome-project composer require laravel/sanctum"))
	fmt.Println()
	fmt.Printf("%s\n", Yellow("Node Commands:"))
	fmt.Printf("  %s Run a Node command\n", Green("hard node ..."))
	fmt.Printf("  %s\n", Green("hard awesome-project node --version"))
	fmt.Println()
	fmt.Printf("%s\n", Yellow("NPM Commands:"))
	fmt.Printf("  %s Run a npm command\n", Green("hard npm ..."))
	fmt.Printf("  %s\n", Green("hard awesome-project npm run prod"))
	fmt.Println()
	fmt.Printf("%s\n", Yellow("Yarn Commands:"))
	fmt.Printf("  %s Run a Yarn command\n", Green("hard yarn ..."))
	fmt.Printf("  %s\n", Green("hard awesome-project yarn prod"))
	fmt.Println()
	fmt.Printf("%s\n", Yellow("Hard commands:"))
	fmt.Printf("  %s Install Hard CLI\n", Green("hard install"))
	fmt.Printf("  %s Update the Laravel Hard CLI\n", Green("hard update"))
	os.Exit(0)
}
