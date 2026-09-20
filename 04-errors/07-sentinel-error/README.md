# Exercise: sentinel-error

## Goal

呼び出し元が比較できるerrorを公開する。

## Task

Popは先頭を返す。空入力ではErrEmptyを返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `Pop([]int{0, 4})` | 値は `0`、残りは `[]int{4}`、errorは `nil` |
| `Pop(nil)` | errorは `ErrEmpty` |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./04-errors/07-sentinel-error
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
