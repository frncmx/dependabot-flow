package builders

import (
	"github.com/spf13/pflag"
	"go.uber.org/multierr"

	"github.com/frncmx/dependabot-flow/internal/flags"
	"github.com/frncmx/dependabot-flow/internal/flows"
)

var _ Interface[*flows.Triage] = new(Triage)

type Triage struct {
	client    Client
	reviewers flags.Reviewers
	pr        flags.PR
}

func (t *Triage) Init(flags *pflag.FlagSet) {
	t.client.Init(flags)
	t.reviewers.RegisterTo(flags)
	t.pr.RegisterTo(flags)
}

func (t *Triage) Validate() error {
	var err error
	multierr.AppendFunc(&err, t.client.Validate)
	multierr.AppendFunc(&err, t.reviewers.Validate)
	multierr.AppendFunc(&err, t.pr.Validate)
	return err
}

func (t *Triage) Build() *flows.Triage {
	return flows.NewTriage(
		t.client.Build(),
		t.pr.Number(),
		t.reviewers.PickRandomReviewer(),
	)
}
