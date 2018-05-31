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

func TestGetUsers(t *testing.T) {
	tt := []struct {
		name     string
		flagsMap map[string]string
		expect   testResult
	}{
		{
			name: "lookup by username and print in yaml",
			flagsMap: map[string]string{
				"username": "frank.watson",
				"out":      "yaml",
			},
			expect: pass,
		},
		{
			name: "lookup by search and print in json",
			flagsMap: map[string]string{
				"search": "paul",
				"out":    "json",
			},
			expect: pass,
		},
		{
			name: "lookup by active users",
			flagsMap: map[string]string{
				"active": "true",
			},
			expect: pass,
		},
		{
			name: "lookup by blocked users",
			flagsMap: map[string]string{
				"active":  "false",
				"blocked": "true",
			},
			expect: pass,
		},
		{
			name: "lookup by create dates",
			flagsMap: map[string]string{
				"created-after":  "2018",
				"created-before": "2019-05-31",
			},
			expect: pass,
		},
		{
			name: "invalid order-by flag value should fail",
			flagsMap: map[string]string{
				"order-by": "xxx",
			},
			expect: fail,
		},
		{
			name: "invalid sort flag value should fail",
			flagsMap: map[string]string{
				"sort": "xxx",
			},
			expect: fail,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			execT := execTestCmdFlags{
				t:        t,
				cmd:      getUsersCmd,
				flagsMap: tc.flagsMap,
			}
			stdout, execResult := execT.executeCommand()
			fmt.Println(stdout)
			require.Equal(t, tc.expect, execResult,
				printFlagsTable(tc.flagsMap, stdout))
		})
	}
}
