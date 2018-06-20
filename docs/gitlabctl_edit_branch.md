## gitlabctl edit branch

Protect or unprotect a repositort branch

### Synopsis

Protect or unprotect a repositort branch

```
gitlabctl edit branch [flags]
```

### Examples

```
# protect a branch
gitlabctl edit branch master -p devopsctl/gitlabctl --protect

# unprotect a branch
gitlabctl edit branch master -p devopsctl/gitlabctl --unprotect
```

### Options

```
      --dev-can-merge    Flag if developers can merge to the branch (default true)
      --dev-can-push     Flag if developers can push to the branch (default true)
  -h, --help             help for branch
  -p, --project string   The name or ID of the project
      --protect          Protect a branch (default true)
      --unprotect        Remove protection of a branch (default true)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
  -o, --out string      Print the command output to the desired format. (json, yaml, simple) (default "simple")
```

### SEE ALSO

* [gitlabctl edit](gitlabctl_edit.md)	 - Update or patch a Gitlab resource

