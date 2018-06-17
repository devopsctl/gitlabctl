// Copyright © 2018 github.com/devopsctl authors
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal
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

var newBranchCmd = &cobra.Command{
	Use:     "branch",
	Aliases: []string{"b"},
	Short:   "Create a new branch for a specified project",
	Example: `# create a develop branch from master branch for project groupx/myapp
gitlabctl new branch develop --project=groupx/myapp --ref=master`,
	SilenceErrors: true,
	SilenceUsage:  true,
	Args:          cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runNewBranch(cmd, args[0])
	},
}

func init() {
	newCmd.AddCommand(newBranchCmd)
	addProjectFlag(newBranchCmd)
	verifyMarkFlagRequired(newBranchCmd, "project")
	newBranchCmd.Flags().StringP("ref", "r", "",
		"The branch name or commit SHA to create branch from")
	verifyMarkFlagRequired(newBranchCmd, "ref")
}

func runNewBranch(cmd *cobra.Command, branch string) error {
	opts := new(gitlab.CreateBranchOptions)
	opts.Ref = gitlab.String(getFlagString(cmd, "ref"))
	opts.Branch = gitlab.String(branch)
	createdBranch, err := newBranch(getFlagString(cmd, "project"), opts)
	if err != nil {
		return err
	}
	printBranchOut(cmd, createdBranch)
	return nil
}

func newBranch(project string, opts *gitlab.CreateBranchOptions) (*gitlab.Branch, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	branch, _, err := git.Branches.CreateBranch(project, opts)
	if err != nil {
		return nil, err
	}
	return branch, nil
}
