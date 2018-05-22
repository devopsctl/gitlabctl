export GITLAB_USERNAME=root
export GITLAB_PASSWORD=123qwe123
export GITLAB_HTTP_URL=${GITLAB_HTTP_URL:-http://localhost:10080}
export GITLAB_API_HTTP_URL=${GITLAB_API_HTTP_URL:-${GITLAB_HTTP_URL}/api/v4}
./testdata/get_token.sh
export GITLAB_PRIVATE_TOKEN=$(cat token.txt)
export GITLAB_OAUTH_TOKEN=$(./testdata/get_oauth_token.sh)
