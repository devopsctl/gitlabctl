package gitlabctl

import (
	gitlab "github.com/xanzy/go-gitlab"
)

func GroupLs(opts *gitlab.ListGroupsOptions) ([]*gitlab.Group, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	g, _, err := git.Groups.ListGroups(opts)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func SubGroupLs(gid interface{}, opts *gitlab.ListSubgroupsOptions) ([]*gitlab.Group, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	g, _, err := git.Groups.ListSubgroups(gid, opts)
	if err != nil {
		return nil, err
	}
	return g, nil
}
