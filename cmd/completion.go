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
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generates bash completion scripts",
	Long: `To configure your BASH shell to load completions for each session add to your bashrc

# ~/.bashrc or ~/.profile
. <(gitlabctl completion --bash)

To configure your ZSH shell to load completions for each session do the following

gitlabctl completion --zsh > /usr/local/share/zsh/site-functions/_gitlabctl`,
	SilenceErrors: true,
	SilenceUsage:  true,
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if !cmd.Flag("bash").Changed && !cmd.Flag("zsh").Changed {
			return newSetAtLeastOneFlagError("bash", "zsh")
		}
		if cmd.Flag("bash").Changed && cmd.Flag("zsh").Changed {
			return newUsedTooManyFlagError("bash", "zsh")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if getFlagBool(cmd, "bash") {
			return rootCmd.GenBashCompletion(os.Stdout)
		}
		return rootCmd.GenZshCompletion(os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
	completionCmd.Flags().Bool("bash", false, "Generate bash completion script")
	completionCmd.Flags().Bool("zsh", false, "Generate zsh completion script")
}
