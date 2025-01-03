package builders

import (
	"github.com/spf13/pflag"
	"go.uber.org/multierr"

	"github.com/frncmx/dependabot-flow/internal"
	flags2 "github.com/frncmx/dependabot-flow/internal/flags"
)

var _ Interface[*internal.TestCredentials] = new(TestCredentials)

type TestCredentials struct {
	client    Client
	reviewers flags2.Reviewers
	pr        flags2.PR
	label     flags2.Label
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
		t.pr.Number(),
		t.reviewers.PickRandomReviewer(),
		t.label.Name(),
	)
}
