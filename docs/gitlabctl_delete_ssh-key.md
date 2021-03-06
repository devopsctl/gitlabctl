## gitlabctl delete ssh-key

Delete registered ssh keys

### Synopsis

Delete ssh keys of a gitlab user. Use 'gitlabctl get ssh' to get the list of ssh keys to delete.

```
gitlabctl delete ssh-key [flags]
```

### Examples

```
# delete ssh key with id (23) for the current authenticated user
gitlabctl delete ssh 23

# delete ssh key for a user (for admins only)
gitlabctl delete ssh 23 --user=lebron.james
```

### Options

```
  -h, --help          help for ssh-key
  -u, --user string   The user which requires removal of ssh key
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
```

### SEE ALSO

* [gitlabctl delete](gitlabctl_delete.md)	 - Delete a Gitlab resource

