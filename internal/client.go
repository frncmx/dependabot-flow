package internal

import (
	"context"

	"github.com/google/go-github/v68/github"

	"github.com/frncmx/dependabot-flow/internal/config"
	"github.com/frncmx/dependabot-flow/internal/graphql"
)

func NewClient(token config.Secret[string], owner, repo string) Client {
	return Client{
		http:    github.NewClient(nil).WithAuthToken(token.Value()),
		graphql: graphql.NewClient(token),
		owner:   owner,
		repo:    repo,
	}
}

type Client struct {
	http        *github.Client
	graphql     graphql.Client
	owner, repo string
}

func (c Client) ApprovePR(ctx context.Context, number int, comment string) error {
	review := &github.PullRequestReviewRequest{
		Event: github.Ptr("APPROVE"),
		Body:  &comment,
	}
	_, _, err := c.http.PullRequests.CreateReview(ctx, c.owner, c.repo, number, review)
	return err
}

func (c Client) Comment(ctx context.Context, number int, comment string) error {
	// Use Issues API as PR API requires commit ID or Reply-to.
	issueComment := &github.IssueComment{Body: &comment}
	_, _, err := c.http.Issues.CreateComment(ctx, c.owner, c.repo, number, issueComment)
	return err
}

func (c Client) DisableAutoMerge(ctx context.Context, number int) error {
	pr, err := c.graphql.GetPR(ctx, c.owner, c.repo, number)
	if err != nil {
		return err
	}
	return c.graphql.DisableAutoMerge(ctx, pr.ID)
}

func (c Client) EnableAutoMerge(ctx context.Context, number int) error {
	// Note:
	// - contents:write give pr:read permission
	// - the GraphQL ID cannot be acquired via HTTP API
	pr, err := c.graphql.GetPR(ctx, c.owner, c.repo, number)
	if err != nil {
		return err
	}
	return c.graphql.EnableAutoMerge(ctx, pr.ID)
}

func (c Client) GetPR(ctx context.Context, number int) (PullRequest, error) {
	response, _, err := c.http.PullRequests.Get(ctx, c.owner, c.repo, number)
	if err != nil {
		return PullRequest{}, err
	}
	result := PullRequest{
		Title: response.GetTitle(),
		Body:  response.GetBody(),
	}
	for _, reviewer := range response.RequestedReviewers {
		result.Reviewers = append(result.Reviewers, reviewer.GetLogin())
	}
	return result, nil
}

func (c Client) Label(ctx context.Context, number int, label string) error {
	_, _, err := c.http.Issues.AddLabelsToIssue(ctx, c.owner, c.repo, number, []string{label})
	return err
}

func (c Client) RequestReview(ctx context.Context, number int, user string) error {
	request := github.ReviewersRequest{Reviewers: []string{user}}
	_, _, err := c.http.PullRequests.RequestReviewers(ctx, c.owner, c.repo, number, request)
	return err
}
