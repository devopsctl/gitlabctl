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

var newProjectCmd = &cobra.Command{
	Use:        "project",
	Aliases:    []string{"p"},
	SuggestFor: []string{"projects"},
	Short:      "Create a new project",
	Example: `# create a new project
gitlabctl new project ProjectX	

# create a new project under a group
gitlabctl new project ProjectY --namespace=GroupY
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
		return runNewProject(cmd, args[0])
	},
}

func init() {
	newCmd.AddCommand(newProjectCmd)
	addNewProjectFlags(newProjectCmd)
	addOutFlag(newProjectCmd)
}

func runNewProject(cmd *cobra.Command, name string) error {
	opts, err := getCreateProjectOptions(cmd)
	if err != nil {
		return err
	}
	opts.Name = gitlab.String(name)
	opts.Path = gitlab.String(name)
	p, err := newProject(opts)
	if err != nil {
		return err
	}
	printProjectsOut(cmd, p)
	return nil
}

func newProject(opts *gitlab.CreateProjectOptions) (*gitlab.Project, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	p, _, err := git.Projects.CreateProject(opts)
	if err != nil {
		return nil, err
	}
	return p, nil
}
