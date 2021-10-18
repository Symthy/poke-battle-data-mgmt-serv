docker run -d --rm -p 3030:8080 plantuml/plantuml-server:jetty
./node_modules/.bin/prisma-uml ./prisma/schema.prisma -o png -s http://localhost:3030 -f ./dia/erd.png
