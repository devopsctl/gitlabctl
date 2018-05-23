#!/bin/bash

if [ ${#@} -lt 1 ]; then
  echo "usage: $0 groupID1 groupID2 ......"
  exit 1
fi

for gid in ${@}; do
  curl --header "Private-Token: $GITLAB_PRIVATE_TOKEN" -X DELETE "http://localhost:10080/api/v4/groups/$gid"
done
