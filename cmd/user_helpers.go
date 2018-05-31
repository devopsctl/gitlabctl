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
	gitlab "github.com/xanzy/go-gitlab"
)

func createListUsersOptions(cmd *cobra.Command) (*gitlab.ListUsersOptions, error) {
	var opts gitlab.ListUsersOptions
	if cmd.Flag("active").Changed {
		opts.Active = gitlab.Bool(getFlagBool(cmd, "active"))
	}
	if cmd.Flag("blocked").Changed {
		opts.Blocked = gitlab.Bool(getFlagBool(cmd, "blocked"))
	}
	if cmd.Flag("search").Changed {
		opts.Search = gitlab.String(getFlagString(cmd, "search"))
	}
	if cmd.Flag("username").Changed {
		opts.Username = gitlab.String(getFlagString(cmd, "username"))
	}
	if cmd.Flag("external-uid").Changed {
		opts.ExternalUID = gitlab.String(getFlagString(cmd, "external-uid"))
	}
	if cmd.Flag("provider").Changed {
		opts.Provider = gitlab.String(getFlagString(cmd, "provider"))
	}
	if cmd.Flag("created-before").Changed {
		ts := getFlagString(cmd, "created-before")
		t, err := newTimeFromString(ts)
		if err != nil {
			return nil, fmt.Errorf("failed parsing time %s, "+
				"got error: %v", ts, err)
		}
		opts.CreatedBefore = t
	}
	if cmd.Flag("created-after").Changed {
		ts := getFlagString(cmd, "created-after")
		t, err := newTimeFromString(ts)
		if err != nil {
			return nil, fmt.Errorf("failed parsing time %s, "+
				"got error: %v", ts, err)
		}
		opts.CreatedAfter = t
	}
	if cmd.Flag("order-by").Changed {
		opts.OrderBy = gitlab.String(getFlagString(cmd, "order-by"))
	}
	if cmd.Flag("sort").Changed {
		opts.Sort = gitlab.String(getFlagString(cmd, "sort"))
	}
	return &opts, nil
}
