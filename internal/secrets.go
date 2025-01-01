package internal

import (
	"os"
)

const GitHubToken SecretEnv = "GITHUB_TOKEN"

type SecretEnv string

func (s SecretEnv) String() string {
	return string(s) + "=*****"
}

func (s SecretEnv) Secret() string {
	return os.Getenv(string(s))
}

func (s SecretEnv) NotSet() bool {
	_, ok := os.LookupEnv(string(s))
	return !ok
}
