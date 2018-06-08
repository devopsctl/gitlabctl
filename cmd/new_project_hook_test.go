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

func TestNewProjectHook(t *testing.T) {
	tt := []struct {
		name     string
		flagsMap map[string]string
		args     []string
		expect   testResult
	}{
		{
			name: "create a new project hook by project id",
			flagsMap: map[string]string{
				"url": 								"http://www.sample123.com/",
				"push-events": 						"true",
				"issues-events": 					"true",
				"confidential-issues-events":		"false",
				"merge-requests-events": 			"true",
				"tag-push-events": 					"true",
				"note-events": 						"false",
				"job-events": 						"false",
				"pipeline-events": 					"false",
				"wiki-page-events": 				"false",
				"enable-ssl-verification": 			"false",
				"token": 							"LxryAgAdBoGXdOVPshHD",
			},
			args:   []string{"22"},
			expect: pass,
		},
		{
			name: "create a new project hook by project path",
			flagsMap: map[string]string{
				"url": 								"http://www.sample321.com/",
				"push-events": 						"true",
				"issues-events": 					"true",
				"confidential-issues-events":		"false",
				"merge-requests-events": 			"true",
				"tag-push-events": 					"true",
				"note-events": 						"false",
				"job-events": 						"false",
				"pipeline-events": 					"false",
				"wiki-page-events": 				"false",
				"enable-ssl-verification": 			"false",
				"token": 							"GxryAgAdBoJRdOVPshHD",
			},
			args:   []string{"paul.lyman/project23"},
			expect: pass,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			execT := execTestCmdFlags{
				t:        t,
				cmd:      newProjectHookCmd,
				flagsMap: tc.flagsMap,
				args:     tc.args,
			}
			stdout, execResult := execT.executeCommand()
			require.Equal(t, tc.expect, execResult,
				printFlagsTable(tc.flagsMap, stdout))
		})
	}
}
