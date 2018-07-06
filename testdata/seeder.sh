#!/bin/bash -e

echo "=============================================="
echo "#              SEEDER SCRIPT                 #"
echo "=============================================="

# get access token
access_token=$(curl -X POST "${GITLAB_HTTP_URL}/oauth/token?grant_type=password&username=${GITLAB_USERNAME}&password=${GITLAB_PASSWORD}" | jq '.access_token' | tr -d '"')

# create user
echo "Creating a user"
user1_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/users?access_token=${access_token}&name=John+Doe&email=john.doe@gmail.com&username=john.doe&password=123qwe123&skip_confirmation=true" | jq '.id')
user2_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/users?access_token=${access_token}&name=John+Smith&email=john.smith@gmail.com&username=john.smith&password=123qwe123&skip_confirmation=true" | jq '.id')
user3_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/users?access_token=${access_token}&name=Matt+Hunter&email=matt.hunter@gmail.com&username=matt.hunter&password=123qwe123&skip_confirmation=true" | jq '.id')
user4_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/users?access_token=${access_token}&name=Amelia+Walsh&email=amelia.walsh@gmail.com&username=amelia.walsh&password=123qwe123&skip_confirmation=true" | jq '.id')
user5_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/users?access_token=${access_token}&name=Kevin+McLean&email=kevin.mclean@gmail.com&username=kevin.mclean&password=123qwe123&skip_confirmation=true" | jq '.id')
user6_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/users?access_token=${access_token}&name=Kylie+Morrison&email=kylie.morrison@gmail.com&username=kylie.morrison&password=123qwe123&skip_confirmation=true" | jq '.id')
user7_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/users?access_token=${access_token}&name=Rebecca+Gray&email=rebecca.gray@gmail.com&username=rebecca.gray&password=123qwe123&skip_confirmation=true" | jq '.id')
user8_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/users?access_token=${access_token}&name=Simon+Turner&email=simon.turner@gmail.com&username=simon.turner&password=123qwe123&skip_confirmation=true" | jq '.id')
user9_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/users?access_token=${access_token}&name=Olivia+White&email=olivia.white@gmail.com&username=olivia.white&password=123qwe123&skip_confirmation=true" | jq '.id')
user10_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/users?access_token=${access_token}&name=Frank+Watson&email=frank.watson@gmail.com&username=frank.watson&password=123qwe123&skip_confirmation=true" | jq '.id')
user11_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/users?access_token=${access_token}&name=Paul+Lyman&email=paul.lyman@gmail.com&username=paul.lyman&password=123qwe123&skip_confirmation=true" | jq '.id')

# create group
echo "Creating a group"
pgroup1_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/groups?access_token=${access_token}&name=Group1&path=Group1" | jq '.id')
pgroup2_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/groups?access_token=${access_token}&name=Group2&path=Group2" | jq '.id')

# add user to group
echo "Adding user to a group"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/groups/${pgroup1_id}/members?access_token=${access_token}&user_id=${user1_id}&access_level=30"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/groups/${pgroup1_id}/members?access_token=${access_token}&user_id=${user2_id}&access_level=40"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/groups/${pgroup1_id}/members?access_token=${access_token}&user_id=${user3_id}&access_level=50"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/groups/${pgroup2_id}/members?access_token=${access_token}&user_id=${user4_id}&access_level=30"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/groups/${pgroup2_id}/members?access_token=${access_token}&user_id=${user5_id}&access_level=40"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/groups/${pgroup2_id}/members?access_token=${access_token}&user_id=${user6_id}&access_level=50"

# create subgroup
echo "Creating a subgroup"
sgroup1_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/groups?access_token=${access_token}&name=SubGroup1&path=SubGroup1&parent_id=${pgroup1_id}" | jq '.id')
sgroup2_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/groups?access_token=${access_token}&name=SubGroup2&path=SubGroup2&parent_id=${pgroup1_id}" | jq '.id')
sgroup3_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/groups?access_token=${access_token}&name=SubGroup3&path=SubGroup3&parent_id=${pgroup2_id}" | jq '.id')
sgroup4_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/groups?access_token=${access_token}&name=SubGroup4&path=SubGroup4&parent_id=${pgroup2_id}" | jq '.id')

echo sleeping for 5 seconds..
sleep 5

# create group project
echo "Creating a project in group/subgroup"
groupproject1_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project1&namespace_id=${pgroup1_id}" | jq '.id')
groupproject2_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project2&namespace_id=${pgroup1_id}" | jq '.id')
groupproject3_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project3&namespace_id=${pgroup1_id}" | jq '.id')
groupproject4_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project4&namespace_id=${sgroup1_id}" | jq '.id')
groupproject5_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project5&namespace_id=${sgroup1_id}" | jq '.id')
groupproject6_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project6&namespace_id=${sgroup1_id}" | jq '.id')
groupproject7_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project7&namespace_id=${sgroup2_id}" | jq '.id')
groupproject8_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project8&namespace_id=${sgroup2_id}" | jq '.id')
groupproject9_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project9&namespace_id=${sgroup2_id}" | jq '.id')
groupproject10_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project10&namespace_id=${pgroup2_id}" | jq '.id')
groupproject11_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project11&namespace_id=${pgroup2_id}" | jq '.id')
groupproject12_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project12&namespace_id=${pgroup2_id}" | jq '.id')
groupproject13_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project13&namespace_id=${sgroup3_id}" | jq '.id')
groupproject14_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project14&namespace_id=${sgroup3_id}" | jq '.id')
groupproject15_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project15&namespace_id=${sgroup3_id}" | jq '.id')
groupproject16_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project16&namespace_id=${sgroup4_id}" | jq '.id')
groupproject17_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project17&namespace_id=${sgroup4_id}" | jq '.id')
groupproject18_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project18&namespace_id=${sgroup4_id}" | jq '.id')

# create user project
echo "Creating users project"
project1_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects/user/${user7_id}?access_token=${access_token}&name=Project19" | jq '.id')
project2_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects/user/${user8_id}?access_token=${access_token}&name=Project20" | jq '.id')
project3_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects/user/${user9_id}?access_token=${access_token}&name=Project21" | jq '.id')
project4_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects/user/${user10_id}?access_token=${access_token}&name=Project22" | jq '.id')
project5_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects/user/${user11_id}?access_token=${access_token}&name=Project23" | jq '.id')

# add user to project
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects/${groupproject1_id}/members?access_token=${access_token}&user_id=${user1_id}&access_level=30"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects/${groupproject1_id}/members?access_token=${access_token}&user_id=${user2_id}&access_level=40"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects/${groupproject1_id}/members?access_token=${access_token}&user_id=${user3_id}&access_level=50"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects/${groupproject1_id}/members?access_token=${access_token}&user_id=${user4_id}&access_level=30"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects/${groupproject2_id}/members?access_token=${access_token}&user_id=${user5_id}&access_level=40"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects/${groupproject2_id}/members?access_token=${access_token}&user_id=${user6_id}&access_level=50"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects/${groupproject2_id}/members?access_token=${access_token}&user_id=${user7_id}&access_level=40"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects/${groupproject2_id}/members?access_token=${access_token}&user_id=${user8_id}&access_level=40"

# create hooks for projects
echo "Creating hooks for projects"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects/${project1_id}/hooks?access_token=${access_token}&url=http%3A%2F%2Fwww.sample1.com%2F"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects/${project2_id}/hooks?access_token=${access_token}&url=http%3A%2F%2Fwww.sample2.com%2F"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects/${project3_id}/hooks?access_token=${access_token}&url=http%3A%2F%2Fwww.sample3.com%2F"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects/${project4_id}/hooks?access_token=${access_token}&url=http%3A%2F%2Fwww.sample4.com%2F"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects/${project5_id}/hooks?access_token=${access_token}&url=http%3A%2F%2Fwww.sample5.com%2F"
