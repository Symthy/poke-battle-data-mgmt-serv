#./node_modules/.bin/openapi-generator-cli generate -g go-server -i ./apiDoc/PokeRest.v1.yaml -o ./src/autoGen

mkdir -p ./pokeRest/autoGen/server

oapi-codegen -generate "server" -package server ./apiDoc/PokeRest.v1.yaml > ./pokeRest/autogen/server/api.gen.go

oapi-codegen -generate "types" -package server ./apiDoc/PokeRest.v1.yaml > ./pokeRest/autogen/server/types.gen.go
