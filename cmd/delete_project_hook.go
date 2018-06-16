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

var deleteProjectHookCmd = &cobra.Command{
	Use:        "project-hook",
	Aliases:    []string{"h"},
	SuggestFor: []string{"hook"},
	Short:      "Delete a Gitlab project hook by specifying the project's full path or id",
	Example: `# delete a project hook by project's path
gitlabctl delete project-hook 1 --project=GroupX/ProjectX

# delete a project hook by project id
gitlabctl delete project-hook 2 --project=22`,
	Args:          cobra.ExactArgs(1),
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return deleteProjectHook(args[0], getFlagString(cmd, "project"))
	},
}

func init() {
	deleteCmd.AddCommand(deleteProjectHookCmd)
	addProjectFlag(deleteProjectHookCmd)
	verifyMarkFlagRequired(deleteProjectHookCmd, "project")
}

func deleteProjectHook(hook string, project string) error {
	git, err := newGitlabClient()
	if err != nil {
		return err
	}
	hid, err := strconv.Atoi(hook)
	if err != nil {
		return err
	}
	projectInfo, _, err := git.Projects.GetProject(project)
	if err != nil {
		return err
	}
	_, err = git.Projects.DeleteProjectHook(projectInfo.ID, hid)
	if err != nil {
		return err
	}
	fmt.Printf("project hook with id (%d) from (%s) has been deleted\n", hid, projectInfo.PathWithNamespace)
	return nil
}
