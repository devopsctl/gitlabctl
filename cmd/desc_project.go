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

var descProjectCmd = &cobra.Command{
	Use:        "project",
	Aliases:    []string{"p"},
	SuggestFor: []string{"projects"},
	Short:      "Describe a project by specifying the id or project path",
	Example: `# describe a project by path
gitlabctl describe project ProjectX
gitlabctl describe project GroupY/ProjectY

# describe a project with id (23)
gitlabctl describe project 23`,
	Args:              cobra.ExactArgs(1),
	SilenceErrors:     true,
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runDescProject(cmd, args[0])
	},
}

func init() {
	descCmd.AddCommand(descProjectCmd)
}

func runDescProject(cmd *cobra.Command, project string) error {
	projectInfo, err := descProject(project)
	if err != nil {
		return err
	}
	printProjectsOut(getFlagString(cmd, "out"), projectInfo)
	return nil
}

func descProject(project string) (*gitlab.Project, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	projectInfo, _, err := git.Projects.GetProject(project)
	if err != nil {
		return nil, err
	}
	return projectInfo, nil
}
