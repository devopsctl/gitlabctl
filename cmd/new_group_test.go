package cmd

import (
	"fmt"
	"testing"
)

func TestNewGroup(t *testing.T) {
	setBasicAuthEnvs()

	tt := []struct {
		name string
		args map[string]string
	}{
		{
			name: "All flags are used",
			args: map[string]string{
				"name":                   "Test_New_Group_Under_Group1",
				"namespace":              "Group1",
				"visibility":             "private",
				"lfs-enabled":            "false",
				"request-access-enabled": "false",
				"json": "true",
			},
		},
		{
			name: "Without namespace flag",
			args: map[string]string{
				"name":                   "Test_New_Group_Without_Namespace",
				"visibility":             "private",
				"lfs-enabled":            "false",
				"request-access-enabled": "false",
			},
		},
		{
			name: "Using an existing id as namespace",
			args: map[string]string{
				"name":                   "Test_New_Group_Using_Namespace",
				"namespace":              "13",
				"visibility":             "private",
				"lfs-enabled":            "false",
				"request-access-enabled": "false",
			},
		},
	}

	for _, tc := range tt {
		testName := getSubTestNameFromFlagsMap(newGroupCmd, tc.args)
		testName = fmt.Sprintf("[%s]%s", tc.name, testName)
		t.Run(testName, func(t *testing.T) {
			execT := execTestCmdFlags{t, newGroupCmd, tc.args}
			execT.assertNilErr()
		})
		// TODO(@bzon): delete created group
		// deleteGroup(tc.name)
		// deleteGroup("Group1/" + tc.name)
	}
}
