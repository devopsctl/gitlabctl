## gitlabctl delete all-members

Delete all members(except creator) of a project

### Synopsis

Delete all members(except creator) of a project

```
gitlabctl delete all-members [flags]
```

### Examples

```
# remove all members(except creator) from a project
gitlabctl delete all-members --from-project Group1/Project1
```

### Options

```
  -P, --from-project string   Use a project as the target namespace when performing the command
  -h, --help                  help for all-members
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
```

### SEE ALSO

* [gitlabctl delete](gitlabctl_delete.md)	 - Delete a Gitlab resource

