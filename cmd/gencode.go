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
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

var genCodeCmd = &cobra.Command{
	Use:               "gencode",
	Aliases:           []string{"gc"},
	Hidden:            true,
	DisableAutoGenTag: true,
	Short:             "For gitlabctl developers only",
	Run: func(cmd *cobra.Command, args []string) {
		generateCode(cmd)
	},
}

type genCodeCfg struct {
	FileName, Parent, Name, Short, Use, TestFuncName string
	MaxArgs                                          int
}

var genCode genCodeCfg

func init() {
	rootCmd.AddCommand(genCodeCmd)
	genCodeCmd.Flags().StringVarP(&genCode.FileName, "file", "f", "", "command file name")
	genCodeCmd.MarkFlagRequired("file")
	genCodeCmd.Flags().StringVarP(&genCode.Use, "use", "u", "", "command use")
	genCodeCmd.MarkFlagRequired("use")
	genCodeCmd.Flags().StringVarP(&genCode.Parent, "parent", "p", "", "command parent variable name")
	genCodeCmd.MarkFlagRequired("parent")
	genCodeCmd.Flags().StringVarP(&genCode.Name, "name", "n", "", "command variable name")
	genCodeCmd.MarkFlagRequired("name")
	genCodeCmd.Flags().StringVarP(&genCode.Short, "short", "s", "", "command short description")
	genCodeCmd.MarkFlagRequired("short")
	genCodeCmd.Flags().IntVar(&genCode.MaxArgs, "max-args", 0, "command max accepted args")
	genCodeCmd.MarkFlagRequired("max-args")
}

func generateCode(cmd *cobra.Command) {
	tmplFiles := []string{"cmd/new_subcommand.gotmpl", "cmd/new_subcommand_test.gotmpl"}
	genCode.TestFuncName = "Test" + strings.Title(genCode.Name)

	for _, tmplFile := range tmplFiles {
		b, err := ioutil.ReadFile(tmplFile)
		if err != nil {
			glog.Fatal(err)
		}
		tmpl := template.Must(template.New("").Parse(string(b)))

		fileExt := ".go"
		if strings.Contains(tmplFile, "_test.gotmpl") {
			fileExt = "_test.go"
		}
		f, err := os.Create("cmd/" + strings.Replace(genCode.FileName,
			".go", "", -1) + fileExt)
		if err != nil {
			glog.Fatal(err)
		}
		err = tmpl.Execute(f, genCode)
		if err != nil {
			glog.Fatal(err)
		}
	}
}
