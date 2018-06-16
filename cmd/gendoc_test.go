package cmd

import (
	"os"
	"testing"

	"github.com/golang/glog"
)

func TestGenDoc(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		glog.Fatal(err)
	}
	docs := cwd + "/testdocs"
	os.Mkdir(docs, 0755)
	if _, err := executeCommand(rootCmd, "gendoc", "--dir="+docs); err != nil {
		t.Fatal(err)
	}
}
