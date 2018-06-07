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
	"github.com/xanzy/go-gitlab"
)

var getMembersCmd = &cobra.Command{
	Use:           "members",
	Aliases:       []string{"m"},
	SuggestFor:    []string{"member"},
	Args:          cobra.ExactArgs(0),
	SilenceErrors: true,
	SilenceUsage:  true,
	Short:         "List all members of a group/project",
	Example: `# list all members of a groups
gitlabctl get members --from-group Group1[flags]

# list all members of a project
gitlabctl get members --from-project Group1/Project1[flags]
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if getFlagString(cmd, "from-group") != "" &&
			getFlagString(cmd, "from-project") != "" {
			return fmt.Errorf("Set only 1 of the following flags:" +
				"--from-group --from-project")
		}
		if getFlagString(cmd, "from-group") == "" &&
			getFlagString(cmd, "from-project") == "" {
			return fmt.Errorf("Set at least 1 of the following flags:" +
				"--from-group --from-project")
		}
		if getFlagString(cmd, "from-group") != "" {
			return runGetGroupMembers(cmd)
		}
		if getFlagString(cmd, "from-project") != "" {
			return runGetProjectMembers(cmd)
		}
		return fmt.Errorf("error")

	},
}

func init() {
	getCmd.AddCommand(getMembersCmd)
	addFromGroupFlag(getMembersCmd)
	addFromProjectFlag(getMembersCmd)
	addQueryFlag(getMembersCmd)

}

func runGetGroupMembers(cmd *cobra.Command) error {
	group := getFlagString(cmd, "from-group")
	opts := assignListGroupMembersOptions(cmd)
	members, err := getGroupsMembers(group, opts)
	if err != nil {
		return err
	}
	printGroupMembersOut(cmd, members...)
	return err
}
func runGetProjectMembers(cmd *cobra.Command) error {
	project := getFlagString(cmd, "from-project")
	opts := assignListProjectMembersOptions(cmd)
	members, err := getProjectMembers(project, opts)
	if err != nil {
		return err
	}
	printProjectMembersOut(cmd, members...)
	return err
}

func getGroupsMembers(gid interface{},
	opts *gitlab.ListGroupMembersOptions) ([]*gitlab.GroupMember, error) {
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

func getProjectMembers(pid interface{},
	opts *gitlab.ListProjectMembersOptions) ([]*gitlab.ProjectMember, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	g, _, err := git.ProjectMembers.ListProjectMembers(pid, opts)
	if err != nil {
		return nil, err
	}
	return g, nil
}
