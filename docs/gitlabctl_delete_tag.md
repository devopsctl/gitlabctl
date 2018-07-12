## gitlabctl delete tag

Delete a project tag

### Synopsis

Delete a project tag

```
gitlabctl delete tag [flags]
```

### Examples

```
# delete v1.0 tag from project groupx/myapp
gitlabctl delete tag v1.0 --project=groupx/myapp
```

### Options

```
  -h, --help             help for tag
  -p, --project string   The name or ID of the project
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
```

### SEE ALSO

* [gitlabctl delete](gitlabctl_delete.md)	 - Delete a Gitlab resource

