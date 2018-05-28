package cmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewGroup(t *testing.T) {
	setBasicAuthEnvs()

	tt := []struct {
		name     string
		flagsMap map[string]string
	}{
		{
			name: "create group and print in yaml",
			flagsMap: map[string]string{
				"name":                   "Test_New_Group_Under_Group1",
				"namespace":              "Group2",
				"visibility":             "private",
				"lfs-enabled":            "false",
				"request-access-enabled": "false",
				"out": "yaml",
			},
		},
		{
			name: "create group and print in json with no namespace",
			flagsMap: map[string]string{
				"name":                   "Test_New_Group_Without_Namespace",
				"visibility":             "private",
				"lfs-enabled":            "false",
				"request-access-enabled": "false",
				"out": "json",
			},
		},
		{
			name: "create group using id in namespace",
			flagsMap: map[string]string{
				"name":                   "Test_New_Group_Using_Namespace",
				"namespace":              "14", // is Group2
				"visibility":             "private",
				"lfs-enabled":            "false",
				"request-access-enabled": "false",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			execT := execTestCmdFlags{
				t:        t,
				cmd:      newGroupCmd,
				flagsMap: tc.flagsMap,
			}
			stdout, execResult := execT.executeCommand()
			require.Equal(t, execResult, pass,
				printFlagsTable(tc.flagsMap, stdout))
			// TODO : validate the output of the command
			// fmt.Println(stdout)
			_ = stdout
		})

		// Delete created group
		group := tc.flagsMap["name"]
		if _, ok := tc.flagsMap["namespace"]; ok {
			group = "Group2/" + tc.flagsMap["name"]
		}
		if err := deleteGroup(group); err != nil {
			tInfo(fmt.Sprintf("failed deleting group %s, got error: %v",
				group, err))
		}
	}
}
