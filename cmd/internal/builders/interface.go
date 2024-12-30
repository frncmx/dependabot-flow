package builders

import (
	"github.com/spf13/pflag"
)

type Interface[T any] interface {
	Init(flags *pflag.FlagSet)
	Validate() error
	Build() T
}
