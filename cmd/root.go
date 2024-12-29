package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "dependabot-flow",
}

var globalRepository string

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(
		&globalRepository,
		"repository",
		"",
		"target repository in {owner}/{repo} format",
	)
	must(rootCmd.MarkPersistentFlagRequired("repository"))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
