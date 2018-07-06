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

func TestEditProjectHook(t *testing.T) {
	tt := []struct {
		name     string
		flagsMap map[string]string
		args     []string
		expect   testResult
	}{
		{
			name: "edit an existing project hook by project path",
			flagsMap: map[string]string{
				"project":                    "frank.watson/project22",
				"url":                        "http://example1.com/",
				"push-events":                "false",
				"issues-events":              "false",
				"confidential-issues-events": "true",
				"merge-requests-events":      "false",
				"tag-push-events":            "false",
				"note-events":                "true",
				"job-events":                 "true",
				"pipeline-events":            "true",
				"wiki-page-events":           "true",
				"enable-ssl-verification":    "true",
				"token":                      "gAWLwCxxK5ss0z6PzgpS",
			},
			args:   []string{"4"},
			expect: pass,
		},
		{
			name: "edit an existing project hook by project id",
			flagsMap: map[string]string{
				"project":                    "23",
				"url":                        "http://example2.com/",
				"push-events":                "false",
				"issues-events":              "false",
				"confidential-issues-events": "true",
				"merge-requests-events":      "false",
				"tag-push-events":            "false",
				"note-events":                "true",
				"job-events":                 "true",
				"pipeline-events":            "true",
				"wiki-page-events":           "true",
				"enable-ssl-verification":    "true",
				"token":                      "EKdYF9AvxRWKlch0XBKO",
			},
			args:   []string{"5"},
			expect: pass,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			execT := execTestCmdFlags{
				t:        t,
				cmd:      editProjectHookCmd,
				flagsMap: tc.flagsMap,
				args:     tc.args,
			}
			stdout, execResult := execT.executeCommand()
			assertEqualResult(t, execResult, tc.expect, printFlagsTable(tc.flagsMap, stdout))
		})
	}
}
