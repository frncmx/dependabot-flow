package cmd

import (
	"github.com/spf13/cobra"

	"github.com/frncmx/dependabot-flow/internal/builders"
)

var (
	testCredentials builders.TestCredentials
)

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.AddCommand(testCredentialsCmd)
	testCredentials.Init(testCredentialsCmd.Flags())
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test CLI setup",
}

var testCredentialsCmd = &cobra.Command{
	Use: "credentials",
	PreRunE: func(*cobra.Command, []string) error {
		return testCredentials.Validate()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		return testCredentials.Build().Run(cmd.Context())
	},
}
