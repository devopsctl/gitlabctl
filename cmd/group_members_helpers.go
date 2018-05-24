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

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	gitlab "github.com/xanzy/go-gitlab"
)

// printGroupMembersOut prints the group members list/get commands to a table view or json
func printGroupMembersOut(cmd *cobra.Command, groupMembers []*gitlab.GroupMember) {
	if getFlagBool(cmd, "json") {
		printJSON(groupMembers)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	header := []string{
		"ID", "USERNAME", "EMAIL", "NAME", "STATE", "CREATED_AT",
		"ACCESS_LEVEL", "EXPIRES_AT",
	}

	table.SetHeader(header)

	for _, v := range groupMembers {
		row := []string{
			strconv.Itoa(v.ID), v.Username, v.Email, v.Name, v.State,
			gitlab.Stringify(v.CreatedAt), gitlab.Stringify(v.AccessLevel), gitlab.Stringify(v.ExpiresAt),
		}

		table.Append(row)
	}
	table.Render()
}
