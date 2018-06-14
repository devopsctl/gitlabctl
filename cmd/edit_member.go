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

var editMemberCmd = &cobra.Command{
	Use:        "member",
	Aliases:    []string{"m"},
	SuggestFor: []string{"members"},
	Short:      "Edit a member by specifying the member name as the first argument",
	Example: `# create a new group
gitlabctl edit member john.smith --from-group Group2 --access-level 20

# create a subgroup using namespace
gitlabctl edit member john.smith --from-project Project1 --expire-at 2069-06-25`,
	Args:          cobra.ExactArgs(1),
	SilenceErrors: true,
	SilenceUsage:  true,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := validateFromGroupAndProjectFlags(cmd); err != nil {
			return err
		}
		return validateAccessLevelFlag(cmd)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if getFlagString(cmd, "from-group") != "" {
			return runEditGroupMember(cmd, args[0])
		}
		return runEditProjectMember(cmd, args[0])
	},
}

func init() {
	editCmd.AddCommand(editMemberCmd)
	addFromGroupFlag(editMemberCmd)
	addFromProjectFlag(editMemberCmd)
	addAccessLevelFlag(editMemberCmd)
	addExpiresAtFlag(editMemberCmd)
}

func runEditProjectMember(cmd *cobra.Command, user string) error {
	opts, err := assignEditProjectMemberOptions(cmd)
	if err != nil {
		return err
	}
	project := getFlagString(cmd, "from-project")
	member, err := editProjectMember(project, user, opts)
	if err != nil {
		return err
	}
	printProjectMembersOut(cmd, member)
	return err
}

func runEditGroupMember(cmd *cobra.Command, user string) error {
	opts, err := assignEditGroupMemberOptions(cmd)
	if err != nil {
		return err
	}
	group := getFlagString(cmd, "from-group")
	member, err := editGroupMember(group, user, opts)
	if err != nil {
		return err
	}
	printGroupMembersOut(cmd, member)
	return err
}

func editProjectMember(pid interface{}, user string,
	opts *gitlab.EditProjectMemberOptions) (*gitlab.ProjectMember, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}

	foundUser, err := getUserByUsername(user)
	if err != nil {
		return nil, err
	}

	member, _, err := git.ProjectMembers.EditProjectMember(pid, foundUser.ID, opts)
	if err != nil {
		return nil, err
	}
	return member, nil
}

func editGroupMember(gid interface{}, user string,
	opts *gitlab.EditGroupMemberOptions) (*gitlab.GroupMember, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}

	foundUser, err := getUserByUsername(user)
	if err != nil {
		return nil, err
	}
	member, _, err := git.GroupMembers.EditGroupMember(gid, foundUser.ID, opts)
	if err != nil {
		return nil, err
	}
	return member, nil
}
