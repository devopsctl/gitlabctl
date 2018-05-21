package gitlabctl

import (
	"github.com/xanzy/go-gitlab"
)

// GroupLs returns a gitlab.Group.
// This is used by the commmand `group ls`.
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

// SubGroupLs returns a gitlab.Group.
// This is used by the command `group ls-subgroup`.
func SubGroupLs(gid interface{},
	opts *gitlab.ListSubgroupsOptions) ([]*gitlab.Group, error) {

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
