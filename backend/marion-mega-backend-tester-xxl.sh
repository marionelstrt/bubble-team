#!/bin/sh

API_URL=http://localhost:8080/api

case "$1" in
    create_user)
        curl -s $API_URL/account/create -d '{
            "boba": "taro",
            "lastName":"mrn",
            "password":"chat",
            "email":"m",
            "name":"chat"
        }' | jq
        ;;
    *)
        echo unknown command $1 >&2
        ;;
esac
