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

func TestEditUser(t *testing.T) {
	tt := []struct {
		name     string
		flagsMap map[string]string
		args     []string
		expect   testResult
	}{
		{
			name: "edit user using all possible flags",
			flagsMap: map[string]string{
				// TODO: create a new user, and use new-username
				// "username":     "matt.hunter",
				"name":                "matt.swagger",
				"email":               "matt.swagger@example.com",
				"password":            "12345678",
				"external":            "false",
				"admin":               "false",
				"can-create-group":    "false",
				"skip-reconfirmation": "false",
				"skype":               "matt.swagger",
				"linkedin":            "matt.swagger",
				"twitter":             "matt.swagger",
				"website-url":         "matt.swagger.com",
				"org":                 "matt.swagger org",
				"external-uid":        "matt.swagger",
				"provider":            "github",
				"bio":                 "like a sir",
				"location":            "london",
				"projects-limit":      "5",
			},
			args:   []string{"matt.hunter"},
			expect: pass,
		},
		{
			name: "editing a non existent user should fail",
			flagsMap: map[string]string{
				"name": "x.swagger",
			},
			args:   []string{"x.hunter"},
			expect: fail,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			execT := execTestCmdFlags{
				t:        t,
				cmd:      editUserCmd,
				flagsMap: tc.flagsMap,
				args:     tc.args,
			}

			stdout, execResult := execT.executeCommand()
			require.Equal(t, tc.expect, execResult,
				printFlagsTable(tc.flagsMap, stdout))
		})
	}

}
