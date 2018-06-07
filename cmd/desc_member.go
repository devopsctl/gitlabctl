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

var descMemberCmd = &cobra.Command{
	Use:        "member",
	Aliases:    []string{"m"},
	SuggestFor: []string{"members"},
	Short:      "Describe a member by specifying the username and source",
	Example: `# describe a member from a group
gitlabctl describe member john.smith --from-group Group1 -o json

# describe a member from a project
gitlabctl describe member john.smith --from-project Group1/Project1 -o yaml`,
	Args:          cobra.ExactArgs(1),
	SilenceErrors: true,
	SilenceUsage:  true,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return validateFromGroupAndProjectFlags(cmd)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if getFlagString(cmd, "from-group") != "" {
			return runDescGroupMember(cmd, args[0])
		}
		return runDescProjectMember(cmd, args[0])
	},
}

func init() {
	descCmd.AddCommand(descMemberCmd)
	addFromGroupFlag(descMemberCmd)
	addFromProjectFlag(descMemberCmd)
}

func runDescGroupMember(cmd *cobra.Command, username string) error {
	group := getFlagString(cmd, "from-group")
	member, err := descGroupMember(group, username)
	if err != nil {
		return err
	}
	printGroupMembersOut(cmd, member)
	return err
}

func runDescProjectMember(cmd *cobra.Command, username string) error {
	group := getFlagString(cmd, "from-project")
	member, err := descProjectMember(group, username)
	if err != nil {
		return err
	}
	printProjectMembersOut(cmd, member)
	return err
}

func descProjectMember(pid interface{}, user string) (
	*gitlab.ProjectMember, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	uid, err := strconv.Atoi(user)
	if err != nil {
		foundUser, err := getUserByUsername(user)
		if err != nil {
			return nil, err
		}
		uid = foundUser.ID
	}

	member, _, err := git.ProjectMembers.GetProjectMember(pid, uid, nil)
	if err != nil {
		return nil, err
	}
	return member, nil
}

func descGroupMember(gid interface{}, user string) (
	*gitlab.GroupMember, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	uid, err := strconv.Atoi(user)
	if err != nil {
		foundUser, err := getUserByUsername(user)
		if err != nil {
			return nil, err
		}
		uid = foundUser.ID
	}

	members, _, err := git.GroupMembers.GetGroupMember(gid, uid, nil)
	if err != nil {
		return nil, err
	}
	return members, nil
}
