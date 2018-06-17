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
	"strconv"

	"github.com/spf13/cobra"
)

var deleteMemberCmd = &cobra.Command{
	Use:        "member",
	Aliases:    []string{"m"},
	SuggestFor: []string{"members"},
	Short:      "Delete a member by specifying the member name as the first argument",
	Example: `# remove a member from a group
gitlabctl delete member john.smith --from-group Group2 

# remove a member from a project
gitlabctl delete member john.smith --from-project Group1/Project1`,
	Args:              cobra.ExactArgs(1),
	SilenceErrors:     true,
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return validateFromGroupAndProjectFlags(cmd)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if getFlagString(cmd, "from-group") != "" {
			return deleteGroupMember(cmd, args[0])
		}
		return deleteProjectMember(cmd, args[0])
	},
}

func init() {
	deleteCmd.AddCommand(deleteMemberCmd)
	addFromGroupFlag(deleteMemberCmd)
	addFromProjectFlag(deleteMemberCmd)
}

func deleteProjectMember(cmd *cobra.Command, user string) error {
	git, err := newGitlabClient()
	if err != nil {
		return err
	}
	project := getFlagString(cmd, "from-project")
	uid, err := strconv.Atoi(user)
	if err != nil {
		foundUser, err := getUserByUsername(user)
		if err != nil {
			return err
		}
		uid = foundUser.ID
	}
	_, err = git.ProjectMembers.DeleteProjectMember(project, uid, nil)
	if err != nil {
		return err
	}
	fmt.Printf("member (%s) with id (%d) has been removed from project (%s)\n", user, uid, project)
	return nil
}

func deleteGroupMember(cmd *cobra.Command, user string) error {
	git, err := newGitlabClient()
	if err != nil {
		return err
	}
	group := getFlagString(cmd, "from-group")
	uid, err := strconv.Atoi(user)
	if err != nil {
		foundUser, err := getUserByUsername(user)
		if err != nil {
			return err
		}
		uid = foundUser.ID
	}
	_, err = git.GroupMembers.RemoveGroupMember(group, uid, nil)
	if err != nil {
		return err
	}
	fmt.Printf("member (%s) with id (%d) has been removed from group (%s)\n", user, uid, group)
	return nil
}
