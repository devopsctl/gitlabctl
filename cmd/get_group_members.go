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
	"log"

	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
)

var getGroupMembersCmd = &cobra.Command{
	Use:     "group-members",
	Aliases: []string{"group-member"},
	Short:   "List all members of a group",
	Run: func(cmd *cobra.Command, args []string) {
		if err := runGetGroupMembers(cmd); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	getCmd.AddCommand(getGroupMembersCmd)
	addPathFlag(getGroupMembersCmd)
	addJSONFlag(getGroupMembersCmd)
}

func runGetGroupMembers(cmd *cobra.Command) error {
	path := getFlagString(cmd, "path")
	groupsMembers, err := listGroupsMembers(path, nil)
	if err != nil {
		return err
	}
	printGroupMembersOut(cmd, groupsMembers)
	return err
}

func listGroupsMembers(gid interface{}, opts *gitlab.ListGroupMembersOptions) ([]*gitlab.GroupMember, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	g, _, err := git.Groups.ListGroupMembers(gid, opts)
	if err != nil {
		return nil, err
	}
	return g, nil
}
