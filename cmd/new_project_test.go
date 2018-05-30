// Copyright © 2018 github.com/devopsctl authors
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

func TestNewProject(t *testing.T) {
	tt := []struct {
		name     string
		flagsMap map[string]string
		args     []string
		expect   testResult
	}{
		{
			name: "create a new project using all flags",
			flagsMap: map[string]string{
				"namespace":                                   "Group1",
				"desc":                                        "Created by go test",
				"issues-enabled":                              "true",
				"merge-requests-enabled":                      "true",
				"jobs-enabled":                                "true",
				"wiki-enabled":                                "true",
				"snippets-enabled":                            "true",
				"resolve-outdated-diff-discussions":           "true",
				"container-registry-enabled":                  "true",
				"shared-runners-enabled":                      "true",
				"visibility":                                  "private",
				"public-jobs":                                 "true",
				"only-allow-merge-if-pipeline-succeeds":       "true",
				"only-allow-merge-if-discussion-are-resolved": "true",
				"merge-method":                                "rebase_merge",
				"lfs-enabled":                                 "true",
				"request-access-enabled":                      "true",
				"tag-list":                                    "gotest,tdd",
				"printing-merge-request-link-enabled":         "true",
				"ci-config-path":                              "gitlabci.yml",
				"out":                                         "json",
			},
			args:   []string{"ProjectX"},
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
				cmd:      newProjectCmd,
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
				// the output should contain the group to be created
				tInfo(stdout)

				// TODO: ensure flags used are also seen in stdout
				require.Contains(t, stdout, tc.args[0],
					printFlagsTable(tc.flagsMap,
						fmt.Sprintf("Expecting to see %s in stdout", tc.args[0])))

				// TEARDOWN:
				// delete created project
				pPath := tc.args[0]
				if v, ok := tc.flagsMap["namespace"]; ok {
					pPath = v + "/" + tc.args[0]
				}
				deleteProject(pPath)
			}

		})
	}
}
