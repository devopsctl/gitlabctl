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
	gitlab "github.com/xanzy/go-gitlab"
)

func TestDeleteProject(t *testing.T) {
	tt := []struct {
		name      string
		args      []string
		namespace string
		expect    testResult
	}{
		{
			name: "successfully delete a project",
			// NOTE: this will be created under root/
			args:   []string{"DeleteThisProject"},
			expect: pass,
		},
		{
			name:   "deleting a non existent project should fail",
			args:   []string{"UnknownProject"},
			expect: fail,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// NOTE: Setup for positive test
			// Create the project to delete
			if tc.expect == pass {
				_, err := newProject(&gitlab.CreateProjectOptions{
					Name: gitlab.String(tc.args[0]),
					Path: gitlab.String(tc.args[0]),
				})
				if err != nil {
					tInfo(fmt.Sprintf("failed to create project %s", tc.args[0]))
				}
				// NOTE: the test data is created under root/ as noted above
				tc.args[0] = "root/" + tc.args[0]
			}

			execT := execTestCmdFlags{
				t:    t,
				cmd:  deleteProjectCmd,
				args: tc.args,
			}
			stdout, execResult := execT.executeCommand()
			require.Equal(t, tc.expect, execResult, stdout)

			// NOTE: assert positive test
			// Validate the command output
			if tc.expect == pass {
				require.Contains(t, stdout, "has been deleted")
			}
		})
	}
}
