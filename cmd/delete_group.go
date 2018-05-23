package cmd

import "fmt"

func deleteGroup(path string) error {
	git, err := newGitlabClient()
	if err != nil {
		return err
	}
	gid := getGroupID(path)
	_, err = git.Groups.DeleteGroup(gid)
	if err != nil {
		return err
	}
	fmt.Printf("Group (%s) with id (%d) has been deleted\n", path, gid)
	return nil
}
