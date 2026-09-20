# Exercise: 15-slice-map

## Goal

変換関数を受け取り各要素を一対一で変換する。

## Task

`Map`は各要素にfを適用し、順序を保った新しいsliceを返します。fは先頭から一度ずつ呼びます。nil入力はnil、nilでない空入力はnilでない空sliceにします。入力の要素は書き換えません。

`[T, U any]`は入力要素の型Tと出力要素の型Uを表す型パラメータです。例えば`[]int`から`[]string`へ変換できます。fはnilではなく、渡された値の参照先も変更しないものとします。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `Map([]int{2, -3}, strconv.Itoa)` | `[]string{"2", "-3"}` |
| `Map([]int(nil), strconv.Itoa)` | `nil` |
| `Map([]int{}, strconv.Itoa)` | `[]string{}（nilではない）` |

## Constraints

- 標準のfor/rangeとsliceで実装してください。ここでのMapは自作関数で、組み込みの辞書型mapとは別です。
- テストと関数のsignatureを変更しないでください。
- テストは振る舞いを確認します。構文の使い分けは模範解答・解説でも確認してください。

## Run

```bash
go run ./cmd/koans check 02-data/15-slice-map
```

[段階的なヒント](HINTS.md) / [ループとコレクションの案内](../../docs/loops-and-collections.md)
