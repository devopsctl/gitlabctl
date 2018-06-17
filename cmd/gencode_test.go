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
	"os"
	"testing"

	"github.com/golang/glog"
	"github.com/stretchr/testify/require"
)

func TestGenCodeCmd(t *testing.T) {
	execT := execTestCmdFlags{
		t:   t,
		cmd: genCodeCmd,
		flagsMap: map[string]string{
			"file":     "test_gencode",
			"name":     "gitlabCmd",
			"parent":   "getCmd",
			"short":    "A short desc",
			"max-args": "0",
			"use":      "gitlabgencode",
		},
		noParent: true,
	}

	// move to the root directory
	// to detect hardcoded cmd/ file path in generateCode func
	if err := os.Chdir("../"); err != nil {
		glog.Fatal(err)
	}
	stdout, execResult := execT.executeCommand()
	require.Equal(t, pass, execResult, printFlagsTable(execT.flagsMap, stdout))
	if err := os.Remove("cmd/test_gencode.go"); err != nil {
		glog.Fatal(err)
	}
	if err := os.Remove("cmd/test_gencode_test.go"); err != nil {
		glog.Fatal(err)
	}

	// ensure to go back
	err := os.Chdir("cmd")
	tInfo(err)
}
