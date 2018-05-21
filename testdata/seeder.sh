#!/bin/bash -ex

echo "=============================================="
echo "#              SEEDER SCRIPT                 #"
echo "=============================================="

# generated by get_token.sh
personal_access_token=$(cat token.txt)

# create group
echo "Creating a group"
pgroup_id=$(curl --header "Private-Token: $personal_access_token" -X POST "http://localhost:10080/api/v4/groups?name=DevOps&path=DevOps" | jq '.id')

# create subgroup
echo "Creating a subgroup"
sgroup_id=$(curl --header "Private-Token: $personal_access_token" -X POST "http://localhost:10080/api/v4/groups?name=SecOps&path=SecOps&parent_id=${pgroup_id}" | jq '.id')

# display groups in gitlab instance
curl --header "Private-Token: $personal_access_token" "http://localhost:10080/api/v4/groups" | jq '.'

