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

	gitlab "github.com/xanzy/go-gitlab"
)

func TestProtectBranch(t *testing.T) {
	tt := []struct {
		name      string
		flagsMap  map[string]string
		args      []string
		expect    testResult
		expectOut []string
	}{
		{
			name: "protect an existing branch",
			flagsMap: map[string]string{
				"project":       "12",
				"protect":       "true",
				"dev-can-push":  "true",
				"dev-can-merge": "true",
				"out":           "json",
			},
			args:   []string{"release-abc"}, // to be created in SETUP
			expect: pass,
			expectOut: []string{
				`"protected": true`,
				`"developers_can_push": true`,
				`"developers_can_merge": true`,
			},
		},
		{
			name: "protect non existing branch fails",
			flagsMap: map[string]string{
				"project": "12",
				"protect": "true",
			},
			args:      []string{"xxxx"},
			expect:    fail,
			expectOut: []string{`message: 404 Branch Not Found`},
		},
		{
			name: "unprotect an existing branch",
			flagsMap: map[string]string{
				"project":   "12",
				"unprotect": "true",
				"out":       "json",
			},
			args:   []string{"master"},
			expect: pass,
			expectOut: []string{
				`"protected": false`,
				`"developers_can_push": false`,
				`"developers_can_merge": false`,
			},
		},
		{
			name: "invalid usage of flags fails",
			flagsMap: map[string]string{
				"project":       "12",
				"dev-can-merge": "true",
			},
			args:      []string{"master"},
			expect:    fail,
			expectOut: []string{"flag can only be used"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			execT := execTestCmdFlags{
				t:        t,
				cmd:      editBranchCmd,
				flagsMap: tc.flagsMap,
				args:     tc.args,
			}
			// SETUP for positive test
			if tc.expect == pass {
				// force to unchanged the flag due to concurrency issue when testing
				if _, ok := tc.flagsMap["protect"]; ok {
					execT.cmd.Flag("unprotect").Changed = false
				}
				if _, ok := tc.flagsMap["unprotect"]; ok {
					execT.cmd.Flag("protect").Changed = false
				}
				// force to create the branch to protect or unprotect
				if _, err := newBranch(tc.flagsMap["project"],
					&gitlab.CreateBranchOptions{
						Branch: gitlab.String(tc.args[0]),
						Ref:    gitlab.String("master"),
					}); err != nil {
					tInfo(err)
				}
			}
			stdout, execResult := execT.executeCommand()
			assertEqualResult(t, execResult, tc.expect,
				printFlagsTable(tc.flagsMap, stdout))
			fmt.Printf("--- OUTPUT: %s\n", stdout)
			assertOutContains(t, stdout, tc.expectOut...)
		})
	}

}
