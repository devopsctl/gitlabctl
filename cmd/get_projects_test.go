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

func TestGetProjects(t *testing.T) {
	setBasicAuthEnvs()
	tt := []struct {
		name string
		args map[string]string
	}{
		{
			name: "Use with json flag",
			args: map[string]string{
				"json": "true",
			},
		},
		{
			name: "Use without json flag",
			args: map[string]string{
				"archived":                    "true",
				"search":                      "project",
				"simple":                      "false",
				"owned":                       "true",
				"sort":                        "desc",
				"order-by":                    "name",
				"membership":                  "false",
				"starred":                     "false",
				"statistics":                  "false",
				"visibility":                  "private",
				"with-issues-enabled":         "false",
				"with-merge-requests-enabled": "false",
			},
		},
	}

	for _, tc := range tt {
		t.Run(getSubTestNameFromFlagsMap(getProjectsCmd, tc.args),
			func(t *testing.T) {
				execT := execTestCmdFlags{t, getProjectsCmd, tc.args}
				execT.assertNilErr()
			})
	}
}
