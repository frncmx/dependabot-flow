package graphql

import (
	"context"

	"github.com/cli/go-gh/v2/pkg/api"
	graphql "github.com/cli/shurcooL-graphql"

	"github.com/frncmx/dependabot-flow/internal/config"
)

func NewClient(token config.Secret[string]) Client {
	client, err := api.NewGraphQLClient(api.ClientOptions{
		AuthToken: token.Value(),
	})
	if err != nil {
		panic("new graphql client: " + err.Error())
	}
	return Client{
		gh: client,
	}
}

type Client struct {
	gh *api.GraphQLClient
}

func (c Client) GetPR(ctx context.Context, owner, repo string, number int) (PR, error) {
	var query struct {
		Repository struct {
			PullRequest struct {
				ID string
			} `graphql:"pullRequest(number: $number)"`
		} `graphql:"repository(owner: $owner, name: $repo)"`
	}
	variables := map[string]interface{}{
		"owner":  graphql.String(owner),
		"repo":   graphql.String(repo),
		"number": graphql.Int(number),
	}
	err := c.gh.QueryWithContext(ctx, "", &query, variables)
	return PR{ID: query.Repository.PullRequest.ID}, err
}

type PR struct {
	ID string `json:"id"`
}

func (c Client) EnableAutoMerge(ctx context.Context, id string) error {
	var mutation struct {
		EnablePullRequestAutoMerge struct {
			PullRequest struct {
				ID string
			}
		} `graphql:"enablePullRequestAutoMerge(input: { pullRequestId: $pullRequestId })"`
	}
	variables := map[string]interface{}{
		"pullRequestId": graphql.ID(id),
	}
	return c.gh.MutateWithContext(ctx, "", &mutation, variables)
}

func (c Client) DisableAutoMerge(ctx context.Context, id string) error {
	var mutation struct {
		EnablePullRequestAutoMerge struct {
			PullRequest struct {
				ID string
			}
		} `graphql:"disablePullRequestAutoMerge(input: { pullRequestId: $pullRequestId })"`
	}
	variables := map[string]interface{}{
		"pullRequestId": graphql.ID(id),
	}
	return c.gh.MutateWithContext(ctx, "", &mutation, variables)
}
