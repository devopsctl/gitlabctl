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

var descUserCmd = &cobra.Command{
	Use:           "user",
	Aliases:       []string{"u"},
	SuggestFor:    []string{"users"},
	Short:         "Describe a user",
	Example:       `gitlabctl describe user john.smith`,
	Args:          cobra.ExactArgs(0),
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runDescUser(cmd)
	},
}

func init() {
	descCmd.AddCommand(descUserCmd)
	descUserCmd.Flags().Int("id", -1, "The ID of a user")
	if err := descUserCmd.MarkFlagRequired("id"); err != nil {
		er(err)
	}
}

func runDescUser(cmd *cobra.Command) error {
	uid := getFlagInt(cmd, "id")
	u, err := descUser(uid)
	if err != nil {
		return err
	}
	printUsersOut(cmd, u)
	return nil
}

func descUser(uid int) (*gitlab.User, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	u, _, err := git.Users.GetUser(uid)
	if err != nil {
		return nil, err
	}
	return u, nil
}
