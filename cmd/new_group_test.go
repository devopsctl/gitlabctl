package cmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGroup(t *testing.T) {
	setBasicAuthEnvs()

	tt := []struct {
		name string
		args map[string]string
	}{
		{
			name: "Use group name as namespace",
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
			name: "Use without namespace flag",
			args: map[string]string{
				"name":                   "Test_New_Group_Without_Namespace",
				"visibility":             "private",
				"lfs-enabled":            "false",
				"request-access-enabled": "false",
			},
		},
		{
			name: "Use group ID as namespace",
			args: map[string]string{
				"name":                   "Test_New_Group_Using_Namespace",
				"namespace":              "13", // is Group1
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
		group := tc.args["name"]
		if _, ok := tc.args["namespace"]; ok {
			group = "Group1/" + tc.args["name"]
		}
		if err := deleteGroup(group); err != nil {
			assert.Nil(t, err)
		}
	}
}
