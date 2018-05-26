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

var getProjectsCmd = &cobra.Command{
	Use:     "projects",
	Aliases: []string{"project", "p"},
	Short:   "List all projects",
	// SilenceUsage:  true,
	// SilenceErrors: true,
	Example: `
# get all projects with full details in JSON format
gitlabctl get projects --json

# get all projects from a group
gitlabctl get projects --from-group=Group1

# get all projects with simple details in JSON format
gitlabctl get projects --simple --json

# get all projects with issues enabled only in Table format
gitlabctl get projects --with-issues-enabled

# get private projects in Table format
gitlabctl get projects --visibility private`,
	// We use PreRunE instead of PreRun to test validation of flags in testing
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := validateSortFlagValue(cmd); err != nil {
			return err
		}
		if err := validateProjectOrderByFlagValue(cmd); err != nil {
			return err
		}
		if err := validateVisibilityFlagValue(cmd); err != nil {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if getFlagString(cmd, "from-group") != "" {
			if err := runGetProjectsFromGroup(cmd); err != nil {
				er(err)
			}
		}
		if err := runGetProjects(cmd); err != nil {
			er(err)
		}
	},
}

func init() {
	getCmd.AddCommand(getProjectsCmd)
	addGetProjectsFlags(getProjectsCmd)
}

func runGetProjects(cmd *cobra.Command) error {
	opts := getListProjectsOptions(cmd)
	projects, err := getProjects(opts)
	if err != nil {
		return err
	}
	printProjectsOut(cmd, projects...)
	return nil
}

func getProjects(opts *gitlab.ListProjectsOptions) ([]*gitlab.Project, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	g, _, err := git.Projects.ListProjects(opts)
	if err != nil {
		return nil, err
	}
	return g, nil
}

// TODO:
// list group project opts is currently not supported by the go-gitlab client
// https://github.com/xanzy/go-gitlab/issues/407
func runGetProjectsFromGroup(cmd *cobra.Command) error {
	group := getFlagString(cmd, "from-group")
	// opts := getListProjectsOptions(cmd)
	projects, err := getProjectsFromGroup(group, nil)
	if err != nil {
		return err
	}
	printProjectsOut(cmd, projects...)
	return nil
}

func getProjectsFromGroup(group string,
	opts *gitlab.ListGroupProjectsOptions) ([]*gitlab.Project, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	p, _, err := git.Groups.ListGroupProjects(group, opts)
	if err != nil {
		return nil, err
	}
	return p, nil
}
