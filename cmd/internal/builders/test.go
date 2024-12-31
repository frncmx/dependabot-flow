package builders

import (
	"github.com/spf13/pflag"
	"go.uber.org/multierr"

	"github.com/frncmx/dependabot-flow/cmd/internal"
	"github.com/frncmx/dependabot-flow/cmd/internal/flags"
)

var _ Interface[*internal.TestCredentials] = new(TestCredentials)

type TestCredentials struct {
	client    Client
	reviewers flags.Reviewers
	pr        flags.PR
	label     flags.Label
}

func (t *TestCredentials) Init(flags *pflag.FlagSet) {
	t.client.Init(flags)
	t.reviewers.RegisterTo(flags)
	t.pr.RegisterTo(flags)
	t.label.RegisterTo(flags)
}

func (t *TestCredentials) Validate() error {
	var err error
	multierr.AppendFunc(&err, t.client.Validate)
	multierr.AppendFunc(&err, t.reviewers.Validate)
	multierr.AppendFunc(&err, t.pr.Validate)
	multierr.AppendFunc(&err, t.label.Validate)
	return err
}

func (t *TestCredentials) Build() *internal.TestCredentials {
	return internal.NewTestCredentials(
		t.client.Build(),
		t.pr.ID(),
		t.reviewers.PickRandomReviewer(),
		t.label.Value(),
	)
}
