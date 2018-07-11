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
)

func TestNewRelease(t *testing.T) {
	tt := []struct {
		name     string
		flagsMap map[string]string
		args     []string
		expect   testResult
	}{
		{
			name: "create a new release",
			flagsMap: map[string]string{
				"project":      "Group1/project1",
				"tag":          "sample_1.0",
			},
			args:   []string{"Amazing"},
			expect: pass,
		},
		{
			name:   "arg must be required",
			flagsMap: map[string]string{
				"project":      "Group1/project1",
				"tag":          "sample_1.0",
			},
			expect: fail,
		},
		{
			name:   "flags must be required",
			args:   []string{"Amazing"},
			expect: fail,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			execT := execTestCmdFlags{
				t:        t,
				cmd:      newReleaseCmd,
				flagsMap: tc.flagsMap,
				args:     tc.args,
			}
			stdout, execResult := execT.executeCommand()
			fmt.Println(stdout)
			assertEqualResult(t, execResult, tc.expect, printFlagsTable(tc.flagsMap, stdout))
		})
	}
}