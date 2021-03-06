## gitlabctl new user

Create a new user by specifying the username as the first argument

### Synopsis

Create a new user by specifying the username as the first argument

```
gitlabctl new user [flags]
```

### Examples

```
# create a new user
gitlabctl new user john.smith --name="Johhny Smith" --password=12345678 --email=john.smith@example.com --skip-confirmation

# create a new user and send reset password link
gitlabctl new user james --name="james" --password=aaaaaaaa --email=aa@example.com --reset-password
```

### Options

```
      --admin                 User is admin
      --bio string            User's biography
      --can-create-group      User can create groups
      --email string          Email
      --external              Flags the user as external
      --external-uid string   External UID
  -h, --help                  help for user
      --linkedin string       Linkedin account
      --location string       User's location
      --name string           Name
      --org string            Organization name
      --password string       Password
      --projects-limit int    Number of projects user can create (default 5)
      --provider string       External Provider Name
      --reset-password        Send user password reset link?
      --skip-confirmation     Skip confirmation
      --skype string          Skype id
      --twitter string        Twitter account
      --website-url string    Website URL
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
  -o, --out string      Print the command output to the desired format. (json, yaml, simple) (default "simple")
```

### SEE ALSO

* [gitlabctl new](gitlabctl_new.md)	 - Create a Gitlab resource

