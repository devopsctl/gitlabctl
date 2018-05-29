package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Get the correct test data from the users environment variables
// getE is binded to have "GITLAB_" as prefix
var (
	// base this from docker-compose.yml
	testingUser     = getE("USERNAME")
	testingPassword = getE("PASSWORD")
	testingHTTPURL  = getE("HTTP_URL")
	testingAPIURL   = getE("API_HTTP_URL")

	// run ./testdata/get_oauth_token.sh
	testingToken = getE("PRIVATE_TOKEN")

	// run ./testdata/get_token.sh
	testingOAuthToken = getE("OAUTH_TOKEN")
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
		"GITLAB_USERNAME",
		"GITLAB_PASSWORD")
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
			if tc.negativeTest {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
			if !tc.negativeTest {
				_, _, err = gitlabClient.Users.ListUsers(nil)
				assert.Nil(t, err)
			}
		})

	}
}
