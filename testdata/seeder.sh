#!/bin/bash -ex

echo "=============================================="
echo "#              SEEDER SCRIPT                 #"
echo "=============================================="

# create group
echo "Creating a group"
pgroup_id=$(curl --header "Private-Token: $GITLAB_PRIVATE_TOKEN" -X POST "http://localhost:10080/api/v4/groups?name=DevOps&path=DevOps" | jq '.id')

# create subgroup
echo "Creating a subgroup"
sgroup_id=$(curl --header "Private-Token: $GITLAB_PRIVATE_TOKEN" -X POST "http://localhost:10080/api/v4/groups?name=SecOps&path=SecOps&parent_id=${pgroup_id}" | jq '.id')

# display groups in gitlab instance
curl --header "Private-Token: $GITLAB_PRIVATE_TOKEN" "http://localhost:10080/api/v4/groups" | jq '.'

