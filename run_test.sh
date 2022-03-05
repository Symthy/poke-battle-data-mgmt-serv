if [ $# = 0 ]; then
  #go test ./test/... -v
  go1.18beta2 test ./test/pokeRest/... -v
  exit 0
fi

target_path=$1
#go test ./test/${target_path}/... -v
go1.18beta2 test ./test/pokeRest/${target_path}/... -v
