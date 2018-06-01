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

var editGroupCmd = &cobra.Command{
	Use:        "group",
	Aliases:    []string{"g"},
	SuggestFor: []string{"groups"},
	Short:      "Update or patch a group",
	Example: `# edit a group
gitlabctl edit group GroupAZ [flags]

# edit a subgroup using namespace
gitlabctl edit group GroupXB --namespace=ParentGroupXB [flags]
`,
	Args:          cobra.ExactArgs(1),
	SilenceErrors: true,
	SilenceUsage:  true,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return validateVisibilityFlagValue(cmd)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return runEditGroup(cmd, args[0])
	},
}

func init() {
	editCmd.AddCommand(editGroupCmd)
	addEditGroupFlags(editGroupCmd)
}

func runEditGroup(cmd *cobra.Command, name string) error {
	opts, err := assignCreateGroupOptions(cmd)
	if err != nil {
		return err
	}
	gid, err := getGroupID(name)
	if err != nil {
		return fmt.Errorf("couldn't find the id of group %s, got error: %v",
			name, err)
	}
	group, err := editGroup(gid, (*gitlab.UpdateGroupOptions)(opts))
	if err != nil {
		return err
	}
	printGroupsOut(cmd, group)
	return nil
}

func editGroup(gid int, opts *gitlab.UpdateGroupOptions) (*gitlab.Group, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	g, _, err := git.Groups.UpdateGroup(gid, opts)
	if err != nil {
		return nil, err
	}
	return g, nil
}
