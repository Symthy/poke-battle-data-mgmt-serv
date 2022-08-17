if [ $# = 0 ]; then
  #go test ./test/... -v
  go test ./test/internal/... -v
  exit 0
fi

target_path=$1
#go test ./test/${target_path}/... -v
go test ./test/internal/${target_path}/... -v
