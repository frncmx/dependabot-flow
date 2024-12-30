package builders

import (
	"fmt"

	"github.com/spf13/pflag"
	"go.uber.org/multierr"

	"github.com/frncmx/dependabot-flow/cmd/internal"
	"github.com/frncmx/dependabot-flow/cmd/internal/flags"
)

var _ Interface[internal.Client] = new(Client)

type Client struct {
	repo flags.Repository
}

func (b *Client) Init(flags *pflag.FlagSet) {
	b.repo.RegisterTo(flags)
}

func (b *Client) Validate() error {
	var err error
	if internal.GitHubToken.NotSet() {
		err = multierr.Append(err,
			fmt.Errorf("%v must be set", internal.GitHubToken))
	}
	multierr.AppendFunc(&err, b.repo.Validate)
	return err
}

func (b *Client) Build() internal.Client {
	return internal.NewClient(internal.GitHubToken.Secret(), b.repo.Owner(), b.repo.Name())
}