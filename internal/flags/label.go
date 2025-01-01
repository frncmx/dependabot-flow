package flags

import (
	"fmt"

	"github.com/spf13/pflag"
)

var _ Interface = new(Label)

type Label struct {
	value string
}

func (l *Label) FlagName() string {
	return "label"
}

func (l *Label) RegisterTo(flags *pflag.FlagSet) {
	flags.StringVar(&l.value, l.FlagName(), "", "label to apply (must exist)")
}

func (l *Label) Validate() error {
	if l.value == "" {
		return fmt.Errorf("flag %q: must be provided", l.FlagName())
	}
	return nil
}

func (l *Label) Value() string {
	return l.value
}
