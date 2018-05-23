package cmd

import (
	"testing"
)

func TestGroupLsSubgroup(t *testing.T) {
	setBasicAuthEnvs()
	tt := []struct {
		args map[string]string
	}{
		{
			args: map[string]string{
				"path": "Group2",
				"json": "true",
			},
		},
		{
			args: map[string]string{
				"path":          "Group2",
				"all-available": "true",
				"owned":         "true",
				"statistics":    "true",
				"search":        "SubGroup3",
				"sort":          "desc",
				"order-by":      "path",
			},
		},
	}

	for _, tc := range tt {
		t.Run(getSubTestNameFromFlagsMap(groupListSubgroupsCmd, tc.args),
			func(t *testing.T) {
				execT := &execTestCmdFlags{t, groupListSubgroupsCmd, tc.args}
				execT.assertNilErr()
			})
	}
}
