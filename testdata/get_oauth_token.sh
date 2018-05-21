#!/bin/bash -x

curl -X POST "http://localhost:10080/oauth/token?grant_type=password&username=root&password=123qwe123" | jq '.access_token' | tr -d '"'

