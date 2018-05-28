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
	"io"
	"log"
	"os"
	"testing"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type testResult string

const (
	pass testResult = "PASS"
	fail testResult = "FAIL"
)

func setE(k, v string) {
	tInfo(fmt.Sprintf("setting env '%s' to '%s'", k, v))
	if err := os.Setenv(k, v); err != nil {
		log.Fatal(err)
	}
}

func tInfo(msg string) {
	fmt.Println("--- INFO:", msg)
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
		return fmt.Errorf("Unknown flag `%s` for `%s %s`",
			flag, cmd.Parent().Name(), cmd.Name())
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

type execTestCmdFlags struct {
	t        *testing.T
	cmd      *cobra.Command
	flagsMap map[string]string
	args     []string
}

func (execT *execTestCmdFlags) executeCommand() (string, testResult) {
	// build up the command args and flags value
	args := []string{execT.cmd.Parent().Name(), execT.cmd.Name()}
	for _, arg := range execT.args {
		args = append(args, arg)
	}
	for k, v := range execT.flagsMap {
		arg := fmt.Sprintf("--%s=%s", k, v)
		args = append(args, arg)
	}

	// NOTE: program exits with error should be captured here
	stdout, err := executeCommand(rootCmd, args...)
	if err != nil {
		// any errors from the executeCommand should go here
		// and not exit the program with error!
		return err.Error(), fail
	}

	// flags must be set back to their default value
	if err := resetFlagsFromMap(execT.cmd, execT.flagsMap); err != nil {
		return err.Error(), fail
	}
	tInfo(string(pass))
	return stdout, pass
}

// A helper to ignore os.Exit(1) errors when running a cobra Command
// This was copied from https://github.com/spf13/cobra/blob/master/command_test.go
func executeCommand(root *cobra.Command,
	args ...string) (stdout string, err error) {
	for i, arg := range args {
		tInfo(fmt.Sprintf("(%d) %s", i, arg))
	}
	// programs can exit with error here..
	stdout, _, err = executeCommandC(root, args...)
	tInfo("The command successfully returned values for assertion.")
	return stdout, err
}

// A helper to ignore os.Exit(1) errors when running a cobra Command
// This was copied from https://github.com/spf13/cobra/blob/master/command_test.go
func executeCommandC(root *cobra.Command,
	args ...string) (stdout string, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOutput(buf)
	root.SetArgs(args)

	// see https://stackoverflow.com/questions/10473800/in-go-how-do-i-capture-stdout-of-a-function-into-a-string
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	_, err = root.ExecuteC()

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	stdout = <-outC

	return stdout, buf.String(), err
}

func printFlagsTable(flagsMap map[string]string, c string) string {
	buff := new(bytes.Buffer)
	header := []string{"FLAGS", "VALUE"}
	table := tablewriter.NewWriter(buff)
	table.SetHeader(header)
	for k, v := range flagsMap {
		table.Append([]string{k, v})
	}

	headerColor := tablewriter.Colors{
		tablewriter.Bold, tablewriter.BgGreenColor}
	table.SetHeaderColor(headerColor, headerColor)
	table.SetColumnColor(tablewriter.Colors{
		tablewriter.Bold, tablewriter.FgHiRedColor},
		tablewriter.Colors{
			tablewriter.Bold, tablewriter.FgHiBlackColor})

	table.SetCaption(true, c)
	table.Render()
	return buff.String()
}
