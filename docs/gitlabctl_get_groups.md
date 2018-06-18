## gitlabctl get groups

List groups and subgroups

### Synopsis

List groups and subgroups

```
gitlabctl get groups [flags]
```

### Examples

```
# list all groups
gitlabctl get groups

# list all subgroups of GroupX
gitlabctl get groups --from-group=GroupX
```

### Options

```
      --all-available       Show all the groups you have access to (defaults to false for authenticated users, true for admin)
  -G, --from-group string   Use a group as the target namespace when performing the command
  -h, --help                help for groups
      --order-by string     Order groups by name or path. Default is name (default "name")
      --owned               Limit to resources owned by the current user
      --search string       Return the list of resources matching the search criteria
      --sort string         Order resources in asc or desc order. Default is asc (default "asc")
      --statistics          Include resource statistics (admins only)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
  -o, --out string      Print the command output to the desired format. (json, yaml, simple) (default "simple")
      --page int        Page of results to retrieve (default 1)
      --per-page int    The number of results to include per page (default 1)
```

### SEE ALSO

* [gitlabctl get](gitlabctl_get.md)	 - Get Gitlab resources

