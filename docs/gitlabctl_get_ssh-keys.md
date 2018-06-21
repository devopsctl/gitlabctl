## gitlabctl get ssh-keys

List all ssh keys of a user

### Synopsis

List all ssh keys of a user

```
gitlabctl get ssh-keys [flags]
```

### Examples

```
# get a list of currently authenticated user ssh keys
gitlabctl get ssh-keys

# get a list of a specific user ssh keys (admin only)
gitlabctl get ssh-keys --user="lebron.james"
```

### Options

```
  -h, --help          help for ssh-keys
  -u, --user string   The username or ID of a user
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
  -o, --out string      Print the command output to the desired format. (json, yaml, simple) (default "simple")
      --page int        Page of results to retrieve
      --per-page int    The number of results to include per page
```

### SEE ALSO

* [gitlabctl get](gitlabctl_get.md)	 - Get Gitlab resources

