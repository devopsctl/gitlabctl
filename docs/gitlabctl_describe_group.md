## gitlabctl describe group

Describe a group by specifying the id or group path

### Synopsis

Describe a group by specifying the id or group path

```
gitlabctl describe group [flags]
```

### Examples

```
# describe a group
gitlabctl describe group GroupX -o json

# describe a group by id
gitlabctl describe group 13 -o json
```

### Options

```
  -h, --help   help for group
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
  -o, --out string      Print the command output to the desired format. (json, yaml, simple) (default "simple")
```

### SEE ALSO

* [gitlabctl describe](gitlabctl_describe.md)	 - Describe a gitlab resource

