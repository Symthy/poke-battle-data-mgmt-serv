# Memo Stack

(途中から開始)
調べたことを無邪気にスタックするメモ置き場

## Important Reference

- [Effective Go](https://go.dev/doc/effective_go)
- [逆引きGolang](https://ashitani.jp/golangtips/index.html)

- [エキスパートGo](https://www.slideshare.net/takuyaueda967/go-77689475) 
- [今どきの Go の書き方まとめ (2020 年末版)](https://www.m3tech.blog/entry/golang-way-primer)

## Golang

### 基本

- メモリアロケーション(割り当て)の回数が多いと速度低下に繋がる

- [Golangのnew()とmake()の違い](https://cipepser.hatenablog.com/entry/go-new-make)
  - make は slice, map, channelのみ。初期化する
  - new は、任意の型。初期化しない。戻り値はポインタ

- [GoでSet型を実現する場合の選択肢](https://kitakitabauer.hatenablog.com/entry/2017/04/04/204701)
  - map × struct で実現 `map[string]struct{}`
  - ライブラリ  https://github.com/deckarep/golang-set

- [忘れがちなGo言語の書式指定子を例付きでまとめた](https://qiita.com/Sekky0905/items/c9cbda2498a685517ad0)

- [Go のモジュール管理【バージョン 1.17 改訂版】](https://zenn.dev/spiegel/articles/20210223-go-module-aware-mode)

### 命名

[Golangのパッケージ名はどうするのが正解か](https://zenn.dev/_kazuya/scraps/fdc65096b0d1d7)

[パッケージ名 (Package names)](https://www.ymotongpoo.com/works/goblog-ja/post/package-names/)

[他言語プログラマが最低限、気にすべきGoのネーミングルール](https://zenn.dev/keitakn/articles/go-naming-rules)

[Go の命名規則](https://micnncim.com/posts/ja/go-naming-convention)

### エラーハンドリング

#### エラー基礎

[Error handling and Go](https://go.dev/blog/error-handling-and-go)
- 公式情報だが参考にならないかも

スタックトレース不要時のerrorラップ
```
fmt.Errorf("detail: %w", err)
```
[go1.13のerrorのラップ](https://qiita.com/gal1996/items/eceacef3a8453cfdb3bf)



pkg/errorsのスタックトレース出力
```go
import (
	"fmt"
 	"github.com/pkg/errors"
)

err := errors.New("error")
fmt.Printf("%+v\n", err)
```



[[Go 1.13~] errors.Is と errors.As の違いについてお気持ちを理解する](https://qiita.com/hiro_o918/items/fb01014e51354b8bb49f)
- errors.Is: 特定のエラーとの比較(あるエラーが特定のエラーを Wrap したものかを判別するためのもの)
- errors.As: エラーに対する型アサーション(自前でエラーを用意した時に，ビルトインのエラーなのか自前で用意したエラーかを判別する際に使える)

#### 見解

- ログへの出力は上位層1箇所に留める
- カスタムエラーにして、level,code,msg 等を付与するして汎用性を上げる
- 下位から上位に上げる時、必要に応じてWrapして綺麗にスタックレースに出す
- pkg/errors で実装。(スタックとレース`fmt.Printf("%+v\n", err)`で取れる)

- 物が複雑で、エラーの多重ラップがほぼ常に必要なら、xerrors で以下のようにした方が良いかもしれない
★[Goでのオススメエラーハンドリング手法 (2020/10)](https://medium.com/eureka-engineering/golang-errors-wrapping-1531bdf409f8)

- スタックトレース要否で以下方法を取る
  - 不要なら fmt.Errorf でラップ＆ errors.Is でエラーの種類による処理分岐
  - 必要なら(標準のerrorsではスタックとレース保持できないため) github.com/pkg/errors を使う。
    - errors.WithStackを使うと狙ったStackTraceだけを綺麗に出せる(し忘れるとstackTraceが出ない)
    - errors.Wrapは、Wrap毎のStackTraceが全部出る。パフォーマンスに懸念ありだが、調査には100困らない
[今goのエラーハンドリングを無難にしておく方法（2021.09現在）](https://zenn.dev/nekoshita/articles/097e00c6d3d1c9)

- もしくはスタックとレースは`zap.Stack("").String`での取得にまかせる

- RESTfulなら、以下のようにするのが良さそう
  - API用のエラーとApp用のエラーは分けた方が良い（単一責任）
[Error handling in Go HTTP applications](https://www.joeshaw.org/error-handling-in-go-http-applications/)
  
- 独自エラークラスでも、Is/As/Unwrap等名前は揃えよう
[Go1.17で警告されるようになったerror#Is/As/Unwrap](https://future-architect.github.io/articles/20210819b/)

#### reference

[Go Tips 連載5: エラーコードベースの例外ハンドリングの実装＋morikuni/failureサンプル](https://future-architect.github.io/articles/20200522/)
- エラーコードを利用した際に重要なことは、エラーコード外のエラーを発生させないことにある

---

[APIサーバのおけるGoのエラーハンドリングについて考えてみる](https://tutuz-tech.hatenablog.com/entry/2020/03/26/193519)

- 処理部分はerrorを返して、呼び元(Controller相当)でエラーの種類に応じて、Http Status Code を設定する

```go
func (u *UserHandle) CreateHandler(w http.ResponseWriter, r *http.Request) {
    err := u.Create(w, r)
    if err != nil {
        writeLog()
        var xe *XError
        if errors.As(err, &xe) {
            http.Error(w, fmt.Sprintf("...: %w", err) , 400)
            return
        }
        var ye *YError
        if errors.As(err, &ye) {
            http.Error(w, fmt.Sprintf("...: %w", err) , 500)
            return
        }

        // application unknown error
        http.Error(w, fmt.Sprintf("...: %w", err) , 500)
        return
    }
}
```

---

★[今goのエラーハンドリングを無難にしておく方法（2021.09現在）](https://zenn.dev/nekoshita/articles/097e00c6d3d1c9)

- golang.org/x/xerrors はもうメンテされてない (go1.13で一部取り込まれて役目を終えた？)
- pkg/errors を使う
  - stacktraceが不要な場合
    - fmt.Errorfでラップし、errors.Isで判定する
  - stacktraceが必要な場合
    - pkg/errors.Wrapでラップし、pkg/errors.Causeで判定する
    - log.Printf("%+v", err)でstacktraceを出力する
    - 新しいエラーを返す時はpkg/errors.New or pkg/errors.Errorfを使う

```go
switch errors.Cause(err) {
case ErrRecordNotFound:
	log.Print("ユーザーが存在しなかった場合のハンドリングをする")
default:
	return err
}

// エラー発生箇所でいずれかを行う
// スタックトレースは綺麗に出る。ただし、外部ライブラリから返ってきたエラーに対してはpkg/errors.WithStackを絶対にしなければいけない。漏れるとトレースが出ない
return errors.WithStack(ErrRecordNotFound)
// どんなエラーを返すときにも、絶対にpkg/errors.Wrapを使う。ただしWrapした分全て出るため見づらい
return errors.Wrap(ErrRecordNotFound, "failed to query")
```

---

[運用を意識したGo言語でのエラーハンドリング/ロギング](https://qiita.com/nayuneko/items/dea02377b797c2a52053)
- エラーはWrapして必要な情報を付与して呼び出し元に返す
  - 発生した箇所を特定できるようメッセージを付与
  - データが原因になりうる場合はキー情報など最低限の情報も付与
- カスタムエラー型を作りエラー特定に必要な情報を持たせる
- ログを出す箇所をまとめる
  - 特定の層のみでログ出力
  - echoではエラーハンドリング関数HTTPErrorHandlerがあるためそこで

★[Goでのオススメエラーハンドリング手法 (2020/10)](https://medium.com/eureka-engineering/golang-errors-wrapping-1531bdf409f8)

- errorsパッケージを作成して独自のエラー構造体を定義
- エラーは全ての箇所でラップして綺麗にスタックトレースを出力
- エラーレベルやエラーメッセージを付与して汎用性を高める

---

★[【Go】エラーハンドリング&ログ出力にまとめて向き合う](https://zenn.dev/yagi_eng/articles/go-error-handling)
- カスタムエラーを活用 (zap.Stack("").Stringでスタックトレース取得できる)
```go
type MyError struct {
  Code       string
  Msg        string
  StackTrace string
}

func (me *MyError) Error() string {
  return fmt.Sprintf("my error: code[%s], message[%s]", me.Code, me.Msg)
}

func New(code string, msg string) *MyError {
  stack := zap.Stack("").String
  return &MyError{
    Code:       code,
    Msg:        msg,
    StackTrace: stack,
  }
}
```
- controllerで複数のエラーを区別する(例：エラーコードでHttpステータス分岐)
- 下位から上位に戻すに連れてエラーメッセージを追加可能()
```go
func (me *MyError) WrapMessage(msg string) {
	me.Msg = fmt.Sprintf("%s %s", msg, me.Msg)
}
// または
type MyError struct {
  Code       int
  Msg        string
  StackTrace string
  err        error  // エラーごとwrapできるようにする（こっちの方が良いかも？）
}
```
- Custom HTTP Error Handlerの活用（echo：echo.NewHTTPErrorをreturnするとe.HTTPErrorHandlerに設定しておいた関数が呼び出される）

- [golangの高速な構造化ログライブラリ「zap」の使い方](https://qiita.com/emonuh/items/28dbee9bf2fe51d28153)

---

[Goでスタックトレースを構造化して取り扱う](https://developers.freee.co.jp/entry/2018/12/23/213000)
- 独自の スタックトレースを作成する例
- モチベーション：ロギングサービス？に独自フォーマットで整形したスタックトレースを送りたかったから

[Goで独自エラー型を定義する](https://qiita.com/minoritea/items/5785f9c8394c7e62bec6)
- 例1: エラーコードを持たせる
- 例2: エラーに追加情報を加える
- 例3: パッケージ間の依存性をなくす
- 例4: 元のエラーを取り出す
- 例5: 複数回ラップしたとき、目当ての情報を取り出す
- 例6: エラー発生時の状態を元にハンドリングする

[Wrap(err) in our production #golang](https://www.wantedly.com/companies/wantedly/post_articles/139554)
- Wantedly の本番 Web アプリケーション上でどのようにしてエラーをハンドリングしているか

---

[Go1.13のError wrappingを触ってみる](https://cipepser.hatenablog.com/entry/go1.13-error-wrapping)
---

[サードパーティのパッケージ](https://zenn.dev/spiegel/books/error-handling-in-golang/viewer/third-party-errors)

- golang.org/x/xerrors < pkg/errors


[Goでスタックトレースを上書きせずにエラーをラップする方法](https://tech.liquid.bio/entry/2021/07/02/135816?utm_source=feed)


[Goエラーハンドリング戦略](https://zenn.dev/nobonobo/articles/0b722c9c2b18d5)

## ロギング

[Goのロギングライブラリ 2021年冬](https://zenn.dev/moriyoshi/articles/1af0659e29d727#go%E3%81%AE%E3%83%AD%E3%82%AE%E3%83%B3%E3%82%B0%E3%83%A9%E3%82%A4%E3%83%96%E3%83%A9%E3%83%AA) 

[Goでのログ出力に標準logとcologを使う](https://qiita.com/kmtr/items/406073651d7a12aab9c6)
- 標準以外のログライブラリを使うと以下の問題がある
  - 標準logを含め、利用するログライブラリが混在しないように気を使う
    - 共通ライブラリを作った場合、そこでも同じログライブラリを使わないと面倒
    - 強い言い方をすると関連するもの全てがログライブラリに汚染されてしまう
  - そのログライブラリの使い方を覚える必要がある
    - チーム開発の場合、全員に周知する必要がある

[colog](https://github.com/comail/colog)
- 文字列の先頭で log level を出し分け出来る手軽なライブラリ

[lumberjack](https://github.com/natefinch/lumberjack)
- ログローテーション

### Zap

[zap]
を使うのがよさそう。gorm でも使える ref: [Frequently Asked Questions](https://github.com/uber-go/zap/blob/master/FAQ.md)
- zap + gorm: [github.com/moul/zapgorm](https://github.com/moul/zapgorm2) 
- zap + echo: [github.com/brpaz/echozap](https://github.com/brpaz/echozap)

ref:
- [logging zap: レベル毎ファイル出力](https://qiita.com/junk616/items/bd642a9712ff8b728978)
- [golangの高速な構造化ログライブラリ「zap」の使い方](https://qiita.com/emonuh/items/28dbee9bf2fe51d28153)
- [Golangの高速なロガーzapとlumberjackでログを出力してrotateさせる](https://www.sambaiz.net/article/104/)
- [GoのロギングライブラリzapのTips](https://christina04.hatenablog.com/entry/golang-zap-tips)
  - 書き込み先をio.Writerで自由に設定したい
  - テストで時刻を固定値にしたい
  - ログレベルによって標準出力、標準エラー出力を分けたい
  - GCP Cloud Loggingのフォーマットでログ出力したい

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

### Custom Middleware

- [Try Golang! EchoでオリジナルのMiddlewareを作ろう！](https://medium.com/veltra-engineering/echo-middleware-in-golang-90e1d301eb27)
```golang
func customMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        // before
        err := next(c)
        // after
        return err
    }
}
```

### Custom Http Error Handler

[【Go】エラーハンドリング&ログ出力にまとめて向き合う](https://zenn.dev/yagi_eng/articles/go-error-handling) 


### echoの ログ出力

echo のロギングは以下２つある。どちらも labstack の gommon/log が使われている
- logger middleware
- echo.Logger
[かけだし Gopher におくる Golang 製 Web Framework echo の logging について](https://goodpatch.com/blog/happy-logging)

標準の出力先は os.Stdout。変更するには以下？
```golang
Echo#Logger.SetOutput(io.Writer)

// middlewareがこれでいいか分からない
fp, err := os.OpenFile("/path/to/log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
if err != nil {
  panic(err)
}
e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
    Output: file,
}))
```

- [ref：echoのログとcologの併用サンプル](https://github.com/kysnm/echo_log_practice/blob/master/main.go)


- [golang echoでRequest/Response Bodyをログに出力](https://qiita.com/fushikky/items/d53b9abcf6fb49c07eda)

```go
 e.Use(middleware.BodyDump(bodyDumpHandler))
```

- echo のロギングを logrus に変える方法 
[echo-logrus](https://github.com/plutov/echo-logrus)

※ echo v5 で echo の logger がどうなるか分からない（変更or削除される可能性がある）
https://github.com/labstack/echo/pull/1555


#### echo loggerでログローテション

標準では搭載してないため追加が必要。lumberjackがメジャー？

logrusへのlumberjackの差し込みが以下でできるので

```golang
// Set the Lumberjack logger
lumberjackLogger := &lumberjack.Logger{
  Filename:   "/var/log/misc.log",
  MaxSize:    10,
  MaxBackups: 3,
  MaxAge:     3,
  LocalTime:  true,
}
logrus.SetOutput(lumberjackLogger)
```

echo の logger にも SetOutput があるので、以下でlumberjackを差し込めそう
```golang
  if l, ok := e.Logger.(*_labstacklog.Logger); ok {
		l.SetHeader("${time_rfc3339} ${level}")
		l.SetOutput(lumberjackLogger1)
	} else {
		e.Logger.Fatalf("failure logging settings. start abort.")
	}
	e.Use(middleware.Logger()) // Log all requests
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: lumberjackLogger2,
	}))
```

## テスト（Testify)

[Testify の assert を使ってテストを書く](https://qiita.com/egawata/items/7c18bbc54dc353034f5f)

[DMMにおけるユーザーレビュー基盤の変革（Goのテスト技法編）](https://inside.dmm.com/entry/2019/03/28/review-go-test)

## その他

[Go言語で.envファイルを環境ごとに読み込む godotenv](https://shi-mo-web.com/golang/gogodotenv-environment-setting-file/)

[Go path/filepathでファイルパスを操作する](https://takuroooooo.hatenablog.com/entry/2020/08/15/Go_path/filepath)

[パッケージ fmt](https://xn--go-hh0g6u.com/pkg/fmt/)

[Go言語のGormを実践投入する時に最低限知っておくべきことのまとめ【ORM】](https://qiita.com/ttiger55/items/3606b8dd570637c12387)
