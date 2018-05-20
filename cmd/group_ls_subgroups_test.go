package cmd

import (
	"testing"
)

func TestGroupLsSubGroupCmd(t *testing.T) {
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

	cmd := groupLsSubGroupCmd
	for _, tc := range tt {
		t.Run(getSubTestNameFromFlagsMap(tc.args), func(t *testing.T) {
			execT := &execTestCmdFlags{t, cmd, tc.args}
			execT.assertNilErr()
		})
	}
}
