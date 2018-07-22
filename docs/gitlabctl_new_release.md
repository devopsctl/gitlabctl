## gitlabctl new release

Create a new release for the specified project's tag

### Synopsis

Create a new release for the specified project's tag

```
gitlabctl new release [flags]
```

### Examples

```
# ensure to create the tag where the release will be created from
gitlabctl new tag v1.0 --ref=master --project=groupx/myapp

# create the release
gitlabctl new release v1.0 --project=groupx/myapp --description="Sample Release Note"
```

### Options

```
  -d, --description string   The release note or description
  -h, --help                 help for release
  -p, --project string       The name or ID of the project
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
  -o, --out string      Print the command output to the desired format. (json, yaml, simple) (default "simple")
```

### SEE ALSO

* [gitlabctl new](gitlabctl_new.md)	 - Create a Gitlab resource

