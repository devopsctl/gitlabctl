package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBasicAuthClient(t *testing.T) {
	tt := []struct {
		name, user, pass, url string
		negativeTest          bool
	}{
		{
			name:         "success basic auth login",
			user:         "root",
			pass:         "123qwe123",
			url:          "http://localhost:10080",
			negativeTest: false,
		},
		{
			name:         "401 basic auth login",
			user:         "unknown",
			pass:         "11111",
			url:          "http://localhost:10080",
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
		t.Run(tc.name, func(t *testing.T) {
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

		// TODO - use private token for testing
		// but if we enable this, we should also have travisci or other devs use the same token
		// we need to come up first with a strategy how to properly test this

		// {
		// 	name:         "success private token",
		// 	privateToken: "d1eQSbsjsXzfBuUhEVt1",
		// 	apiURL:       "http://localhost:10080/api/v4",
		// 	negativeTest: false,
		// },
		{
			name:         "token is wrong",
			privateToken: "invalidTokenxxxHehe",
			apiURL:       "http://localhost:10080/api/v4",
			negativeTest: true,
		},
		{
			name:         "api url is wrong",
			privateToken: "d1eQSbsjsXzfBuUhEVt1",
			apiURL:       "http://localhost:10080",
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
		t.Run("gitlab new client test "+tc.name, func(t *testing.T) {
			unsetEnv("GITLAB_USERNAME", "GITLAB_PASSWORD")
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
