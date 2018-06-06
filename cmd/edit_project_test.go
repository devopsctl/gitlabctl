//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEditProject(t *testing.T) {
	tt := []struct {
		name     string
		flagsMap map[string]string
		args     []string
		expect   testResult
	}{
		{
			name: "edit an existing project using all flags",
			flagsMap: map[string]string{
				"desc":                                        "Updated by go test",
				"issues-enabled":                              "false",
				"merge-requests-enabled":                      "false",
				"jobs-enabled":                                "false",
				"wiki-enabled":                                "false",
				"snippets-enabled":                            "false",
				"resolve-outdated-diff-discussions":           "false",
				"container-registry-enabled":                  "false",
				"shared-runners-enabled":                      "false",
				"visibility":                                  "private",
				"public-jobs":                                 "false",
				"only-allow-merge-if-pipeline-succeeds":       "true",
				"only-allow-merge-if-discussion-are-resolved": "true",
				"merge-method":                                "ff",
				"lfs-enabled":                                 "false",
				"request-access-enabled":                      "false",
				"tag-list":                                    "patched,gotest,tdd",
				"printing-merge-request-link-enabled":         "true",
				"ci-config-path":                              "xgitlabci.yml",
				"out":                                         "json",
			},
			args:   []string{"Group2/project12"},
			expect: pass,
		},
		{
			name: "update a project by id",
			flagsMap: map[string]string{
				"desc": "updated by using id",
				"name": "Project 4",
				"out":  "json",
			},
			args:   []string{"4"},
			expect: pass,
		},
		{
			name: "invalid visibility value should fail",
			flagsMap: map[string]string{
				"visibility": "xxxx",
			},
			args:   []string{"ProjectY"},
			expect: fail,
		},
		{
			name: "invalid merge method value should fail",
			flagsMap: map[string]string{
				// fix the visibility value from above tc
				"visibility":   "private",
				"merge-method": "xxxx",
			},
			args:   []string{"ProjectZ"},
			expect: fail,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			execT := execTestCmdFlags{
				t:        t,
				cmd:      editProjectCmd,
				flagsMap: tc.flagsMap,
				args:     tc.args,
			}
			stdout, execResult := execT.executeCommand()
			if tc.expect == fail {
				tInfo(fmt.Sprintf("negative test result: %s", stdout))
			}
			require.Equal(t, tc.expect, execResult,
				printFlagsTable(tc.flagsMap, stdout))
			if tc.expect == pass {
				tInfo(stdout)
				require.Contains(t, stdout, tc.flagsMap["desc"],
					printFlagsTable(tc.flagsMap,
						fmt.Sprintf("Expecting to see '%s' in stdout",
							tc.flagsMap["desc"])))
			}
		})
	}
}
