package internal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/frncmx/dependabot-flow/internal"
)

func TestCommentWithQuote_Markdown(t *testing.T) {
	tests := []struct {
		name    string
		message string
		quote   string
		want    string
	}{
		{
			name:    "empty",
			message: "",
			quote:   "",
			want:    "",
		},
		{
			name:    "no quote",
			message: "foo",
			quote:   "",
			want:    "foo",
		},
		{
			name:    "no quote with newline",
			message: "foo\n",
			quote:   "",
			want:    "foo",
		},
		{
			name:    "single quote",
			message: "foo",
			quote:   "bar",
			want:    "foo\n\n> bar",
		},
		{
			name:    "single quote with newlines",
			message: "foo\n\n\n",
			quote:   "bar\n\n",
			want:    "foo\n\n> bar",
		},
		{
			name:    "multiline quote",
			message: "foo",
			quote:   "bar 1\n \nbar 2",
			want:    "foo\n\n> bar 1\n>\n> bar 2",
		},
		{
			name:    "leading space",
			message: " \nfoo",
			quote:   "\n bar",
			want:    "foo\n\n> bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := internal.CommentWithQuote{
				Message: tt.message,
				Quote:   tt.quote,
			}.Markdown()

			assert.Equal(t, tt.want, got)
		})
	}

	t.Run("called twice", func(t *testing.T) {
		want := "foo"
		in := internal.CommentWithQuote{Message: want, Quote: ""}
		_ = in.Markdown()
		got := in.Markdown()
		require.Equal(t, want, got)
	})
}
