package builders

import (
	"fmt"

	"github.com/spf13/pflag"
	"go.uber.org/multierr"

	internal2 "github.com/frncmx/dependabot-flow/internal"
	"github.com/frncmx/dependabot-flow/internal/flags"
)

var _ Interface[internal2.Client] = new(Client)

type Client struct {
	repo flags.Repository
}

func (b *Client) Init(flags *pflag.FlagSet) {
	b.repo.RegisterTo(flags)
}

func (b *Client) Validate() error {
	var err error
	if internal2.GitHubToken.NotSet() {
		err = multierr.Append(err,
			fmt.Errorf("%v must be set", internal2.GitHubToken))
	}
	multierr.AppendFunc(&err, b.repo.Validate)
	return err
}

func (b *Client) Build() internal2.Client {
	return internal2.NewClient(internal2.GitHubToken.Secret(), b.repo.Owner(), b.repo.Name())
}
