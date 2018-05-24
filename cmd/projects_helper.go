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
