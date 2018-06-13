package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/howeyc/gopass"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:           "login",
	Short:         "Login to gitlab",
	Long:          `This command authenticates you to a Gitlab server, retrieves your OAuth Token and then save it to $HOME/.gitlabctl.yaml file.`,
	Example:       `gitlabctl login -H http://localhost:8080`,
	Args:          cobra.ExactArgs(0),
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runLogin(cmd)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("host-url", "H", "", "Gitlab host url")
	loginCmd.Flags().StringP("username", "u", "", "Username")
	loginCmd.Flags().StringP("password", "p", "", "Password")
}

func runLogin(cmd *cobra.Command) error {
	host := getFlagString(cmd, "host-url")
	if host == "" {
		h, err := promptStringInput("gitlab host url")
		if err != nil {
			return err
		}
		host = h
	}
	username := getFlagString(cmd, "username")
	if username == "" {
		u, err := promptStringInput("gitlab username")
		if err != nil {
			return err
		}
		username = u
	}
	password := getFlagString(cmd, "password")
	if password == "" {
		pw, err := promptPasswordInput()
		if err != nil {
			return err
		}
		password = pw
	}

	uri := fmt.Sprintf("/oauth/token?grant_type=password&username=%s&password=%s", username, password)
	resp, err := http.Post(host+uri, "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var cfgMap map[string]interface{}
	if err := json.Unmarshal(b, &cfgMap); err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Login failed: %s", cfgMap["error_description"])
	}

	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	cfgFile = home + "/.gitlabctl.yaml"

	// add host_url and user to config file
	cfgMap["host_url"] = host
	cfgMap["user"] = username

	b, err = yaml.Marshal(cfgMap)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(cfgFile, b, 0600); err != nil {
		return err
	}

	fmt.Printf("%s file has been created by login command\n", cfgFile)
	return nil
}

func promptPasswordInput() (string, error) {
	fmt.Print(">> Enter gitlab password: ")
	password, err := gopass.GetPasswd()
	return strings.TrimSpace(string(password)), err
}

func promptStringInput(askFor string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(">> Enter %s: ", askFor)
	username, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	// convert CRLF to LF
	return strings.Replace(username, "\n", "", -1), nil
}
