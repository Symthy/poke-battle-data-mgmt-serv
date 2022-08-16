#./node_modules/.bin/openapi-generator-cli generate -g go-server -i ./docs/PokeRest.v1.yaml -o ./src/autoGen

shell_path=$(dirname $(readlink -f $0))

mkdir -p "${shell_path}/pokeRest/adapters/rest/autogen/server"

oapi-codegen -generate "server" -package server "${shell_path}/docs/PokeRest.v1.yaml" > "${shell_path}/pokeRest/adapters/rest/autogen/server/api.gen.go"

oapi-codegen -generate "types" -package server "${shell_path}/docs/PokeRest.v1.yaml" > "${shell_path}/pokeRest/adapters/rest/autogen/server/types.gen.go"

wire "${shell_path}/pokerest/adapters/di/wire.go"
