# Exercise: transaction

## Goal

複数更新を一つのtransactionにする。

## Task

RenameBothはu1とu2をまとめて改名する。二つ目が失敗したら一つ目も元に戻る。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。
- PostgreSQLが必要です。[準備](../../docs/database.md)を済ませてください。

## Run

```bash
go test ./10-database/08-transaction
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
