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
)

var descTagCmd = &cobra.Command{
	Use:               "tag",
	Aliases:           []string{"t"},
	Short:             "Describe a tag of a specified project",
	Example:           "gitlabctl describe tag v5.0 --project=devopsctl/gitlabctl",
	SilenceErrors:     true,
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	Args:              cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return descTag(
			getFlagString(cmd, "project"),
			args[0],
			getFlagString(cmd, "out"),
		)
	},
}

func init() {
	descCmd.AddCommand(descTagCmd)
	addProjectFlag(descTagCmd)
	verifyMarkFlagRequired(descTagCmd, "project")
}

func descTag(project, tag, out string) error {
	git, err := newGitlabClient()
	if err != nil {
		return err
	}
	tagInfo, _, err := git.Tags.GetTag(project, tag)
	if err != nil {
		return err
	}
	printTagsOut(out, tagInfo)
	return nil
}
