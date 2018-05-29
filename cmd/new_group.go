package cmd

import (
	"github.com/spf13/cobra"
	gitlab "github.com/xanzy/go-gitlab"
)

var newGroupCmd = &cobra.Command{
	Use:   "group",
	Short: "Create a new group",
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
		if err := validateVisibilityFlagValue(cmd); err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := runNewGroup(cmd, args[0]); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	newCmd.AddCommand(newGroupCmd)
	addNewGroupFlags(newGroupCmd)
}

func runNewGroup(cmd *cobra.Command, name string) error {
	group, err := newGroup(cmd, name)
	if err != nil {
		return err
	}
	printGroupsOut(cmd, group)
	return nil
}

func newGroup(cmd *cobra.Command, name string) (*gitlab.Group, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	opts := getCreateGroupOptions(cmd)
	opts.Path = gitlab.String(name)
	opts.Name = gitlab.String(name)
	g, _, err := git.Groups.CreateGroup(opts)
	if err != nil {
		return nil, err
	}
	return g, nil
}
