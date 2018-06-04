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
	"time"

	"github.com/araddon/dateparse"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	gitlab "github.com/xanzy/go-gitlab"
	yaml "gopkg.in/yaml.v2"
)

const (
	// JSON is used as a constant of word "json" for out flag
	JSON = "json"
	// YAML is used as a constant of word "yaml" for out flag
	YAML = "yaml"
)

func getNamespaceID(id string) (int, error) {
	git, err := newGitlabClient()
	if err != nil {
		return -1, err
	}
	ns, _, err := git.Namespaces.GetNamespace(id)
	if err != nil {
		return -1, err
	}
	return ns.ID, nil
}

func getGroupID(path string) (int, error) {
	g, err := descGroup(path)
	if err != nil {
		return -1, err
	}
	return g.ID, err
}

func newTimeFromString(s string) (*time.Time, error) {
	t, err := dateparse.ParseAny(s)
	if err != nil {
		return nil, err
	}
	p := new(time.Time)
	*p = t
	return p, nil
}

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
	case JSON:
		printJSON(groups)
	case YAML:
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
	case JSON:
		printJSON(projects)
	case YAML:
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
	case JSON:
		printJSON(members)
	case YAML:
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

func printUsersOut(cmd *cobra.Command, users ...*gitlab.User) {
	switch getFlagString(cmd, "out") {
	case JSON:
		printJSON(users)
	case YAML:
		printYAML(users)
	default:
		header := []string{"ID", "USERNAME", "EMAIL", "NAME", "LAST SIGN IN AT"}
		var rows [][]string
		for _, v := range users {
			rows = append(rows, []string{strconv.Itoa(v.ID),
				v.Username, v.Email, v.Name, fmt.Sprintf("%v", v.LastSignInAt)})
		}
		printTable(header, rows)
	}
}
