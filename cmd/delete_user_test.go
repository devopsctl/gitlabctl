package cmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	gitlab "github.com/xanzy/go-gitlab"
)

func TestDeleteUser(t *testing.T) {
	tt := []struct {
		name   string
		args   []string
		expect testResult
	}{
		{
			name:   "successfully delete a user",
			args:   []string{"john.wick"},
			expect: pass,
		},
		{
			name:   "delete a non existent user fails",
			args:   []string{"john.missing"},
			expect: fail,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// SETUP for passing test
			// Create the user before deleting it
			if tc.expect == pass {
				_, err := newUser(&gitlab.CreateUserOptions{
					Name:     gitlab.String(tc.args[0]),
					Username: gitlab.String(tc.args[0]),
					Password: gitlab.String(tc.args[0]),
					Email:    gitlab.String(tc.args[0] + "@example.com"),
				})
				if err != nil {
					tInfo(fmt.Sprintf("setup error: %v", err))
				}
			}
			execT := execTestCmdFlags{
				t:    t,
				cmd:  deleteUserCmd,
				args: tc.args,
			}
			stdout, execResult := execT.executeCommand()
			tInfo(stdout)
			require.Equal(t, tc.expect, execResult, stdout)
		})
	}
}
