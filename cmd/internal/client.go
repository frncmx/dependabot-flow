package internal

import (
	"context"

	"github.com/google/go-github/v68/github"
)

func NewClient(token, owner, repo string) Client {
	return Client{
		gh:    github.NewClient(nil).WithAuthToken(token),
		owner: owner,
		repo:  repo,
	}
}

type Client struct {
	gh          *github.Client
	owner, repo string
}

func (c Client) Approve(ctx context.Context, id int, comment string) error {
	review := &github.PullRequestReviewRequest{
		Event: github.Ptr("APPROVE"),
		Body:  &comment,
	}
	_, _, err := c.gh.PullRequests.CreateReview(ctx, c.owner, c.repo, id, review)
	return err
}

func (c Client) CommentIssue(ctx context.Context, id int, comment string) error {
	// Use Issues API as PR API requires commit ID or Reply-to.
	issueComment := &github.IssueComment{Body: &comment}
	_, _, err := c.gh.Issues.CreateComment(ctx, c.owner, c.repo, id, issueComment)
	return err
}

func (c Client) RequestReview(ctx context.Context, id int, user string) error {
	request := github.ReviewersRequest{Reviewers: []string{user}}
	_, _, err := c.gh.PullRequests.RequestReviewers(ctx, c.owner, c.repo, id, request)
	return err
}

func (c Client) Label(ctx context.Context, pr int, label string) error {
	_, _, err := c.gh.Issues.AddLabelsToIssue(ctx, c.owner, c.repo, pr, []string{label})
	return err
}
