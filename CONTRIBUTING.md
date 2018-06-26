# Developers Contribution Guide

<!-- vim-markdown-toc GFM -->

* [Development Setup](#development-setup)
  * [Requirements](#requirements)
  * [Environment Variables and Test Data](#environment-variables-and-test-data)
  * [Refresh Test Data](#refresh-test-data)
  * [Issue Tracker](#issue-tracker)
  * [Branching](#branching)
  * [Code Generator](#code-generator)
  * [Good Commit Message Examples](#good-commit-message-examples)
  * [Before Pushing your Commit](#before-pushing-your-commit)
  * [Tips and Tricks](#tips-and-tricks)
* [Custom Packages we are using](#custom-packages-we-are-using)
* [Test Driven Development](#test-driven-development)
* [How the Commands Authenticate](#how-the-commands-authenticate)
* [Commands Pattern](#commands-pattern)
* [Releasing to Github](#releasing-to-github)

<!-- vim-markdown-toc -->

Our goal is to create a gitlab cli written in Go that is simple to use and
easy to maintain. The code must be simple and flags must be patterned with the
gitlab client package https://godoc.org/github.com/xanzy/go-gitlab.

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

### Issue Tracker

[![Throughput Graph](http://graphs.waffle.io/devopsctl/gitlabctl/throughput.svg)](https://waffle.io/devopsctl/gitlabctl/metrics)

* Ensure that you are working on an [Issue](https://github.com/devopsctl/gitlabctl/issues) 
and is visible in [Waffle Issue Tracker](https://waffle.io/devopsctl/gitlabctl).
* Ensure to create a branch for your Issue.

### Branching 

* Branch name should be in the format of `{{issueNumber}}-{{shortTaskName}}`. 
Example: `19-add-group-get-cmd`.
* Always run a rebase pull when master or remote branch is updated. 
Use `git pull --rebase origin master` or `git pull --rebase origin branchName` 
as much as possible.

### Code Generator

Once you get the basics of creating a gitlabctl command, you will find out that most code are repeatable. 
This is where the `gitlabctl gencode` commands can be of help.

Please use it in for your own convenience. Feedback about this command is also appreciated!

### Good Commit Message Examples

* `new_group: fix known issue when printing output` - when patching an existing command
* `Add new command get labels` - when adding a new command

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

### Tips and Tricks

Use `grm cmd subcmd` for ease of testing a command.
* How: `alias grm='go run main.go'`
* Example Usage: `grm get groups`

Use `gt` for ease of running `go test`.

* How: `alias gt='go test -v ./...'`
* Example Usage: `gt -run NewGroup`


## Custom Packages we are using

* Gitlab api client - https://godoc.org/github.com/xanzy/go-gitlab 
* Commandline flags - https://github.com/spf13/cobra 
* Table writer - https://github.com/olekukonko/tablewriter

## Test Driven Development

This project may grow big in the future so the definition of done for every 
commands should be tested against a local gitlab instance. 

## How the Commands Authenticate

Authenticate using environment variables.

* Basic authentication - `GITLAB_USERNAME`, `GITLAB_PASSWORD` and `GITLAB_HTTP_URL`
* Private token authentication - `GITLAB_PRIVATE_TOKEN` and `GITLAB_API_HTTP_URL`
* OAuth2 token authentication - `GITLAB_OAUTH_TOKEN` and `GITLAB_API_HTTP_URL`

Authenticate using `gitlabctl login` command.

## Commands Pattern

The command chain format is inspired from `kubectl` or `oc` Verb -> Subject -> Flags.

## Releasing to Github

* Build and create the version tag

```bash
export VERSION=nextVersion
make test
make all # create biniaries in bin/
git tag -a "v${VERSION}" -m "release description"
git push origin v${VERSION}
```

* Create a github release page
* Upload binaries to github release page
