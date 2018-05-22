package cmd

import (
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	gitlab "github.com/xanzy/go-gitlab"
)

// printGroupsOut prints the group list/get commands to a table view or json
func printGroupsOut(cmd *cobra.Command, groups ...*gitlab.Group) {
	if getFlagBool(cmd, "json") {
		printJSON(groups)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	header := []string{
		"ID", "NAME", "PATH", "PARENT_ID", "VISIBILITY",
		"REQUEST ACCESS ENABLED", "LFS ENABLED",
	}
	statsFlag := cmd.Flag("statistics")
	if statsFlag != nil {
		if cmd.Flag("statistics").Changed {
			header = append(header,
				"STORAGE SIZE", "REPO SIZE", "LFS SIZE", "JOB ARTIFACT SIZE",
			)
		}
	}
	table.SetHeader(header)

	for _, v := range groups {
		row := []string{
			strconv.Itoa(v.ID), v.Name, v.FullPath,
			strconv.Itoa(v.ParentID),
			strings.Replace(gitlab.Stringify(v.Visibility), `"`, "", -1),
			strconv.FormatBool(v.RequestAccessEnabled),
			strconv.FormatBool(v.LFSEnabled),
		}
		if statsFlag != nil {
			if getFlagBool(cmd, "statistics") {
				row = append(row, strconv.FormatInt(v.Statistics.StorageSize, 10),
					strconv.FormatInt(v.Statistics.RepositorySize, 10),
					strconv.FormatInt(v.Statistics.LfsObjectsSize, 10),
					strconv.FormatInt(v.Statistics.JobArtifactsSize, 10))
			}
		}
		table.Append(row)
	}
	table.Render()
}
