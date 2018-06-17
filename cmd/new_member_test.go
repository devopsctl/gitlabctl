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

func TestNewMember(t *testing.T) {
	tt := []struct {
		name      string
		args      []string
		flagsMap  map[string]string
		expect    testResult
		expectOut string
	}{
		{
			name: "add a member using username to a project",
			args: []string{"kevin.mclean"},
			flagsMap: map[string]string{
				"from-project": "Group1/Project1",
				"expires-at":   "2069-06-09",
				"access-level": "10",
			},
			expect: pass,
		},
		{
			name: "add a member using username to a group",
			args: []string{"amelia.walsh"},
			flagsMap: map[string]string{
				"from-group":   "Group1",
				"expires-at":   "2069-06-09",
				"access-level": "30",
			},
			expect: pass,
		},
		{
			name:      "missing required flag fails",
			args:      []string{"amelia.walsh"},
			expect:    fail,
			expectOut: fmt.Sprint(setAtLeastOneFlagError),
		},
		{
			name: "using invalid access level flag value fails",
			args: []string{"john.doe"},
			flagsMap: map[string]string{
				"from-group":   "Group2",
				"access-level": "21",
			},
			expect:    fail,
			expectOut: "not a recognized value of 'access-level' flag",
		},
		{
			name: "using both from-group and from-project fails",
			args: []string{"amelia.walsh"},
			flagsMap: map[string]string{
				"from-group":   "Group2",
				"from-project": "Group2",
				"access-level": "30",
			},
			expect:    fail,
			expectOut: fmt.Sprint(usedMoreThanOneFlagError),
		},
	}

	for _, tc := range tt {
		// SETUP
		// Ensure username is not a member of the group
		if tc.expect == pass {
			if g, ok := tc.flagsMap["from-group"]; ok {
				if err := deleteGroupMember(g, tc.args[0]); err != nil {
					tInfo(err)
				}
			}
			if p, ok := tc.flagsMap["from-project"]; ok {
				if err := deleteProjectMember(p, tc.args[0]); err != nil {
					tInfo(err)
				}
			}
		}
		t.Run(tc.name, func(t *testing.T) {
			execT := execTestCmdFlags{
				t:        t,
				args:     tc.args,
				cmd:      newMemberCmd,
				flagsMap: tc.flagsMap,
			}
			stdout, execResult := execT.executeCommand()
			if tc.expect == fail {
				require.Contains(t, stdout, tc.expectOut,
					printFlagsTable(tc.flagsMap, stdout))
			} else {
				require.Equal(t, tc.expect, execResult,
					printFlagsTable(tc.flagsMap, stdout))
			}
		})
	}
}
