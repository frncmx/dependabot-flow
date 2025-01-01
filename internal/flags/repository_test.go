package flags_test

import (
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"

	flags2 "github.com/frncmx/dependabot-flow/internal/flags"
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
			var got flags2.Repository
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

func initializeFlag[T flags2.Interface](t *testing.T, flag T, arg string) {
	t.Helper()
	flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
	flag.RegisterTo(flagSet)
	err := flagSet.Parse([]string{"--" + flag.FlagName(), arg})
	if err != nil {
		t.Fatal("flag parsing:", err)
	}
}
