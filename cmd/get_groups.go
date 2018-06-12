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

var getGroupsCmd = &cobra.Command{
	Use:           "groups",
	Aliases:       []string{"g"},
	SuggestFor:    []string{"group"},
	SilenceErrors: true,
	SilenceUsage:  true,
	Short:         "List groups and subgroups",
	Example: `# list all groups
gitlabctl get groups

# list all subgroups of GroupX
gitlabctl get groups --from-group=GroupX`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := validateGroupOrderByFlagValue(cmd); err != nil {
			return err
		}
		return validateSortFlagValue(cmd)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if getFlagString(cmd, "from-group") != "" {
			return runGetSubgroups(cmd)
		}
		return runGetGroups(cmd)
	},
}

func init() {
	getCmd.AddCommand(getGroupsCmd)
	addGetGroupsFlags(getGroupsCmd)
	addFromGroupFlag(getGroupsCmd)
}

func runGetGroups(cmd *cobra.Command) error {
	opts := assignListGroupOptions(cmd)
	groups, err := getGroups(opts)
	if err != nil {
		return err
	}
	printGroupsOut(cmd, groups...)
	return nil
}

func getGroups(opts *gitlab.ListGroupsOptions) ([]*gitlab.Group, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	groups, _, err := git.Groups.ListGroups(opts)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func runGetSubgroups(cmd *cobra.Command) error {
	// to reuse the same opts mapping from groupLsCmd for groupLsSubGroup
	// convert gitlab.ListGroupsOptions to gitlab.ListSubgroupsOptions
	opts := (*gitlab.ListSubgroupsOptions)(assignListGroupOptions(cmd))
	path := getFlagString(cmd, "from-group")
	groups, err := getSubgroups(path, opts)
	if err != nil {
		return err
	}
	printGroupsOut(cmd, groups...)
	return nil
}

func getSubgroups(gid interface{},
	opts *gitlab.ListSubgroupsOptions) ([]*gitlab.Group, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	groups, _, err := git.Groups.ListSubgroups(gid, opts)
	if err != nil {
		return nil, err
	}
	return groups, nil
}
