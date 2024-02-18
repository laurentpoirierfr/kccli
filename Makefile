

api-codegen:
	mkdir -p keycloak
	oapi-codegen -package=keycloak -generate=types,client,spec -o=keycloak/keycloak-client.go ./openapi/22.0.0.json
	go mod tidy


