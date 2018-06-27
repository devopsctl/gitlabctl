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
	"strconv"
	"testing"

	gitlab "github.com/xanzy/go-gitlab"
)

func TestDeleteProjectHook(t *testing.T) {
	tt := []struct {
		name     string
		args     []string
		flagsMap map[string]string
		expect   testResult
	}{
		{
			name: "successfully delete a project hook",
			args: []string{""}, // to be filled in test setup
			flagsMap: map[string]string{
				"project": "23",
			},
			expect: pass,
		},
		{
			name: "deleting a non existent project hook should fail",
			args: []string{"55"},
			flagsMap: map[string]string{
				"project": "23",
			},
			expect: fail,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// SETUP
			// Create the project, get it's id for testing
			if tc.expect == pass {
				sampleHookURL := "http://example.com/"
				hook, err := newProjectHook("23", &gitlab.AddProjectHookOptions{
					URL: &sampleHookURL,
				})
				if err != nil {
					tInfo(fmt.Sprintf("failed to create project hook %s", tc.args[0]))
				}
				tc.args[0] = strconv.Itoa(hook.ID)
			}

			execT := execTestCmdFlags{
				t:        t,
				cmd:      deleteProjectHookCmd,
				args:     tc.args,
				flagsMap: tc.flagsMap,
			}
			stdout, execResult := execT.executeCommand()
			assertEqualResult(t, execResult, tc.expect, stdout)
			// NOTE: assert positive test
			// Validate the command output
			if tc.expect == pass {
				assertOutContains(t, stdout, "has been deleted")
			}
		})
	}
}
