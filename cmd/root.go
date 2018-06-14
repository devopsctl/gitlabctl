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
	"fmt"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var authDoc = `
There are two options to authenticate the command-line client to Gitlab interface:

1.) Using the 'login' command by passing the host url, username and password.

$ gitlabctl login

The login token will be saved in $HOME/.gitlabctl.yaml file.

2.) Using Environment variables.

* Basic Authentication (if using a username and password)
    - GITLAB_USERNAME
    - GITLAB_PASSWORD
    - GITLAB_HTTP_URL

* Private Token (if using a private token)
    - GITLAB_PRIVATE_TOKEN
    - GITLAB_API_HTTP_URL

* OAuth Token (if using an oauth token)
    - GITLAB_OAUTH_TOKEN
    - GITLAB_API_HTTP_URL`

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: "alpha",
	Use:     "gitlabctl",
	Short:   "Gitlab command-line interface",
	Long: `gitlabctl is a Gitlab client for the command-line.

This client helps you view, update, create, and delete Gitlab resources from the 
command-line interface.` + "\n" + authDoc,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gitlabctl.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			er(err)
		}

		// Search config in home directory with name ".gitlabctl" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".gitlabctl")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		// NOTE: the config file is not required to exists
		// raise an error if error is other than config file not found
		if !strings.Contains(err.Error(), `Config File ".gitlabctl" Not Found`) {
			er(err)
		}
	}
}
