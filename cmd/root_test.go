package cmd

import (
	"testing"
)

func TestRootCmd(t *testing.T) {
	if err := rootCmd.Execute(); err != nil {
		t.Fatal(err)
	}
}
