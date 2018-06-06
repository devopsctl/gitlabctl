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

var editProjectCmd = &cobra.Command{
	Use:        "project",
	Aliases:    []string{"p"},
	SuggestFor: []string{"projects"},
	Short:      "Edit a project by specifying the project id or path and using flags for fields to modify",
	Example: `# update a project by path
gitlabctl edit project ProjectX --desc="A go project"
gitlabctl edit project GroupX/ProjectX --merge-method=rebase_merge 

# update a project with id (23)
gitlabctl edit project 3 --desc="A go project"`,
	Args:          cobra.ExactArgs(1),
	SilenceErrors: true,
	SilenceUsage:  true,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := validateVisibilityFlagValue(cmd); err != nil {
			return err
		}
		return validateMergeMethodValue(cmd)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return runEditProject(cmd, args[0])
	},
}

func init() {
	editCmd.AddCommand(editProjectCmd)
	addEditProjectFlags(editProjectCmd)
}

func runEditProject(cmd *cobra.Command, project string) error {
	opts, err := assignCreateProjectOptions(cmd)
	if err != nil {
		return err
	}
	editedProject, err := editProject(project, (*gitlab.EditProjectOptions)(opts))
	if err != nil {
		return err
	}
	printProjectsOut(cmd, editedProject)
	return nil
}

func editProject(project string, opts *gitlab.EditProjectOptions) (*gitlab.Project, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	editedProject, _, err := git.Projects.EditProject(project, opts)
	if err != nil {
		return nil, err
	}
	return editedProject, nil
}
