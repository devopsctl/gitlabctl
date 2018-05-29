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

func TestGetGroups(t *testing.T) {
	setBasicAuthEnvs()
	tt := []struct {
		name     string
		flagsMap map[string]string
		expect   testResult
	}{
		{
			name: "print in yaml",
			flagsMap: map[string]string{
				"out": "yaml",
			},
			expect: pass,
		},
		{
			name: "print in json and use all flags",
			flagsMap: map[string]string{
				"all-available": "true",
				"owned":         "true",
				"statistics":    "true",
				"search":        "Group1",
				"sort":          "asc",
				"order-by":      "path",
				"out":           "json",
			},
			expect: pass,
		},
		{
			name:   "print in simple view with no flags",
			expect: pass,
		},
		{
			name: "get subgroups using from-group flag",
			flagsMap: map[string]string{
				"from-group": "Group1",
			},
			expect: pass,
		},
		{
			name: "invalid sort value must fail",
			flagsMap: map[string]string{
				"sort": "xxx",
			},
			expect: fail,
		},
		{
			name: "invalid order-by value must fail",
			flagsMap: map[string]string{
				// NOTE(@bzon): i had to set sort to correct value
				// it is not being reset to default..
				"sort":     "asc",
				"order-by": "xxx",
			},
			expect: fail,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			execT := execTestCmdFlags{
				t:        t,
				cmd:      getGroupsCmd,
				flagsMap: tc.flagsMap,
			}
			stdout, execResult := execT.executeCommand()
			require.Equal(t, tc.expect, execResult,
				printFlagsTable(tc.flagsMap, stdout))
			// TODO : validate the output of the command
			// fmt.Println(stdout)
			_ = stdout
		})
	}
}
