package internal

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"
)

func NewTestCredentials(client Client, targetPR int) TestCredentials {
	return TestCredentials{
		client: client,
		pr:     targetPR,
		output: os.Stdout,
		failed: false,
	}
}

type TestCredentials struct {
	client Client
	pr     int
	output io.Writer
	failed bool
}

func (t TestCredentials) Run(ctx context.Context) error {
	t.commenting(ctx)

	if t.failed {
		return fmt.Errorf("there were some errors, credentials might not be set up correctly")
	}
	return nil
}

func (t TestCredentials) commenting(ctx context.Context) {
	t.printf("commenting (pr:write)...")
	_, err := t.client.CommentIssue(ctx, t.pr, "test credentials "+t.now())
	if err != nil {
		t.printf("failed:\n%v\n", err)
		t.failed = true
	}
	t.printf("ok\n")
}

func (t TestCredentials) printf(format string, a ...any) {
	_, _ = fmt.Fprintf(t.output, format, a...)
}

func (t TestCredentials) now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
