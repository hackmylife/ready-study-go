# Exercise: query

## Goal

複数行を読み、Rowsを閉じる。

## Task

Namesはid順に全ユーザー名を返す。空なら非nilの空slice。Rows.Errを確認する。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| usersに `u1/Aki` と `u2/Ren` がある状態で `Names(ctx, db)` | `[]string{"Aki", "Ren"}` と `nil` error（id順） |
| usersが空 | nilではない空sliceと `nil` error |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。
- PostgreSQLが必要です。[準備](../../docs/database.md)を済ませてください。

## Run

```bash
go test ./10-database/03-query
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
