# Memo Stack

(途中から開始)
調べたことを無邪気にスタックするメモ置き場

## Important Reference

- [Effective Go](https://go.dev/doc/effective_go)
- [逆引きGolang](https://ashitani.jp/golangtips/index.html)

## Golang

- [Golangのnew()とmake()の違い](https://cipepser.hatenablog.com/entry/go-new-make)
  - make は slice, map, channelのみ。初期化する
  - new は、任意の型。初期化しない。戻り値はポインタ

- [GoでSet型を実現する場合の選択肢](https://kitakitabauer.hatenablog.com/entry/2017/04/04/204701)
  - map × struct で実現 `map[string]struct{}`
  - ライブラリ  https://github.com/deckarep/golang-set


## echo
### JWT認証


- [【Go】EchoによるJWT認証の実装](https://log.include.co.jp/2021/02/03/golang_5/)
- [Go言語でEchoを用いて認証付きWebアプリの作成](https://qiita.com/x-color/items/24ff2491751f55e866cf)

ログイン
- パスワードチェック
- custom claims 生成
- claimsを元にトークン生成。jwt.NewWithClaims()
- jwt.Token.SignedString(秘密鍵) でレスポンスに含めるトークン生成

サインアップ
- 

### ログ出力

- [golang echoでRequest/Response Bodyをログに出力](https://qiita.com/fushikky/items/d53b9abcf6fb49c07eda)

```go
 e.Use(middleware.BodyDump(bodyDumpHandler))
```
