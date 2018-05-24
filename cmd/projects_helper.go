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
	gitlab "github.com/xanzy/go-gitlab"
)

// getListProjectsOptions maps the cmd flags to gitlab.ListProjectsOptions struct.
// It also ensures that the struct field that is associated with the command
// flag does not use the flag default value.
func getListProjectsOptions(cmd *cobra.Command) *gitlab.ListProjectsOptions {
	var opts gitlab.ListProjectsOptions
	if cmd.Flag("archived").Changed {
		opts.Archived = gitlab.Bool(getFlagBool(cmd, "archived"))
	}
	if cmd.Flag("order-by").Changed {
		opts.OrderBy = gitlab.String(getFlagString(cmd, "order-by"))
	}
	if cmd.Flag("sort").Changed {
		opts.Sort = gitlab.String(getFlagString(cmd, "sort"))
	}
	if cmd.Flag("search").Changed {
		opts.Search = gitlab.String(getFlagString(cmd, "search"))
	}
	if cmd.Flag("simple").Changed {
		opts.Simple = gitlab.Bool(getFlagBool(cmd, "simple"))
	}
	if cmd.Flag("owned").Changed {
		opts.Owned = gitlab.Bool(getFlagBool(cmd, "owned"))
	}
	if cmd.Flag("membership").Changed {
		opts.Membership = gitlab.Bool(getFlagBool(cmd, "membership"))
	}
	if cmd.Flag("starred").Changed {
		opts.Starred = gitlab.Bool(getFlagBool(cmd, "starred"))
	}
	if cmd.Flag("statistics").Changed {
		opts.Statistics = gitlab.Bool(getFlagBool(cmd, "statistics"))
	}
	if cmd.Flag("visibility").Changed {
		v, err := getFlagVisibility(cmd)
		if err != nil {
			er(err)
		}
		opts.Visibility = v
	}
	if cmd.Flag("with-issues-enabled").Changed {
		opts.WithIssuesEnabled = gitlab.Bool(getFlagBool(cmd, "with-issues-enabled"))
	}
	if cmd.Flag("with-merge-requests-enabled").Changed {
		opts.WithMergeRequestsEnabled = gitlab.Bool(getFlagBool(cmd, "with-merge-requests-enabled"))
	}
	return &opts
}

// printProjectsOut prints the project list/get commands to a table view or json
func printProjectsOut(cmd *cobra.Command, projects ...*gitlab.Project) {
	if getFlagBool(cmd, "json") {
		printJSON(projects)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	header := []string{
		"ID", "NAME", "PATH", "URL", "VISIBILITY",
		"REQUEST ACCESS ENABLED", "LFS ENABLED",
	}
	table.SetHeader(header)

	for _, v := range projects {
		row := []string{
			strconv.Itoa(v.ID), v.Name, v.PathWithNamespace, v.WebURL,
			strings.Replace(gitlab.Stringify(v.Visibility), `"`, "", -1),
			strconv.FormatBool(v.RequestAccessEnabled),
			strconv.FormatBool(v.LFSEnabled),
		}
		table.Append(row)
	}
	table.Render()
}
