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

var editBranchCmd = &cobra.Command{
	Use:     "branch",
	Aliases: []string{"b"},
	Short:   "Protect or unprotect a repositort branch",
	Example: `# protect a branch
gitlabctl edit branch master -p devopsctl/gitlabctl --protect

# unprotect a branch
gitlabctl edit branch master -p devopsctl/gitlabctl --unprotect`,
	SilenceErrors:     true,
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	Args:              cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := validateFlagCombination(cmd, "protect", "dev-can-merge", "dev-can-push"); err != nil {
			return err
		}
		if cmd.Flag("protect").Changed && cmd.Flag("unprotect").Changed {
			return newUsedTooManyFlagError("protect", "unprotect")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Flag("protect").Changed {
			return runProtectBranch(cmd, args[0])
		}
		return runUnProtectBranch(args[0], getFlagString(cmd, "project"), getFlagString(cmd, "out"))
	},
}

func init() {
	editCmd.AddCommand(editBranchCmd)
	addProjectFlag(editBranchCmd)
	verifyMarkFlagRequired(editBranchCmd, "project")
	editBranchCmd.Flags().Bool("unprotect", false,
		"Remove protection of a branch")
	editBranchCmd.Flags().Bool("protect", false,
		"Protect a branch")
	editBranchCmd.Flags().Bool("dev-can-push", false,
		"Used with '--protect'. Flag if developers can push to the branch")
	editBranchCmd.Flags().Bool("dev-can-merge", false,
		"Used with '--protect'. Flag if developers can merge to the branch")
}

func runProtectBranch(cmd *cobra.Command, branch string) error {
	opts := new(gitlab.ProtectBranchOptions)
	if cmd.Flag("dev-can-push").Changed {
		opts.DevelopersCanPush = gitlab.Bool(getFlagBool(cmd, "dev-can-push"))
	}
	if cmd.Flag("dev-can-merge").Changed {
		opts.DevelopersCanMerge = gitlab.Bool(getFlagBool(cmd, "dev-can-merge"))
	}
	branchInfo, err := protectBranch(branch, getFlagString(cmd, "project"), opts)
	if err != nil {
		return err
	}
	printBranchOut(getFlagString(cmd, "out"), branchInfo)
	return nil
}

func protectBranch(branch, project string, opts *gitlab.ProtectBranchOptions) (*gitlab.Branch, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	protectedBranch, _, err := git.Branches.ProtectBranch(project, branch, opts)
	if err != nil {
		return nil, err
	}
	return protectedBranch, nil
}

func runUnProtectBranch(branch, project, out string) error {
	branchInfo, err := unProtectBranch(branch, project)
	if err != nil {
		return err
	}
	printBranchOut(out, branchInfo)
	return nil
}

func unProtectBranch(branch, project string) (*gitlab.Branch, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	branchInfo, _, err := git.Branches.UnprotectBranch(project, branch)
	if err != nil {
		return nil, err
	}
	return branchInfo, nil
}
