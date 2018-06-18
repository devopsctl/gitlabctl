## gitlabctl edit project-hook

Edit a project hook by specifying the project id or path and using flags for fields to modify

### Synopsis

Edit a project hook by specifying the project id or path and using flags for fields to modify

```
gitlabctl edit project-hook [flags]
```

### Examples

```
# update a project hook by project path
gitlabctl edit project-hook 1 --project=ProjectX --url="http://www.sample123.com/"
gitlabctl edit project-hook 2 --project=GroupX/ProjectX --tag-push-events=false  

# update a project hook by project id
gitlabctl edit project-hook 3 --project=3 --url="http://www.sample321.com/" --issues-events
```

### Options

```
      --confidential-issues-events   Trigger hook on confidential issues events
      --enable-ssl-verification      Do SSL verification when triggering the hook
  -h, --help                         help for project-hook
      --issues-events                Trigger hook on issues events
      --job-events                   Trigger hook on job events
      --merge-requests-events        Trigger hook on merge requests events
      --note-events                  Trigger hook on note events
      --pipeline-events              Trigger hook on pipeline events
  -p, --project string               The name or ID of the project
      --push-events                  Trigger hook on push events
      --tag-push-events              Trigger hook on tag push events
      --token string                 Secret token to validate received payloads;this will not be returned in the response
      --url string                   The hook URL
      --wiki-page-events             Trigger hook on wiki events
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
  -o, --out string      Print the command output to the desired format. (json, yaml, simple) (default "simple")
```

### SEE ALSO

* [gitlabctl edit](gitlabctl_edit.md)	 - Update or patch a Gitlab resource

