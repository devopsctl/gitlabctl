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

func TestGetTags(t *testing.T) {
	tt := []struct {
		name     string
		flagsMap map[string]string
		args     []string
		expect   testResult
	}{
		{
			name: "list all tags by specifying project's id",
			args:   []string{"1"},
			expect: pass,
		},
		{
			name: "list all tags by specifying project's path",
			args:   []string{"Group1/project1"},
			expect: pass,
		},
		{
			name: "list all tags and print in json",
			flagsMap: map[string]string{
				"out":          "json",
			},
			args:   []string{"Group1/project2"},
			expect: pass,
		},
		{
			name: "list all tags and print in yaml",
			flagsMap: map[string]string{
				"out":          "yaml",
			},
			args:   []string{"Group1/project3"},
			expect: pass,
		},
		{
			name: "list non existing project tags",
			args:   []string{"4"},
			expect: pass,
		},
		{
			name: "no project id or path as argument should fail",
			expect: fail,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			execT := execTestCmdFlags{
				t:        t,
				cmd:      getTagsCmd,
				flagsMap: tc.flagsMap,
				args:     tc.args,
			}
			stdout, execResult := execT.executeCommand()
			fmt.Println(stdout)
			require.Equal(t, tc.expect, execResult,
				printFlagsTable(tc.flagsMap, stdout))
		})
	}
}
