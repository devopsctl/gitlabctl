#!/bin/bash -ex

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
pgroup1_sgroup1_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/groups?access_token=${access_token}&name=ParentGroup1SubGroup1&path=ParentGroup1SubGroup1&parent_id=${pgroup1_id}" | jq '.id')
pgroup1_sgroup2_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/groups?access_token=${access_token}&name=ParentGroup1SubGroup2&path=ParentGroup1SubGroup2&parent_id=${pgroup1_id}" | jq '.id')
pgroup2_sgroup1_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/groups?access_token=${access_token}&name=ParentGroup2SubGroup1&path=ParentGroup2SubGroup1&parent_id=${pgroup2_id}" | jq '.id')
pgroup2_sgroup2_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/groups?access_token=${access_token}&name=ParentGroup2SubGroup2&path=ParentGroup2SubGroup2&parent_id=${pgroup2_id}" | jq '.id')

# create group project
echo "Creating a project in group/subgroup"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project1&namespace_id=${pgroup1_id}"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project2&namespace_id=${pgroup1_id}"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project3&namespace_id=${pgroup1_id}"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project1&namespace_id=${pgroup1_sgroup1_id}"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project2&namespace_id=${pgroup1_sgroup1_id}"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project3&namespace_id=${pgroup1_sgroup1_id}"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project1&namespace_id=${pgroup1_sgroup2_id}"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project2&namespace_id=${pgroup1_sgroup2_id}"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project3&namespace_id=${pgroup1_sgroup2_id}"

curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project1&namespace_id=${pgroup2_id}"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project2&namespace_id=${pgroup2_id}"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project3&namespace_id=${pgroup2_id}"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project1&namespace_id=${pgroup2_sgroup1_id}"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project2&namespace_id=${pgroup2_sgroup1_id}"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project3&namespace_id=${pgroup2_sgroup1_id}"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project1&namespace_id=${pgroup2_sgroup2_id}"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project2&namespace_id=${pgroup2_sgroup2_id}"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects?access_token=${access_token}&name=Project3&namespace_id=${pgroup2_sgroup2_id}"

# create user project
echo "Creating users project"
project1_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects/user/${user7_id}?access_token=${access_token}&name=User7Project" | jq '.id')
project2_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects/user/${user8_id}?access_token=${access_token}&name=User8Project" | jq '.id')
project3_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects/user/${user9_id}?access_token=${access_token}&name=User9Project" | jq '.id')
project4_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects/user/${user10_id}?access_token=${access_token}&name=User10Project" | jq '.id')
project5_id=$(curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects/user/${user11_id}?access_token=${access_token}&name=User11Project" | jq '.id')

# create hooks for projects
echo "Creating hooks for projects"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects/${project1_id}/hooks?access_token=${access_token}&url=http%3A%2F%2Fwww.sample.com%2F"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects/${project2_id}/hooks?access_token=${access_token}&url=http%3A%2F%2Fwww.sample.com%2F"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects/${project3_id}/hooks?access_token=${access_token}&url=http%3A%2F%2Fwww.sample.com%2F"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects/${project4_id}/hooks?access_token=${access_token}&url=http%3A%2F%2Fwww.sample.com%2F"
curl -X POST "${GITLAB_HTTP_URL}/api/v4/projects/${project5_id}/hooks?access_token=${access_token}&url=http%3A%2F%2Fwww.sample.com%2F"
