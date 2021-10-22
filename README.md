# PokeRest

Poke REST API

Use OpenApi Generator

## API Specification File

apiDoc/PokeRest.v1.yaml

## Code Auto Generate

```
./autoGen.sh
```

## Use Libraries

- ent: https://entgo.io/ja/docs/getting-started/
```
go get entgo.io/ent/cmd/ent
cd pokeRest/adapters/orm
go run entgo.io/ent/cmd/ent init --target ./schema Users Pokemons Abilities Moves Types Items Forms
cd autogen/ent/
go run entgo.io/ent/cmd/entc generate --target . ../../schema
```
- wire

## Use tools

### Create API spec tool

Stoplight Studio

https://stoplight.io/studio/

### Migration tool

prisma: 
- https://github.com/prisma/prisma
- https://www.prisma.io/docs/getting-started/setup-prisma/start-from-scratch-typescript-postgres

### Create ER Diagram tool

- https://www.npmjs.com/package/prisma-uml
- https://www.npmjs.com/package/prisma-erd-generator
