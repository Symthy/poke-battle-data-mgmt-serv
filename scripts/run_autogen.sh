#./node_modules/.bin/openapi-generator-cli generate -g go-server -i ./docs/PokeRest.v1.yaml -o ./src/autoGen

scripts_path=$(dirname $(readlink -f $0))
root_path = "${scripts_path}/.."

mkdir -p "${root_path}/internal/adapters/rest/autogen/server"

oapi-codegen -generate "server" -package server "${root_path}/docs/PokeRest.v1.yaml" > "${root_path}/internal/adapters/rest/autogen/server/api.gen.go"

oapi-codegen -generate "types" -package server "${root_path}/docs/PokeRest.v1.yaml" > "${root_path}/internal/adapters/rest/autogen/server/types.gen.go"

wire "${root_path}/internal/adapters/di/wire.go"

# go generate "${root_path}/tools/codegen/..."
# cp 