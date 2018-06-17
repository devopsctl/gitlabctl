package cmd

import (
	"fmt"

	gitlab "github.com/xanzy/go-gitlab"

	"github.com/spf13/viper"
)

func newBasicAuthClient(username, password,
	basehttpURL string) (*gitlab.Client, error) {
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
	switch {
	case getCfg("access_token") != "" && getCfg("host_url") != "":
		gitlabClient, err :=
			newOAuthClient(getCfg("access_token"), getCfg("host_url")+"/api/v4")
		if err != nil {
			return nil, err
		}
		return gitlabClient, nil

	case getCfg("USERNAME") != "" &&
		getCfg("PASSWORD") != "" &&
		getCfg("HTTP_URL") != "":
		gitlabClient, err := newBasicAuthClient(getCfg("USERNAME"),
			getCfg("PASSWORD"),
			getCfg("HTTP_URL"),
		)
		if err != nil {
			return nil, err
		}
		return gitlabClient, nil

	case getCfg("PRIVATE_TOKEN") != "" && getCfg("API_HTTP_URL") != "":
		gitlabClient, err :=
			newClient(getCfg("PRIVATE_TOKEN"), getCfg("API_HTTP_URL"))
		if err != nil {
			return nil, err
		}
		return gitlabClient, nil

	case getCfg("PRIVATE_TOKEN") != "" && getCfg("API_HTTP_URL") != "":
		gitlabClient, err :=
			newClient(getCfg("PRIVATE_TOKEN"), getCfg("API_HTTP_URL"))
		if err != nil {
			return nil, err
		}
		return gitlabClient, nil

	case getCfg("OAUTH_TOKEN") != "" && getCfg("API_HTTP_URL") != "":
		gitlabClient, err :=
			newOAuthClient(getCfg("OAUTH_TOKEN"), getCfg("API_HTTP_URL"))
		if err != nil {
			return nil, err
		}
		return gitlabClient, nil

	default:
		return nil, fmt.Errorf("no client was created. "+
			"gitlab configuration was not set properly. \n %s", authDoc)
	}
}

// getCfg retrieves the value of a token from the config '--config' file
// in 'key=value' format or from any environment variable that starts
// with 'GITLAB' as prefix; e.g: GITLAB_USERNAME
func getCfg(e string) string {
	viper.SetEnvPrefix("GITLAB")
	if err := viper.BindEnv(e); err != nil {
		panic(err)
	}
	return viper.GetString(e)
}
