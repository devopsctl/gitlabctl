## gitlabctl edit release

Update the release note of a project's release

### Synopsis

Update the release note of a project's release

```
gitlabctl edit release [flags]
```

### Examples

```
gitlabctl edit release v1.0 --project=groupx/myapp --description="Updated Release Note"
```

### Options

```
  -d, --description string   Release note
  -h, --help                 help for release
  -p, --project string       The name or ID of the project
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
  -o, --out string      Print the command output to the desired format. (json, yaml, simple) (default "simple")
```

### SEE ALSO

* [gitlabctl edit](gitlabctl_edit.md)	 - Update or patch a Gitlab resource

