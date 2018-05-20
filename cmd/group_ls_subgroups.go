package cmd

import (
	"log"

	gitlabctl "github.com/devopsctl/gitlabctl/gitlab"
	"github.com/spf13/cobra"
	gitlab "github.com/xanzy/go-gitlab"
)

var groupLsSubGroupCmd = &cobra.Command{
	Use:   "ls-subgroup",
	Short: "List all the projects of a group",
	Run: func(cmd *cobra.Command, args []string) {
		if err := runGroupLsSubGroup(cmd); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	groupCmd.AddCommand(groupLsSubGroupCmd)
	addPathFlag(groupLsSubGroupCmd)
	addJSONFlag(groupLsSubGroupCmd)
	addGroupLsFlags(groupLsSubGroupCmd)
}

func runGroupLsSubGroup(cmd *cobra.Command) error {
	// convert gitlab.ListGroupsOptions to gitlab.ListSubgroupsOptions
	opts := (*gitlab.ListSubgroupsOptions)(getGroupLsCmdOpts(cmd))
	path := getFlagString(cmd, "path")
	groups, err := gitlabctl.SubGroupLs(path, opts)
	if err != nil {
		return err
	}
	printGroupLsOut(cmd, groups)
	return nil
}
