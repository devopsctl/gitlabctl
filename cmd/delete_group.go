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
	"fmt"

	"github.com/spf13/cobra"
)

var deleteGroupCmd = &cobra.Command{
	Use:        "group",
	Aliases:    []string{"g"},
	SuggestFor: []string{"groups"},
	Short:      "Delete a Gitlab group",
	Example: `
# delete a Group named GroupX
gitlabctl delete group GroupX

# delete a Subgroup named GroupY under GroupX
gitlabctl delete group GroupX/GroupY
`,
	Args:          cobra.ExactArgs(1),
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return deleteGroup(args[0])
	},
}

func init() {
	deleteCmd.AddCommand(deleteGroupCmd)
}

func deleteGroup(path string) error {
	git, err := newGitlabClient()
	if err != nil {
		return err
	}
	gid, err := getGroupID(path)
	if err != nil {
		return err
	}
	_, err = git.Groups.DeleteGroup(gid)
	if err != nil {
		return err
	}
	fmt.Printf("Group (%s) with id (%d) has been deleted\n", path, gid)
	return nil
}
