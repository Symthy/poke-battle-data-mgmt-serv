#!/bin/bash
scripts_path=$(dirname $(readlink -f $0))
pushd "${scripts_path}/../tools/schema" > /dev/null

docker-compose up

popd > /dev/null