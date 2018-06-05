package cmd

import (
	"fmt"
	"strconv"

	gitlab "github.com/xanzy/go-gitlab"
)

// deleteUser can accept a username or user id and deletes it
func deleteUser(username string) error {
	git, err := newGitlabClient()
	if err != nil {
		return err
	}
	id, err := strconv.Atoi(username)
	// if err is nil, the username is a string.
	// therefore search by username and get the id
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
