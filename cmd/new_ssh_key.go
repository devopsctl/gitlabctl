package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	gitlab "github.com/xanzy/go-gitlab"
)

var newSSHKeyCmd = &cobra.Command{
	Use:   "ssh-key",
	Short: "Upload or create ssh key for a gitlab user",
	Example: `# upload a public ssh key for the current user
gitlabctl -f=~/path/to/sshkey.pub -t"my ssh key"

# upload ssh key for another user (only for admin)
gitlabctl -f=/path/to/sshkey.pub -u="lebron.james" -t="the GOAT ssh key"

# upload ssh key for another user with user id 23
gitlabctl -f=path/to/sshkey.pub -u="23" -t="the GOAT ssh key"`,
	Aliases:    []string{"ssh"},
	SuggestFor: []string{"ssh-keys"},
	Args:       cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runNewSSHKey(cmd)
	},
}

func init() {
	newCmd.AddCommand(newSSHKeyCmd)
	newSSHKeyCmd.Flags().StringP("title", "t", "uploaded by gitlabctl",
		"SSH Key's title")
	newSSHKeyCmd.Flags().StringP("user", "u", "",
		"Upload the ssh key file to the specified user")
	newSSHKeyCmd.Flags().StringP("keyfile", "f", "",
		"Public SSH key file")
	if err := newSSHKeyCmd.MarkFlagRequired("keyfile"); err != nil {
		er(err)
	}
}

func runNewSSHKey(cmd *cobra.Command) error {
	opts, err := assignAddSSHKeyOptions(cmd)
	if err != nil {
		return err
	}
	if cmd.Flag("user").Changed {
		user := getFlagString(cmd, "user")
		// if the passed user string is a number, use it immediately
		if uid, err := strconv.Atoi(user); err == nil {
			fmt.Println("got here")
			return newSSHKey(uid, opts)
		}
		// get the user info of the passed user and use its id
		userInfo, err := getUserByUsername(user)
		if err != nil {
			return err
		}
		return newSSHKey(userInfo.ID, opts)
	}
	return newSSHKey(-1, opts)
}

// newSSHKey creates a new ssh key for the current user if uid is -1.
// If a uid greater than -1 is passed, it will upload the ssh key for that user.
func newSSHKey(uid int, opts *gitlab.AddSSHKeyOptions) error {
	git, err := newGitlabClient()
	if err != nil {
		return err
	}
	if uid == -1 {
		_, _, err := git.Users.AddSSHKey(opts)
		if err != nil {
			return err
		}
		return nil
	}
	_, _, err = git.Users.AddSSHKeyForUser(uid, opts)
	if err != nil {
		return err
	}
	return nil
}
