# Exercise: database-crud

## Goal

SQLでCRUDを一通り実装する。

## Task

usersの作成・取得・改名・削除を実装する。空ID/空白名を拒否する。重複作成はerror、対象がない取得・更新・削除はsql.ErrNoRows。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。
- PostgreSQLが必要です。[準備](../../docs/database.md)を済ませてください。

## Run

```bash
go test -race ./11-practical/04-database-crud
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
