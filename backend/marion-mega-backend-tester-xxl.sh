#!/bin/sh

API_URL=http://localhost:8080/api

case "$1" in
    create_user)
        curl -s $API_URL/account/create -d '{
            "boba": "taro",
            "lastName":"mrn",
            "password":"chat",
            "email":"mariontarento@gmail.com",
            "name":"chat"
        }' | jq
        ;;
    verify_user)
        curl -s $API_URL/account/verify -d '{
            "email":"m",
            "code":"'$2'"
        }' | jq
        ;;
    *)
        echo unknown command $1 >&2
        ;;
esac
