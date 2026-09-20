# Exercise: 16-range-integer

## Goal

整数へのrangeで回数を指定する。

## Task

`Indices`は0以上n未満の整数を昇順で返します。nが0以下ならnilを返します。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `Indices(4)` | `[]int{0, 1, 2, 3}` |
| `Indices(1)` | `[]int{0}` |
| `Indices(0) または Indices(-2)` | `nil` |

## Constraints

- `for i := range n`を使ってください。
- テストと関数のsignatureを変更しないでください。
- テストは振る舞いを確認します。構文の使い分けは模範解答・解説でも確認してください。

## Run

```bash
go run ./cmd/koans check 01-language/16-range-integer
```

[段階的なヒント](HINTS.md) / [ループとコレクションの案内](../../docs/loops-and-collections.md)
