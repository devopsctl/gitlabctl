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
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	gitlab "github.com/xanzy/go-gitlab"
)

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func printJSON(v interface{}) {
	b, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		er(fmt.Sprintf("failed printing to json: %v", err))
	}
	fmt.Println(string(b))
}

func printYAML(v interface{}) {
	b, err := yaml.Marshal(v)
	if err != nil {
		er(fmt.Sprintf("failed printing to yaml: %v", err))
	}
	fmt.Println(string(b))
}

func printTable(header []string, rows [][]string) {
	if len(header) > 5 {
		panic("maximum allowed length of a table header is only 5.")
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	for _, row := range rows {
		table.Append(row)
	}
	table.SetCaption(true,
		"Note: Use --out=json or --out=yaml to get more resource details.")
	table.Render()
}

func printGroupsOut(cmd *cobra.Command, groups ...*gitlab.Group) {
	switch getFlagString(cmd, "out") {
	case "json":
		printJSON(groups)
	case "yaml":
		printYAML(groups)
	default:
		header := []string{"ID", "PATH", "URL", "PARENT ID", "PROJECTS COUNT"}
		var rows [][]string
		for _, v := range groups {
			rows = append(rows, []string{strconv.Itoa(v.ID), v.FullPath,
				v.WebURL, strconv.Itoa(v.ParentID),
				strconv.Itoa(len(v.Projects))})
		}
		printTable(header, rows)
	}
}

func printProjectsOut(cmd *cobra.Command, projects ...*gitlab.Project) {
	switch getFlagString(cmd, "out") {
	case "json":
		printJSON(projects)
	case "yaml":
		printYAML(projects)
	default:
		header := []string{"ID", "PATH", "URL", "ISSUES COUNT", "TAGS"}
		var rows [][]string
		for _, v := range projects {
			rows = append(rows, []string{strconv.Itoa(v.ID),
				v.PathWithNamespace, v.HTTPURLToRepo,
				strconv.Itoa(v.OpenIssuesCount),
				strings.Join(v.TagList, ",")})
		}
		printTable(header, rows)
	}
}

func printGroupMembersOut(cmd *cobra.Command, members ...*gitlab.GroupMember) {
	switch getFlagString(cmd, "out") {
	case "json":
		printJSON(members)
	case "yaml":
		printYAML(members)
	default:
		header := []string{"ID", "USERNAME", "EMAIL", "ACCESS_LEVEL"}
		var rows [][]string
		for _, v := range members {
			rows = append(rows, []string{strconv.Itoa(v.ID),
				v.Username, v.Email, gitlab.Stringify(v.AccessLevel)})
		}
		printTable(header, rows)
	}
}
