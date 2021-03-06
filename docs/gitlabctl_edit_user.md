## gitlabctl edit user

Modify a user by specifying the id or username and using flags for fields to modify

### Synopsis

Modify a user by specifying the id or username and using flags for fields to modify

```
gitlabctl edit user [flags]
```

### Examples

```
# modify a user bio using username
gitlabctl edit user john.smith --bio="frontend devloper"

# modify a user with id (23) 
gitlabctl edit user 23 --bio="King james is GOAT"
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
      --skip-reconfirmation   Skip reconfirmation
      --skype string          Skype id
      --twitter string        Twitter account
      --username string       New username
      --website-url string    Website URL
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
  -o, --out string      Print the command output to the desired format. (json, yaml, simple) (default "simple")
```

### SEE ALSO

* [gitlabctl edit](gitlabctl_edit.md)	 - Update or patch a Gitlab resource

