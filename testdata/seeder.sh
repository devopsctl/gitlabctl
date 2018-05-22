#!/bin/bash -ex

echo "=============================================="
echo "#              SEEDER SCRIPT                 #"
echo "=============================================="

#set gitlab base url
base_url="${1}/"

# get access token
access_token=$(curl -X POST $base_url"oauth/token?grant_type=password&username=${2}&password=${3}" | jq '.access_token' | tr -d '"')

# create user
echo "Creating a user"
curl -X POST $base_url"api/v4/users?access_token=$access_token&name=John+Doe&email=john.doe@gmail.com&username=john.doe&password=123qwe123&skip_confirmation=true"

# create group
echo "Creating a group"
pgroup_id=$(curl -X POST $base_url"api/v4/groups?access_token=$access_token&name=DevOps&path=DevOps" | jq '.id')

# create group project
echo "Creating a group project"
curl -X POST $base_url"api/v4/projects?name=DevOpsProject&namespace_id=$pgroup_id&access_token=$access_token"

# create subgroup
echo "Creating a subgroup"
sgroup_id=$(curl -X POST $base_url"api/v4/groups?access_token=$access_token&name=SecOps&path=SecOps&parent_id=$pgroup_id" | jq '.id')

# create subgroup project
echo "Creating a subgroup project"
curl -X POST $base_url"api/v4/projects?access_token=$access_token&name=SecOpsProject&namespace_id=$sgroup_id"

