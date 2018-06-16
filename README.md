[![Build Status](https://travis-ci.org/devopsctl/gitlabctl.svg?branch=master)](https://travis-ci.org/devopsctl/gitlabctl)
[![codecov](https://codecov.io/gh/devopsctl/gitlabctl/branch/master/graph/badge.svg)](https://codecov.io/gh/devopsctl/gitlabctl)
[![Go Report Card](https://goreportcard.com/badge/github.com/devopsctl/gitlabctl)](https://goreportcard.com/report/github.com/devopsctl/gitlabctl)
[![Waffle board](https://badge.waffle.io/devopsctl/gitlabctl.png?columns=all)](https://waffle.io/devopsctl/gitlabctl)

# [`gitlabctl`](https://devopsctl.github.io/gitlabctl/)

Be a rockstar and efficiently manage your team's gitlab.org or [self-hosted Gitlab](https://about.gitlab.com/installation/) projects, groups, users and other resources.

## Getting Started

Please see [Github site](https://devopsctl.github.io/gitlabctl/) for a nice documentation of this project.

## Installing

Download the binary from [releases](https://github.com/devopsctl/gitlabctl/releases) or use Go Get, `go get github.com/devopsctl/gitlabctl`.

Enable auto complete for __bash__ or __zsh__ shell. ❤️

```bash
# follow the instructions from the command output
gitlabctl completion -h
```

## Contributing

Contributors are welcomed with love! Please read [CONTRIBUTING.md](./CONTRIBUTION.md) for the process for submitting pull requests to us.

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

