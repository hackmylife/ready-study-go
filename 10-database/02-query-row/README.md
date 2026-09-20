# Exercise: query-row

## Goal

単一行の値を読む。

## Task

Nameはusersからidに一致するnameを返す。未登録はsql.ErrNoRows。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。
- PostgreSQLが必要です。[準備](../../docs/database.md)を済ませてください。

## Run

```bash
go test ./10-database/02-query-row
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
