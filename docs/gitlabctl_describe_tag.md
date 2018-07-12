## gitlabctl describe tag

Describe a tag of a specified project

### Synopsis

Describe a tag of a specified project

```
gitlabctl describe tag [flags]
```

### Examples

```
gitlabctl describe tag v5.0 --project=devopsctl/gitlabctl
```

### Options

```
  -h, --help             help for tag
  -p, --project string   The name or ID of the project
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
  -o, --out string      Print the command output to the desired format. (json, yaml, simple) (default "simple")
```

### SEE ALSO

* [gitlabctl describe](gitlabctl_describe.md)	 - Describe a gitlab resource

