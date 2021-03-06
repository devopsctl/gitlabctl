## gitlabctl get project-hooks

List all project hooks of a specified project

### Synopsis

List all project hooks of a specified project

```
gitlabctl get project-hooks [flags]
```

### Examples

```
# get project hooks of projectX
gitlabctl get project-hooks projectX

# get project hooks of project with id (23)
gitlabctl get project-hooks 23
```

### Options

```
  -h, --help   help for project-hooks
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

