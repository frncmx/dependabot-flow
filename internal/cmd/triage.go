package cmd

import (
	"github.com/spf13/cobra"

	"github.com/frncmx/dependabot-flow/internal/builders"
)

var (
	triage builders.Triage
)

func init() {
	rootCmd.AddCommand(triageCmd)
	triage.Init(triageCmd.Flags())
}

var triageCmd = &cobra.Command{
	Use:   "triage",
	Short: "triage a PR, merge or ask for review",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return triage.Validate()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		return triage.Build().Run(cmd.Context())
	},
}
