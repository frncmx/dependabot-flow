package internal

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"
)

const (
	repoPRWrite permission = "repo:pr:write"
)

type permission string

func NewTestCredentials(client Client, targetPR int, reviewer string) *TestCredentials {
	return &TestCredentials{
		client:   client,
		pr:       targetPR,
		reviewer: reviewer,
		output:   os.Stdout,
		failed:   false,
	}
}

type TestCredentials struct {
	client   Client
	pr       int
	reviewer string
	output   io.Writer
	failed   bool
}

func (t *TestCredentials) Run(ctx context.Context) error {
	t.comment(ctx, repoPRWrite)
	t.requestReview(ctx, repoPRWrite)

	if t.failed {
		return fmt.Errorf("there were some errors, credentials might not be set up correctly")
	}

	return nil
}

func (t *TestCredentials) comment(ctx context.Context, permissions ...permission) {
	t.operation("commenting", permissions...)
	err := t.client.CommentIssue(ctx, t.pr, "test credentials "+t.timestamp())
	t.finalize(err)
}

func (t *TestCredentials) operation(name string, permissions ...permission) {
	t.printf("%v %v%v", name, permissions, "...")
}

func (t *TestCredentials) printf(format string, a ...any) {
	_, _ = fmt.Fprintf(t.output, format, a...)
}

func (t *TestCredentials) timestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (t *TestCredentials) finalize(err error) {
	if err != nil {
		t.printf("failed:\n%v\n", err)
		t.failed = true
	}
	t.printf("ok\n")
}

func (t *TestCredentials) requestReview(ctx context.Context, permissions ...permission) {
	t.operation("review request", permissions...)
	err := t.client.RequestReview(ctx, t.pr, t.reviewer)
	t.finalize(err)
}
