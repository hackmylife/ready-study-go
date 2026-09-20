# Exercise: exec

## Goal

更新件数を確認する。

## Task

Renameはnameを更新し、行がなければsql.ErrNoRowsを返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| usersにu1があり、`Rename(ctx, db, "u1", "Nao")` | 戻り値は `nil`。u1のnameが `"Nao"` になる |
| 同じ更新を存在しないidへ行う | `sql.ErrNoRows` |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。
- PostgreSQLが必要です。[準備](../../docs/database.md)を済ませてください。

## Run

```bash
go test ./10-database/05-exec
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
