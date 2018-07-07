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

func verifyMarkFlagRequired(cmd *cobra.Command, fName string) {
	if err := cmd.MarkFlagRequired(fName); err != nil {
		glog.Fatalf("error marking %s flag as required for command %s: %v",
			fName, cmd.Name(), err)
	}
}

// addGetGroupsFlags adds common flags for `get groups` and `get subgroups` commands
// Flags usage reference:
// https://docs.gitlab.com/ce/api/groups.html#list-groups
// https://docs.gitlab.com/ce/api/groups.html#list-a-groups-s-subgroups
func addGetGroupsFlags(cmd *cobra.Command) {
	addAllAvailableFlag(cmd)
	addGroupOrderByFlag(cmd)
	addOwnedFlag(cmd)
	addSortFlag(cmd)
	addStatisticsFlag(cmd)
	addSearchFlag(cmd)
}

// addGetUsersFlags adds flags for `get users` command
// Flags usage reference:
// https://docs.gitlab.com/ce/api/users.html#for-user
func addGetUsersFlags(cmd *cobra.Command) {
	addActiveFlag(cmd)
	addBlockedFlag(cmd)
	addSearchFlag(cmd)
	addUsernameFlag(cmd)
	addExternalUIDFlag(cmd)
	addProviderFlag(cmd)
	addCreatedBefore(cmd)
	addCreatedAfter(cmd)
	addUserOrderByFlag(cmd)
	addSortFlag(cmd)
}

// addNewUserFlags adds flags for `new user` command
// Flags usage reference:
// https://docs.gitlab.com/ee/api/users.html#user-creation
func addNewUserFlags(cmd *cobra.Command) {
	addNewUserEditUserFlags(cmd)
	verifyMarkFlagRequired(cmd, "name")
	verifyMarkFlagRequired(cmd, "email")
	cmd.Flags().Bool("reset-password", false, "Send user password reset link?")
	cmd.Flags().Bool("skip-confirmation", false, "Skip confirmation")
}

// addEditUserFlags adds flags for `edit user` command
// Flags usage reference:
// https://docs.gitlab.com/ce/api/users.html#user-modification
func addEditUserFlags(cmd *cobra.Command) {
	addNewUserEditUserFlags(cmd)
	cmd.Flags().String("username", "", "New username")
	cmd.Flags().Bool("skip-reconfirmation", false, "Skip reconfirmation")
}

func addNewUserEditUserFlags(cmd *cobra.Command) {
	cmd.Flags().String("name", "", "Name")
	cmd.Flags().String("email", "", "Email")
	cmd.Flags().String("password", "", "Password")
	cmd.Flags().String("skype", "", "Skype id")
	cmd.Flags().String("linkedin", "", "Linkedin account")
	cmd.Flags().String("twitter", "", "Twitter account")
	cmd.Flags().String("website-url", "", "Website URL")
	cmd.Flags().String("org", "", "Organization name")
	cmd.Flags().String("external-uid", "", "External UID")
	cmd.Flags().String("provider", "", "External Provider Name")
	cmd.Flags().String("bio", "", "User's biography")
	cmd.Flags().String("location", "", "User's location")
	cmd.Flags().Bool("external", false, "Flags the user as external")
	cmd.Flags().Bool("admin", false, "User is admin")
	cmd.Flags().Bool("can-create-group", false, "User can create groups")
	cmd.Flags().Int("projects-limit", 5, "Number of projects user can create")
}

// addGetProjectsFlags adds flags for `get projects` command
// Flags usage reference:
// https://docs.gitlab.com/ce/api/groups.html#list-a-group-39-s-projects
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

func addProjectFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("project", "p", "", "The name or ID of the project")
}

// addNewGroupFlags add the required flags for creating a new group
// Flag usage reference: https://docs.gitlab.com/ce/api/groups.html#new-group
func addNewGroupFlags(cmd *cobra.Command) {
	addNamespaceFlag(cmd)
	addDescriptionFlag(cmd)
	addLFSenabled(cmd)
	addRequestAccessEnabledFlag(cmd)
	addVisibilityFlag(cmd)
}

// addEditGroupFlags add the required flags for updating an existing group
// Flag usage reference: https://docs.gitlab.com/ce/api/groups.html#update-group
func addEditGroupFlags(cmd *cobra.Command) {
	cmd.Flags().String("name", "", "New group name")
	cmd.Flags().String("path", "", "New group path")
	addDescriptionFlag(cmd)
	addLFSenabled(cmd)
	addRequestAccessEnabledFlag(cmd)
	addVisibilityFlag(cmd)
}

// addEditProjectFlags add the required flags for creating a new project
// Flag usage reference:
// https://docs.gitlab.com/ce/api/projects.html#edit-project
func addEditProjectFlags(cmd *cobra.Command) {
	addNewProjectFlags(cmd)
	cmd.Flags().String("name", "", "New project name")
	cmd.Flags().String("path", "", "New project path")
	cmd.Flags().String("default-branch", "master", "The default branch")
}

// addNewProjectFlags add the required flags for creating a new project
// Flag usage reference: https://docs.gitlab.com/ce/api/projects.html#create-project
func addNewProjectFlags(cmd *cobra.Command) {
	addDescriptionFlag(cmd)
	addLFSenabled(cmd)
	addRequestAccessEnabledFlag(cmd)
	addVisibilityFlag(cmd)
	// unique flags for projects
	cmd.Flags().Bool("issues-enabled", true, "Enable issues")
	cmd.Flags().Bool("merge-requests-enabled", true, "Enable merge requests")
	cmd.Flags().Bool("jobs-enabled", true, "Enable jobs")
	cmd.Flags().Bool("wiki-enabled", true, "Enable wiki")
	cmd.Flags().Bool("snippets-enabled", true, "Enable snippets")
	cmd.Flags().Bool("resolve-outdated-diff-discussions", false,
		"Automatically resolve merge request diffs discussions on lines "+
			"changed with a push")
	cmd.Flags().Bool("container-registry-enabled", false,
		"Enable container registry for this project")
	cmd.Flags().Bool("shared-runners-enabled", false,
		"Enable shared runners for this project")
	cmd.Flags().Bool("public-jobs", false, "If true, jobs can be viewed "+
		"by non-project-members")
	cmd.Flags().Bool("only-allow-merge-if-pipeline-succeeds", false,
		"Set whether merge requests can only be merged with successful jobs")
	cmd.Flags().Bool("only-allow-merge-if-discussion-are-resolved", false,
		"Set whether merge requests can only be merged "+
			"when all the discussions are resolved")
	cmd.Flags().String("merge-method", "merge",
		"Set the merge method used. (available: 'merge', 'rebase_merge', 'ff')")
	cmd.Flags().StringSlice("tag-list", []string{},
		"The list of tags for a project; put array of tags, "+
			"that should be finally assigned to a project.\n"+
			"Example: --tag-list='tag1,tag2'")
	cmd.Flags().Bool("printing-merge-request-link-enabled", true,
		"Show link to create/view merge request "+
			"when pushing from the command line")
	cmd.Flags().String("ci-config-path", "", "The path to CI config file")
}

// addNewProjectHookFlags add the required flags for creating a new project hook
// Flag usage reference: https://docs.gitlab.com/ce/api/projects.html#add-project-hook
func addNewProjectHookFlags(cmd *cobra.Command) {
	addNewProjectHookEditProjectHookFlags(cmd)
	verifyMarkFlagRequired(cmd, "url")
}

// addEditProjectHookFlags add the required flags for editing a project hook
// Flag usage reference: https://docs.gitlab.com/ce/api/projects.html#edit-project-hook
func addEditProjectHookFlags(cmd *cobra.Command) {
	addProjectFlag(cmd)
	addNewProjectHookEditProjectHookFlags(cmd)
	verifyMarkFlagRequired(cmd, "url")
	verifyMarkFlagRequired(cmd, "project")
}

func addNewProjectHookEditProjectHookFlags(cmd *cobra.Command) {
	cmd.Flags().String("url", "", "The hook URL")
	cmd.Flags().Bool("push-events", false, "Trigger hook on push events")
	cmd.Flags().Bool("issues-events", false, "Trigger hook on issues events")
	cmd.Flags().Bool("confidential-issues-events", false, "Trigger hook on confidential issues events")
	cmd.Flags().Bool("merge-requests-events", false, "Trigger hook on merge requests events")
	cmd.Flags().Bool("tag-push-events", false, "Trigger hook on tag push events")
	cmd.Flags().Bool("note-events", false, "Trigger hook on note events")
	cmd.Flags().Bool("job-events", false, "Trigger hook on job events")
	cmd.Flags().Bool("pipeline-events", false, "Trigger hook on pipeline events")
	cmd.Flags().Bool("wiki-page-events", false, "Trigger hook on wiki events")
	cmd.Flags().Bool("enable-ssl-verification", false, "Do SSL verification when triggering the hook")
	cmd.Flags().String("token", "", "Secret token to validate received payloads;"+
		"this will not be returned in the response")
}

func validateMergeMethodValue(cmd *cobra.Command) error {
	return validateFlagStringValue(
		[]string{"merge", "ff", "rebase_merge"},
		cmd, "merge-method")
}

func addDescriptionFlag(cmd *cobra.Command) {
	cmd.Flags().String("desc", "", "The description of the resource")
}

func addFromGroupFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("from-group", "G", "",
		"Use a group as the target namespace when performing the command")
}

func addFromProjectFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("from-project", "P", "",
		"Use a project as the target namespace when performing the command")
}

func validateFromGroupAndProjectFlags(cmd *cobra.Command) error {
	if getFlagString(cmd, "from-group") != "" &&
		getFlagString(cmd, "from-project") != "" {
		return newUsedTooManyFlagError("from-group", "from-project")
	}
	if getFlagString(cmd, "from-group") == "" &&
		getFlagString(cmd, "from-project") == "" {
		return newSetAtLeastOneFlagError("from-group", "from-project")
	}
	return nil
}

func validateFromProjectFlag(cmd *cobra.Command) error {
	if getFlagString(cmd, "from-project") == "" {
		return newSetAtLeastOneFlagError("from-project")
	}
	return nil
}

func validateAccessLevelFlag(cmd *cobra.Command) error {
	return validateFlagIntValue([]int{10, 20, 30, 40, 50},
		cmd, "access-level")
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
	return validateFlagStringValue([]string{"path", "name"},
		cmd, "order-by")
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
	return validateFlagStringValue([]string{"asc", "desc"},
		cmd, "sort")
}

func addActiveFlag(cmd *cobra.Command) {
	cmd.Flags().Bool("active", true, "Lookup users with active status")
}

func addBlockedFlag(cmd *cobra.Command) {
	cmd.Flags().Bool("blocked", true, "Lookup users with blocked status")
}

func addUsernameFlag(cmd *cobra.Command) {
	cmd.Flags().String("username", "", "Lookup users by username")
}

func addExternalUIDFlag(cmd *cobra.Command) {
	cmd.Flags().String("external-uid", "", "Lookup users by external uid."+
		"Best combined with --provider flag.")
}

func addProviderFlag(cmd *cobra.Command) {
	cmd.Flags().String("provider", "", "Lookup users by provider. "+
		"Best combined with --external-uid flag.")
}

func addCreatedBefore(cmd *cobra.Command) {
	cmd.Flags().String("created-before", "", "Lookup users that are "+
		"created before the specified value.\n"+
		"Example: gitlabctl get users --created-before=2001-01-02T00:00:00.060Z\n"+
		"Acceptable Format Reference: https://github.com/devopsctl/gitlabctl/blob/master/docs/dateparse.md\n")
}

func addCreatedAfter(cmd *cobra.Command) {
	cmd.Flags().String("created-after", "", "Lookup users that are "+
		"created after the specified value.\n"+
		"Example: gitlabctl get users --created-after=2001-01-02T00:00:00.060Z\n"+
		"Acceptable Format Reference: https://github.com/devopsctl/gitlabctl/blob/master/docs/dateparse.md\n")
}

// TODO: not supported by go-gitlab client yet
// func addTwoFactorFlagValue(cmd *cobra.Command) {
// 	getUsersCmd.Flags().String("two-factor", "disabled",
// 		"Filter users by Two-factor authentication. "+
// 			"Filter values are enabled or disabled. "+
// 			"By default it returns all users")
// }

// func validateTwoFactorFlagValue(cmd *cobra.Command) error {
// 	return validateFlagStringValue([]string{"enabled", "disabled"},
// 		cmd, "two-factor")
// }

func addUserOrderByFlag(cmd *cobra.Command) {
	cmd.Flags().String("order-by", "id",
		"Return projects ordered by id, name, username, created_at, updated_at "+
			" fields. Default is created_at")
}

func validateUserOrderByFlagValue(cmd *cobra.Command) error {
	return validateFlagStringValue([]string{"id", "name", "username",
		"created_at", "updated_at"},
		cmd, "order-by")
}

func addProjectOrderByFlag(cmd *cobra.Command) {
	cmd.Flags().String("order-by", "created_at",
		"Return projects ordered by id, name, path, created_at, updated_at, "+
			"or last_activity_at fields. Default is created_at")
}

func validateProjectOrderByFlagValue(cmd *cobra.Command) error {
	return validateFlagStringValue([]string{"id", "name", "path",
		"created_at", "updated_at", "last_activity_at"},
		cmd, "order-by")
}

func addNamespaceFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("namespace", "n", "",
		"This can be the parent namespace ID, group path, or user path. "+
			"(defaults to current user namespace)")
}

func addQueryFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("query", "q", "",
		"A query string to search for members"+
			"(defaults to blank)")
}

func addExpiresAtFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("expires-at", "e", "",
		"A date string in the format YEAR-MONTH-DAY"+
			"(defaults to blank)")
}

func addAccessLevelFlag(cmd *cobra.Command) {
	cmd.Flags().IntP("access-level", "a", 30,
		"Access level of member"+
			"(defaults to 30)")
}

func addVisibilityFlag(cmd *cobra.Command) {
	cmd.Flags().String("visibility", "private", "public, internal or private")
}

func validateVisibilityFlagValue(cmd *cobra.Command) error {
	return validateFlagStringValue([]string{"public", "private", "internal"},
		cmd, "visibility")
}

func addRequestAccessEnabledFlag(cmd *cobra.Command) {
	cmd.Flags().Bool("request-access-enabled", false, "Enable request access")
}

func addLFSenabled(cmd *cobra.Command) {
	cmd.Flags().Bool("lfs-enabled", false, "Enable LFS")
}

func addOutFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("out", "o", "simple",
		"Print the command output to the "+
			"desired format. (json, yaml, simple)")
}

func validateOutFlagValue(cmd *cobra.Command) error {
	return validateFlagStringValue([]string{JSON, YAML, "simple"},
		cmd, "out")
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

func validateFlagIntValue(intSlice []int,
	cmd *cobra.Command, fName string) error {
	fValue := getFlagInt(cmd, fName)
	for _, v := range intSlice {
		if fValue == v {
			return nil
		}
	}
	return fmt.Errorf("'%v' is not a recognized value of '%s' flag. "+
		"Please choose from: [%s]\n",
		fValue, fName, strings.Trim(strings.Join(strings.Split(fmt.Sprint(intSlice), " "), ","), "[]"))
}

func addPaginationFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().Int("page", 0, "Page of results to retrieve")
	cmd.PersistentFlags().Int("per-page", 0, "The number of results to include per page")
}

//
// NOTE(@bzon): All getFlag* helpers should be added below
//

// getFlagVisibility converts the string flag visiblity to gitlab.VisibilityValue.
func getFlagVisibility(cmd *cobra.Command) *gitlab.VisibilityValue {
	v := getFlagString(cmd, "visibility")
	return gitlab.Visibility(gitlab.VisibilityValue(v))
}

// getFlagMergeMethod converts the string flag merge-method to gitlab.MergeMethod
func getFlagMergeMethod(cmd *cobra.Command) *gitlab.MergeMethodValue {
	v := getFlagString(cmd, "merge-method")
	return gitlab.MergeMethod(gitlab.MergeMethodValue(v))
}

func getFlagStringSlice(cmd *cobra.Command, flag string) []string {
	s, err := cmd.Flags().GetStringSlice(flag)
	if err != nil {
		glog.Fatalf("error accessing flag %s for command %s: %v",
			flag, cmd.Name(), err)
	}
	return s
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

func getFlagInt(cmd *cobra.Command, flag string) int {
	i, err := cmd.Flags().GetInt(flag)
	if err != nil {
		glog.Fatalf("error accessing flag %s for command %s: %v",
			flag, cmd.Name(), err)
	}
	return i
}

func getFlagAccessLevel(cmd *cobra.Command, flag string) gitlab.AccessLevelValue {
	i, err := cmd.Flags().GetInt(flag)
	if err != nil {
		glog.Fatalf("error accessing flag %s for command %s: %v",
			flag, cmd.Name(), err)
	}
	return gitlab.AccessLevelValue(i)
}
