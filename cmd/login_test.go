package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {
	stdout, err := executeCommand(rootCmd, "login", "-H=http://localhost:10080",
		"-p=123qwe123", "-u=root")
	tInfo(stdout)
	require.Nil(t, err, stdout)
	require.Contains(t, stdout,
		"gitlabctl.yaml file has been created by login command")
}
