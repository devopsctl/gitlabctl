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

var editProjectHookCmd = &cobra.Command{
	Use:        "project-hook",
	Aliases:    []string{"h"},
	SuggestFor: []string{"hook"},
	Short:      "Edit a project hook by specifying the project id or path and using flags for fields to modify",
	Example: `# update a project hook by path
gitlabctl edit project-hook ProjectX --hook-id=1 --url="http://www.sample123.com/"
gitlabctl edit project-hook GroupX/ProjectX --hook-id=2 --tag-push-events=false  

# update a project hook with id
gitlabctl edit project-hook 3 --hook-id=3 --url="http://www.sample321.com/" --issues-events`,
	Args:          cobra.ExactArgs(1),
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runEditProjectHook(cmd, args[0])
	},
}

func init() {
	editCmd.AddCommand(editProjectHookCmd)
	addEditProjectHookFlags(editProjectHookCmd)
}

func runEditProjectHook(cmd *cobra.Command, project string) error {
	opts, err := assignEditProjectHookOptions(cmd)
	if err != nil {
		return err
	}
	hook := getFlagInt(cmd, "hook-id")
	editedProjectHook, err := editProjectHook(project, hook, (*gitlab.EditProjectHookOptions)(opts))
	if err != nil {
		return err
	}
	printProjectHooksOut(cmd, editedProjectHook)
	return nil
}

func editProjectHook(project string, hook int, opts *gitlab.EditProjectHookOptions) (*gitlab.ProjectHook, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	editedProjectHook, _, err := git.Projects.EditProjectHook(project, hook, opts)
	if err != nil {
		return nil, err
	}
	return editedProjectHook, nil
}
