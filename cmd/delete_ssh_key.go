package cmd

import (
	"fmt"
	"strconv"
)

func deleteSSHKey(key int) error {
	git, err := newGitlabClient()
	if err != nil {
		return err
	}
	_, err = git.Users.DeleteSSHKey(key)
	if err != nil {
		return err
	}
	return nil
}

func deleteSSHKeyForUser(uid, key int) error {
	git, err := newGitlabClient()
	if err != nil {
		return err
	}
	_, err = git.Users.DeleteSSHKeyForUser(uid, key)
	if err != nil {
		return err
	}
	fmt.Printf("Deleted ssh key with id (%d) for user (%d)\n", key, uid)
	return nil
}

func deleteAllSSHKeyForUser(user string) error {
	uid, err := strconv.Atoi(user)
	// if user is not a number, find the user's id and use it as uid
	if err != nil {
		userInfo, err := getUserByUsername(user)
		if err != nil {
			return err
		}
		uid = userInfo.ID
	}
	userKeys, err := getSSHKeysForUser(uid, nil)
	if err != nil {
		return err
	}
	for _, key := range userKeys {
		if err := deleteSSHKeyForUser(uid, key.ID); err != nil {
			return err
		}
	}
	return nil
}
