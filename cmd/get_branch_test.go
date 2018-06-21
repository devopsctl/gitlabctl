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
	"testing"
)

func TestGetBranchCmd(t *testing.T) {
	tt := []struct {
		name      string
		args      []string
		expect    testResult
		expectOut []string
	}{
		{
			name:      "get branches of a project",
			args:      []string{"12"},
			expect:    pass,
			expectOut: []string{"master"},
		},
		{
			name:      "get branches of a non existent project fails",
			args:      []string{"xxx"},
			expect:    fail,
			expectOut: []string{`message: 404 Project Not Found`},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			execT := execTestCmdFlags{
				t:    t,
				cmd:  getBranchCmd,
				args: tc.args,
			}
			stdout, execResult := execT.executeCommand()
			assertEqualResult(t, execResult, tc.expect, stdout)
			assertOutContains(t, stdout, tc.expectOut...)
		})
	}

}
