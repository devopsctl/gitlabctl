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
	Short:      "Update or patch a project",
	Example: `# update a project description
gitlabctl edit project ProjectX --desc="A go project"

# update a project within a group or subgroup using 'namespace'
gitlabctl edit project ProjectX --namespace=GroupX --merge-method=rebase_merge 
`,
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

func runEditProject(cmd *cobra.Command, path string) error {
	opts, err := getCreateProjectOptions(cmd)
	if err != nil {
		return err
	}
	p, err := editProject(path, (*gitlab.EditProjectOptions)(opts))
	if err != nil {
		return err
	}
	printProjectsOut(cmd, p)
	return nil
}

func editProject(path string, opts *gitlab.EditProjectOptions) (*gitlab.Project, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	p, _, err := git.Projects.EditProject(path, opts)
	if err != nil {
		return nil, err
	}
	return p, nil
}
