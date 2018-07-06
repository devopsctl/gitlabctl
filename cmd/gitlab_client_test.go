package cmd

import (
	"os"
	"testing"

	gitlab "github.com/xanzy/go-gitlab"
)

// Get the correct test data from the users environment variables
// getE is binded to have "GITLAB_" as prefix
var (
	// base this from docker-compose.yml
	testingUser     = getCfg("USERNAME")
	testingPassword = getCfg("PASSWORD")
	testingHTTPURL  = getCfg("HTTP_URL")
	testingAPIURL   = getCfg("API_HTTP_URL")

	// run ./testdata/get_oauth_token.sh
	testingToken = getCfg("PRIVATE_TOKEN")

	// run ./testdata/get_token.sh
	testingOAuthToken = getCfg("OAUTH_TOKEN")
)

func setBasicAuthEnvs() {
	setE("GITLAB_USERNAME", testingUser)
	setE("GITLAB_PASSWORD", testingPassword)
	setE("GITLAB_HTTP_URL", testingHTTPURL)
}

func setPrivateTokenEnvs() {
	setE("GITLAB_PRIVATE_TOKEN", testingToken)
	setE("GITLAB_API_HTTP_URL", testingAPIURL)
}

func setOAuthTokenEnvs() {
	setE("GITLAB_OAUTH_TOKEN", testingOAuthToken)
	setE("GITLAB_API_HTTP_URL", testingAPIURL)
}

func unsetAuthEnvs() {
	unsetE("GITLAB_PRIVATE_TOKEN",
		"GITLAB_OAUTH_TOKEN",
		"GITLAB_API_HTTP_URL",
		"GITLAB_HTTP_URL",
		"GITLAB_USERNAME",
		"GITLAB_PASSWORD")
	// ensure to remove the config file
	// this causes testing missing variable scenario in gitlab_client_test.go
	if err := os.Remove(cfgFile); err != nil {
		tInfo(err)
	}
}

func TestGitlabNewClient(t *testing.T) {
	tt := []struct {
		name         string
		setEnv       func()
		negativeTest bool
	}{
		{"BASIC_AUTH_OK", setBasicAuthEnvs, false},
		{"OAUTH_TOKEN_OK", setOAuthTokenEnvs, false},
		{"PRIVATE_TOKEN_OK", setPrivateTokenEnvs, false},
		{"MISSING_ENVS_FAILS", unsetAuthEnvs, true},
	}

	for _, tc := range tt {
		// Setup before each test case
		unsetAuthEnvs()
		tc.setEnv()

		// Run the test
		t.Run(tc.name, func(t *testing.T) {
			gitlabClient, err := newGitlabClient()
			if err != nil && !tc.negativeTest {
				t.Fatalf("gitlab client test is expected to pass: %+v", err)
			}
			if err == nil && tc.negativeTest {
				t.Fatalf("gitlab client test is expected to fail: %+v", err)
			}
			if !tc.negativeTest {
				_, _, err = gitlabClient.Users.ListUsers(&gitlab.ListUsersOptions{})
				if err != nil {
					t.Fatalf("gitlab client test is expected to pass: %+v", err)
				}
			}
		})
	}
}
