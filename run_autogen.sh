#./node_modules/.bin/openapi-generator-cli generate -g go-server -i ./apiDoc/PokeRest.v1.yaml -o ./src/autoGen

mkdir -p ./pokeRest/adapters/rest/autogen/server

oapi-codegen -generate "server" -package server ./apiDoc/PokeRest.v1.yaml > ./pokeRest/adapters/rest/autogen/server/api.gen.go

oapi-codegen -generate "types" -package server ./apiDoc/PokeRest.v1.yaml > ./pokeRest/adapters/rest/autogen/server/types.gen.go

wire ./pokerest/adapters/di/wire.go
