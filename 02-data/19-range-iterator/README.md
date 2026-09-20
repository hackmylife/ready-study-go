# Exercise: 19-range-iterator

## Goal

iteratorをrangeし、必要な件数で走査を終了する。

## Task

`Take`はseqから先頭n件までを順番に取得してsliceにします。n件に満たなければ取得できた分だけ返します。nが0以下ならseqを呼ばずnilを返します。空のseqの結果もnilです。

n件を取得した時点で終了し、余分な要素を要求しません。seqはnilではなく、yieldがfalseなら終了するiteratorです。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `Take(slices.Values([]int{4, 5, 6}), 2)` | `[]int{4, 5}` |
| `Take(slices.Values([]int{4}), 3)` | `[]int{4}` |
| `Take(slices.Values([]int{4}), 0)` | `nil。seqを呼ばない` |

## Constraints

- `for value := range seq`と`break`を使ってください。goroutineは不要です。
- テストと関数のsignatureを変更しないでください。
- テストは振る舞いを確認します。構文の使い分けは模範解答・解説でも確認してください。

## Run

```bash
go run ./cmd/koans check 02-data/19-range-iterator
```

[段階的なヒント](HINTS.md) / [ループとコレクションの案内](../../docs/loops-and-collections.md)
