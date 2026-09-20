# Exercise: transaction

## Goal

残高と履歴の永続化を一つの単位にする。

## Task

Transferは送金とtransfer_eventsへの記録を同一transactionで行う。記録が失敗したら残高も戻す。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。
- PostgreSQLが必要です。[準備](../../docs/database.md)を済ませてください。

## Run

```bash
go test -race ./11-practical/05-transaction
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
