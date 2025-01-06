package flags

import (
	"fmt"

	"github.com/spf13/pflag"
)

var _ Interface = new(Label)

type Label struct {
	name  string
	value string
}

func (l *Label) FlagName() string {
	if l.name == "" {
		return "label"
	}
	return l.name
}

func (l *Label) RegisterTo(flags *pflag.FlagSet) {
	flags.StringVar(&l.value, l.FlagName(), "", "label to apply (must exist)")
}

func (l *Label) SetFlagName(name string) {
	l.name = name
}

func (l *Label) Validate() error {
	if l.value == "" {
		return fmt.Errorf("flag %q: must be provided", l.FlagName())
	}
	return nil
}

func (l *Label) Name() string {
	return l.value
}
