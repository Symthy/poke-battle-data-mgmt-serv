# PokeRest

Pokemon REST API

構成は以下を参考

- https://github.com/golang-standards/project-layout

## API Specification File

docs/PokeRest.v1.yaml

### Search

以下のような検索用の独自 QL は設けない

https://github.com/cbrand/go-filterparams

全文検索かける程でもなく検索項目も少ない、項目間は AND 条件のみため。
ただし、一部項目に OR/AND 条件指定が必要なため。一部パラメータでは以下により指定

| 条件 | 指定記号 | 例(encode 無し)   |
| :--- | :------- | :---------------- | -------- | --- | ---- |
| AND  | `&`      | xxx='aaa&bbb&ccc' |
| OR   | `        | `                 | xxx='aaa | bbb | ccc' |

※高度な検索が必要になるなら別途検討

## Code Auto Generate

```sh
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
go install github.com/google/wire/cmd/wire@latest
./run_autogen.sh
```

## Use Libraries

### Lint tool

- golangci: https://github.com/golangci/golangci-lint#macos

※ go-lint は非推奨になった

参考： [golangci を理解する](https://zenn.dev/sanpo_shiho/books/61bc1e1a30bf27/viewer)

### OpenAPI server code generate

- oapi-codegen: https://github.com/deepmap/oapi-codegen

参考：

- [OpenAPI スキーマから生成されるコードに任意の型を指定する](https://times.hrbrain.co.jp/entry/2020/12/18/openapi-x-go-type)

### Framework

- Echo

### DI

- wire
  https://github.com/google/wire

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
cd internal/adapters/orm/entio
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

- ZAP

### Test/Mock

ref: https://github.com/bmuschko/go-testing-frameworks

- testify
  テストのセットアップ・終了時の処理のコードが重複しないテストコードを書ける+ mock helper 有

https://github.com/stretchr/testify

- go-sqlmock
  https://github.com/DATA-DOG/go-sqlmock

```
go get -u github.com/DATA-DOG/go-sqlmock
```

gorm + postgresql + go-sqlmock の場合、通常のやり方では INSERT のモックがうまく行かないらしい。以下参照。

- https://github.com/DATA-DOG/go-sqlmock/issues/118
- https://betterprogramming.pub/how-to-unit-test-a-gorm-application-with-sqlmock-97ee73e36526
- https://simple-minds-think-alike.moritamorie.com/entry/go-sqlmock-gorm

- array の contains
  https://zenn.dev/glassonion1/articles/7c7830a269909c

### Other

- go-funk: 戻り値が interface{}になるため使用は限定的にしておく。※generics が来れば Wrap して使えるはず
  https://github.com/thoas/go-funk

- gonstructor: コンストラクタ(ビルダー)自動生成ツール。自前で用意する時間がないため暫定使用。ただし`gorm:～` の指定入れている物には使えず
  https://github.com/moznion/gonstructor

- golang-set：使用していないがメモ。凝った Set 型を使いたければ以下を使う
  https://github.com/deckarep/golang-set

## Use tools

### Create API spec tool

Stoplight Studio

https://stoplight.io/studio/

### Migration tool

[golang-migrate](https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md]

### Create ER Diagram tool

#### SchemaSpy

そのままでは schemaspy コンテナ内臓の driver が古いためか？ postgres に 以下エラーで接続失敗するため

`Failed to connect to database URL [jdbc:postgresql://postgres:5432/pokedb] The authentication type 10 is not supported. Check that you have configured the pg_hba.conf file to include the client's IP address or subnet, and that it is using an authentication scheme supported by the driver.`

tools/drivers に postgres の driver を配置すること（動作実績：postgresql-42.3.5.jar）

```sh
cd tools
docker-compose up
```

ref:

- [PostgreSQL の ER 図をコマンド一発で生成したい](https://zenn.dev/ucwork/articles/a42121e85451be)
- [【さらば ER 図設計作業】docker-compose で「SchemaSpy」でデータベースのドキュメントを自動生成するやり方](https://blogenist.jp/2019/04/20/8075/)
- docker-compose.yml 作成時に参考：
- [複数の docker-compose 間を接続するには networks を設定する必要があります](https://amateur-engineer.com/docker-compose-network-share/)
