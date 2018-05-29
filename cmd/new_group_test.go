package cmd

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewGroup(t *testing.T) {
	setBasicAuthEnvs()
	tt := []struct {
		name     string
		flagsMap map[string]string
		args     []string
		expect   testResult
	}{
		{
			name: "create group and print in yaml",
			flagsMap: map[string]string{
				"namespace":              "Group2",
				"visibility":             "private",
				"lfs-enabled":            "false",
				"request-access-enabled": "false",
				"out": "yaml",
			},
			args:   []string{"Test_New_Group_Under_Group1"},
			expect: pass,
		},
		{
			name: "create group and print in json with no namespace",
			flagsMap: map[string]string{
				"visibility":             "private",
				"lfs-enabled":            "false",
				"request-access-enabled": "false",
				"out": "json",
			},
			args:   []string{"Test_New_Group_Without_Namespace"},
			expect: pass,
		},
		{
			name: "create group using id in namespace",
			flagsMap: map[string]string{
				"namespace":              "14", // is Group2
				"visibility":             "private",
				"lfs-enabled":            "false",
				"request-access-enabled": "false",
			},
			args:   []string{"Test_New_Group_Using_Namespace"},
			expect: pass,
		},
		{
			name: "invalid visibility value must fail",
			flagsMap: map[string]string{
				"visibility": "xxx",
			},
			args:   []string{"Dummy_Group"},
			expect: fail,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			execT := execTestCmdFlags{
				t:        t,
				cmd:      newGroupCmd,
				flagsMap: tc.flagsMap,
				args:     tc.args,
			}
			stdout, execResult := execT.executeCommand()
			if strings.Contains(stdout, "{error:") {
				tInfo(fmt.Sprintf("got error: %v", stdout))
			}
			require.Equal(t, execResult, tc.expect,
				printFlagsTable(tc.flagsMap, stdout))
		})

		// Delete created group if test is meant to pass
		if tc.expect == pass {
			g := tc.args[0]
			if _, ok := tc.flagsMap["namespace"]; ok {
				g = "Group2/" + tc.args[0]
			}
			tInfo(fmt.Sprintf("Running teardown code for created group %s", g))
			if err := deleteGroup(g); err != nil {
				tInfo(fmt.Sprintf("failed deleting group %s, got error: %v",
					g, err))
			}
		}
	}
}
