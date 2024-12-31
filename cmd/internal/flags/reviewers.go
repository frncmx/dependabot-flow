package flags

import (
	"fmt"
	"math/rand"

	"github.com/spf13/pflag"
)

var _ Interface = new(Reviewers)

type Reviewers struct {
	value []string
}

func (f *Reviewers) FlagName() string {
	return "reviewers"
}

func (f *Reviewers) RegisterTo(flags *pflag.FlagSet) {
	const usage = "list of reviewers to select from (groups not supported)"
	flags.StringSliceVar(&f.value, f.FlagName(), nil, usage)
}

func (f *Reviewers) Validate() error {
	if len(f.value) < 1 {
		return fmt.Errorf("flag %q: cannot be empty", f.FlagName())
	}
	return nil
}

func (f *Reviewers) PickRandomReviewer() string {
	return f.value[rand.Intn(len(f.value))]
}

func (f *Reviewers) Reviewers() []string {
	return f.value
}
