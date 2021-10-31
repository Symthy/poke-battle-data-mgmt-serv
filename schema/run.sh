#docker run --rm -v "$PWD/output" --net="host" schemaspy/schemaspy:latest -t pgsql11 -host postgres:5432 -db pokedb -u postgres -p postgres

docker run -v "$PWD/output:/output" -v "$PWD/schemaspy.properties:/schemaspy.properties" schemaspy/schemaspy:latest
