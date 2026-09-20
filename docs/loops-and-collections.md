# ループとコレクション

Goでは反復を`for`で書きます。初期化・条件・更新を並べる形式のほか、条件だけ、無限ループ、rangeがあります。既存の演習番号はそのままに、追加演習を章の末尾に置いています。

## forの使い分け

| 形式 | 使いどころ・例 | 演習 |
|---|---|---|
| `for i := 0; i < n; i++` | 添字や刻み幅を制御する。n=3ならiは0, 1, 2 | [05-for](../01-language/05-for/README.md) |
| `for n > 0` | while相当。nを10で割り、120なら120 → 12 → 1 → 0 | [14-for-condition](../01-language/14-for-condition/README.md) |
| `for { ... }` | ループ内で終了を判断する。`break`で終了、`continue`で次へ進む | [15-for-break-continue](../01-language/15-for-break-continue/README.md) |
| `for i, value := range values` | sliceやarrayの添字と値を取り出す | [06-range](../01-language/06-range/README.md) |
| `for i := range values` | 元の要素を`values[i]`で更新する | [17-range-index](../01-language/17-range-index/README.md) |
| `for i := range n` | 0からn-1まで。`for range n`なら値を受け取らず反復する | [16-range-integer](../01-language/16-range-integer/README.md) |
| `for key, value := range m` | mapを列挙する。順序は保証されない | [18-map-keys](../02-data/18-map-keys/README.md) |
| `for offset, r := range text` | UTF-8のバイト位置とruneを取り出す。`"Aあ"`なら(0, 'A'), (1, 'あ') | [03-rune](../02-data/03-rune/README.md) |
| `for value := range ch` | channelがcloseされ、値を受信し終わるまで反復する | [05-close-channel](../06-concurrency/05-close-channel/README.md) |
| `for value := range seq` | iteratorから値を受け取り、breakで途中終了する | [19-range-iterator](../02-data/19-range-iterator/README.md) |

使わない添字を捨てるときは`for _, value := range values`と書きます。`range`で受け取る値はコピーなので、元のsliceを更新するなら添字を使います。ネストしたループから抜ける場合にはラベル付き`break`も使えます。詳細は[Go仕様のfor文](https://go.dev/ref/spec#For_statements)と[break文](https://go.dev/ref/spec#Break_statements)を参照してください。

## map・filter・flatMap・reduceとの対応

ここでいう**mapは要素を変換する操作**です。Goの辞書型`map[K]V`や標準の`maps`パッケージとは区別してください。

| 操作 | 入力 → 出力の例 | 演習 |
|---|---|---|
| map | 各整数を二倍にする：`[2, -3] → [4, -6]` | [06-slice-range](../02-data/06-slice-range/README.md) |
| map（変換関数を渡す） | strconv.Itoaで変換：`[2, -3] → ["2", "-3"]` | [15-slice-map](../02-data/15-slice-map/README.md) |
| filter | 正の数だけ残す：`[-1, 3, 0, 2] → [3, 2]` | [07-slice-filter](../02-data/07-slice-filter/README.md) |
| flatMap | 単語に分けて連結：`["Go is", "fun"] → ["Go", "is", "fun"]` | [16-slice-flat-map](../02-data/16-slice-flat-map/README.md) |
| reduce / fold | 初期値10に加算：`[1, 2, 3] → 16` | [17-slice-reduce](../02-data/17-slice-reduce/README.md) |
| groupBy | メンバーをTeamで分類：`[{Name:"A", Team:"red"}, {Name:"B", Team:"blue"}, {Name:"C", Team:"red"}] → {"red": ["A", "C"], "blue": ["B"]}` | [14-map-grouping](../02-data/14-map-grouping/README.md) |

Go 1.27.1の標準パッケージには汎用の`Map`・`FlatMap`・`Reduce`関数はありません。この教材ではfor/rangeとappendで実装します。`[T, U any]`を使う演習では、入力要素の型Tと出力要素の型Uを分けて、型が変わる変換も扱います。例えば`Map([]int{2}, strconv.Itoa)`は`[]string{"2"}`を返します。

一箇所の処理なら具体的なループで書けます。変換処理を共通化したいときに関数を受け取る形を検討してください。エラーを返す変換なら、エラー時に停止するか、結果と一緒に蓄積するかも契約として決める必要があります。この追加演習の変換関数はエラーを返しません。

## 標準ライブラリでできること

[`slices`](https://pkg.go.dev/slices)には検索、整列、コピー、iteratorからの収集などがあります。[`maps`](https://pkg.go.dev/maps)は辞書のコピーやキー・値の列挙などを扱います。[`iter`](https://pkg.go.dev/iter)はiteratorの型と操作を提供します。

例えば`slices.Sorted(maps.Keys(map[string]int{"pear": 1, "apple": 9}))`は`[]string{"apple", "pear"}`を返します。`maps.Keys`の列挙順に頼らず、キーを明示的に整列しています。

`slices.Values`はsliceからiteratorを作り、`slices.Collect`はiteratorの全要素をsliceに集めます。`Take`の演習では全件を集めず、必要な件数で終了します。これらは標準APIであり、この教材で実装する`Map`・`FlatMap`とは別です。
