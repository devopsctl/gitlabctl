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
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"log"

	gitlabctl "github.com/devopsctl/gitlabctl/gitlab"
	"github.com/spf13/cobra"
	gitlab "github.com/xanzy/go-gitlab"
)

var groupLsSubGroupCmd = &cobra.Command{
	Use:   "ls-subgroup",
	Short: "List all the projects of a group",
	Run: func(cmd *cobra.Command, args []string) {
		if err := runGroupLsSubGroup(cmd); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	groupCmd.AddCommand(groupLsSubGroupCmd)
	addPathFlag(groupLsSubGroupCmd)
	addJSONFlag(groupLsSubGroupCmd)
	addGroupLsFlags(groupLsSubGroupCmd)
}

func runGroupLsSubGroup(cmd *cobra.Command) error {
	// convert gitlab.ListGroupsOptions to gitlab.ListSubgroupsOptions
	opts := (*gitlab.ListSubgroupsOptions)(getGroupLsCmdOpts(cmd))
	path := getFlagString(cmd, "path")
	groups, err := gitlabctl.SubGroupLs(path, opts)
	if err != nil {
		return err
	}
	printGroupLsOut(cmd, groups)
	return nil
}
