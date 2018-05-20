package gitlabctl

import (
	"fmt"

	"github.com/xanzy/go-gitlab"

	"github.com/spf13/viper"
)

func newBasicAuthClient(username, password, basehttpURL string) (*gitlab.Client, error) {
	gitlabClient, err := gitlab.NewBasicAuthClient(nil,
		basehttpURL,
		username,
		password,
	)
	if err != nil {
		return nil, err
	}
	return gitlabClient, nil
}

func newClient(privateToken, apihttpURL string) (*gitlab.Client, error) {
	gitlabClient := gitlab.NewClient(nil, privateToken)
	if err := gitlabClient.SetBaseURL(apihttpURL); err != nil {
		return nil, err
	}
	return gitlabClient, nil
}

func newOAuthClient(oAuthToken, apihttpURL string) (*gitlab.Client, error) {
	gitlabClient := gitlab.NewOAuthClient(nil, oAuthToken)
	if err := gitlabClient.SetBaseURL(apihttpURL); err != nil {
		return nil, err
	}
	return gitlabClient, nil
}

func newGitlabClient() (*gitlab.Client, error) {
	if get("USERNAME") != "" && get("PASSWORD") != "" && get("HTTP_URL") != "" {
		gitlabClient, err := newBasicAuthClient(get("USERNAME"),
			get("PASSWORD"),
			get("HTTP_URL"),
		)
		if err != nil {
			return nil, err
		}
		return gitlabClient, nil
	}
	if get("PRIVATE_TOKEN") != "" && get("API_HTTP_URL") != "" {
		gitlabClient, err := newClient(get("PRIVATE_TOKEN"), get("API_HTTP_URL"))
		if err != nil {
			return nil, err
		}
		return gitlabClient, nil
	}
	if get("OAUTH_TOKEN") != "" && get("API_HTTP_URL") != "" {
		gitlabClient, err := newOAuthClient(get("OAUTH_TOKEN"), get("API_HTTP_URL"))
		if err != nil {
			return nil, err
		}
		return gitlabClient, nil
	}
	return nil, fmt.Errorf("no clients were created. GITLAB variables may not be set properly")
}

func get(e string) string {
	viper.SetEnvPrefix("GITLAB")
	if err := viper.BindEnv(e); err != nil {
		panic(err)
	}
	return viper.GetString(e)
}
