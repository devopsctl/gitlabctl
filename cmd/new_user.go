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

	"github.com/spf13/cobra"
	gitlab "github.com/xanzy/go-gitlab"
)

var newUserCmd = &cobra.Command{
	Use:        "user",
	Aliases:    []string{"u"},
	SuggestFor: []string{"users"},
	Short:      "Create a new user by specifying the username as the first argument",
	Example: `# create a new user
gitlabctl new user john.smith --name="Johhny Smith" --password=12345678 --email=john.smith@example.com --skip-confirmation

# create a new user and send reset password link
gitlabctl new user james --name="james" --password=aaaaaaaa --email=aa@example.com --reset-password`,
	Args:          cobra.ExactArgs(1),
	SilenceErrors: true,
	SilenceUsage:  true,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if getFlagBool(cmd, "reset-password") == false &&
			getFlagString(cmd, "password") == "" {
			return fmt.Errorf("password, reset-password are missing, " +
				"at least one parameter must be provided")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return runNewUser(cmd, args[0])
	},
}

func init() {
	newCmd.AddCommand(newUserCmd)
	addNewUserFlags(newUserCmd)
}

func runNewUser(cmd *cobra.Command, username string) error {
	opts, err := assignCreateUserOptions(cmd)
	if err != nil {
		return err
	}
	opts.Username = gitlab.String(username)
	user, err := newUser(opts)
	if err != nil {
		return err
	}
	printUsersOut(cmd, user)
	return nil
}

func newUser(opt *gitlab.CreateUserOptions) (*gitlab.User, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	user, _, err := git.Users.CreateUser(opt, nil)
	if err != nil {
		return nil, err
	}
	return user, nil
}
