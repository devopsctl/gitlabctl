## gitlabctl login

Login to gitlab

### Synopsis

This command authenticates you to a Gitlab server, retrieves your OAuth Token and then save it to $HOME/.gitlabctl.yaml file.

```
gitlabctl login [flags]
```

### Examples

```
gitlabctl login -H http://localhost:8080
```

### Options

```
  -h, --help              help for login
  -H, --host-url string   Gitlab host url
  -p, --password string   Password
  -u, --username string   Username
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
```

### SEE ALSO

* [gitlabctl](gitlabctl.md)	 - Gitlab command-line interface

