![](https://talks.golang.org/2012/waza/gophercomplex1.jpg)

[![Build Status](https://travis-ci.org/devopsctl/gitlabctl.svg?branch=master)](https://travis-ci.org/devopsctl/gitlabctl)
[![codecov](https://codecov.io/gh/devopsctl/gitlabctl/branch/master/graph/badge.svg)](https://codecov.io/gh/devopsctl/gitlabctl)
[![Go Report Card](https://goreportcard.com/badge/github.com/devopsctl/gitlabctl)](https://goreportcard.com/report/github.com/devopsctl/gitlabctl)
[![Waffle board](https://badge.waffle.io/devopsctl/gitlabctl.png?columns=all)](https://waffle.io/devopsctl/gitlabctl)
# Gitlab Control (Work in Progress)

<!-- vim-markdown-toc GFM -->

* [Completed Commands Documentation](#completed-commands-documentation)
* [Development Setup](#development-setup)
	* [Requirements](#requirements)
	* [Environment Variables and Test Data](#environment-variables-and-test-data)
	* [Refresh Test Data](#refresh-test-data)
	* [Branching](#branching)
	* [Issue Tracker](#issue-tracker)
	* [Before Pushing your Commit](#before-pushing-your-commit)
* [Custom Packages](#custom-packages)
* [Test Driven Development](#test-driven-development)
* [How the Commands Authenticate](#how-the-commands-authenticate)
* [Commands Pattern](#commands-pattern)
* [Group Commands](#group-commands)
	* [Get details of a group - `describe group`](#get-details-of-a-group---describe-group)
	* [Get all groups - `get groups`](#get-all-groups---get-groups)
	* [List all the subgroups of a group - `get subgroups`](#list-all-the-subgroups-of-a-group---get-subgroups)
	* [Get all the projects of a group - `get group-projects`](#get-all-the-projects-of-a-group---get-group-projects)
	* [Remove a group - `remove group`](#remove-a-group---remove-group)
	* [Add a new group - `new group`](#add-a-new-group---new-group)
	* [Edit a group - `edit group`](#edit-a-group---edit-group)
	* [Get details of a member - `describe group-member`](#get-details-of-a-member---describe-group-member)
	* [List all members of a group - `get group-members`](#list-all-members-of-a-group---get-group-members)
	* [Remove a group member - `remove group-member`](#remove-a-group-member---remove-group-member)
	* [Add a group member - `new member`](#add-a-group-member---new-member)
	* [Remove all group members - `remove all-members`](#remove-all-group-members---remove-all-members)
* [Project Commands](#project-commands)
	* [Get details of a project - `describe project`](#get-details-of-a-project---describe-project)
	* [List all projects - `get projects`](#list-all-projects---get-projects)
	* [Delete a project - `remove project`](#delete-a-project---remove-project)
	* [Create a project - `new project`](#create-a-project---new-project)
	* [Edit a project - `edit project`](#edit-a-project---edit-project)
	* [Get details of a project member - `describe project-member`](#get-details-of-a-project-member---describe-project-member)
	* [Get all members of a project - `get project-members`](#get-all-members-of-a-project---get-project-members)
	* [Remove a project member - `remove project-member`](#remove-a-project-member---remove-project-member)
	* [Add a project member - `new project-member`](#add-a-project-member---new-project-member)
	* [Remove all project members - `remove all-project-members`](#remove-all-project-members---remove-all-project-members)
	* [Describe a project hook - `describe project-hook`](#describe-a-project-hook---describe-project-hook)
	* [List all hooks of a project - `get project-hooks`](#list-all-hooks-of-a-project---get-project-hooks)
	* [Add a project hook - `new project-hook`](#add-a-project-hook---new-project-hook)
	* [Edit a project hook - `edit project-hook`](#edit-a-project-hook---edit-project-hook)
	* [Delete a project hook - `remove project-hook`](#delete-a-project-hook---remove-project-hook)
	* [Delete all hooks in a project - `remove all-project-hooks`](#delete-all-hooks-in-a-project---remove-all-project-hooks)

<!-- vim-markdown-toc -->

Our goal is to create a gitlab cli written in Go that is simple to use and
easy to maintain. The code must be simple and flags must be patterned with the
gitlab client package https://godoc.org/github.com/xanzy/go-gitlab.

## Completed Commands Documentation

Please read the [commands documentation](./docs/gitlabctl.md). 
This is auto-generated using `gitlabctl gendocs`.

## Development Setup

### Requirements

* Install the latest stable version of go (1.10.1 as of this writing).
* Install docker.
* Install local gitlab instance using the [docker-compose](./docker-compose.yml) 
file. Run `docker-compose up -d`.

### Environment Variables and Test Data

To be in the same local environment setup, all developers including Travis CI 
must have the same Gitlab credentials.

To set the environment variables as credentials:

```bash
source testdata/credentials.sh # set the environment variables
env | grep GITLAB # check GITLAB variables
```

To seed your local gitlab instance:

```bash
source testdata/credentials.sh
testdata/seeder.sh 
```

### Refresh Test Data

```bash
docker-compose down # delete all containers
docker volume prune # delete existing docker volumes
docker-compose up # wait for gitlab to be up and running
source testdata/credentials.sh
testdata/seeder.sh
```


### Branching 

* Branch name should be in the format of `{{issueNumber}}-{{shortTaskName}}`. 
Example: `19-add-group-get-cmd`.
* Always run a rebase pull when master or remote branch is updated. 
Use `git pull --rebase origin master` or `git pull --rebase origin branchName` 
as much as possible.

### Issue Tracker

* Ensure that you are working on an [Issue](https://github.com/devopsctl/gitlabctl/issues) 
and is visible in [Waffle Issue Tracker](https://waffle.io/devopsctl/gitlabctl).
* Ensure to create a branch for your Issue.

### Before Pushing your Commit

* Run `gometalinter -v ./...`. Ask for help with issues found that can't be solved.
* Run `go test -v ./...`. Everything must pass the test. There will be an issue
with private token testing, as this is unique on each developer installation.
* Don't hesitate to ask questions! [Gophers Slack](https://gophers.slack.com) 
community may be able to answer your questions.
* Ensure to have a Pull Request for your branch before asking for Code Review.
* Ask for Code Review if your Issue is ready for Merging.
* If you are not changing code (e.g: updating docs or adding test data), 
use __ci skip__ in commit message to [Skip TravisCI build](https://docs.travis-ci.com/user/customizing-the-build/#Skipping-a-build)

## Custom Packages

* Gitlab api client - https://godoc.org/github.com/xanzy/go-gitlab 
* Commandline flags - https://github.com/spf13/cobra 

## Test Driven Development

This project may grow big in the future so the definition of done for every 
commands should be tested against a local gitlab instance. 

## How the Commands Authenticate

Authenticate using environment variables.

* Basic authentication - `GITLAB_USERNAME`, `GITLAB_PASSWORD` and `GITLAB_HTTP_URL`
* Private token authentication - `GITLAB_PRIVATE_TOKEN` and `GITLAB_API_HTTP_URL`
* OAuth2 token authentication - `GITLAB_OAUTH_TOKEN` and `GITLAB_API_HTTP_URL`

## Commands Pattern

The command chain format is inspired from `kubectl` or `oc` Verb -> Subject -> Flags.

* get/fetch
	* groups
	* subgroups
	* group-projects
	* group-members
	* projects
	* project-members
	* project-hooks 
	* users
* desc/describe
	* group
	* group-member
	* project
	* project-member
	* user
* new
	* group
	* group-member
	* project
	* project-member
	* user
	* project-hook
* edit/patch
	* group
	* project
	* user
	* project-hook
* remove
	* group
	* group-member
	* project
	* project-member
	* user
	* project-hook

## Group Commands

API doc: https://godoc.org/github.com/xanzy/go-gitlab#GroupService

### Get details of a group - `describe group`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#GroupsService.GetGroup

| Flag  | Type   | Description                                                    | Required? | Default |
| :---- | :---   | :----------                                                    | :-------- | :------ |
| path  | string | The group name, id or full the path including the parent group | yes       |         |
| json  | bool   | Print the command output to json                               | no        | false   |

### Get all groups - `get groups`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#GroupsService.ListGroups

| Flag          | Type   | Description                                                                                        | Required? | Default |
| :----         | :---   | :----------                                                                                        | :-------- | :------ |
| all-available | bool   | Show all the groups you have access to (defaults to false for authenticated users, true for admin) | no        |         |
| order-by      | string | Order groups by name or path. Default is name                                                      | no        |         |
| owned         | bool   | Limit to groups owned by the current user                                                          | no        |         |
| search        | string | Return the list of authorized groups matching the search criteria                                  | no        |         |
| sort          | string | Order groups in asc or desc order. Default is asc                                                  | no        |         |
| statistics    | bool   | Include group statistics (admins only)                                                             | no        |         |
| json          | bool   | Print the command output to json                                                                   | no        | false   |

### List all the subgroups of a group - `get subgroups`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#GroupsService.ListSubgroups

| Flag          | Type   | Description                                                                                        | Required? | Default |
| :----         | :---   | :----------                                                                                        | :-------- | :------ |
| path          | string | The group name, id or full the path including the parent group                                     | yes       |
| all-available | bool   | Show all the groups you have access to (defaults to false for authenticated users, true for admin) | no        |         |
| order-by      | string | Order groups by name or path. Default is name                                                      | no        |         |
| owned         | bool   | Limit to groups owned by the current user                                                          | no        |         |
| search        | string | Return the list of authorized groups matching the search criteria                                  | no        |         |
| sort          | string | Order groups in asc or desc order. Default is asc                                                  | no        |         |
| statistics    | bool   | Include group statistics (admins only)                                                             | no        |         |
| json          | bool   | Print the command output to json                                                                   | no        | false   |

### Get all the projects of a group - `get group-projects`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#GroupsService.ListGroupProjects

| Flag  | Type   | Description                                                    | Required? | Default |
| :---- | :---   | :----------                                                    | :-------- | :------ |
| path  | string | The group name, id or full the path including the parent group | yes       |
| json  | bool   | Print the command output to json                               | no        | false   |

### Remove a group - `remove group`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#GroupsService.DeleteGroup

| Flag  | Type   | Description                                                    | Required? |
| :---- | :---   | :----------                                                    | :-------- |
| path  | string | The group name, id or full the path including the parent group | yes       |

### Add a new group - `new group`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#GroupsService.CreateGroup

| Flag                   | Type   | Description                                               | Required? | Default |
| :----                  | :---   | :----------                                               | :-------- | :------ |
| name                   | string | The group name                                            | yes       |         |
| namespace              | string | The parent group id or group path if creating a subgroup. | no        |         |
| visibility             | string | public, internal or private                               | no        |
| lfs-enabled            | bool   | Enable LFS                                                | no        |
| request-access-enabled | bool   | Enable Request Access                                     | no        |
| json                   | bool   | Print the command output to json                          | no        | false   |

Custom flag validation:

* If optional or non-required flags are not set, do not use or ignore the default value.

### Edit a group - `edit group`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#GroupsService.UpdateGroup

| Flag                   | Type   | Description                                                    | Required? | Default |
| :----                  | :---   | :----------                                                    | :-------- | :------ |
| path                   | string | The group name, id or full the path including the parent group | yes       |         |
| visibility             | string | public, internal or private                                    | no        |         |
| lfs-enabled            | bool   | Enable LFS                                                     | no        |         |
| request-access-enabled | bool   | Enable Request Access                                          | no        |         |
| json                   | bool   | Print the command output to json                               | no        | false   |

Custom flag validation:

* If optional or non-required flags are not set, do not use or ignore the default value.
* Command requires at least 1 optional flag to be set.

API doc: https://godoc.org/github.com/xanzy/go-gitlab#GroupMembersService

### Get details of a member - `describe group-member`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#GroupMembersService.GetGroupMember

| Flag     | Type   | Description                                                    | Required? | Default |
| :----    | :---   | :----------                                                    | :-------- | :------ |
| path     | string | The group name, id or full the path including the parent group | yes       |         |
| username | string | username to describe                                           | yes       |         |
| json     | bool   | Print the command output to json                               | no        | false   |

### List all members of a group - `get group-members`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#GroupsService.ListGroupMembers

| Flag  | Type   | Description                                                    | Required? | Default |
| :---- | :---   | :----------                                                    | :-------- | :------ |
| path  | string | The group name, id or full the path including the parent group | yes       |         |
| json  | bool   | Print the command output to json                               | no        | false   |

### Remove a group member - `remove group-member`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#GroupMembersService.RemoveGroupMember

| Flag     | Type   | Description                                                    | Required? | Default |
| :----    | :---   | :----------                                                    | :-------- | :------ |
| path     | string | The group name, id or full the path including the parent group | yes       |         |
| username | string | username to remove                                             | yes       |         |

### Add a group member - `new member`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#GroupMembersService.AddGroupMember

| Flag         | Type   | Description                                                                                                           | Required? | Default |
| :----        | :---   | :----------                                                                                                           | :-------- | :------ |
| path         | string | The group name, id or full the path including the parent group                                                        | yes       |         |
| username     | string | User's username                                                                                                       | yes       |         |
| access-level | int    | Member group access level (0, 10, 20, 30, 40, 50). Reference: https://docs.gitlab.com/ce/permissions/permissions.html | no        | 10      |
| json         | bool   | Print the command output to json                                                                                      | no        | false   |

### Remove all group members - `remove all-members`

A wrapper of listing all group members and deleting them all.

| Flag         | Type   | Description                                                                                                           | Required? | Default |
| :----        | :---   | :----------                                                                                                           | :-------- | :------ |
| path         | string | The group name, id or full the path including the parent group                                                        | yes       |         |
| username     | string | User's username                                                                                                       | yes       |         |

## Project Commands

API doc: https://godoc.org/github.com/xanzy/go-gitlab#ProjectsService

### Get details of a project - `describe project`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#ProjectsService.GetProject

| Flag  | Type   | Description                                                                          | Required? | Default |
| :---- | :---   | :----------                                                                          | :-------- | :------ |
| path  | string | The project name, id or full the path including the parent group - (path/to/project) | yes       |         |
| json  | bool   | Print the command output to json                                                     | no        | false   |

### List all projects - `get projects`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#ProjectsService.ListProjects

| Flag  | Type   | Description                                                                          | Required? | Default |
| :---- | :---   | :----------                                                                          | :-------- | :------ |
| path  | string | The project name, id or full the path including the parent group - (path/to/project) | yes       |         |
| json  | bool   | Print the command output to json                                                     | no        | false   |

### Delete a project - `remove project`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#ProjectsService.DeleteProject

| Flag  | Type   | Description                                                                          | Required? | Default |
| :---- | :---   | :----------                                                                          | :-------- | :------ |
| path  | string | The project name, id or full the path including the parent group - (path/to/project) | yes       |         |

### Create a project - `new project`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#ProjectsService.CreateProject

| Flag                                        | Type   | Description                                                                                     | Required? | Default |
| :----                                       | :---   | :----------                                                                                     | :-------- | :------ |
| name                                        | string | The Project name                                                                                | yes       |         |
| namespace                                   | string | The parent group id or group path if creating a subgroup.                                       | no        |         |
| description                                 | string | Project description                                                                             | no        |         |
| issues-enabled                              | bool   | Enable issues                                                                                   | no        |
| merge-requests-enabled                      | bool   | Enable merge requests                                                                           | no        |
| jobs-enabled                                | bool   | Enable jobs                                                                                     | no        |
| wiki-enabled                                | bool   | Enable wikis                                                                                    | no        |
| snippets-enabled                            | bool   | Enable snippets                                                                                 | no        |
| resolve-outdated-diff-discussions           | bool   | Resolve outdated diff discussions                                                               | no        |
| container-registry-enabled                  | bool   | Enable container registry                                                                       | no        |
| shared-runners-enabled                      | bool   | Enable shared runners                                                                           | no        |
| visibility                                  | string | Project visibility (public, internal, private)                                                  | no        | public  |
| public-jobs                                 | bool   | If true, jobs can be viewed by non-project-members                                              | no        |
| only-allow-merge-if-pipeline-succeeds       | bool   | Set whether merge requests can only be merged with successful jobs                              | no        |
| only-allow-merge-if-discussion-are-resolved | bool   | Set whether merge requests can only be merged when all the discussions are resolved             | no        |
| merge-method                                | string | Set the merge method used                                                                       | no        |
| lfs-enabled                                 | bool   | Enable lfs                                                                                      | no        |
| request-access-enabled                      | bool   | Allow users to request member access                                                            | no        |
| tag-list                                    | string | The list of tags for a project; put array of tags, that should be finally assigned to a Project | no        |
| printing-merge-request-link-enabled         | bool   | Show link to create/view merge request when pushing from the command line                       | no        |
| ci-config-path                              | string | The path to ci config file                                                                      | no        |
| json                                        | bool   | Print the command output to json                                                                | no        | false   |

Custom flag validation:

* If optional or non-required flags are not set, do not use or ignore the default value.

### Edit a project - `edit project`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#ProjectsService.EditProject

| Flag                                        | Type   | Description                                                                                     | Required? | Default |
| :----                                       | :---   | :----------                                                                                     | :-------- | :------ |
| path                                        | string | The project name, id or full the path including the parent group - (path/to/project)            | yes       |         |
| description                                 | string | Project description                                                                             | no        |         |
| issues-enabled                              | bool   | Enable issues                                                                                   | no        |
| merge-requests-enabled                      | bool   | Enable merge requests                                                                           | no        |
| jobs-enabled                                | bool   | Enable jobs                                                                                     | no        |
| wiki-enabled                                | bool   | Enable wikis                                                                                    | no        |
| snippets-enabled                            | bool   | Enable snippets                                                                                 | no        |
| resolve-outdated-diff-discussions           | bool   | Resolve outdated diff discussions                                                               | no        |
| container-registry-enabled                  | bool   | Enable container registry                                                                       | no        |
| shared-runners-enabled                      | bool   | Enable shared runners                                                                           | no        |
| visibility                                  | string | Project visibility (public, internal, private)                                                  | no        | public  |
| public-jobs                                 | bool   | If true, jobs can be viewed by non-project-members                                              | no        |
| only-allow-merge-if-pipeline-succeeds       | bool   | Set whether merge requests can only be merged with successful jobs                              | no        |
| only-allow-merge-if-discussion-are-resolved | bool   | Set whether merge requests can only be merged when all the discussions are resolved             | no        |
| merge-method                                | string | Set the merge method used                                                                       | no        |
| lfs-enabled                                 | bool   | Enable lfs                                                                                      | no        |
| request-access-enabled                      | bool   | Allow users to request member access                                                            | no        |
| tag-list                                    | string | The list of tags for a project; put array of tags, that should be finally assigned to a Project | no        |
| printing-merge-request-link-enabled         | bool   | Show link to create/view merge request when pushing from the command line                       | no        |
| ci-config-path                              | string | The path to ci config file                                                                      | no        |
| json                                        | bool   | Print the command output to json                                                                | no        | false   |

Custom flag validation:

* If optional or non-required flags are not set, do not use or ignore the default value.
* Command requires at least 1 optional flag to be set.

### Get details of a project member - `describe project-member`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#ProjectMembersService.GetProjectMember

| Flag     | Type   | Description                                                                          | Required? | Default |
| :----    | :---   | :----------                                                                          | :-------- | :------ |
| path     | string | The project name, id or full the path including the parent group - (path/to/project) | yes       |         |
| username | string | Member username                                                                      | yes       |         |
| json     | bool   | Print the command output to json                                                     | no        | false   |

### Get all members of a project - `get project-members`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#ProjectMembersService.ListProjectMembers

| Flag  | Type   | Description                                                                          | Required? | Default |
| :---- | :---   | :----------                                                                          | :-------- | :------ |
| path  | string | The project name, id or full the path including the parent group - (path/to/project) | yes       |         |
| json  | bool   | Print the command output to json                                                     | no        | false   |

### Remove a project member - `remove project-member`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#ProjectMembersService.AddProjectMember

| Flag     | Type   | Description                                                                          | Required? | Default |
| :----    | :---   | :----------                                                                          | :-------- | :------ |
| path     | string | The project name, id or full the path including the parent group - (path/to/project) | yes       |         |
| username | string | Member username                                                                      | yes       |         |

### Add a project member - `new project-member`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#ProjectMembersService.AddProjectMember

| Flag         | Type   | Description                                                                                                           | Required? | Default |
| :----        | :---   | :----------                                                                                                           | :-------- | :------ |
| path         | string | The project name, id or full the path including the parent group - (path/to/project)                                  | yes       |         |
| username     | string | Member username                                                                                                       | yes       |         |
| access-level | int    | Member group access level (0, 10, 20, 30, 40, 50). Reference: https://docs.gitlab.com/ce/permissions/permissions.html | no        | 10      |
| json         | bool   | Print the command output to json                                                                                      | no        | false   |

### Remove all project members - `remove all-project-members`

A wrapper of listing all project members and removing them all.

| Flag     | Type   | Description                                                                          | Required? | Default |
| :----    | :---   | :----------                                                                          | :-------- | :------ |
| path     | string | The project name, id or full the path including the parent group - (path/to/project) | yes       |         |

### Describe a project hook - `describe project-hook`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#ProjectsService.GetProjectHook

| Flag  | Type   | Description                                                                          | Required? | Default |
| :---- | :---   | :----------                                                                          | :-------- | :------ |
| path  | string | The project name, id or full the path including the parent group - (path/to/project) | yes       |         |
| id    | int    | Hook ID                                                                              | yes       |         |
| json  | bool   | Print the command output to json                                                     | no        | false   |

### List all hooks of a project - `get project-hooks`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#ProjectsService.ListProjectHooks

| Flag  | Type   | Description                                                                          | Required? | Default |
| :---- | :---   | :----------                                                                          | :-------- | :------ |
| path  | string | The project name, id or full the path including the parent group - (path/to/project) | yes       |         |
| json  | bool   | Print the command output to json                                                     | no        | false   |

### Add a project hook - `new project-hook`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#ProjectsService.AddProjectHook

| Flag                       | Type   | Description                                                                          | Required? | Default |
| :----                      | :---   | :----------                                                                          | :-------- | :------ |
| path                       | string | The project name, id or full the path including the parent group - (path/to/project) | yes       |         |
| url                        | string | The hook URL                                                                         | yes       |         |
| push-events                | bool   | Trigger hook on push events                                                          |           |         |
| issues-events              | bool   | Trigger hook on issues events                                                        |           |         |
| confidential-issues-events | bool   | Trigger hook on confidential issues events                                           |           |         |
| merge-requests-events      | bool   | Trigger hook on merge requests events                                                |           |         |
| tag-push-events            | bool   | Trigger hook on tag push events                                                      |           |         |
| note-events                | bool   | Trigger hook on note events                                                          |           |         |
| job-events                 | bool   | Trigger hook on wiki events                                                          |           |         |
| pipeline-events            | bool   | Trigger hook on pipeline events                                                      |           |         |
| wiki-page-events           | bool   | Trigger hook on wiki events                                                          |           |         |
| enable-ssl-verification    | bool   | Do SSL verification when triggering the hook                                         |           |         |
| token                      | string | Secret token to validate received payloads                                           |           |         |
| json                       | bool   | Print the command output to json                                                     | no        | false   |

Custom flag validation:

* If optional or non-required flags are not set, do not use or ignore the default value.

### Edit a project hook - `edit project-hook`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#ProjectsService.EditProjectHook

| Flag                       | Type   | Description                                                                          | Required? | Default |
| :----                      | :---   | :----------                                                                          | :-------- | :------ |
| id                         | int    | Hook ID                                                                              | yes       |         |
| path                       | string | The project name, id or full the path including the parent group - (path/to/project) | yes       |         |
| url                        | string | The hook URL                                                                         | yes       |         |
| push-events                | bool   | Trigger hook on push events                                                          |           |         |
| issues-events              | bool   | Trigger hook on issues events                                                        |           |         |
| confidential-issues-events | bool   | Trigger hook on confidential issues events                                           |           |         |
| merge-requests-events      | bool   | Trigger hook on merge requests events                                                |           |         |
| tag-push-events            | bool   | Trigger hook on tag push events                                                      |           |         |
| note-events                | bool   | Trigger hook on note events                                                          |           |         |
| job-events                 | bool   | Trigger hook on wiki events                                                          |           |         |
| pipeline-events            | bool   | Trigger hook on pipeline events                                                      |           |         |
| wiki-page-events           | bool   | Trigger hook on wiki events                                                          |           |         |
| enable-ssl-verification    | bool   | Do SSL verification when triggering the hook                                         |           |         |
| token                      | string | Secret token to validate received payloads                                           |           |         |

Custom flag validation:

* If optional or non-required flags are not set, do not use or ignore the default value.
* Command requires at least 1 optional flag to be set.

### Delete a project hook - `remove project-hook`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#ProjectsService.DeleteProjectHook

| Flag  | Type   | Description                                                                          | Required? | Default |
| :---- | :---   | :----------                                                                          | :-------- | :------ |
| id    | int    | Hook ID                                                                              | yes       |         |
| path  | string | The project name, id or full the path including the parent group - (path/to/project) | yes       |         |

### Delete all hooks in a project - `remove all-project-hooks`

A wrapper of listing all project hooks and deleting all of them.

| Flag  | Type   | Description                                                                          | Required? | Default |
| :---- | :---   | :----------                                                                          | :-------- | :------ |
| path  | string | The project name, id or full the path including the parent group - (path/to/project) | yes       |         |


