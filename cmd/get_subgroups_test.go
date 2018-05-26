package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSubgroups(t *testing.T) {
	setBasicAuthEnvs()
	tt := []struct {
		name     string
		flagsMap map[string]string
	}{
		{
			name: "print in yaml",
			flagsMap: map[string]string{
				"path": "Group2",
				"out":  "yaml",
			},
		},
		{
			name: "print in json and use all flags",
			flagsMap: map[string]string{
				"path":          "Group2",
				"all-available": "true",
				"owned":         "true",
				"statistics":    "true",
				"search":        "SubGroup3",
				"sort":          "desc",
				"order-by":      "path",
				"out":           "json",
			},
		},
		{
			name: "print in simple view",
			flagsMap: map[string]string{
				"path": "Group2",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			execT := &execTestCmdFlags{t, getSubgroupsCmd, tc.flagsMap}
			stdout, execResult := execT.executeCommand()
			require.Equal(t, execResult, pass,
				printFlagsTable(tc.flagsMap, stdout))
			// TODO : validate the output of the command
			// fmt.Println(stdout)
			_ = stdout
		})
	}
}
