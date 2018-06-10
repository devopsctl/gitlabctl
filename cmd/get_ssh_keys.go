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
	"strconv"

	"github.com/spf13/cobra"
	gitlab "github.com/xanzy/go-gitlab"
)

var getSSHKeysCmd = &cobra.Command{
	Use:     "ssh-keys",
	Aliases: []string{"ssh", "ssh-key"},
	Example: `# get a list of currently authenticated user ssh keys
gitlabctl get ssh-keys

# get a list of a specific user ssh keys (admin only)
gitlabctl get ssh-keys --user="lebron.james"`,
	Args:          cobra.ExactArgs(0),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Flag("user").Changed {
			return runGetSSHKeyForUser(cmd)
		}
		return runGetSSHKey(cmd)
	},
}

func init() {
	getCmd.AddCommand(getSSHKeysCmd)
	getSSHKeysCmd.Flags().StringP("user", "u", "", "The username or ID of a user")
}

func runGetSSHKeyForUser(cmd *cobra.Command) error {
	opts := new(gitlab.ListSSHKeysForUserOptions)
	if cmd.Flag("page").Changed {
		opts.Page = getFlagInt(cmd, "page")
	}
	if cmd.Flag("per-page").Changed {
		opts.PerPage = getFlagInt(cmd, "per-page")
	}
	user := getFlagString(cmd, "user")
	uid, err := strconv.Atoi(user)
	if err != nil {
		userInfo, err := getUserByUsername(user)
		if err != nil {
			return err
		}
		uid = userInfo.ID
	}
	sshKeys, err := getSSHKeysForUser(uid, opts)
	if err != nil {
		return err
	}
	printSSHKeysOut(cmd, sshKeys...)
	return nil
}

func runGetSSHKey(cmd *cobra.Command) error {
	sshKeys, err := getSSHKeys()
	if err != nil {
		return err
	}
	printSSHKeysOut(cmd, sshKeys...)
	return nil
}

func getSSHKeys() ([]*gitlab.SSHKey, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	keys, _, err := git.Users.ListSSHKeys(nil)
	if err != nil {
		return nil, err
	}
	return keys, nil
}

func getSSHKeysForUser(uid int, opts *gitlab.ListSSHKeysForUserOptions) ([]*gitlab.SSHKey, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	keys, _, err := git.Users.ListSSHKeysForUser(uid, opts)
	if err != nil {
		return nil, err
	}
	return keys, nil
}
