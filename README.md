[![Build Status](https://travis-ci.org/devopsctl/gitlabctl.svg?branch=master)](https://travis-ci.org/devopsctl/gitlabctl)
[![codecov](https://codecov.io/gh/devopsctl/gitlabctl/branch/master/graph/badge.svg)](https://codecov.io/gh/devopsctl/gitlabctl)
[![Go Report Card](https://goreportcard.com/badge/github.com/devopsctl/gitlabctl)](https://goreportcard.com/report/github.com/devopsctl/gitlabctl)
[![Waffle board](https://badge.waffle.io/devopsctl/gitlabctl.png?columns=all)](https://waffle.io/devopsctl/gitlabctl)

# [`gitlabctl`](https://devopsctl.github.io/gitlabctl/)

Be a rockstar and efficiently manage your team's gitlab.org or [self-hosted Gitlab](https://about.gitlab.com/installation/) projects, groups, users and other resources.

Tested on __Gitlab 10.7.4__

## Getting Started

Please see [Commands Manual](https://devopsctl.github.io/gitlabctl/) for a nice documentation of this project.

## Installing

Get the download link of your preferred platform binary from [RELEASES](https://github.com/devopsctl/gitlabctl/releases) and ensure that it is downloaded in one of your `$PATH` directories.

```bash
# This is an example manual installation
wget https://github.com/devopsctl/gitlabctl/releases/download/v0.0.1/gitlabctl-darwin-amd64
mv gitlabctl-darwin-amd64 /usr/local/bin/gitlabctl && chmod +x /usr/local/bin/gitlabctl

# Test the command
gitlabctl --version
```

For __Golang__ developers, you can get the latest binary using `go get github.com/devopsctl/gitlabctl`.

Enable auto complete for __bash__ or __zsh__ shell. ❤️

```bash
# follow the instructions from the command output
gitlabctl completion -h
```

## Quickstart

### Using the help commands

```bash
# get commands is used to list resources
gitlabctl get [subcommand] -h

# new commands is used to create a single resource
gitlabctl new [subcommand] -h

# edit commands is used to update a single resource
gitlabctl edit [subcommand] -h

# delete commands is used to delete a single resource
gitlabctl delete [subcommand] -h

# describe command is used to describe a single resource
gitlabctl describe [subcommand]  -h
```

### Usage Examples

Fetching resources with using `--output, -o` formatter flag.

```bash
>> gitlabctl get groups
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

```bash
>> gitlabctl get groups -o json
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

```bash
>> gitlabctl get groups -o yaml
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

Creating resources.

```bash
# create a group
>> gitlabctl new group devopsctl

# create a project under devopsctl group
>> gitlabctl new project gitlab-cli --namespace=devopsctl

# create a new user with username john
>> gitlabctl new user john --name="John Smith" --password="john123456" --email=john@example.com --reset-password 
```

## Contributing

Contributors are welcomed with love! Please read [CONTRIBUTING.md](./CONTRIBUTING.md) for the process for submitting pull requests to us.

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

### Project Branch

* [ ] `get branch [project id or project path] [flags]`
* [ ] `describe branch [project id or project path] [flags]`
* [x] `new branch [project id or project path] [flags]`
* [ ] `edit branch [project id or project path] [--protect] [flags]`
* [ ] `edit branch [project id or project path] [--unprotect] [flags]`

### Project Hooks

* [x] `get project-hooks [project id or project path]`
* [x] `new project-hook [project id or project path] [flags]`
* [x] `edit project-hook [project id or project path] [flags]`
* [ ] `delete project-hook [hook id]`

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
* [ ] `remove all-members --from-group`
* [ ] `remove all-members --from-project`

