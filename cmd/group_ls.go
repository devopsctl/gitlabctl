package cmd

import (
	"log"
	"os"
	"strconv"

	"github.com/devopsctl/gitlabctl/gitlabctl"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	gitlab "github.com/xanzy/go-gitlab"
)

var groupLsCmd = &cobra.Command{
	Use: "ls", Short: "List all groups",
	Run: func(cmd *cobra.Command, args []string) {
		runGroupLs(cmd)
	},
}

func init() {
	groupCmd.AddCommand(groupLsCmd)
	addJsonFlag(groupLsCmd)
	addGroupLsFlags(groupLsCmd)
}

func addGroupLsFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("all-available", false, "Show all the groups you have access to (defaults to false for authenticated users, true for admin)")
	cmd.Flags().Bool("owned", false, "Limit to groups owned by the current user")
	cmd.Flags().Bool("statistics", false, "Include group statistics (admins only)")
	cmd.Flags().String("sort", "asc", "Order groups in asc or desc order. Default is asc")
	cmd.Flags().String("search", "", "Return the list of authorized groups matching the search criteria")
	cmd.Flags().String("order-by", "name", "Order groups by name or path. Default is name")
}

// maps the cmd flags to gitlab.ListGroupsOptions struct
// It ensures that the struct field that is associated with the command flag does not use the flag default value
func getGroupLsCmdOpts(cmd *cobra.Command) *gitlab.ListGroupsOptions {
	var opts gitlab.ListGroupsOptions
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

// runGroupLs calls gitlabctl.GroupLs to return a group list
func runGroupLs(cmd *cobra.Command) {
	opts := getGroupLsCmdOpts(cmd)
	groups, err := gitlabctl.GroupLs(opts)
	if err != nil {
		log.Fatal(err)
	}
	printGroupLsOut(cmd, groups)
}

func printGroupLsOut(cmd *cobra.Command, groups []*gitlab.Group) {
	if cmd.Flag("json").Changed {
		printJson(groups)
		return
	}
	table := tablewriter.NewWriter(os.Stdout)
	header := []string{"ID", "NAME", "PATH", "VISIBILITY", "LFS ENABLED", "PARENT_ID"}
	if cmd.Flag("statistics").Changed {
		header = append(header, "STORAGE SIZE", "REPO SIZE", "LFS SIZE", "JOB ARTIFACT SIZE")
	}
	table.SetHeader(header)
	for _, v := range groups {
		row := []string{strconv.Itoa(v.ID), v.Name, v.Path, gitlab.Stringify(v.Visibility),
			strconv.FormatBool(v.LFSEnabled), strconv.Itoa(v.ParentID)}
		if cmd.Flag("statistics").Changed {
			row = append(row, strconv.FormatInt(v.Statistics.StorageSize, 10),
				strconv.FormatInt(v.Statistics.RepositorySize, 10),
				strconv.FormatInt(v.Statistics.LfsObjectsSize, 10),
				strconv.FormatInt(v.Statistics.JobArtifactsSize, 10))
		}
		table.Append(row)
	}
	table.Render()
}
