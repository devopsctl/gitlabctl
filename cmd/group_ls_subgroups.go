package cmd

import (
	"log"

	"github.com/devopsctl/gitlabctl/gitlabctl"
	"github.com/spf13/cobra"
	gitlab "github.com/xanzy/go-gitlab"
)

var groupLsSubGroupCmd = &cobra.Command{
	Use:   "ls-subgroup",
	Short: "List all the projects of a group",
	Run: func(cmd *cobra.Command, args []string) {
		runGroupLsSubGroup(cmd)
	},
}

func init() {
	groupCmd.AddCommand(groupLsSubGroupCmd)
	addPathFlag(groupLsSubGroupCmd)
	addJsonFlag(groupLsSubGroupCmd)
	addGroupLsFlags(groupLsSubGroupCmd)
}

// maps the cmd flags to gitlab.ListSubgroupsOptions struct
// It ensures that the struct field that is associated with the command flag does not use the flag default value
func getSubGroupLsCmdOpts(cmd *cobra.Command) *gitlab.ListSubgroupsOptions {
	var opts gitlab.ListSubgroupsOptions
	if cmd.Flag("all-available").Changed {
		opts.AllAvailable = gitlab.Bool(getFlagBool(cmd, "all-available"))
	}
	if cmd.Flag("owned").Changed {
		opts.Owned = gitlab.Bool(getFlagBool(cmd, "owned"))
	}
	if cmd.Flag("statistics").Changed {
		opts.Statistics = gitlab.Bool(getFlagBool(cmd, "statistics"))
	}
	if cmd.Flag("sort").Changed {
		opts.Sort = gitlab.String(getFlagString(cmd, "sort"))
	}
	if cmd.Flag("search").Changed {
		opts.Search = gitlab.String(getFlagString(cmd, "search"))
	}
	if cmd.Flag("order-by").Changed {
		opts.OrderBy = gitlab.String(getFlagString(cmd, "order-by"))
	}
	return &opts
}

func runGroupLsSubGroup(cmd *cobra.Command) {
	opts := getSubGroupLsCmdOpts(cmd)
	path := getFlagString(cmd, "path")
	groups, err := gitlabctl.SubGroupLs(path, opts)
	if err != nil {
		log.Fatal(err)
	}
	printGroupLsOut(cmd, groups)
}
