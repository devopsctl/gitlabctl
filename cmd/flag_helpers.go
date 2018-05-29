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
	"strings"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	gitlab "github.com/xanzy/go-gitlab"
)

// addGetGroupsFlags adds common flags for `get groups` and `get subgroups` commands
func addGetGroupsFlags(cmd *cobra.Command) {
	addAllAvailableFlag(cmd)
	addGroupOrderByFlag(cmd)
	addOwnedFlag(cmd)
	addSortFlag(cmd)
	addStatisticsFlag(cmd)
	addSearchFlag(cmd)
}

// addGetProjectsFlags adds common flags for `get projects` commands
func addGetProjectsFlags(cmd *cobra.Command) {
	addFromGroupFlag(cmd)
	addProjectOrderByFlag(cmd)
	addSortFlag(cmd)
	addSearchFlag(cmd)
	addStatisticsFlag(cmd)
	addVisibilityFlag(cmd)
	addOwnedFlag(cmd)
	cmd.Flags().Bool("archived", false,
		"Limit by archived status")
	cmd.Flags().Bool("simple", false,
		"Return only the ID, URL, name, and path of each project")
	cmd.Flags().Bool("membership", false,
		"Limit by projects that the current user is a member of")
	cmd.Flags().Bool("starred", false,
		"Limit by projects starred by the current user")
	cmd.Flags().Bool("with-issues-enabled", false,
		"Limit by enabled issues feature")
	cmd.Flags().Bool("with-merge-requests-enabled", false,
		"Limit by enabled merge requests feature")
}

// addNewGroupFlags adds common flags for `new group` or `edit group` commands
func addNewGroupFlags(cmd *cobra.Command) {
	addNamespaceFlag(cmd)
	addLFSenabled(cmd)
	addRequestAccessEnabledFlag(cmd)
	addVisibilityFlag(cmd)
}

func addFromGroupFlag(cmd *cobra.Command) {
	cmd.Flags().String("from-group", "",
		"Use a group as the target namespace when performing the command")
}

func addAllAvailableFlag(cmd *cobra.Command) {
	cmd.Flags().Bool("all-available", false,
		"Show all the groups you have access to "+
			"(defaults to false for authenticated users, true for admin)")
}

func addOwnedFlag(cmd *cobra.Command) {
	cmd.Flags().Bool("owned", false,
		"Limit to resources owned by the current user")
}

func addGroupOrderByFlag(cmd *cobra.Command) {
	cmd.Flags().String("order-by", "name",
		"Order groups by name or path. Default is name")
}

func validateGroupOrderByFlagValue(cmd *cobra.Command) error {
	if err := validateFlagStringValue([]string{"path", "name"},
		cmd, "order-by"); err != nil {
		return err
	}
	return nil
}

func addSearchFlag(cmd *cobra.Command) {
	cmd.Flags().String("search", "",
		"Return the list of resources matching the search criteria")
}

func addStatisticsFlag(cmd *cobra.Command) {
	cmd.Flags().Bool("statistics", false,
		"Include resource statistics (admins only)")
}

func addSortFlag(cmd *cobra.Command) {
	cmd.Flags().String("sort", "asc",
		"Order resources in asc or desc order. Default is asc")
}

func validateSortFlagValue(cmd *cobra.Command) error {
	if err := validateFlagStringValue([]string{"asc", "desc"},
		cmd, "sort"); err != nil {
		return err
	}
	return nil
}

func addProjectOrderByFlag(cmd *cobra.Command) {
	cmd.Flags().String("order-by", "created_at",
		"Return projects ordered by id, name, path, created_at, updated_at, "+
			"or last_activity_at fields. Default is created_at")
}

func validateProjectOrderByFlagValue(cmd *cobra.Command) error {
	if err := validateFlagStringValue([]string{"id", "name", "path",
		"created_at", "updated_at", "last_activity_at"},
		cmd, "order-by"); err != nil {
		return err
	}
	return nil
}

//
// NOTE(@bzon): All addFlags* helpers should be added below
//

func addNamespaceFlag(cmd *cobra.Command) {
	cmd.Flags().String("namespace", "",
		"The parent group id or group path if creating a subgroup. "+
			"(defaults to current user namespace)")
}

func addVisibilityFlag(cmd *cobra.Command) {
	cmd.Flags().String("visibility", "private", "public, internal or private")
}

func validateVisibilityFlagValue(cmd *cobra.Command) error {
	if err := validateFlagStringValue([]string{"public", "private", "internal"},
		cmd, "visibility"); err != nil {
		return err
	}
	return nil
}

func addRequestAccessEnabledFlag(cmd *cobra.Command) {
	cmd.Flags().Bool("request-access-enabled", false, "enable request access")
}

func addLFSenabled(cmd *cobra.Command) {
	cmd.Flags().Bool("lfs-enabled", false, "enable LFS")
}

func addPathFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("path", "p", "",
		"the group name, id or full the path "+
			"including the parent group (path/to/group)")
	if err := cmd.MarkFlagRequired("path"); err != nil {
		panic(err)
	}
}

func addOutFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("out", "o", "simple",
		"Print the command output to the "+
			"desired format. (json, yaml, simple)")
}

func validateOutFlagValue(cmd *cobra.Command) error {
	if err := validateFlagStringValue([]string{"json", "yaml", "simple"},
		cmd, "out"); err != nil {
		return err
	}
	return nil
}

func validateFlagStringValue(stringSlice []string,
	cmd *cobra.Command, fName string) error {
	fValue := getFlagString(cmd, fName)
	for _, v := range stringSlice {
		if fValue == v {
			return nil
		}
	}
	return fmt.Errorf("'%s' is not a recognized value of '%s' flag. "+
		"Please choose from: [%s]\n",
		fValue, fName, strings.Join(stringSlice, ", "))
}

//
// NOTE(@bzon): All getFlag* helpers should be added below
//

// getFlagVisibility converts the string flag visiblity to gitlab.VisibilityValue.
func getFlagVisibility(cmd *cobra.Command) *gitlab.VisibilityValue {
	// TODO(@bzon): can this be improved?
	v := getFlagString(cmd, "visibility")
	return gitlab.Visibility(gitlab.VisibilityValue(v))
}

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

// TODO - to be used soon
// func getFlagInt(cmd *cobra.Command, flag string) int {
// 	i, err := cmd.Flags().GetInt(flag)
// 	if err != nil {
// 		glog.Fatalf("error accessing flag %s for command %s: %v",
// 			flag, cmd.Name(), err)
// 	}
// 	return i
// }
