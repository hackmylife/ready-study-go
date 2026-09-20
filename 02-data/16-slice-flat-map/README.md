# Exercise: 16-slice-flat-map

## Goal

各要素の変換結果を一段平らに連結する。

## Task

`FlatMap`は先頭から各要素へfを一度ずつ適用し、返されたsliceを一段だけ連結します。入力の順序と、fが返す要素の順序・重複を保ちます。結果が空ならnilです。入力やfが返したsliceは書き換えず、結果の要素を書き換えてもそれらのsliceに影響しないようにします。

`T`は入力要素、`U`は出力要素の型です。fはnilではなく、入力の参照先も変更しないものとします。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `FlatMap([]string{"Go is", "", "fun Go"}, strings.Fields)` | `[]string{"Go", "is", "fun", "Go"}` |
| `FlatMap([]string{"", "  "}, strings.Fields)` | `nil` |
| `FlatMap([]string(nil), strings.Fields)` | `nil` |

## Constraints

- 標準のfor/rangeとappendで実装してください。
- テストと関数のsignatureを変更しないでください。
- テストは振る舞いを確認します。構文の使い分けは模範解答・解説でも確認してください。

## Run

```bash
go run ./cmd/koans check 02-data/16-slice-flat-map
```

[段階的なヒント](HINTS.md) / [ループとコレクションの案内](../../docs/loops-and-collections.md)
