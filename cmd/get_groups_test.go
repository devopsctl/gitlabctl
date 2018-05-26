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
		result   testResult
	}{
		{
			name: "print in yaml",
			flagsMap: map[string]string{
				"out": "yaml",
			},
			result: pass,
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
			result: pass,
		},
		{
			name:   "print in simple view with no flags",
			result: pass,
		},
		// TODO: test if flags are validated properly
		// if the progam exits with error, you must put a validator
		// on this command's or its parent a Command.PersistentPreRun
		// {
		// 	name: "invalid sort value returns an error",
		// 	flagsMap: map[string]string{
		// 		"sort": "xxx",
		// 	},
		// 	result: fail,
		// },
		// {
		// 	name: "invalid order-by value returns an error",
		// 	flagsMap: map[string]string{
		// 		"order-by": "name",
		// 	},
		// 	result: fail,
		// },
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			execT := execTestCmdFlags{t, getGroupsCmd, tc.flagsMap}
			stdout, execResult := execT.executeCommand()
			require.Equal(t, tc.result, execResult,
				printFlagsTable(tc.flagsMap, stdout))
			// TODO : validate the output of the command
			// fmt.Println(stdout)
			_ = stdout
		})
	}
}
