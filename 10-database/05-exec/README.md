# Exercise: exec

## Goal

更新件数を確認する。

## Task

Renameはnameを更新し、行がなければsql.ErrNoRowsを返す。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。
- PostgreSQLが必要です。[準備](../../docs/database.md)を済ませてください。

## Run

```bash
go test ./10-database/05-exec
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
