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

var newTagCmd = &cobra.Command{
	Use:     "tag",
	Aliases: []string{"t"},
	Short:   "Create a new tag for a specified project",
	Example: `# create tag from master branch for project groupx/myapp
gitlabctl new tag v2.0 --project=groupx/myapp --ref=master

# create a tag and create a release from it
gitlabctl new tag v2.1 --project=groupx/myapp --ref=master --description="Released!"

# NOTE: You can also use 'gitlabctl new release' to create a release separately.`,
	SilenceErrors:     true,
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	Args:              cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runNewTag(cmd, args[0])
	},
}

func init() {
	newCmd.AddCommand(newTagCmd)
	addProjectFlag(newTagCmd)
	verifyMarkFlagRequired(newTagCmd, "project")
	newTagCmd.Flags().StringP("ref", "r", "",
		"The branch name or commit SHA to create branch from")
	verifyMarkFlagRequired(newTagCmd, "ref")
	newTagCmd.Flags().StringP("message", "m", "",
		"Creates annotated tag")
	newTagCmd.Flags().StringP("description", "d", "",
		"Create a release from the git tag with the description as the release note")
}

func runNewTag(cmd *cobra.Command, tag string) error {
	opts := new(gitlab.CreateTagOptions)
	opts.Ref = gitlab.String(getFlagString(cmd, "ref"))
	opts.TagName = gitlab.String(tag)
	opts.Message = gitlab.String(getFlagString(cmd, "message"))
	opts.ReleaseDescription = gitlab.String(getFlagString(cmd, "description"))
	createdTag, err := newTag(getFlagString(cmd, "project"), opts)
	if err != nil {
		return err
	}
	printTagsOut(getFlagString(cmd, "out"), createdTag)
	return nil
}

func newTag(project string, opts *gitlab.CreateTagOptions) (*gitlab.Tag, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	tag, _, err := git.Tags.CreateTag(project, opts)
	if err != nil {
		return nil, err
	}
	return tag, nil
}
