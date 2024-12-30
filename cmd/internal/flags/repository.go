package flags

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
)

var _ Interface = new(Repository)

type Repository struct {
	value string
}

func (f *Repository) RegisterTo(flags *pflag.FlagSet) {
	const usage = "target repository in {owner}/{repo} format"
	flags.StringVar(&f.value, f.FlagName(), "", usage)
}

func (f *Repository) FlagName() string {
	return "repo"
}

func (f *Repository) Validate() error {
	parts := strings.Split(f.value, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return fmt.Errorf("flag %q: %q does not follow {owner}/{repo} format", f.FlagName(), f.value)
	}
	return nil
}

func (f *Repository) Owner() string {
	return strings.Split(f.value, "/")[0]
}

func (f *Repository) Name() string {
	return strings.Split(f.value, "/")[1]
}
