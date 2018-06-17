// Copyright © 2018 github.com/devopsctl authors
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
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

var descGroupCmd = &cobra.Command{
	Use:        "group",
	Aliases:    []string{"g"},
	SuggestFor: []string{"groups"},
	Short:      "Describe a group by specifying the id or group path",
	Example: `# describe a group
gitlabctl describe group GroupX -o json

# describe a group by id
gitlabctl describe group 13 -o json`,
	Args:              cobra.ExactArgs(1),
	SilenceErrors:     true,
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runDescGroup(cmd, args[0])
	},
}

func init() {
	descCmd.AddCommand(descGroupCmd)
}

func runDescGroup(cmd *cobra.Command, group string) error {
	groupInfo, err := descGroup(group)
	if err != nil {
		return err
	}
	printGroupsOut(getFlagString(cmd, "out"), groupInfo)
	return nil
}

func descGroup(group string) (*gitlab.Group, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	groupInfo, _, err := git.Groups.GetGroup(group)
	if err != nil {
		return nil, err
	}
	return groupInfo, nil
}
