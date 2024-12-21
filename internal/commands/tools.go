package commands

import (
	"github.com/clebsonsh/hard/pkg/docker"
	"github.com/spf13/cobra"
)

var (
	bashCmd = &cobra.Command{
		Use:   "bash",
		Short: "Start a shell session within the app container",
		Run: func(cmd *cobra.Command, args []string) {
			docker.RunInContainer("app", "bash", args...)
		},
	}

	laravelCmd = &cobra.Command{
		Use:                "laravel",
		Short:              "Run Laravel commands",
		DisableFlagParsing: true,
		Run: func(cmd *cobra.Command, args []string) {
			docker.RunInContainer("app", "laravel", args...)
		},
	}

	phpCmd = &cobra.Command{
		Use:                "php",
		Short:              "Run PHP commands",
		DisableFlagParsing: true,
		Run: func(cmd *cobra.Command, args []string) {
			docker.RunInProject("php", args)
		},
	}

	composerCmd = &cobra.Command{
		Use:   "composer",
		Short: "Run Composer commands",
		Run: func(cmd *cobra.Command, args []string) {
			docker.RunInProject("composer", args)
		},
	}

	nodeCmd = &cobra.Command{
		Use:   "node",
		Short: "Run Node commands",
		Run: func(cmd *cobra.Command, args []string) {
			docker.RunInProject("node", args)
		},
	}

	npmCmd = &cobra.Command{
		Use:                "npm",
		Short:              "Run npm commands",
		DisableFlagParsing: true,
		Run: func(cmd *cobra.Command, args []string) {
			docker.RunInProject("npm", args)
		},
	}

	yarnCmd = &cobra.Command{
		Use:                "yarn",
		Short:              "Run Yarn commands",
		DisableFlagParsing: true,
		Run: func(cmd *cobra.Command, args []string) {
			docker.RunInProject("yarn", args)
		},
	}
)
