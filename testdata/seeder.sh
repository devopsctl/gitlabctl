#!/bin/bash -x

echo "=============================================="
echo "#              SEEDER SCRIPT                 #"
echo "=============================================="


./get_token.sh

personal_access_token=$(cat token.txt)

echo "Testing token"
curl --header "Private-Token: ${personal_access_token}" http://localhost:10080/api/v4/users
