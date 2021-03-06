## gitlabctl describe project

Describe a project by specifying the id or project path

### Synopsis

Describe a project by specifying the id or project path

```
gitlabctl describe project [flags]
```

### Examples

```
# describe a project by path
gitlabctl describe project ProjectX
gitlabctl describe project GroupY/ProjectY

# describe a project with id (23)
gitlabctl describe project 23
```

### Options

```
  -h, --help   help for project
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
  -o, --out string      Print the command output to the desired format. (json, yaml, simple) (default "simple")
```

### SEE ALSO

* [gitlabctl describe](gitlabctl_describe.md)	 - Describe a gitlab resource

