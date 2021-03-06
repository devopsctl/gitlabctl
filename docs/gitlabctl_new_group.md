## gitlabctl new group

Create a new group by specifying the group name as the first argument

### Synopsis

Create a new group by specifying the group name as the first argument

```
gitlabctl new group [flags]
```

### Examples

```
# create a new group
gitlabctl new group GroupAZ

# create a subgroup using namespace
gitlabctl new group GroupXB --namespace=ParentGroupXB
```

### Options

```
      --desc string              The description of the resource
  -h, --help                     help for group
      --lfs-enabled              Enable LFS
  -n, --namespace string         This can be the parent namespace ID, group path, or user path. (defaults to current user namespace)
      --request-access-enabled   Enable request access
      --visibility string        public, internal or private (default "private")
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
  -o, --out string      Print the command output to the desired format. (json, yaml, simple) (default "simple")
```

### SEE ALSO

* [gitlabctl new](gitlabctl_new.md)	 - Create a Gitlab resource

