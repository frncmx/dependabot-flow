package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"go.uber.org/multierr"

	"github.com/frncmx/dependabot-flow/cmd/internal"
)

var (
	reviewers internal.ReviewersFlag
	repo      internal.RepositoryFlag
	pr        internal.PRFlag
)

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.AddCommand(testCredentialsCmd)
	reviewers.RegisterTo(testCredentialsCmd.Flags())
	repo.RegisterTo(testCredentialsCmd.Flags())
	pr.RegisterTo(testCredentialsCmd.Flags())
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test CLI setup",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("use subcommands")
	},
}

var testCredentialsCmd = &cobra.Command{
	Use: "credentials",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		if internal.GitHubToken.NotSet() {
			err = multierr.Append(err,
				fmt.Errorf("%v must be set", internal.GitHubToken))
		}
		multierr.AppendFunc(&err, reviewers.Validate)
		multierr.AppendFunc(&err, repo.Validate)
		multierr.AppendFunc(&err, pr.Validate)
		return err
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		ok := true
		client := internal.NewClient(internal.GitHubToken.Secret(), repo.Owner(), repo.Name())
		ok = ok && testCommenting(cmd.Context(), client, pr.ID())
		if ok {
			return nil
		}
		return fmt.Errorf("there were some errors, credentials might not be set up correctly")
	},
}

func testCommenting(ctx context.Context, client internal.Client, id int) bool {
	fmt.Print("commenting (pr:write)...")
	_, err := client.CommentIssue(ctx, id, "test "+time.Now().String())
	if err != nil {
		fmt.Println("failed:\n", err)
		return false
	}
	fmt.Println("ok")
	return true
}
