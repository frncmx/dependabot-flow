package config

import (
	"os"
)

const GitHubToken SecretEnv = "GITHUB_TOKEN"

type SecretEnv string

func (s SecretEnv) NotSet() bool {
	_, ok := os.LookupEnv(string(s))
	return !ok
}

func (s SecretEnv) Secret() Secret[string] {
	return Secret[string]{
		value: os.Getenv(string(s)),
	}

}

func (s SecretEnv) String() string {
	return string(s) + "=*****"
}

type Secret[T any] struct {
	value T
}

func (s Secret[T]) Value() T {
	return s.value
}
