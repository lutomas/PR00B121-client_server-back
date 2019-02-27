#!/usr/bin/env bash

pushd ../
swagger generate server -t ./internal/app/consult-api-server -A consult-api-server -f ./api/consult-api.swagger.yaml

cp -r internal/app/consult-api-server/cmd/consult-api-server-server/main.go cmd/consult-api-server
rm -r internal/app/consult-api-server/cmd
mv internal/app/consult-api-server/restapi/configure_consult_api_server.go internal/app/consult-api-server/restapi/configure_consult_api_server.go.swagger
popd