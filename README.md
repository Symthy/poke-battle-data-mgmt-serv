# PokeRest

Poke REST API

## API Specification File

apiDoc/PokeRest.v1.yaml

## Code Auto Generate

```
./autoGen.sh
```

## Use Libraries

### OpenAPI server code generate

- oapi-codegen: https://github.com/deepmap/oapi-codegen

### Framework

- Echo

### DI 

- wire
https://github.com/google/wire
```
go get github.com/google/wire/cmd/wire
```

### ORM 
- gorm: 本使用
https://gorm.io/ja_JP/docs/index.html
```sh
go get gorm.io/gorm
go get gorm.io/driver/postgres
```

- ent：使用見送り
https://entgo.io/ja/docs/getting-started/
```sh
go get entgo.io/ent/cmd/ent
cd pokeRest/adapters/orm/entio
go run entgo.io/ent/cmd/ent init --target ./schema Pokemons Forms Abilities Moves TypeCompatibility HeldItems Users TrainedPokemons Party Tag BattleRecord BattleOpponentParty
cd autogen/ent/
go run entgo.io/ent/cmd/entc generate --target ./autogen/ent ./autogen/ent/schema
```

### Other

- go-funk: 戻り値がinterface{}になるため使用は限定的にする。genericsが来ればWrapして使える
https://github.com/thoas/go-funk

## Use tools

### Create API spec tool

Stoplight Studio

https://stoplight.io/studio/

### Migration tool

prisma: 
- https://github.com/prisma/prisma
- https://www.prisma.io/docs/getting-started/setup-prisma/start-from-scratch-typescript-postgres

### Create ER Diagram tool

#### For Prisma
- https://www.npmjs.com/package/prisma-uml
- https://www.npmjs.com/package/prisma-erd-generator

#### SchemaSpy

コンテナのpostgresに繋がらずうまくいかず

フォルダ： ./schema

ref: [PostgreSQLのER図をコマンド一発で生成したい](https://zenn.dev/ucwork/articles/a42121e85451be)
