package main

import (
	"os"
	"testing"
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
				if err == nil {
					t.Fatal(err)
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
			}
		})
		t.Run(tc.name, func(t *testing.T) {
			// setup environment variables
			os.Setenv("GITLAB_USERNAME", tc.user)
			os.Setenv("GITLAB_PASSWORD", tc.pass)
			os.Setenv("GITLAB_HTTP_URL", tc.url)

			// login using the environment variables
			_, err := newGitlabClient()
			if tc.negativeTest {
				if err == nil {
					t.Fatal(err)
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
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
			name:         "success",
			privateToken: "d1eQSbsjsXzfBuUhEVt1",
			apiURL:       "http://localhost:10080/api/v4",
			negativeTest: false,
		},
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
			if err != nil {
				t.Fatal(err)
			}
			// test a quick api call
			_, _, err = gitClient.Users.ListUsers(nil)
			if tc.negativeTest {
				if err == nil {
					t.Fatal(err)
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
			}
		})
		t.Run("gitlab new client test "+tc.name, func(t *testing.T) {
			// setup environment variables
			os.Unsetenv("GITLAB_USERNAME")
			os.Unsetenv("GITLAB_PASSWORD")
			os.Setenv("GITLAB_PRIVATE_TOKEN", tc.privateToken)
			os.Setenv("GITLAB_API_HTTP_URL", tc.apiURL)

			gitClient, err := newGitlabClient()
			if err != nil {
				t.Fatal(err)
			}
			// test a quick api call
			_, _, err = gitClient.Users.ListUsers(nil)
			if tc.negativeTest {
				if err == nil {
					t.Fatal(err)
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}
