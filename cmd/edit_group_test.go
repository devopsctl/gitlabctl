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

	"github.com/stretchr/testify/require"
)

func TestEditGroup(t *testing.T) {
	tt := []struct {
		name     string
		flagsMap map[string]string
		args     []string
		expect   testResult
	}{
		{
			name: "edit an existing group",
			flagsMap: map[string]string{
				"desc":                   "Updated by go test",
				"visibility":             "internal",
				"lfs-enabled":            "false",
				"request-access-enabled": "false",
			},
			args:   []string{"Group1"},
			expect: pass,
		},
		{
			name: "edit an existing subgroup",
			flagsMap: map[string]string{
				"desc": "Updated by go test",
				"out":  "yaml",
			},
			args:   []string{"Group1/SubGroup1"},
			expect: pass,
		},
		{
			name: "change a group name and path",
			flagsMap: map[string]string{
				"desc":        "Updated by go test",
				"change-name": "Renamed",
				// TODO(@bzon):
				// we need to be able to rename the path back to
				// its original path name
				// https://gitlab.com/gitlab-com/support-forum/issues/1068
				// "change-path": "RenamedPath",
				"out": "json",
			},
			args:   []string{"Group1/SubGroup2"},
			expect: pass,
		},
		{
			name: "editing a non existent group must fail",
			flagsMap: map[string]string{
				"desc": "Updated by go test",
			},
			args:   []string{"GroupXXX"},
			expect: fail,
		},
		{
			name: "invalid visibility value must fail",
			flagsMap: map[string]string{
				"visibility": "xxxx",
			},
			args:   []string{"Group1"},
			expect: fail,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			execT := execTestCmdFlags{
				t:        t,
				cmd:      editGroupCmd,
				flagsMap: tc.flagsMap,
				args:     tc.args,
			}
			stdout, execResult := execT.executeCommand()
			require.Equal(t, tc.expect, execResult,
				printFlagsTable(tc.flagsMap, stdout))
		})
	}
}
