package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/golang/glog"

	"github.com/spf13/cobra"
)

// addPathFlag adds the --path flag
func addPathFlag(cmd *cobra.Command) {
	cmd.Flags().String("path", "", "the group name, id or full the path including the parent group (path/to/group)")
	if err := cmd.MarkFlagRequired("path"); err != nil {
		panic(err)
	}
}

// addJSONFlag adds the --json flag for printing output to Json to a command
func addJSONFlag(cmd *cobra.Command) {
	cmd.Flags().Bool("json", false, "print the command output to json")
}

// printJSON prints the gitlab struct output to json format
func printJSON(v interface{}) {
	b, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		glog.Fatalf("error marshalling on printJson: %v", err)
	}
	fmt.Println(string(b))
}

// helper from https://github.com/kubernetes/kubernetes/blob/master/pkg/kubectl/cmd/util/helpers.go
func getFlagString(cmd *cobra.Command, flag string) string {
	s, err := cmd.Flags().GetString(flag)
	if err != nil {
		glog.Fatalf("error accessing flag %s for command %s: %v", flag, cmd.Name(), err)
	}
	return s
}

// helper from https://github.com/kubernetes/kubernetes/blob/master/pkg/kubectl/cmd/util/helpers.go
func getFlagBool(cmd *cobra.Command, flag string) bool {
	b, err := cmd.Flags().GetBool(flag)
	if err != nil {
		glog.Fatalf("error accessing flag %s for command %s: %v", flag, cmd.Name(), err)
	}
	return b
}

// helper from https://github.com/kubernetes/kubernetes/blob/master/pkg/kubectl/cmd/util/helpers.go
// Assumes the flag has a default value.
func getFlagInt(cmd *cobra.Command, flag string) int {
	i, err := cmd.Flags().GetInt(flag)
	if err != nil {
		glog.Fatalf("error accessing flag %s for command %s: %v", flag, cmd.Name(), err)
	}
	return i
}
