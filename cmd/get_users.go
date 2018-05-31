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
	"github.com/spf13/cobra"
	gitlab "github.com/xanzy/go-gitlab"
)

// getUsersCmd references:
// Gitlab API doc: https://docs.gitlab.com/ce/api/users.html#list-users
// Go Client doc: https://godoc.org/github.com/xanzy/go-gitlab#UsersService.ListUsers
var getUsersCmd = &cobra.Command{
	Use:           "users",
	Aliases:       []string{"u"},
	SuggestFor:    []string{"user"},
	Short:         "List all Gitlab users",
	Args:          cobra.ExactArgs(0),
	SilenceErrors: true,
	SilenceUsage:  true,
	Example:       `gitlabctl get users --out json`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := validateSortFlagValue(cmd); err != nil {
			return err
		}
		return validateUserOrderByFlagValue(cmd)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return runGetUsers(cmd)
	},
}

func init() {
	getCmd.AddCommand(getUsersCmd)
	addGetUsersFlags(getUsersCmd)
}

func runGetUsers(cmd *cobra.Command) error {
	opts, err := createListUsersOptions(cmd)
	if err != nil {
		return err
	}
	users, err := getUsers(opts)
	if err != nil {
		return err
	}
	printUsersOut(cmd, users...)
	return nil
}

func getUsers(opts *gitlab.ListUsersOptions) ([]*gitlab.User, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	users, _, err := git.Users.ListUsers(opts)
	if err != nil {
		return nil, err
	}
	return users, nil
}
