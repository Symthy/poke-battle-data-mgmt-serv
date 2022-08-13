# AST & Jennifer

モチベーション：

ref: [entity からコード自動生成した話](https://tech.mfkessai.co.jp/2019/09/ebgen/)

ベースとなるソースファイルから AST 取得し必要情報抽出 ⇒ jennifer にてソース自動生成

- その他 refs
  - [コードの自動生成は高いレバレッジを達成できる手段 Go におけるコマンド、モデル、テンプレートを組み合わせた実装](https://logmi.jp/tech/articles/326270#s4)
  - [go 言語の AST の全てのコメントは、\*ast.File 以下の子孫ノードから集める事はできないという話](https://pod.hatenablog.com/entry/2018/03/06/104109)
- ※template を使用する場合は、以下が参考になる
  - [Go でひたすら運用を楽にするためのコード生成をする](https://medium.com/eureka-engineering/advent-calendar-2020-vim-go-generate-1a2a11cf0cef)
  - https://github.com/kaneshin/go-generate-sample

## AST（抽象構文木） 取得/解析コード

ソースファイルの情報を取得可能

```golang
import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func Parse(filename string) error {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, parser.Mode(0))
	if err != nil {
		return err
	}
	// for _, d := range f.Decls {  // 1ファイルの全情報表示
	// 	ast.Print(fset, d)
	// 	fmt.Println()
	// }

	fieldNames := []string{}
	ast.Inspect(f, func(node ast.Node) bool {
		t, ok := node.(*ast.TypeSpec)
		if !ok {
			return true
		}

		st, ok := t.Type.(*ast.StructType)
		if !ok {
			return true
		}
		for _, field := range st.Fields.List {
			ast.Print(fset, field)
			fmt.Println()

			for _, nameNode := range field.Names {  // フィールド名取得
				fieldNames = append(fieldNames, nameNode.Name)
			}
		}
		return true
	})
	fmt.Print(fieldNames)
	return nil
}
```

## 解析結果

```golang
type Move struct {
    id                 identifier.MoveId
    name               string
	description        *string
    effects            *battles.Effects
	usedMember         []*Member
    leanableCharacters []Character
    mixin.UpdateTimes
}
```

- id identifier.MoveId

```
0  *ast.Field {
1  .  Names: []*ast.Ident (len = 1) {
2  .  .  0: *ast.Ident {
3  .  .  .  NamePos: <full path>\move.go:12:2
4  .  .  .  Name: "id"
5  .  .  .  Obj: *ast.Object {
6  .  .  .  .  Kind: var
7  .  .  .  .  Name: "id"
8  .  .  .  .  Decl: *(obj @ 0)
9  .  .  .  }
10  .  .  }
11  .  }
12  .  Type: *ast.SelectorExpr {
13  .  .  X: *ast.Ident {
14  .  .  .  NamePos: <full path>\move.go:12:16
15  .  .  .  Name: "identifier"
16  .  .  }
17  .  .  Sel: *ast.Ident {
18  .  .  .  NamePos: <full path>\move.go:12:27
19  .  .  .  Name: "HeldItemId"
20  .  .  }
21  .  }
22  }
```

- name string

```
0  *ast.Field {
1  .  Names: []*ast.Ident (len = 1) {
2  .  .  0: *ast.Ident {
3  .  .  .  NamePos: <full path>\move.go:13:2
4  .  .  .  Name: "name"
5  .  .  .  Obj: *ast.Object {
6  .  .  .  .  Kind: var
7  .  .  .  .  Name: "name"
8  .  .  .  .  Decl: *(obj @ 0)
9  .  .  .  }
10  .  .  }
11  .  }
12  .  Type: *ast.Ident {
13  .  .  NamePos: <full path>\move.go:13:16
14  .  .  Name: "string"
15  .  }
16  }
```

- description \*string

```
     0  *ast.Field {
     1  .  Names: []*ast.Ident (len = 1) {
     2  .  .  0: *ast.Ident {
     3  .  .  .  NamePos: <full path>/move.go:14:2
     4  .  .  .  Name: "description"
     5  .  .  .  Obj: *ast.Object {
     6  .  .  .  .  Kind: var
     7  .  .  .  .  Name: "description"
     8  .  .  .  .  Decl: *(obj @ 0)
     9  .  .  .  }
    10  .  .  }
    11  .  }
    12  .  Type: *ast.StarExpr {
    13  .  .  Star: <full path>/move.go:14:14
    14  .  .  X: *ast.Ident {
    15  .  .  .  NamePos: <full path>/move.go:14:15
    16  .  .  .  Name: "string"
    17  .  .  }
    18  .  }
    19  }
```

- effects \*battles.Effects

```
0  *ast.Field {
1  .  Names: []*ast.Ident (len = 1) {
2  .  .  0: *ast.Ident {
3  .  .  .  NamePos: <full path>\move.go:15:2
4  .  .  .  Name: "battleEffects"
5  .  .  .  Obj: *ast.Object {
6  .  .  .  .  Kind: var
7  .  .  .  .  Name: "battleEffects"
8  .  .  .  .  Decl: *(obj @ 0)
9  .  .  .  }
10  .  .  }
11  .  }
12  .  Type: *ast.StarExpr {
13  .  .  Star: <full path>\move.go:15:16
14  .  .  X: *ast.SelectorExpr {
15  .  .  .  X: *ast.Ident {
16  .  .  .  .  NamePos: <full path>\move.go:15:17
17  .  .  .  .  Name: "battles"
18  .  .  .  }
19  .  .  .  Sel: *ast.Ident {
20  .  .  .  .  NamePos: <full paht>\move.go:15:25
21  .  .  .  .  Name: "BattleEffects"
22  .  .  .  }
23  .  .  }
24  .  }
25  }
```

- ポインタ配列 usedMember []\*Member

```
0  *ast.Field {
1  .  Names: []*ast.Ident (len = 1) {
2  .  .  0: *ast.Ident {
3  .  .  .  NamePos: <full path>\move.go:4:2
4  .  .  .  Name: "items"
5  .  .  .  Obj: *ast.Object {
6  .  .  .  .  Kind: var
7  .  .  .  .  Name: "items"
8  .  .  .  .  Decl: *(obj @ 0)
9  .  .  .  }
10  .  .  }
11  .  }
12  .  Type: *ast.ArrayType {
13  .  .  Lbrack: <full path>\move.go:4:8
14  .  .  Elt: *ast.StarExpr {
15  .  .  .  Star: <full path>\move.go:4:10
16  .  .  .  X: *ast.Ident {
17  .  .  .  .  NamePos: <full path>\move.go:4:11
18  .  .  .  .  Name: "User"
19  .  .  .  }
20  .  .  }
21  .  }
22  }
```

- 配列 leanableCharacters []Character

```
0  *ast.Field {
1  .  Names: []*ast.Ident (len = 1) {
2  .  .  0: *ast.Ident {
3  .  .  .  NamePos: <full path>\move.go:14:2
4  .  .  .  Name: "leanablePokemons"
5  .  .  .  Obj: *ast.Object {
6  .  .  .  .  Kind: var
7  .  .  .  .  Name: "leanablePokemons"
8  .  .  .  .  Decl: *(obj @ 0)
9  .  .  .  }
10  .  .  }
11  .  }
12  .  Type: *ast.ArrayType {
13  .  .  Lbrack: <full path>\move.go:14:20
14  .  .  Elt: *ast.Ident {
15  .  .  .  NamePos: <full path>\move.go:14:22
16  .  .  .  Name: "Pokemon"
17  .  .  }
18  .  }
19  }
```

- ポインタ配列

```
0  *ast.Field {
1  .  Names: []*ast.Ident (len = 1) {
2  .  .  0: *ast.Ident {
3  .  .  .  NamePos: <full path>\move.go:16:2
4  .  .  .  Name: "leanablePokemons"
5  .  .  .  Obj: *ast.Object {
6  .  .  .  .  Kind: var
7  .  .  .  .  Name: "leanablePokemons"
8  .  .  .  .  Decl: *(obj @ 0)
9  .  .  .  }
10  .  .  }
11  .  }
12  .  Type: *ast.ArrayType {
13  .  .  Lbrack: <full path>\move.go:16:19
14  .  .  Elt: *ast.StarExpr {
15  .  .  .  Star: <full path>\move.go:16:21
16  .  .  .  X: *ast.SelectorExpr {
17  .  .  .  .  X: *ast.Ident {
18  .  .  .  .  .  NamePos: <full path>\move.go:16:22
19  .  .  .  .  .  Name: "identifier"
20  .  .  .  .  }
21  .  .  .  .  Sel: *ast.Ident {
22  .  .  .  .  .  NamePos: <full path>\move.go:16:33
23  .  .  .  .  .  Name: "PokemonId"
24  .  .  .  .  }
25  .  .  .  }
26  .  .  }
27  .  }
28  }
```

- 埋め込み mixin.UpdateTimes

```
0  *ast.Field {
1  .  Type: *ast.SelectorExpr {
2  .  .  X: *ast.Ident {
3  .  .  .  NamePos: <full path>\move.go:16:2
4  .  .  .  Name: "mixin"
5  .  .  }
6  .  .  Sel: *ast.Ident {
7  .  .  .  NamePos: <full path>\move.go:16:8
8  .  .  .  Name: "UpdateTimes"
9  .  .  }
10  .  }
11  }
```
