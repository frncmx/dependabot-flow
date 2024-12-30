package flags

import "github.com/spf13/pflag"

type Interface interface {
	FlagName() string
	RegisterTo(flags *pflag.FlagSet)
	Validate() error
}
