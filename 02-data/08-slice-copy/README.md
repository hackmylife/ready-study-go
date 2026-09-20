# Exercise: slice-copy

## Goal

共有しないsliceのコピーを作る。

## Task

Cloneは独立したコピーを返す。nilはnil、空sliceは非nilの空sliceとして返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `in := []int{2, 4}` に対して `out := Clone(in)` | `out` は `[]int{2, 4}` |
| 続けて `out[0] = 9` とする | `out` は `[]int{9, 4}`、`in` は `[]int{2, 4}` のまま |
| `Clone(nil)` / `Clone([]int{})` | それぞれ `nil` / nilではない空slice |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./02-data/08-slice-copy
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
