package cmd

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewSSHKey(t *testing.T) {
	tt := []struct {
		name     string
		flagsMap map[string]string
		expect   testResult
	}{
		{
			name: "add ssh key for current user",
			flagsMap: map[string]string{
				"title":   "gotest ssh key",
				"keyfile": "../testdata/root_ssh_key.pub",
			},
			expect: pass,
		},
		{
			name: "add ssh key for another user",
			flagsMap: map[string]string{
				"title":   "gotest ssh key",
				"keyfile": "../testdata/matt_ssh_key.pub",
				"user":    "matt.hunter",
			},
			expect: pass,
		},
		{
			name: "add ssh key for another user by id",
			flagsMap: map[string]string{
				"title":   "gotest ssh key",
				"keyfile": "../testdata/paul_ssh_key.pub",
				"user":    "12", // paul.lyman
			},
			expect: pass,
		},
		{
			name: "add non existing ssh key file fails",
			flagsMap: map[string]string{
				"title":   "gotest ssh key",
				"keyfile": "../testdata/xxxx.pub",
			},
			expect: fail,
		},
	}

	// SETUP
	// Ensure to delete all ssh keys attached to users
	users := []string{"root", "12", "matt.hunter"}
	for _, user := range users {
		if err := deleteAllSSHKeyForUser(user); err != nil {
			tInfo(fmt.Sprintf("teardown failure, don't mind me: %v", err))
		}
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			execT := execTestCmdFlags{
				t:        t,
				cmd:      newSSHKeyCmd,
				flagsMap: tc.flagsMap,
			}
			stdout, execResult := execT.executeCommand()
			require.Equal(t, tc.expect, execResult, stdout)

		})
	}
}

// deleteAllSSHKeyForUser is currently used for test teardown
func deleteAllSSHKeyForUser(user string) error {
	uid, err := strconv.Atoi(user)
	if err != nil {
		uid, err = getUserIDbyUsername(user)
		if err != nil {
			return err
		}
	}
	userKeys, err := getSSHKeysForUser(uid, nil)
	if err != nil {
		return err
	}
	for _, key := range userKeys {
		if err := deleteSSHKeyForUser(uid, key.ID); err != nil {
			return err
		}
	}
	return nil
}
