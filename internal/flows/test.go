package flows

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/frncmx/dependabot-flow/internal"
)

const (
	repoContentsWrite permission = "repo:contents:write"
	repoPRRead        permission = "repo:pr:read"
	repoPRWrite       permission = "repo:pr:write"
)

type permission string

var _ Interface = new(TestCredentials)

func NewTestCredentials(client internal.Client, targetPR int, reviewer, label string) *TestCredentials {
	return &TestCredentials{
		client:   client,
		pr:       targetPR,
		reviewer: reviewer,
		label:    label,
		output:   os.Stdout,
		failed:   false,
	}
}

type TestCredentials struct {
	client   internal.Client
	pr       int
	reviewer string
	label    string
	output   io.Writer
	failed   bool
}

func (t *TestCredentials) Run(ctx context.Context) error {
	t.getPR(ctx, repoPRRead)
	t.comment(ctx, repoPRWrite)
	t.applyLabel(ctx, repoPRWrite)
	t.requestReview(ctx, repoPRWrite)
	t.enableThenDisableAutoMerge(ctx, repoContentsWrite)
	t.approve(ctx, repoPRWrite)

	if t.failed {
		return fmt.Errorf("there were some errors, credentials might not be set up correctly")
	}

	return nil
}

func (t *TestCredentials) getPR(ctx context.Context, permissions ...permission) {
	t.operation("get PR", permissions...)
	_, err := t.client.GetPR(ctx, t.pr)
	t.finalize(err)
}

func (t *TestCredentials) operation(name string, permissions ...permission) {
	t.printf("%v %v%v", name, permissions, "...")
}

func (t *TestCredentials) printf(format string, a ...any) {
	_, _ = fmt.Fprintf(t.output, format, a...)
}

func (t *TestCredentials) finalize(err error) {
	if err != nil {
		t.printf("failed:\n%v\n", err)
		t.failed = true
		return
	}
	t.printf("ok\n")
}

func (t *TestCredentials) comment(ctx context.Context, permissions ...permission) {
	t.operation("commenting", permissions...)
	err := t.client.Comment(ctx, t.pr, "test commenting "+t.timestamp())
	t.finalize(err)
}

func (t *TestCredentials) timestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (t *TestCredentials) applyLabel(ctx context.Context, permissions ...permission) {
	t.operation("labeling", permissions...)
	err := t.client.Label(ctx, t.pr, t.label)
	t.finalize(err)
}

func (t *TestCredentials) requestReview(ctx context.Context, permissions ...permission) {
	t.operation("review request", permissions...)
	err := t.client.RequestReview(ctx, t.pr, t.reviewer)
	t.finalize(err)
}

func (t *TestCredentials) enableThenDisableAutoMerge(ctx context.Context, permissions ...permission) {
	t.operation("enable auto-merge", permissions...)
	err := t.client.EnableAutoMerge(ctx, t.pr)
	t.finalize(err)

	if err == nil {
		t.operation("disable auto-merge (to avoid merge)", permissions...)
		err = t.client.DisableAutoMerge(ctx, t.pr)
		t.finalize(err)
	}
}

func (t *TestCredentials) approve(ctx context.Context, permissions ...permission) {
	t.operation("approval", permissions...)
	err := t.client.ApprovePR(ctx, t.pr, "test approval "+t.timestamp())
	t.finalize(err)
}
