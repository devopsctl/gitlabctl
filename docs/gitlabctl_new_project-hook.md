## gitlabctl new project-hook

Create a new project hook by specifying the project id or project path as the first argument

### Synopsis

Create a new project hook by specifying the project id or project path as the first argument

```
gitlabctl new project-hook [flags]
```

### Examples

```
# create a new project hook by project path
gitlabctl new project-hook GroupX/ProjectX --url="http://www.sample.com/"

# create a new project hook by project id
gitlabctl new project-hook 123 --url="http://www.sample.com/"

# create a new project hook with merge request events trigger enabled
gitlabctl new project-hook 123 --url="http://www.sample.com/" --merge-requests-events
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

* [gitlabctl new](gitlabctl_new.md)	 - Create a Gitlab resource

