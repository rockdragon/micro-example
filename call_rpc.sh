#!/usr/bin/env bash

exec curl -d 'service=me.moye.srv.business' \
    -d 'method=API.Call' \
    -d 'request={"name": "Jack"}' \
    http://localhost:8080/rpc