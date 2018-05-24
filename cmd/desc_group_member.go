package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
)

var descGroupMembersCmd = &cobra.Command{
	Use:     "group-member",
	Aliases: []string{"group-members"},
	Short:   "Describe a group member",
	Run: func(cmd *cobra.Command, args []string) {
		if err := runDescGroupMembers(cmd); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	descCmd.AddCommand(descGroupMembersCmd)
	addPathFlag(descGroupMembersCmd)
	addUsernameFlag(descGroupMembersCmd)
	addJSONFlag(descGroupMembersCmd)
}

func runDescGroupMembers(cmd *cobra.Command) error {
	path := getFlagString(cmd, "path")
	username := getFlagString(cmd, "username")

	groupsMember, err := descGroupMember(path, username)
	printGroupMembersOut(cmd, groupsMember)
	return err
}

func descGroupMember(gid interface{}, uid string) (*gitlab.GroupMember, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	opts := &gitlab.ListUsersOptions{
		Username: &uid,
	}
	userlist, _, err := git.Users.ListUsers(opts)
	if err != nil {
		return nil, err
	}
	if len(userlist) > 2 {
		return nil, fmt.Errorf("Fetching uid failed, Username: %v is not unique", uid)
	}
	if len(userlist) == 0 {
		return nil, fmt.Errorf("Fetching uid failed, Username: %v returned 0 results ", uid)

	}
	g, _, err := git.GroupMembers.GetGroupMember(gid, userlist[0].ID, nil)
	if err != nil {
		return nil, err
	}
	return g, nil
}
