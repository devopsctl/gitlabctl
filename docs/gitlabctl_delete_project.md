## gitlabctl delete project

Delete a Gitlab project by specifying the full path

### Synopsis

Delete a Gitlab project by specifying the full path

```
gitlabctl delete project [flags]
```

### Examples

```
# delete a project
gitlabctl delete project ProjectX

# delete a project under a group
gitlabctl delete project GroupX/ProjectX
```

### Options

```
  -h, --help   help for project
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
```

### SEE ALSO

* [gitlabctl delete](gitlabctl_delete.md)	 - Delete a Gitlab resource

