## gitlabctl delete project-hook

Delete a Gitlab project hook by specifying the project's full path or id

### Synopsis

Delete a Gitlab project hook by specifying the project's full path or id

```
gitlabctl delete project-hook [flags]
```

### Examples

```
# delete a project hook by project's path
gitlabctl delete project-hook 1 --project=GroupX/ProjectX

# delete a project hook by project id
gitlabctl delete project-hook 2 --project=22
```

### Options

```
  -h, --help             help for project-hook
  -p, --project string   The name or ID of the project
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
```

### SEE ALSO

* [gitlabctl delete](gitlabctl_delete.md)	 - Delete a Gitlab resource

