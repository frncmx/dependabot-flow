package internal

import (
	"bufio"
	"regexp"
	"strings"
)

var (
	breaking      = regexp.MustCompile("(?i)breaking|breaks|major|incompatible")
	filteredWords = []string{
		// When the package does not follow semantic import versioning:
		"+incompatible",
		"majority",
	}
)

type PullRequest struct {
	Number    int
	URL       string
	Body      string
	Reviewers []string
}

func (pr PullRequest) HasReviewer() bool {
	return len(pr.Reviewers) > 0
}

func (pr PullRequest) Safe() bool {
	return pr.SuspiciousText() == ""
}

func (pr PullRequest) SuspiciousText() string {
	var result stringsBuilder
	scanner := bufio.NewScanner(strings.NewReader(pr.Body))
	for scanner.Scan() {
		line := scanner.Text()
		if isDependabotHelp(line) {
			continue
		}
		filtered := filterWords(line, filteredWords)
		if breaking.MatchString(filtered) {
			result.write(line, nl, nl)
		}
	}
	return strings.TrimSuffix(result.string(), nl)
}

func (pr PullRequest) SuspiciousPattern() string {
	return breaking.String()
}

func isDependabotHelp(s string) bool {
	return strings.HasPrefix(s, "- `@dependabot")
}

func filterWords(s string, words []string) string {
	for _, word := range words {
		s = strings.ReplaceAll(s, word, "")
	}
	return s
}
