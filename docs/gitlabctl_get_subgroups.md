## gitlabctl get subgroups

List all the projects of a group

### Synopsis

List all the projects of a group

```
gitlabctl get subgroups [flags]
```

### Options

```
      --all-available     Show all the groups you have access to (defaults to false for authenticated users, true for admin)
  -h, --help              help for subgroups
      --order-by string   Order groups by name or path. Default is name (default "name")
      --owned             Limit to resources owned by the current user
  -p, --path string       the group name, id or full the path including the parent group (path/to/group)
      --search string     Return the list of resources matching the search criteria
      --sort string       Order resources in asc or desc order. Default is asc (default "asc")
      --statistics        Include resource statistics (admins only)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
  -o, --out string      Print the command output to the desired format. (json, yaml, simple) (default "simple")
```

### SEE ALSO

* [gitlabctl get](gitlabctl_get.md)	 - Get Gitlab resources

###### Auto generated by spf13/cobra on 28-May-2018
