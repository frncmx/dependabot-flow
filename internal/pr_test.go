package internal_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/frncmx/dependabot-flow/internal"
)

func TestPullRequest_SuspiciousText(t *testing.T) {
	tests := []struct {
		name          string
		inputFileName string
		safe          bool
	}{
		{
			name:          "dependabot help",
			inputFileName: "pr/dependabot-help",
			safe:          true,
		},
		{
			name:          "does not follow semantic import versioning",
			inputFileName: "pr/version+incompatible",
			safe:          true,
		},
		{
			name:          "breaking changes",
			inputFileName: "pr/breaking-changes",
			safe:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := internal.PullRequest{
				Body: readAll(t, tt.inputFileName),
			}
			if tt.safe {
				assert.True(t, in.Safe(), "expect PR body to be safe")
				assert.Empty(t, in.SuspiciousText(), "should have no text")
			} else {
				want := readAll(t, tt.inputFileName+".want")
				assert.False(t, in.Safe(), "expect PR body to be suspicious")
				assert.Equal(t, want, in.SuspiciousText(), "text should match")
			}
		})
	}
}

func readAll(t *testing.T, fileName string) string {
	t.Helper()
	file, err := os.ReadFile(filepath.Join("testdata", fileName+".txt"))
	require.NoError(t, err, "read all")
	return string(file)
}
