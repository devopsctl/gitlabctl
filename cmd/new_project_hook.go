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

var newProjectHookCmd = &cobra.Command{
	Use:        "project-hook",
	Aliases:    []string{"h"},
	SuggestFor: []string{"hook"},
	Short:      "Create a new project hook by specifying the project id or project path as the first argument",
	Example: `# create a new project hook by project path
gitlabctl new project-hook GroupX/ProjectX --url="http://www.sample.com/"

# create a new project hook by project id
gitlabctl new project-hook 123 --url="http://www.sample.com/"

# create a new project hook with merge request events trigger enabled
gitlabctl new project-hook 123 --url="http://www.sample.com/" --merge-requests-events`,
	Args:              cobra.ExactArgs(1),
	SilenceErrors:     true,
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runNewProjectHook(cmd, args[0])
	},
}

func init() {
	newCmd.AddCommand(newProjectHookCmd)
	addNewProjectHookFlags(newProjectHookCmd)
}

func runNewProjectHook(cmd *cobra.Command, name string) error {
	opts, err := assignAddProjectHookOptions(cmd)
	if err != nil {
		return err
	}
	hook, errr := newProjectHook(name, opts)
	if errr != nil {
		return errr
	}
	printProjectHooksOut(cmd, hook)
	return nil
}

func newProjectHook(project string, opts *gitlab.AddProjectHookOptions) (*gitlab.ProjectHook, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	p, _, err := git.Projects.AddProjectHook(project, opts)
	if err != nil {
		return nil, err
	}
	return p, nil
}
