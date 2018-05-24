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

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	gitlab "github.com/xanzy/go-gitlab"
)

// addGetGroupsFlags adds common flags for `get groups` and `get subgroups` commands
func addGetGroupsFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("all-available", false,
		"Show all the groups you have access to "+
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

// addNewGroupFlags adds common flags for `new group` or `edit group` commands
func addNewGroupFlags(cmd *cobra.Command) {
	addNameFlag(cmd)
	addNamespaceFlag(cmd)
	addLFSenabled(cmd)
	addRequestsAccessenabledFlag(cmd)
	addVisibilityFlag(cmd)
	addJSONFlag(cmd)
	if err := cmd.MarkFlagRequired("name"); err != nil {
		er(err)
	}
}

//
// NOTE(@bzon): All addFlags* helpers should be added below
//

func addNameFlag(cmd *cobra.Command) {
	cmd.Flags().String("name", "", "resource name")
}

func addNamespaceFlag(cmd *cobra.Command) {
	cmd.Flags().String("namespace", "",
		"The parent group id or group path if creating a subgroup. "+
			"(defaults to current user namespace)")
}

func addVisibilityFlag(cmd *cobra.Command) {
	cmd.Flags().String("visibility", "", "public, internal or private")
}

func addRequestsAccessenabledFlag(cmd *cobra.Command) {
	cmd.Flags().Bool("request-access-enabled", false, "enable request access")
}

func addLFSenabled(cmd *cobra.Command) {
	cmd.Flags().Bool("lfs-enabled", false, "enable LFS")
}

func addPathFlag(cmd *cobra.Command) {
	cmd.Flags().String("path", "",
		"the group name, id or full the path "+
			"including the parent group (path/to/group)")
	if err := cmd.MarkFlagRequired("path"); err != nil {
		panic(err)
	}
}

//
// TODO(@migueltanada) Update Desciprtion of username for all commands
//
func addUsernameFlag(cmd *cobra.Command) {
	cmd.Flags().String("username", "",
		"username of user/member")
	if err := cmd.MarkFlagRequired("username"); err != nil {
		panic(err)
	}
}

func addJSONFlag(cmd *cobra.Command) {
	cmd.Flags().Bool("json", false, "print the command output to json")
}

//
// NOTE(@bzon): All getFlag* helpers should be added below
//

func getFlagString(cmd *cobra.Command, flag string) string {
	s, err := cmd.Flags().GetString(flag)
	if err != nil {
		glog.Fatalf("error accessing flag %s for command %s: %v",
			flag, cmd.Name(), err)
	}
	return s
}

func getFlagBool(cmd *cobra.Command, flag string) bool {
	b, err := cmd.Flags().GetBool(flag)
	if err != nil {
		glog.Fatalf("error accessing flag %s for command %s: %v",
			flag, cmd.Name(), err)
	}
	return b
}

func getFlagInt(cmd *cobra.Command, flag string) int {
	i, err := cmd.Flags().GetInt(flag)
	if err != nil {
		glog.Fatalf("error accessing flag %s for command %s: %v",
			flag, cmd.Name(), err)
	}
	return i
}

// getFlagVisibility converts the string flag visiblity to gitlab.VisibilityValue.
// It will exit with error if the vibility flag given by the user is not valid.
func getFlagVisibility(cmd *cobra.Command) (*gitlab.VisibilityValue, error) {
	// TODO(@bzon): can this be improved?
	v := getFlagString(cmd, "visibility")
	gv := gitlab.VisibilityValue(v)
	switch gv {
	case gitlab.PrivateVisibility:
		return gitlab.Visibility(gv), nil
	case gitlab.PublicVisibility:
		return gitlab.Visibility(gv), nil
	case gitlab.InternalVisibility:
		return gitlab.Visibility(gv), nil
	default:
		// user error
		return nil, fmt.Errorf("visibility valid values are only (public, private, internal)")
	}
}
