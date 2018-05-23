package cmd

func deleteGroup(path string) error {
	git, err := newGitlabClient()
	if err != nil {
		return err
	}
	id := getGroupID(path)
	_, err = git.Groups.DeleteGroup(id)
	if err != nil {
		return err
	}
	return nil
}
