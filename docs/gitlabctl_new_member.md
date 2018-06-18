## gitlabctl new member

Create a new member by specifying the member name as the first argument

### Synopsis

Create a new member by specifying the member name as the first argument

```
gitlabctl new member [flags]
```

### Examples

```
# create a new group
gitlabctl new member john.smith --from-group Group2 

# create a subgroup using namespace
gitlabctl new member john.smith --from-project Project1
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

* [gitlabctl new](gitlabctl_new.md)	 - Create a Gitlab resource

