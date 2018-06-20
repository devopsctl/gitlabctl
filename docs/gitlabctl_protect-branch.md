## gitlabctl protect-branch

Protect a repository branch

### Synopsis

Protect a repository branch

```
gitlabctl protect-branch [flags]
```

### Examples

```
gitlabctl protect-branch master --project=devopsctl/gitlabctl
```

### Options

```
      --dev-can-merge    Flag if developers can merge to the branch (default true)
      --dev-can-push     Flag if developers can push to the branch (default true)
  -h, --help             help for protect-branch
  -o, --out string       Print the command output to the desired format. (json, yaml, simple) (default "simple")
  -p, --project string   The name or ID of the project
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
```

### SEE ALSO

* [gitlabctl](gitlabctl.md)	 - Gitlab command-line interface

