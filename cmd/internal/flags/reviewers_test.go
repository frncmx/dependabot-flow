package flags_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/frncmx/dependabot-flow/cmd/internal/flags"
)

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
			var got flags.Reviewers
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
