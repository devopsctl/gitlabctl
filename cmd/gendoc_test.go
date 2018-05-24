package cmd

import "testing"

func TestGenDoc(t *testing.T) {
	if _, err := executeCommand(rootCmd, "gendoc", "--dir=../docs"); err != nil {
		t.Fatal(err)
	}
}
