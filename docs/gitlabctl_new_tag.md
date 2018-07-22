## gitlabctl new tag

Create a new tag for a specified project

### Synopsis

Create a new tag for a specified project

```
gitlabctl new tag [flags]
```

### Examples

```
# create tag from master branch for project groupx/myapp
gitlabctl new tag v2.0 --project=groupx/myapp --ref=master

# create a tag and create a release from it
gitlabctl new tag v2.1 --project=groupx/myapp --ref=master --description="Released!"

# NOTE: You can also use 'gitlabctl new release' to create a release separately.
```

### Options

```
  -d, --description string   Create a release from the git tag with the description as the release note
  -h, --help                 help for tag
  -m, --message string       Creates annotated tag
  -p, --project string       The name or ID of the project
  -r, --ref string           The branch name or commit SHA to create branch from
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
  -o, --out string      Print the command output to the desired format. (json, yaml, simple) (default "simple")
```

### SEE ALSO

* [gitlabctl new](gitlabctl_new.md)	 - Create a Gitlab resource

