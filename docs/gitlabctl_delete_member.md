## gitlabctl delete member

Delete a member by specifying the member name as the first argument

### Synopsis

Delete a member by specifying the member name as the first argument

```
gitlabctl delete member [flags]
```

### Examples

```
# remove a member from a group
gitlabctl delete member john.smith --from-group Group2 

# remove a member from a project
gitlabctl delete member john.smith --from-project Group1/Project1
```

### Options

```
  -G, --from-group string     Use a group as the target namespace when performing the command
  -P, --from-project string   Use a project as the target namespace when performing the command
  -h, --help                  help for member
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
```

### SEE ALSO

* [gitlabctl delete](gitlabctl_delete.md)	 - Delete a Gitlab resource

