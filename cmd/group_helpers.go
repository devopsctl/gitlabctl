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
	"strconv"

	"github.com/spf13/cobra"
	gitlab "github.com/xanzy/go-gitlab"
)

// getListGroupsOptions maps the cmd flags to gitlab.ListGroupsOptions struct.
// It also ensures that the struct field that is associated with the command
// flag does not use the flag default value.
func getListGroupsOptions(cmd *cobra.Command) *gitlab.ListGroupsOptions {
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

// getCreateGroupOptions maps the cmd flags to gitlab.CreateGroupOptions struct.
// It also ensures that the struct field that is associated with the command
// flag does not use the flag default value.
func getCreateGroupOptions(cmd *cobra.Command) (*gitlab.CreateGroupOptions, error) {
	var opts gitlab.CreateGroupOptions

	// change-name is only required when editing a group
	if f := cmd.Flag("change-name"); f != nil {
		if f.Changed {
			opts.Name = gitlab.String(getFlagString(cmd, "change-name"))
		}
	}

	// change-path is only required when editing a group
	if f := cmd.Flag("change-path"); f != nil {
		if f.Changed {
			opts.Path = gitlab.String(getFlagString(cmd, "change-path"))
		}
	}

	// namespace is only required when creating a new group
	if f := cmd.Flag("namespace"); f != nil {
		if f.Changed {
			ns := getFlagString(cmd, "namespace")
			id, err := strconv.Atoi(ns)
			// if not nil take the given number
			if err == nil {
				opts.ParentID = &id
			}
			// find the group as string and get it's id
			gid, err := getGroupID(getFlagString(cmd, "namespace"))
			if err != nil {
				return nil, err
			}
			opts.ParentID = gitlab.Int(gid)
		}
	}

	if cmd.Flag("desc").Changed {
		opts.Description = gitlab.String(getFlagString(cmd, "desc"))
	}

	if cmd.Flag("visibility").Changed {
		v := getFlagVisibility(cmd)
		opts.Visibility = v
	}
	if cmd.Flag("lfs-enabled").Changed {
		opts.LFSEnabled = gitlab.Bool(getFlagBool(cmd, "lfs-enabled"))
	}
	if cmd.Flag("request-access-enabled").Changed {
		opts.RequestAccessEnabled = gitlab.Bool(
			getFlagBool(cmd, "request-access-enabled"))
	}
	return &opts, nil
}

// get the groupID of a group
func getGroupID(path string) (int, error) {
	g, err := descGroup(path)
	if err != nil {
		return -1, err
	}
	return g.ID, err
}
