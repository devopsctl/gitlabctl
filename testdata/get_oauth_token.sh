#!/bin/bash -x

curl -X POST "${GITLAB_HTTP_URL}/oauth/token?grant_type=password&username=${GITLAB_USERNAME}&password=${GITLAB_PASSWORD}" | jq '.access_token' | tr -d '"'

