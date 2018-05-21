package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TEST DATA
const (
	testingUser     = "root"
	testingPassword = "123qwe123"
	testingHTTPURL  = "http://localhost:10080"
	testingAPIURL   = "http://localhost:10080/api/v4"
)

var testingToken = os.Getenv("GITLAB_PRIVATE_TOKEN")
var testingOAuthToken = os.Getenv("GITLAB_OAUTH_TOKEN")

// The values here should be what is set in the docker-compose.yml file
func setupGitlabEnvVars() {
	setEnv("GITLAB_USERNAME", testingUser)
	setEnv("GITLAB_PASSWORD", testingPassword)
	setEnv("GITLAB_HTTP_URL", testingHTTPURL)
}

func TestEnvVariablesNotMetErr(t *testing.T) {
	unsetEnv("GITLAB_HTTP_URL", "GITLAB_API_HTTP_URL")
	_, err := newGitlabClient()
	assert.NotNil(t, err)
}

func TestOAuthClient(t *testing.T) {
	tt := []struct {
		name, oAuthToken, apiURL string
		negativeTest             bool
	}{
		{
			name:         "CORRECT_TOKEN_OK",
			oAuthToken:   testingOAuthToken,
			apiURL:       testingAPIURL,
			negativeTest: false,
		},
		{
			name:         "INVALID_TOKEN_FAILS",
			oAuthToken:   "xxxx-token-xxxx",
			apiURL:       testingAPIURL,
			negativeTest: true,
		},
	}
	for _, tc := range tt {
		t.Run("["+tc.name+"][WITH_TOKEN="+tc.oAuthToken+"]", func(t *testing.T) {
			// Ensure that oAuth authentication is used by
			// setting the required env variables
			// for other authentication type to blank
			unsetEnv("GITLAB_USERNAME", "GITLAB_PRIVATE_TOKEN")
			setEnv("GITLAB_OAUTH_TOKEN", tc.oAuthToken)
			setEnv("GITLAB_API_HTTP_URL", tc.apiURL)
			gitClient, err := newGitlabClient()
			assert.Nil(t, err)
			// test quick api call
			_, _, err = gitClient.Users.ListUsers(nil)
			if tc.negativeTest {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestNewBasicAuthClient(t *testing.T) {
	tt := []struct {
		name, user, pass, url string
		negativeTest          bool
	}{
		{
			name:         "CORRECT_CREDENTIALS_OK",
			user:         testingUser,
			pass:         testingPassword,
			url:          testingHTTPURL,
			negativeTest: false,
		},
		{
			name:         "INVALID_CREDENTIALS_FAILS",
			user:         "unknown",
			pass:         "11111",
			url:          testingHTTPURL,
			negativeTest: true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := newBasicAuthClient(tc.user, tc.pass, tc.url)
			if tc.negativeTest {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
		t.Run("["+tc.name+"][WITH_USER="+tc.user+",WITH_PASS="+tc.pass+"]",
			func(t *testing.T) {
				// setup environment variables
				setEnv("GITLAB_USERNAME", tc.user)
				setEnv("GITLAB_PASSWORD", tc.pass)
				setEnv("GITLAB_HTTP_URL", tc.url)

				// login using the environment variables
				_, err := newGitlabClient()
				if tc.negativeTest {
					assert.NotNil(t, err)
				} else {
					assert.Nil(t, err)
				}
			})
	}
}

func TestNewClient(t *testing.T) {
	tt := []struct {
		name, privateToken, apiURL string
		negativeTest               bool
	}{
		{
			name:         "CORRECT_TOKEN_OK",
			privateToken: testingToken,
			apiURL:       testingAPIURL,
			negativeTest: false,
		},
		{
			name:         "INVALID_TOKEN_FAILS",
			privateToken: "invalidTokenxxxHehe",
			apiURL:       testingAPIURL,
			negativeTest: true,
		},
		{
			name:         "INVALID_API_URL_FAILS",
			privateToken: testingToken,
			apiURL:       testingHTTPURL,
			negativeTest: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			gitClient, err := newClient(tc.privateToken, tc.apiURL)
			assert.Nil(t, err)
			// test a quick api call
			_, _, err = gitClient.Users.ListUsers(nil)
			if tc.negativeTest {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
		t.Run("["+tc.name+
			"][WITH_TOKEN="+tc.privateToken+"]",
			func(t *testing.T) {
				unsetEnv("GITLAB_USERNAME", "GITLAB_PASSWORD", "GITLAB_HTTP_URL")
				setEnv("GITLAB_PRIVATE_TOKEN", tc.privateToken)
				setEnv("GITLAB_API_HTTP_URL", tc.apiURL)

				gitClient, err := newGitlabClient()
				assert.Nil(t, err)
				// test a quick api call
				_, _, err = gitClient.Users.ListUsers(nil)
				if tc.negativeTest {
					assert.NotNil(t, err)
				} else {
					assert.Nil(t, err)
				}
			})
	}
}
