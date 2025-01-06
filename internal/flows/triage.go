package flows

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/frncmx/dependabot-flow/internal"
)

var _ Interface = new(Triage)

func NewTriage(client internal.Client, number int, reviewer string) *Triage {
	return &Triage{
		client:   client,
		pr:       internal.PullRequest{Number: number},
		reviewer: reviewer,
		output:   os.Stdout,
	}
}

type Triage struct {
	client   internal.Client
	pr       internal.PullRequest
	reviewer string
	output   io.Writer
	err      error
}

func (t *Triage) Run(ctx context.Context) error {
	t.fetch(ctx)
	if t.safe() {
		t.approve(ctx)
		t.enableAutoMerge(ctx)
	} else {
		t.comment(ctx)
		t.askForReview(ctx)
	}
	return t.err
}

func (t *Triage) fetch(ctx context.Context) {
	pr, err := t.client.GetPR(ctx, t.pr.Number)
	if err != nil {
		t.err = err
		return
	}
	t.pr = pr
	t.printf("fetched: %v\n", pr.URL)
}

func (t *Triage) printf(format string, a ...any) {
	_, _ = fmt.Fprintf(t.output, format, a...)
}

func (t *Triage) safe() bool {
	safe := t.pr.Safe()
	if safe {
		t.printf("looks safe\n")
	} else {
		t.printf("looks suspicious\n")
	}
	return safe
}

func (t *Triage) approve(ctx context.Context) {
	if t.errored() {
		return
	}
	comment := fmt.Sprintf("LGTM\n(PR body scan pattern: `%v`)", t.pr.SuspiciousPattern())
	t.err = t.client.ApprovePR(ctx, t.pr.Number, comment)
	if !t.errored() {
		t.printf("approved\n")
	}
}

func (t *Triage) errored() bool {
	return t.err != nil
}

func (t *Triage) enableAutoMerge(ctx context.Context) {
	if t.errored() {
		return
	}
	t.err = t.client.EnableAutoMerge(ctx, t.pr.Number)
	if !t.errored() {
		t.printf("enabled auto-merge\n")
	}
}

func (t *Triage) comment(ctx context.Context) {
	if t.errored() {
		return
	}
	comment := internal.CommentWithQuote{
		Message: fmt.Sprintf(
			"@%v please, look into this!\n\n:x: Suspicious PR body! Regexp pattern:\n\n`%v`",
			t.reviewer,
			t.pr.SuspiciousPattern(),
		),
		Quote: t.pr.SuspiciousText(),
	}
	separator := strings.Repeat("#", 30)
	t.printf("%v\n%v\n%v\n", separator, comment, separator)
	t.err = t.client.Comment(ctx, t.pr.Number, comment.Markdown())
	if !t.errored() {
		t.printf("commented on PR\n")
	}
}

func (t *Triage) askForReview(ctx context.Context) {
	if t.errored() {
		return
	}
	t.err = t.client.RequestReview(ctx, t.pr.Number, t.reviewer)
	if !t.errored() {
		t.printf("requested review from %q\n", t.reviewer)
	}
}
