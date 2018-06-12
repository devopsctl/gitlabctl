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

var getProjectHooksCmd = &cobra.Command{
	Use:           "project-hooks",
	Aliases:       []string{"h"},
	SuggestFor:    []string{"project-hook"},
	Short:         "List all project hooks of a specified project",
	Args:          cobra.ExactArgs(1),
	SilenceErrors: true,
	SilenceUsage:  true,
	Example: `# get project hooks of projectX
gitlabctl get project-hooks projectX

# get project hooks of project with id (23)
gitlabctl get project-hooks 23`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runGetProjectHooks(cmd, args[0])
	},
}

func init() {
	getCmd.AddCommand(getProjectHooksCmd)
}

func runGetProjectHooks(cmd *cobra.Command, project string) error {
	opts := new(gitlab.ListProjectHooksOptions)
	if cmd.Flag("page").Changed {
		opts.Page = getFlagInt(cmd, "page")
	}
	if cmd.Flag("per-page").Changed {
		opts.PerPage = getFlagInt(cmd, "per-page")
	}
	hooks, err := getProjectHooks(project, opts)
	if err != nil {
		return err
	}
	printProjectHooksOut(cmd, hooks...)
	return nil
}

func getProjectHooks(project string, opts *gitlab.ListProjectHooksOptions) ([]*gitlab.ProjectHook, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	hooks, _, err := git.Projects.ListProjectHooks(project, opts)
	if err != nil {
		return nil, err
	}
	return hooks, nil
}
