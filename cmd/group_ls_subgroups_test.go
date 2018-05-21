package cmd

import (
	"testing"
)

func TestGroupLsSubgroup(t *testing.T) {
	setupGitlabEnvVars()
	tt := []struct {
		args map[string]string
	}{
		{
			args: map[string]string{
				"path": "boom",
				"json": "true",
			},
		},
		{
			args: map[string]string{
				"path":          "boom",
				"all-available": "true",
				"owned":         "true",
				"statistics":    "true",
				"search":        " ",
				"sort":          "desc",
				"order-by":      "path",
			},
		},
	}

	cmd := groupListSubgroupsCmd
	for _, tc := range tt {
		t.Run(getSubTestNameFromFlagsMap(cmd, tc.args), func(t *testing.T) {
			execT := &execTestCmdFlags{t, cmd, tc.args}
			execT.assertNilErr()
		})
	}
}
