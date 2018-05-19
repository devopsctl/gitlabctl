package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO standardized the test data for groups to use
func TestGroupLsSubgroupCmd(t *testing.T) {
	setup()
	_, err := executeCommand(rootCmd, "group", "ls-subgroup", "--path=exb")
	assert.Nil(t, err)
}
