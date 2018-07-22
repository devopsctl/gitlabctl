[![Build Status](https://travis-ci.org/devopsctl/gitlabctl.svg?branch=master)](https://travis-ci.org/devopsctl/gitlabctl)
[![codecov](https://codecov.io/gh/devopsctl/gitlabctl/branch/master/graph/badge.svg)](https://codecov.io/gh/devopsctl/gitlabctl)
[![Go Report Card](https://goreportcard.com/badge/github.com/devopsctl/gitlabctl)](https://goreportcard.com/report/github.com/devopsctl/gitlabctl)
[![Waffle board](https://badge.waffle.io/devopsctl/gitlabctl.png?columns=all)](https://waffle.io/devopsctl/gitlabctl)

# [`gitlabctl`](https://devopsctl.github.io/gitlabctl/)

Be a rockstar and efficiently manage your team's gitlab.org or [self-hosted Gitlab](https://about.gitlab.com/installation/) projects, groups, users and other resources.

Tested with [__Gitlab 11.1__](https://about.gitlab.com/2018/07/22/gitlab-11-1-released/).

## Getting Started

Complete the [installation](#installation) guide below and then follow the [quickstart](#quickstart) guide.

If you are hooked, please visit the awesome [Project Github site](https://devopsctl.github.io/gitlabctl/) to read the full manual of each command.

If you just want to quickly check what are the available commands, go to [Gitlab Commands Available](#gitlab-commands-available) in this page.

## Installation

Get the download link of your preferred platform binary from [RELEASES](https://github.com/devopsctl/gitlabctl/releases).

### OSX

```bash
GITLABCTL_BIN_DIR=$HOME/.gitlabctl/bin

mkdir -p $GITLABCTL_BIN_DIR && curl -Lo gitlabctl \
 https://github.com/devopsctl/gitlabctl/releases/download/v0.3.0/gitlabctl-darwin-amd64 \
 && chmod +x gitlabctl \
 && mv gitlabctl $GITLABCTL_BIN_DIR

export PATH=$PATH:$GITLABCTL_BIN_DIR

# You can also append this export command in $HOME/.profile or similar SHELL file
# to make sure it is included in your `$PATH` variable 
# automatically upon starting a new terminal session.
```

### Linux

```bash
GITLABCTL_BIN_DIR=$HOME/.gitlabctl/bin

mkdir -p $GITLABCTL_BIN_DIR && curl -Lo gitlabctl \
 https://github.com/devopsctl/gitlabctl/releases/download/v0.3.0/gitlabctl-linux-amd64 \
 && chmod +x gitlabctl \
 && mv gitlabctl $GITLABCTL_BIN_DIR

export PATH=$PATH:$GITLABCTL_BIN_DIR

# You can also append this export command in $HOME/.profile or similar SHELL file
# to make sure it is included in your `$PATH` variable 
# automatically upon starting a new terminal session.
```

### Windows

Download the gitlabctl-windows-amd64.exe file, rename it to gitlabctl.exe and add it to your Windows System Environment Variable `%PATH%`.

### Auto Complete

Enable auto complete for __bash__ or __zsh__ shell. ❤️

```bash
# follow the instructions from the command output
gitlabctl completion -h
```

## Quickstart

### Authenticating to Gitlab server

Using `gitlabctl login` to fetch personal access token

```bash
gitlabctl login

>> Enter gitlab host url: http://localhost:10080
>> Enter gitlab username: root
>> Enter gitlab password: *****
/Users/jb/.gitlabctl.yaml file has been created by login command
```

__Windows Users:__ password masking does not work in Git bash.

Using environment variables. See `gitlabctl -h`

### Using the help commands

Use __-h__ flag when possible. 

```bash
gitlabctl [command] -h
```

or

```bash
gitlabctl [command] [subcommand] -h
```

### Usage Examples

Fetching resources with using `--output, -o` formatter flag.

```bash
gitlabctl get groups

+----+------------------+------------------------------------------------+-----------+
| ID |       PATH       |                      URL                       | PARENT ID |
+----+------------------+------------------------------------------------+-----------+
| 13 | Group1           | http://localhost:10080/groups/Group1           |         0 |
| 14 | Group2           | http://localhost:10080/groups/Group2           |         0 |
| 16 | Group1/SubGroup2 | http://localhost:10080/groups/Group1/SubGroup2 |        13 |
| 15 | Group1/SubGroup1 | http://localhost:10080/groups/Group1/SubGroup1 |        13 |
| 17 | Group2/SubGroup3 | http://localhost:10080/groups/Group2/SubGroup3 |        14 |
| 18 | Group2/SubGroup4 | http://localhost:10080/groups/Group2/SubGroup4 |        14 |
+----+------------------+------------------------------------------------+-----------+
```

View command outout in __json__.

```bash
gitlabctl get groups -o json

[
 {
  "id": 13,
  "name": "Group1",
  "path": "Group1",
  "description": "Updated by go test by id",
  "visibility": "private",
  "lfs_enabled": false,
  "avatar_url": "",
  "web_url": "http://localhost:10080/groups/Group1",
  "request_access_enabled": false,
  "full_name": "Group1",
  "full_path": "Group1",
  "parent_id": 0,
  "projects": null,
  "statistics": null
 },
 {
  "id": 14,
  "name": "Group2",
  "path": "Group2",
  "description": "",
  "visibility": "private",
  "lfs_enabled": true,
  "avatar_url": "",
  "web_url": "http://localhost:10080/groups/Group2",
  "request_access_enabled": false,
  "full_name": "Group2",
  "full_path": "Group2",
  "parent_id": 0,
  "projects": null,
  "statistics": null
 },
 ]
```

View the command output in __yaml__.

```bash
gitlabctl get groups -o yaml

- avatar_url: ""
  description: Updated by go test by id
  full_name: Group1
  full_path: Group1
  id: 13
  lfs_enabled: false
  name: Group1
  parent_id: 0
  path: Group1
  projects: null
  request_access_enabled: false
  statistics: null
  visibility: private
  web_url: http://localhost:10080/groups/Group1
- avatar_url: ""
  description: ""
  full_name: Group2
  full_path: Group2
  id: 14
  lfs_enabled: true
  name: Group2
  parent_id: 0
  path: Group2
  projects: null
  request_access_enabled: false
  statistics: null
  visibility: private
  web_url: http://localhost:10080/groups/Group2
```

Create gitlab resources.

```bash
# create a group
gitlabctl new group devopsctl

# create a project under devopsctl group
gitlabctl new project gitlab-cli --namespace=devopsctl

# create a new user with username john
gitlabctl new user john --name="John Smith" --password="john123456" --email=john@example.com --reset-password 
```

## Gitlab Commands Available 
### Authentication 

* [x] Authentication through environment variables.
* [x] Authentication using `gitlabctl login` command.

### Completion

* [x] `completion --bash`
* [x] `completion --zsh`

### Group

* [x] `get groups [flags]`
* [x] `get groups --from-group [flags]`
* [x] `describe group [group id or group path] [flags]`
* [x] `new group [group name] [flags]`
* [x] `delete group [group id or group path]`
* [x] `edit group [group id or group path] [flags]`

### Project

* [x] `get projects [flags]`
* [x] `get projects --from-group  [flags]`
* [x] `describe project [project id or project path]`
* [x] `new project [project name] [flags]`
* [x] `edit project [project id or project path] [flags]`
* [x] `delete project [project id or project path]`

### Project/Repository Branch

* [x] `get branch [project id or project path] [flags]`
* [x] `describe branch [branch name] [--project] [flags]`
* [x] `new branch [branch name] [--project] [flags]`
* [x] `delete branch [branch name] [--project]`
* [x] `edit branch [branch name] [--project] [--protect] [flags]`
* [x] `edit branch [branch name] [--project] [--unprotect] [flags]`

### Project/Repository Tags and Releases

* [x] `get tags [project id or project path] [flags]`
* [x] `describe tag [tag name] [--project] [flags]`
* [x] `new tag [tag name] [--project] [flags]`
* [x] `delete tag [tag name] [--project]`
* [x] `new release [tag name] [--project] [flags]`
* [x] `edit release [tag name] [--project] [flags]`

### Project Hooks

* [x] `get project-hooks [project id or project path] [flags]`
* [x] `new project-hook [project id or project path] [flags]`
* [x] `edit project-hook [hook id] [--project] [flags]`
* [x] `delete project-hook [hook id] [--project]`

### Users 

* [x] `get users [flags]`
* [x] `describe user [user id or username]`
* [x] `new user [username] [flags]`
* [x] `delete user [user id or username]`
* [x] `edit user [user id or username] [flags]`

### Users SSH Keys 

* [x] `get ssh-keys [flags]`
* [x] `new ssh-key [flags]`
* [x] `delete ssh-key [flags]`

### Members

* [x] `get members --from-group [flags]`
* [x] `get members --from-project [flags]`
* [x] `describe member [username] --from-group [flags]`
* [x] `describe member [username] --from-project [flags]`
* [x] `new member [username] --from-group [flags]`
* [x] `new member [username] --from-project [flags]`
* [x] `delete member [username] --from-group`
* [x] `delete member [username] --from-project`
* [x] `edit member [username] --from-group [flags]`
* [x] `edit member [username] --from-project [flags]`
* [x] `delete all-members --from-project`

## Contributing

Contributors are welcomed with love! Please read [CONTRIBUTING.md](./CONTRIBUTING.md) for the process for submitting pull requests to us.

