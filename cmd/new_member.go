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

var newMemberCmd = &cobra.Command{
	Use:        "member",
	Aliases:    []string{"m"},
	SuggestFor: []string{"members"},
	Short:      "Create a new member by specifying the member name as the first argument",
	Example: `# create a new group
gitlabctl new member john.smith --from-group Group2 

# create a subgroup using namespace
gitlabctl new member john.smith --from-project Project1`,
	Args:              cobra.ExactArgs(1),
	SilenceErrors:     true,
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := validateFromGroupAndProjectFlags(cmd); err != nil {
			return err
		}
		return validateAccessLevelFlag(cmd)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if getFlagString(cmd, "from-group") != "" {
			return runNewGroupMember(cmd, args[0])
		}
		return runNewProjectMember(cmd, args[0])
	},
}

func init() {
	newCmd.AddCommand(newMemberCmd)
	addFromGroupFlag(newMemberCmd)
	addFromProjectFlag(newMemberCmd)
	addAccessLevelFlag(newMemberCmd)
	addExpiresAtFlag(newMemberCmd)
}

func runNewProjectMember(cmd *cobra.Command, user string) error {
	opts, err := assignAddProjectMemberOptions(cmd)
	if err != nil {
		return err
	}
	project := getFlagString(cmd, "from-project")
	member, err := newProjectMember(project, user, opts)
	if err != nil {
		return err
	}
	printProjectMembersOut(cmd, member)
	return err
}

func runNewGroupMember(cmd *cobra.Command, user string) error {
	opts, err := assignAddGroupMemberOptions(cmd)
	if err != nil {
		return err
	}
	group := getFlagString(cmd, "from-group")
	member, err := newGroupMember(group, user, opts)
	if err != nil {
		return err
	}
	printGroupMembersOut(cmd, member)
	return err
}

func newProjectMember(pid interface{}, user string,
	opts *gitlab.AddProjectMemberOptions) (*gitlab.ProjectMember, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}

	foundUser, err := getUserByUsername(user)
	if err != nil {
		return nil, err
	}
	opts.UserID = gitlab.Int(foundUser.ID)

	member, _, err := git.ProjectMembers.AddProjectMember(pid, opts)
	if err != nil {
		return nil, err
	}
	return member, nil
}

func newGroupMember(gid interface{}, user string,
	opts *gitlab.AddGroupMemberOptions) (*gitlab.GroupMember, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}

	foundUser, err := getUserByUsername(user)
	if err != nil {
		return nil, err
	}
	opts.UserID = gitlab.Int(foundUser.ID)

	member, _, err := git.GroupMembers.AddGroupMember(gid, opts)
	if err != nil {
		return nil, err
	}
	return member, nil
}
