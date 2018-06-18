## gitlabctl delete branch

Delete a project branch

### Synopsis

Delete a project branch

```
gitlabctl delete branch [flags]
```

### Examples

```
# delete a develop branch from project groupx/myapp
gitlabctl delete branch develop --project=groupx/myapp
```

### Options

```
  -h, --help             help for branch
  -p, --project string   The name or ID of the project
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
```

### SEE ALSO

* [gitlabctl delete](gitlabctl_delete.md)	 - Delete a Gitlab resource

