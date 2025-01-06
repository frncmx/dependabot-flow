package internal

import (
	"strings"
)

const nl = "\n"

type CommentWithQuote struct {
	Message, Quote string
}

func (c CommentWithQuote) String() string {
	return c.Markdown()
}

func (c CommentWithQuote) Markdown() string {
	c.Message = strings.TrimSpace(c.Message)
	c.Quote = strings.TrimSpace(c.Quote)

	var result stringsBuilder
	result.write(c.Message)
	if c.Quote == "" {
		return result.string()
	}
	result.write(nl, nl)
	quote := strings.Split(c.Quote, nl)
	for _, line := range quote {
		result.write(">")
		if strings.TrimSpace(line) != "" {
			result.write(" ", line)
		}
		result.write(nl)
	}
	return strings.TrimSuffix(result.string(), nl)
}

type stringsBuilder struct {
	b strings.Builder
}

func (s *stringsBuilder) write(v ...string) {
	for _, str := range v {
		_, _ = s.b.WriteString(str)
	}
}

func (s *stringsBuilder) string() string {
	return s.b.String()
}
