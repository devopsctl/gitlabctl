package cmd

import (
	"github.com/spf13/cobra"
	gitlab "github.com/xanzy/go-gitlab"
)

// addGetGroupsFlags adds common flags for listing groups and subgroups commands
func addGetGroupsFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("all-available", false,
		"Show all the groups you have access to "+
			"(defaults to false for authenticated users, true for admin)")
	cmd.Flags().Bool("owned", false,
		"Limit to groups owned by the current user")
	cmd.Flags().Bool("statistics", false,
		"Include group statistics (admins only)")
	cmd.Flags().String("sort", "asc",
		"Order groups in asc or desc order. Default is asc")
	cmd.Flags().String("search", "",
		"Return the list of authorized groups matching the search criteria")
	cmd.Flags().String("order-by", "name",
		"Order groups by name or path. Default is name")
}

// getListGroupsOptions maps the cmd flags to gitlab.ListGroupsOptions struct.
// It also ensures that the struct field that is associated with the command
// flag does not use the flag default value.
func getListGroupsOptions(cmd *cobra.Command) *gitlab.ListGroupsOptions {
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
