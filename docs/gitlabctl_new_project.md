## gitlabctl new project

Create a new project by specifying the project name as the first argument

### Synopsis

Create a new project by specifying the project name as the first argument

```
gitlabctl new project [flags]
```

### Examples

```
# create a new project
gitlabctl new project ProjectX --desc="Project X is party!"

# create a new project under a group
gitlabctl new project ProjectY --namespace=GroupY
```

### Options

```
      --ci-config-path string                         The path to CI config file
      --container-registry-enabled                    Enable container registry for this project
      --desc string                                   The description of the resource
  -h, --help                                          help for project
      --issues-enabled                                Enable issues (default true)
      --jobs-enabled                                  Enable jobs (default true)
      --lfs-enabled                                   Enable LFS
      --merge-method string                           Set the merge method used. (available: 'merge', 'rebase_merge', 'ff') (default "merge")
      --merge-requests-enabled                        Enable merge requests (default true)
  -n, --namespace string                              This can be the parent namespace ID, group path, or user path. (defaults to current user namespace)
      --only-allow-merge-if-discussion-are-resolved   Set whether merge requests can only be merged when all the discussions are resolved
      --only-allow-merge-if-pipeline-succeeds         Set whether merge requests can only be merged with successful jobs
      --printing-merge-request-link-enabled           Show link to create/view merge request when pushing from the command line (default true)
      --public-jobs                                   If true, jobs can be viewed by non-project-members
      --request-access-enabled                        Enable request access
      --resolve-outdated-diff-discussions             Automatically resolve merge request diffs discussions on lines changed with a push
      --shared-runners-enabled                        Enable shared runners for this project
      --snippets-enabled                              Enable snippets (default true)
      --tag-list strings                              The list of tags for a project; put array of tags, that should be finally assigned to a project.
                                                      Example: --tag-list='tag1,tag2'
      --visibility string                             public, internal or private (default "private")
      --wiki-enabled                                  Enable wiki (default true)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
  -o, --out string      Print the command output to the desired format. (json, yaml, simple) (default "simple")
```

### SEE ALSO

* [gitlabctl new](gitlabctl_new.md)	 - Create a Gitlab resource

