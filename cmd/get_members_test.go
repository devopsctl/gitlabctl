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

func TestGetMembers(t *testing.T) {
	tt := []struct {
		name      string
		flagsMap  map[string]string
		expect    testResult
		expectOut string
	}{
		{
			name: "from project print in yaml",
			flagsMap: map[string]string{
				"out":          "yaml",
				"from-project": "Group1/Project1",
				"query":        "amelia",
			},
			expect: pass,
		},
		{
			name: "from project print in json",
			flagsMap: map[string]string{
				"out":          "json",
				"from-project": "Group1/Project1",
			},
			expect: pass,
		},
		{
			name: "from project print in table",
			flagsMap: map[string]string{
				"from-project": "Group1/Project1",
			},
			expect: pass,
		},
		{
			name: "from group print in yaml",
			flagsMap: map[string]string{
				"out":        "yaml",
				"from-group": "Group2",
				"query":      "amelia",
			},
			expect: pass,
		},
		{
			name: "from group print in json",
			flagsMap: map[string]string{
				"out":        "json",
				"from-group": "Group2",
			},
			expect: pass,
		},
		{
			name: "from group print in table",
			flagsMap: map[string]string{
				"from-group": "Group2",
			},
			expect: pass,
		},
		{
			name:      "nothing is passed should fail",
			expect:    fail,
			expectOut: fmt.Sprint(setAtLeastOneFlagError),
		},
		{
			name: "from group and from project is passed should fail",
			flagsMap: map[string]string{
				"from-group":   "Group2",
				"from-project": "Group2",
			},
			expect:    fail,
			expectOut: fmt.Sprint(usedMoreThanOneFlagError),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			execT := execTestCmdFlags{
				t:        t,
				cmd:      getMembersCmd,
				flagsMap: tc.flagsMap,
			}
			stdout, execResult := execT.executeCommand()
			if tc.expect == fail {
				tInfo(stdout)
				require.Contains(t, stdout, tc.expectOut,
					printFlagsTable(tc.flagsMap, stdout))
			} else {
				require.Equal(t, tc.expect, execResult,
					printFlagsTable(tc.flagsMap, stdout))
			}
		})
	}
}
