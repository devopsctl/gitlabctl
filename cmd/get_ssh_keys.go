package cmd

import gitlab "github.com/xanzy/go-gitlab"

func getSSHKeysForUser(uid int, opts *gitlab.ListSSHKeysForUserOptions) ([]*gitlab.SSHKey, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	keys, _, err := git.Users.ListSSHKeysForUser(uid, opts)
	if err != nil {
		return nil, err
	}
	return keys, nil
}
