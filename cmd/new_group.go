package cmd

import (
	"github.com/spf13/cobra"
	gitlab "github.com/xanzy/go-gitlab"
)

var newGroupCmd = &cobra.Command{
	Use:        "group",
	Aliases:    []string{"g"},
	SuggestFor: []string{"groups"},
	Short:      "Create a new group",
	Example: `
# create a new group
gitlabctl new group GroupAZ [flags]

# create a subgroup using namespace
gitlabctl new group GroupXB --namespace=ParentGroupXB [flags]
`,
	Args:          cobra.ExactArgs(1),
	SilenceErrors: true,
	SilenceUsage:  true,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return validateVisibilityFlagValue(cmd)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return runNewGroup(cmd, args[0])
	},
}

func init() {
	newCmd.AddCommand(newGroupCmd)
	addNewGroupFlags(newGroupCmd)
}

func runNewGroup(cmd *cobra.Command, name string) error {
	opts := getCreateGroupOptions(cmd)
	opts.Path = gitlab.String(name)
	opts.Name = gitlab.String(name)
	group, err := newGroup(opts)
	if err != nil {
		return err
	}
	printGroupsOut(cmd, group)
	return nil
}

func newGroup(opts *gitlab.CreateGroupOptions) (*gitlab.Group, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	g, _, err := git.Groups.CreateGroup(opts)
	if err != nil {
		return nil, err
	}
	return g, nil
}
