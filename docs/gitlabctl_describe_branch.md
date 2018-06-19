## gitlabctl describe branch

Describe a branch of a specified project

### Synopsis

Describe a branch of a specified project

```
gitlabctl describe branch [flags]
```

### Examples

```
gitlabctl describe master --project=devopsctl/gitlabctl
```

### Options

```
  -h, --help             help for branch
  -p, --project string   The name or ID of the project
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
  -o, --out string      Print the command output to the desired format. (json, yaml, simple) (default "simple")
```

### SEE ALSO

* [gitlabctl describe](gitlabctl_describe.md)	 - Describe a gitlab resource

