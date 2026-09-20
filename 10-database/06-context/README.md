# Exercise: context

## Goal

DB操作にキャンセルを伝える。

## Task

Pingは受け取ったcontextでDBを検証する。キャンセル後にBackgroundでやり直さない。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| 利用可能なDBとキャンセルされていないctxで `Ping` | `nil` |
| 同じDBとキャンセル済みctxで `Ping` | `context.Canceled` を原因に含むerror |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。
- PostgreSQLが必要です。[準備](../../docs/database.md)を済ませてください。

## Run

```bash
go test ./10-database/06-context
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
