# Exercise: slice-filter

## Goal

条件を満たす要素だけを残す。

## Task

Positiveは順序を保ち正の値だけを返す。該当なしはnil。入力を変更しない。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `Positive([]int{-1, 3, 0, 2, 3})` | `[]int{3, 2, 3}`。順序・重複を保ち、入力は変わらない |
| `Positive([]int{0, -1})` | `nil` |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./02-data/07-slice-filter
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
