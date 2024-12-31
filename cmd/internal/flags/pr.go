package flags

import (
	"fmt"

	"github.com/spf13/pflag"
)

var _ Interface = new(PR)

type PR struct {
	value int
}

func (f *PR) FlagName() string {
	return "pr"
}

func (f *PR) RegisterTo(flags *pflag.FlagSet) {
	const usage = "the ID of target PR"
	flags.IntVar(&f.value, f.FlagName(), 0, usage)
}

func (f *PR) Validate() error {
	if f.value == 0 {
		return fmt.Errorf("flag %q: must be set to a non-zero number", f.FlagName())
	}
	return nil
}

func (f *PR) ID() int {
	return f.value
}
