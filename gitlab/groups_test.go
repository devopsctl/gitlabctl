package gitlabctl

import (
	"os"
)

// The values here should be what is set in the docker-compose.yml file
func setupGitlabEnvVars() {
	os.Setenv("GITLAB_USERNAME", "root")
	os.Setenv("GITLAB_PASSWORD", "123qwe123")
	os.Setenv("GITLAB_HTTP_URL", "http://localhost:10080")
}
