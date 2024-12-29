package internal

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/spf13/pflag"
)

type RepositoryFlag struct {
	value string
}

func (f *RepositoryFlag) RegisterTo(flags *pflag.FlagSet) {
	const usage = "target repository in {owner}/{repo} format"
	flags.StringVar(&f.value, f.FlagName(), "", usage)
}

func (f *RepositoryFlag) FlagName() string {
	return "repository"
}

func (f *RepositoryFlag) Validate() error {
	parts := strings.Split(f.value, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return fmt.Errorf("flag %v: %q does not follow {owner}/{repo} format", f.FlagName(), f.value)
	}
	return nil
}

func (f *RepositoryFlag) Owner() string {
	return strings.Split(f.value, "/")[0]
}

func (f *RepositoryFlag) Repo() string {
	return strings.Split(f.value, "/")[1]
}

type ReviewersFlag struct {
	value []string
}

func (f *ReviewersFlag) RegisterTo(flags *pflag.FlagSet) {
	const usage = "list of reviewers to select from (groups not supported)"
	flags.StringSliceVar(&f.value, f.FlagName(), nil, usage)
}

func (f *ReviewersFlag) FlagName() string {
	return "reviewers"
}

func (f *ReviewersFlag) Validate() error {
	if len(f.value) < 1 {
		return fmt.Errorf("flag %v: cannot be empty", f.FlagName())
	}
	return nil
}

func (f *ReviewersFlag) PickRandomReviewer() string {
	return f.value[rand.Intn(len(f.value))]
}

func (f *ReviewersFlag) Reviewers() []string {
	return f.value
}
