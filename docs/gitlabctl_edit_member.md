## gitlabctl edit member

Edit a member by specifying the member name as the first argument

### Synopsis

Edit a member by specifying the member name as the first argument

```
gitlabctl edit member [flags]
```

### Examples

```
# edit member of a group
gitlabctl edit member john.smith --from-group Group2 --access-level 20

# edit member of a project
gitlabctl edit member john.smith --from-project Project1 --expire-at 2069-06-25
```

### Options

```
  -a, --access-level int      Access level of member(defaults to 30) (default 30)
  -e, --expires-at string     A date string in the format YEAR-MONTH-DAY(defaults to blank)
  -G, --from-group string     Use a group as the target namespace when performing the command
  -P, --from-project string   Use a project as the target namespace when performing the command
  -h, --help                  help for member
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
  -o, --out string      Print the command output to the desired format. (json, yaml, simple) (default "simple")
```

### SEE ALSO

* [gitlabctl edit](gitlabctl_edit.md)	 - Update or patch a Gitlab resource

