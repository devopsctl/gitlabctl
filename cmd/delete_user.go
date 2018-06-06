package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	gitlab "github.com/xanzy/go-gitlab"
)

var deleteUserCmd = &cobra.Command{
	Use:        "user",
	Aliases:    []string{"u"},
	SuggestFor: []string{"users"},
	Short:      "Delete a Gitlab user by specifying the username",
	Example: `# delete a user by username
gitlabctl delete user john.smith

# delete a user with user id (15)
gitlabctl delete user 15`,
	Args:          cobra.ExactArgs(1),
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return deleteUser(args[0])
	},
}

func init() {
	deleteCmd.AddCommand(deleteUserCmd)
}

// deleteUser can accept a username or user id and deletes it
func deleteUser(username string) error {
	git, err := newGitlabClient()
	if err != nil {
		return err
	}
	id, err := strconv.Atoi(username)
	// if there is an error, the username is not a number
	// therefore, search the username's user id, and then assign it to id
	if err != nil {
		users, _, err := git.Users.ListUsers(&gitlab.ListUsersOptions{
			Username: gitlab.String(username),
		})
		if err != nil {
			return err
		}
		if len(users) < 1 {
			return fmt.Errorf("username %s not found", username)
		}
		id = users[0].ID
	}
	_, err = git.Users.DeleteUser(id)
	if err != nil {
		return err
	}
	fmt.Printf("User (%s) with id (%d) has been deleted\n", username, id)
	return nil
}
