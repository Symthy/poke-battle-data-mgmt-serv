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
```
wire pokeRest/adapters/di/wire.go 
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
go run entgo.io/ent/cmd/entc generate --target ./autogen/ent ../../orm/entio/schema
```

### Logging
- ログローテション
https://github.com/natefinch/lumberjack
```golang
log.SetOutput(&lumberjack.Logger{
    Filename:   "/var/log/myapp/foo.log",
    MaxSize:    500, // megabytes
    MaxBackups: 3,
    MaxAge:     28, //days
    Compress:   true, // disabled by default
})
```


### Test/Mock
ref: https://github.com/bmuschko/go-testing-frameworks

- testify
テストのセットアップ・終了時の処理のコードが重複しないテストコードを書ける+ mock helper有

https://github.com/stretchr/testify

- go-sqlmock
https://github.com/DATA-DOG/go-sqlmock
```
go get -u github.com/DATA-DOG/go-sqlmock
```
gorm + postgresql + go-sqlmock の場合、通常のやり方ではINSERTのモックがうまく行かないらしい。以下参照。
- https://github.com/DATA-DOG/go-sqlmock/issues/118
- https://betterprogramming.pub/how-to-unit-test-a-gorm-application-with-sqlmock-97ee73e36526
- https://simple-minds-think-alike.moritamorie.com/entry/go-sqlmock-gorm

- array の contains
https://zenn.dev/glassonion1/articles/7c7830a269909c

### Other

- go-funk: 戻り値がinterface{}になるため使用は限定的にしておく。※genericsが来ればWrapして使えるはず
https://github.com/thoas/go-funk

- gonstructor: コンストラクタ(ビルダー)自動生成ツール。自前で用意する時間がないため暫定使用。ただし`gorm:～` の指定入れている物には使えず
https://github.com/moznion/gonstructor

- golang-set：使用していないがメモ。凝ったSet型を使いたければ以下を使う
https://github.com/deckarep/golang-set

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

コンテナのpostgresに繋げられず使用中断

フォルダ： ./schema

ref: [PostgreSQLのER図をコマンド一発で生成したい](https://zenn.dev/ucwork/articles/a42121e85451be)
