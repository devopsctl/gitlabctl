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

var newGroupCmd = &cobra.Command{
	Use:        "group",
	Aliases:    []string{"g"},
	SuggestFor: []string{"groups"},
	Short:      "Create a new group by specifying the group name as the first argument",
	Example: `# create a new group
gitlabctl new group GroupAZ

# create a subgroup using namespace
gitlabctl new group GroupXB --namespace=ParentGroupXB`,
	Args:          cobra.ExactArgs(1),
	SilenceErrors: true,
	SilenceUsage:  true,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return validateVisibilityFlagValue(cmd)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return runNewGroup(cmd, args[0])
	},
}

func init() {
	newCmd.AddCommand(newGroupCmd)
	addNewGroupFlags(newGroupCmd)
}

func runNewGroup(cmd *cobra.Command, name string) error {
	opts, err := assignCreateGroupOptions(cmd)
	if err != nil {
		return err
	}
	opts.Path = gitlab.String(name)
	opts.Name = gitlab.String(name)
	group, err := newGroup(opts)
	if err != nil {
		return err
	}
	printGroupsOut(cmd, group)
	return nil
}

func newGroup(opts *gitlab.CreateGroupOptions) (*gitlab.Group, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	g, _, err := git.Groups.CreateGroup(opts)
	if err != nil {
		return nil, err
	}
	return g, nil
}
