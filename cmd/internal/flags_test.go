package internal_test

import (
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"

	"github.com/frncmx/dependabot-flow/cmd/internal"
)

func TestRepositoryFlag(t *testing.T) {
	tests := []struct {
		name      string
		in        string
		wantValid bool
		wantOwner string
		wantRepo  string
	}{
		{
			name:      "valid",
			in:        "foo/bar",
			wantValid: true,
			wantOwner: "foo",
			wantRepo:  "bar",
		},
		{
			name:      "empty",
			in:        "",
			wantValid: false,
		},
		{
			name:      "no slash",
			in:        "foo",
			wantValid: false,
		},
		{
			name:      "just slash",
			in:        "/",
			wantValid: false,
		},
		{
			name:      "no owner",
			in:        "/bar",
			wantValid: false,
		},
		{
			name:      "no repo",
			in:        "foo/",
			wantValid: false,
		},
		{
			name:      "multiple slashes",
			in:        "foo/bar/",
			wantValid: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got internal.RepositoryFlag
			initializeFlag(t, &got, tt.in)
			err := got.Validate()
			if err != nil {
				if tt.wantValid {
					assert.NoError(t, err, "validate repository flag")
				}
			} else {
				assert.Equal(t, tt.wantOwner, got.Owner(), "owner")
				assert.Equal(t, tt.wantRepo, got.Name(), "repo")
			}
		})
	}
}

func initializeFlag[T Flag](t *testing.T, flag T, arg string) {
	t.Helper()
	flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
	flag.RegisterTo(flagSet)
	err := flagSet.Parse([]string{"--" + flag.FlagName(), arg})
	if err != nil {
		t.Fatal("flag parsing:", err)
	}
}

type Flag interface {
	FlagName() string
	RegisterTo(flags *pflag.FlagSet)
}

func TestReviewersFlag(t *testing.T) {
	tests := []struct {
		name          string
		in            string
		wantValid     bool
		wantReviewers []string
	}{
		{
			name:          "singe",
			in:            "foo",
			wantValid:     true,
			wantReviewers: []string{"foo"},
		},
		{
			name:          "multiple",
			in:            "foo,bar",
			wantValid:     true,
			wantReviewers: []string{"foo", "bar"},
		},
		{
			name:      "empty",
			in:        "",
			wantValid: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got internal.ReviewersFlag
			initializeFlag(t, &got, tt.in)
			err := got.Validate()
			if err != nil {
				if tt.wantValid {
					assert.NoError(t, err, "validate repository flag")
				}
			} else {
				assert.Equal(t, tt.wantReviewers, got.Reviewers(), "reviewers")
			}
		})
	}
}
