package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var getMembersCmd = &cobra.Command{
	Use:   "members",
	Short: "List all members of a group or a project",
	Example: `
# get all members of a group
gitlabctl get members --from-group=GroupX

# get all members of a project
gitlabctl get members --from-project=ProjectX
`,
	Args: fromFlagValidator,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := runGetMembers(cmd); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	getCmd.AddCommand(getMembersCmd)
	getMembersCmd.Flags().String("from-group", "", "List members of a group.")
	getMembersCmd.Flags().String("from-project", "", "List members of a project.")
}

func fromFlagValidator(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		return errors.New("this command does not need a non-flag arguments")
	}
	fromGroupF := cmd.Flag("from-group")
	fromProjectF := cmd.Flag("from-project")
	if fromGroupF.Changed && fromProjectF.Changed {
		return fmt.Errorf(`invalid usage of flag(s); only one from "from-group" and "from-project" must be used`)
	}
	if !fromGroupF.Changed && !fromProjectF.Changed {
		return fmt.Errorf(`required flag(s) "from-group" or "from-project" not set`)
	}
	return nil
}

func runGetMembers(cmd *cobra.Command) error {
	fmt.Println("run get member")
	return nil
}
