package cmd

import (
	"github.com/spf13/cobra"
	gitlab "github.com/xanzy/go-gitlab"
)

var newGroupCmd = &cobra.Command{
	Use:   "group",
	Short: "Create a new group",
	PreRun: func(cmd *cobra.Command, args []string) {
		validateVisibilityFlagValue(cmd)
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := runNewGroup(cmd); err != nil {
			er(err)
		}
	},
}

func init() {
	newCmd.AddCommand(newGroupCmd)
	addNewGroupFlags(newGroupCmd)
}

func runNewGroup(cmd *cobra.Command) error {
	group, err := newGroup(cmd)
	if err != nil {
		return err
	}
	printGroupsOut(cmd, group)
	return nil
}

func newGroup(cmd *cobra.Command) (*gitlab.Group, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	opts := getCreateGroupOptions(cmd)
	g, _, err := git.Groups.CreateGroup(opts)
	if err != nil {
		return nil, err
	}
	return g, nil
}
