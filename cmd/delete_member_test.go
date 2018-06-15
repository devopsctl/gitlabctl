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
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	gitlab "github.com/xanzy/go-gitlab"
)

func TestDeleteMember(t *testing.T) {
	tg := []struct {
		name      string
		args      []string
		flagsMap  map[string]string
		expect    testResult
		expectOut string
	}{
		{
			name: "delete from a group",
			args: []string{"kevin.mclean"},
			flagsMap: map[string]string{
				"from-group": "Group1",
			},
			expect: pass,
		},
		{
			name: "delete from a project",
			args: []string{"kylie.morrison"},
			flagsMap: map[string]string{
				"from-project": "Group1/Project1",
			},
			expect: pass,
		},
	}

	t.Run(tg[0].name, func(t *testing.T) {
		// SETUP:
		// create the test member to be removed
		if tg[0].expect == pass {
			if _, err := newGroupMember(tg[0].flagsMap["from-group"],
				tg[0].args[0],
				&gitlab.AddGroupMemberOptions{
					AccessLevel: gitlab.AccessLevel(40),
				}); err != nil {
				tInfo("Test data setup failure for delete group member")
				os.Exit(1)
			}
			tInfo("Test member to be removed is added")
		}

		execT := execTestCmdFlags{
			t:        t,
			cmd:      deleteMemberCmd,
			args:     tg[0].args,
			flagsMap: tg[0].flagsMap,
		}
		stdout, execResult := execT.executeCommand()
		fmt.Println(stdout)
		require.Equal(t, tg[0].expect, execResult, stdout)
	})

	t.Run(tg[1].name, func(t *testing.T) {
		// SETUP:
		// create the test member to be removed
		if tg[0].expect == pass {
			if _, err := newProjectMember(tg[1].flagsMap["from-project"],
				tg[1].args[0],
				&gitlab.AddProjectMemberOptions{
					AccessLevel: gitlab.AccessLevel(40),
				}); err != nil {
				tInfo("Test data setup failure for delete project member")
				os.Exit(1)
			}
			tInfo("Test member to be removed is added")
		}

		execT := execTestCmdFlags{
			t:        t,
			cmd:      deleteMemberCmd,
			args:     tg[1].args,
			flagsMap: tg[1].flagsMap,
		}
		stdout, execResult := execT.executeCommand()
		fmt.Println(stdout)
		require.Equal(t, tg[1].expect, execResult, stdout)
	})

	// negative tests
	tt := []struct {
		name      string
		args      []string
		flagsMap  map[string]string
		expect    testResult
		expectOut string
	}{
		{
			name:      "fails when from-group and from-project is not used",
			args:      []string{"amelia.walsh"},
			expect:    fail,
			expectOut: fmt.Sprint(setAtLeastOneFlagError),
		},
		{
			name: "fails when from-group and from-project are both used",
			args: []string{"amelia.walsh"},
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
				cmd:      deleteMemberCmd,
				args:     tc.args,
				flagsMap: tc.flagsMap,
			}
			stdout, _ := execT.executeCommand()
			fmt.Println(stdout)
			require.Contains(t, stdout, tc.expectOut,
				printFlagsTable(tc.flagsMap, stdout))
		})
	}
}
