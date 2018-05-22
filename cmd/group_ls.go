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
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
)

var groupListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all groups",
	Run: func(cmd *cobra.Command, args []string) {
		if err := runGroupLs(cmd); err != nil {
			er(err)
		}
	},
}

func init() {
	groupCmd.AddCommand(groupListCmd)
	addJSONFlag(groupListCmd)
	addGroupLsFlags(groupListCmd)
}

func addGroupLsFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("all-available", false,
		"Show all the groups you have access to"+
			"(defaults to false for authenticated users, true for admin)")
	cmd.Flags().Bool("owned", false,
		"Limit to groups owned by the current user")
	cmd.Flags().Bool("statistics", false,
		"Include group statistics (admins only)")
	cmd.Flags().String("sort", "asc",
		"Order groups in asc or desc order. Default is asc")
	cmd.Flags().String("search", "",
		"Return the list of authorized groups matching the search criteria")
	cmd.Flags().String("order-by", "name",
		"Order groups by name or path. Default is name")
}

// getGroupLsCmdOpts maps the cmd flags to gitlab.ListGroupsOptions struct.
// It also ensures that the struct field that is associated with the command
// flag does not use the flag default value.
func getGroupLsCmdOpts(cmd *cobra.Command) *gitlab.ListGroupsOptions {
	var opts gitlab.ListGroupsOptions
	if cmd.Flag("all-available").Changed {
		opts.AllAvailable = gitlab.Bool(getFlagBool(cmd, "all-available"))
	}
	if cmd.Flag("owned").Changed {
		opts.Owned = gitlab.Bool(getFlagBool(cmd, "owned"))
	}
	if cmd.Flag("statistics").Changed {
		opts.Statistics = gitlab.Bool(getFlagBool(cmd, "statistics"))
	}
	if cmd.Flag("sort").Changed {
		opts.Sort = gitlab.String(getFlagString(cmd, "sort"))
	}
	if cmd.Flag("search").Changed {
		opts.Search = gitlab.String(getFlagString(cmd, "search"))
	}
	if cmd.Flag("order-by").Changed {
		opts.OrderBy = gitlab.String(getFlagString(cmd, "order-by"))
	}
	return &opts
}

func runGroupLs(cmd *cobra.Command) error {
	opts := getGroupLsCmdOpts(cmd)
	groups, err := listGroups(opts)
	if err != nil {
		return err
	}
	printGroupLsOut(cmd, groups...)
	return err
}

func listGroups(opts *gitlab.ListGroupsOptions) ([]*gitlab.Group, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	g, _, err := git.Groups.ListGroups(opts)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func printGroupLsOut(cmd *cobra.Command, groups ...*gitlab.Group) {
	if getFlagBool(cmd, "json") {
		printJSON(groups)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	header := []string{
		"ID", "NAME", "PATH", "PARENT_ID", "VISIBILITY",
		"REQUEST ACCESS ENABLED", "LFS ENABLED",
	}
	statsFlag := cmd.Flag("statistics")
	if statsFlag != nil {
		if cmd.Flag("statistics").Changed {
			header = append(header,
				"STORAGE SIZE", "REPO SIZE", "LFS SIZE", "JOB ARTIFACT SIZE",
			)
		}
	}
	table.SetHeader(header)

	for _, v := range groups {
		row := []string{
			strconv.Itoa(v.ID), v.Name, v.FullPath,
			strconv.Itoa(v.ParentID),
			strings.Replace(gitlab.Stringify(v.Visibility), `"`, "", -1),
			strconv.FormatBool(v.RequestAccessEnabled),
			strconv.FormatBool(v.LFSEnabled),
		}
		if statsFlag != nil {
			if getFlagBool(cmd, "statistics") {
				row = append(row, strconv.FormatInt(v.Statistics.StorageSize, 10),
					strconv.FormatInt(v.Statistics.RepositorySize, 10),
					strconv.FormatInt(v.Statistics.LfsObjectsSize, 10),
					strconv.FormatInt(v.Statistics.JobArtifactsSize, 10))
			}
		}
		table.Append(row)
	}
	table.Render()
}
