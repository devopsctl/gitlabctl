## gitlabctl new release

Create a new release for a specified project

### Synopsis

Create a new release for a specified project

```
gitlabctl new release [flags]
```

### Examples

```
# create release from v1.0 tag of project groupx/myapp
gitlabctl new release sample --project=groupx/myapp --tag=v1.0
```

### Options

```
  -h, --help             help for release
  -p, --project string   The name or ID of the project
  -t, --tag string       The name of a tag
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
  -o, --out string      Print the command output to the desired format. (json, yaml, simple) (default "simple")
```

### SEE ALSO

* [gitlabctl new](gitlabctl_new.md)	 - Create a Gitlab resource

