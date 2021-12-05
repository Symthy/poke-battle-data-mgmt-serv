if [ $# = 0 ]; then
  go test ./test/... -v
fi

target_path=$1
go test ./test/${target_path}/... -v
