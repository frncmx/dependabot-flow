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

func (c Client) CommentIssue(ctx context.Context, id int, comment string) (int64, error) {
	// Use Issues API as PR API requires commit ID or Reply-to.
	ic := &github.IssueComment{Body: &comment}
	response, _, err := c.gh.Issues.CreateComment(ctx, c.owner, c.repo, id, ic)
	if err != nil {
		return 0, err
	}
	return response.GetID(), nil
}
