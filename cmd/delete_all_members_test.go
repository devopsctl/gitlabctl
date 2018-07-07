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

func TestDeleteAllMembers(t *testing.T) {
	tg := []struct {
		name     string
		args     []string
		flagsMap map[string]string
		expect   testResult
	}{
		{
			name: "delete all from a project",
			flagsMap: map[string]string{
				"from-project": "TestDeleteAllMembersProject",
			},
			expect: pass,
		},
	}
	t.Run(tg[0].name, func(t *testing.T) {
		// SETUP:
		// create the test project
		gid, _ := getNamespaceID("Group1")
		if _, err := newProject(&gitlab.CreateProjectOptions{
			Path:        gitlab.String(tg[0].flagsMap["from-project"]),
			Name:        gitlab.String(tg[0].flagsMap["from-project"]),
			NamespaceID: gitlab.Int(gid),
		}); err != nil {
			tInfo(err)
		}

		tg[0].flagsMap["from-project"] = "Group1/" + tg[0].flagsMap["from-project"]
		// add project members
		if _, err := newProjectMember(tg[0].flagsMap["from-project"],
			"john.smith",
			&gitlab.AddProjectMemberOptions{
				AccessLevel: gitlab.AccessLevel(40),
			}); err != nil {
			tInfo(err)
		}

		if _, err := newProjectMember(tg[0].flagsMap["from-project"],
			"john.doe",
			&gitlab.AddProjectMemberOptions{
				AccessLevel: gitlab.AccessLevel(40),
			}); err != nil {
			tInfo(err)
		}

		execT := execTestCmdFlags{
			t:        t,
			cmd:      deleteAllMembersCmd,
			flagsMap: tg[0].flagsMap,
		}
		stdout, execResult := execT.executeCommand()
		assertEqualResult(t, tg[0].expect, execResult, stdout)

		// delete test project
		if err := deleteProject(tg[0].flagsMap["from-project"]); err != nil {
			tInfo("Removal of test data failure for delete all-members test")
		}

	})

	// negative tests
	tt := []struct {
		name      string
		flagsMap  map[string]string
		expect    testResult
		expectOut string
	}{
		{
			name: "fails when no flag is used",
			flagsMap: map[string]string{
				"from-project": "",
			},
			expect:    fail,
			expectOut: fmt.Sprint(setAtLeastOneFlagError),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			execT := execTestCmdFlags{
				t:        t,
				cmd:      deleteAllMembersCmd,
				flagsMap: tc.flagsMap,
			}
			stdout, _ := execT.executeCommand()
			assertOutContains(t, stdout, tc.expectOut)
		})
	}
}
