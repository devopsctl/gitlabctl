#
# From: https://stackoverflow.com/questions/47948887/login-to-gitlab-using-curl
#

# curl for the login page to get a session cookie and the sources with the auth tokens
body_header=$(curl -c cookies.txt -i "${GITLAB_HTTP_URL}/users/sign_in" -s)

# grep the auth token for the user login for
#   not sure whether another token on the page will work, too - there are 3 of them
csrf_token=$(echo $body_header | perl -ne 'print "$1\n" if /new_user.*?authenticity_token"[[:blank:]]value="(.+?)"/' | sed -n 1p)

# send login credentials with curl, using cookies and token from previous request
curl -b cookies.txt -c cookies.txt -i "${GITLAB_HTTP_URL}/users/sign_in" \
    --data "user[login]=${GITLAB_USERNAME}&user[password]=${GITLAB_PASSWORD}" \
    --data-urlencode "authenticity_token=${csrf_token}"

# send curl GET request to personal access token page to get auth token
body_header=$(curl -H 'user-agent: curl' -b cookies.txt -i "${GITLAB_HTTP_URL}/profile/personal_access_tokens" -s)
csrf_token=$(echo $body_header | perl -ne 'print "$1\n" if /authenticity_token"[[:blank:]]value="(.+?)"/' | sed -n 1p)

# curl POST request to send the "generate personal access token form"
# the response will be a redirect, so we have to follow using `-L`
body_header=$(curl -L -b cookies.txt "${GITLAB_HTTP_URL}/profile/personal_access_tokens" \
    --data-urlencode "authenticity_token=${csrf_token}" \
    --data 'personal_access_token[name]=golab-generated&personal_access_token[expires_at]=&personal_access_token[scopes][]=api')

# Scrape the personal access token from the response HTML
personal_access_token=$(echo $body_header | perl -ne 'print "$1\n" if /created-personal-access-token"[[:blank:]]value="(.+?)"/' | sed -n 1p)

echo $personal_access_token > token.txt
