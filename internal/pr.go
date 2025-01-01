package internal

type PullRequest struct {
	ID        int
	Title     string
	Body      string
	Reviewers []string
}

func (pr PullRequest) HasReviewer() bool {
	return len(pr.Reviewers) > 0
}
