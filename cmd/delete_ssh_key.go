// Copyright © 2018 github.com/devopsctl authors
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteSSHKeyCmd = &cobra.Command{
	Use:   "ssh-key",
	Short: "Delete registered ssh keys",
	Long:  "Delete ssh keys of a gitlab user. Use 'gitlabctl get ssh' to get the list of ssh keys to delete.",
	Example: `# delete ssh key with id (23) for the current authenticated user
gitlabctl delete ssh 23

# delete ssh key for a user (for admins only)
gitlabctl delete ssh 23 --user=lebron.james`,
	Aliases:           []string{"ssh"},
	SuggestFor:        []string{"ssh-keys"},
	Args:              cobra.ExactArgs(1),
	DisableAutoGenTag: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runDeleteSSHKey(cmd, args[0])
	},
}

func init() {
	deleteCmd.AddCommand(deleteSSHKeyCmd)
	deleteSSHKeyCmd.Flags().StringP("user", "u", "",
		"The user which requires removal of ssh key")
}

func runDeleteSSHKey(cmd *cobra.Command, key string) error {
	keyID, err := strconv.Atoi(key)
	if err != nil {
		return err
	}
	if cmd.Flag("user").Changed {
		user := getFlagString(cmd, "user")
		uid, err := strconv.Atoi(user)
		if err != nil {
			uid, err = getUserIDbyUsername(user)
			if err != nil {
				return err
			}
		}
		return deleteSSHKeyForUser(uid, keyID)
	}
	return deleteSSHKey(keyID)
}

func deleteSSHKey(key int) error {
	git, err := newGitlabClient()
	if err != nil {
		return err
	}
	_, err = git.Users.DeleteSSHKey(key)
	if err != nil {
		return err
	}
	return nil
}

func deleteSSHKeyForUser(uid, key int) error {
	git, err := newGitlabClient()
	if err != nil {
		return err
	}
	_, err = git.Users.DeleteSSHKeyForUser(uid, key)
	if err != nil {
		return err
	}
	fmt.Printf("Deleted ssh key with id (%d) for user (%d)\n", key, uid)
	return nil
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
