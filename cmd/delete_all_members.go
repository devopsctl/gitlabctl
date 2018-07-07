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
)

var deleteAllMembersCmd = &cobra.Command{
	Use:               "all-members",
	SuggestFor:        []string{"all", "all-member"},
	Short:             "Delete all members of a project",
	Example:           `gitlabctl delete all-members --from-project Group1/Project1`,
	Args:              cobra.ExactArgs(0),
	SilenceErrors:     true,
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return validateFromProjectFlag(cmd)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return deleteAllProjectMembers(getFlagString(cmd, "from-project"))
	},
}

func init() {
	deleteCmd.AddCommand(deleteAllMembersCmd)
	addFromProjectFlag(deleteAllMembersCmd)
}

func deleteAllProjectMembers(project string) error {
	git, err := newGitlabClient()
	if err != nil {
		return err
	}

	members, _, err := git.ProjectMembers.ListProjectMembers(project, nil)
	if err != nil {
		return err
	}

	p, _ := descProject(project)

	for _, member := range members {
		if member.ID != p.CreatorID {
			if err := deleteProjectMember(project, strconv.Itoa(member.ID)); err != nil {
				return err
			}
		}
	}
	return nil
}
