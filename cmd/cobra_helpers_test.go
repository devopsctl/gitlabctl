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
	"bytes"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func setE(k, v string) {
	if err := os.Setenv(k, v); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("SETTING ENVIRONMENT %s = %s\n", k, v)
}

func unsetE(k ...string) {
	for _, v := range k {
		if err := os.Unsetenv(v); err != nil {
			log.Fatal(err)
		}
	}
}

// checkFlagExistsInCmd ensures that the command contains the flag that is
// passed from the test suite. The test suite should always be the source of
// truth with regards to which flags must exists in the command.
func checkFlagExistsInCmd(cmd *cobra.Command, flag string) error {
	if cmdFlag := cmd.Flag(flag); cmdFlag == nil {
		return fmt.Errorf("(%s) flag is not defined in (%s %s) command",
			flag, cmd.Parent().Name(), cmd.Name())
	}
	return nil
}

// setFlags assign the flag -> value from a flags map to the command
func setFlags(cmd *cobra.Command, flagsMap map[string]string) error {
	for k, v := range flagsMap {
		if err := checkFlagExistsInCmd(cmd, k); err != nil {
			return err
		}
		if err := cmd.Flags().Set(k, v); err != nil {
			return fmt.Errorf("failed setting flag %s -> %s: %v", k, v, err)
		}
	}
	return nil
}

// resetFlagsFromMap resets each key or flag name from a flags map to it's
// default value. This can be used for test teardown to ensure that the used
// flag goes back to its default value.
func resetFlagsFromMap(cmd *cobra.Command, flagsMap map[string]string) error {
	for k := range flagsMap {
		if err := checkFlagExistsInCmd(cmd, k); err != nil {
			return err
		}
		defValue := cmd.Flag(k).DefValue
		if err := cmd.Flags().Set(k, defValue); err != nil {
			return err
		}
		// NOTE(@bzon): force the Changed state of the flag to be false.
		// Because even if the flag is set to its default value,
		// it is still technically "Changed".
		cmd.Flag(k).Changed = false
	}
	return nil
}

// getSubTestNameFromFlagsMap creates a informative string according to the
// flags map that is passed. This is used to name subtests and provide a nice
// visual output of the flags and their values when running the test suite.
func getSubTestNameFromFlagsMap(cmd *cobra.Command,
	flagsMap map[string]string) string {
	name := "[COMMAND=" + cmd.Parent().Name() +
		" " + cmd.Name() + "][WITH_FLAGS="
	for k, v := range flagsMap {
		name += fmt.Sprintf("(--%s=%v)", k, v)
	}
	name += "]"
	return name
}

type execTestCmdFlags struct {
	t        *testing.T
	cmd      *cobra.Command
	flagsMap map[string]string
}

// assertNilErr asserts that errors are 'nil'
// after setting the command's flags and executing it.
func (execT *execTestCmdFlags) assertNilErr() {
	err := setFlags(execT.cmd, execT.flagsMap)
	assert.Nil(execT.t, err)
	_, err = executeCommand(rootCmd,
		execT.cmd.Parent().Name(),
		execT.cmd.Name())
	assert.Nil(execT.t, err)
	assert.Nil(execT.t, resetFlagsFromMap(execT.cmd, execT.flagsMap))
}

// A helper to ignore os.Exit(1) errors when running a cobra Command
// This was copied from https://github.com/spf13/cobra/blob/master/command_test.go
func executeCommand(root *cobra.Command,
	args ...string) (output string, err error) {

	_, output, err = executeCommandC(root, args...)
	return output, err
}

// A helper to ignore os.Exit(1) errors when running a cobra Command
// This was copied from https://github.com/spf13/cobra/blob/master/command_test.go
func executeCommandC(root *cobra.Command,
	args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	// TODO: not capturing stdout
	root.SetOutput(buf)
	root.SetArgs(args)

	c, err = root.ExecuteC()

	return c, buf.String(), err
}
