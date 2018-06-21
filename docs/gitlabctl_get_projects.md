## gitlabctl get projects

List projects of the authenticated user or of a group

### Synopsis

List projects of the authenticated user or of a group

```
gitlabctl get projects [flags]
```

### Examples

```
# get all projects
gitlabctl get projects

# get all projects from a group
gitlabctl get projects --from-group=Group1
```

### Options

```
      --archived                      Limit by archived status
  -G, --from-group string             Use a group as the target namespace when performing the command
  -h, --help                          help for projects
      --membership                    Limit by projects that the current user is a member of
      --order-by string               Return projects ordered by id, name, path, created_at, updated_at, or last_activity_at fields. Default is created_at (default "created_at")
      --owned                         Limit to resources owned by the current user
      --search string                 Return the list of resources matching the search criteria
      --simple                        Return only the ID, URL, name, and path of each project
      --sort string                   Order resources in asc or desc order. Default is asc (default "asc")
      --starred                       Limit by projects starred by the current user
      --statistics                    Include resource statistics (admins only)
      --visibility string             public, internal or private (default "private")
      --with-issues-enabled           Limit by enabled issues feature
      --with-merge-requests-enabled   Limit by enabled merge requests feature
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
  -o, --out string      Print the command output to the desired format. (json, yaml, simple) (default "simple")
      --page int        Page of results to retrieve
      --per-page int    The number of results to include per page
```

### SEE ALSO

* [gitlabctl get](gitlabctl_get.md)	 - Get Gitlab resources

