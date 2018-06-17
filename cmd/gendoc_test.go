package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenDoc(t *testing.T) {
	_, err := executeCommand(rootCmd, "gendoc", "--dir=../docs")
	require.Nil(t, err, err)
}
