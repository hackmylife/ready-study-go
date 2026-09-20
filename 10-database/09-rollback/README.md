# Exercise: rollback

## Goal

callback失敗時にtransactionを取り消す。

## Task

WithTxはfnが成功した時だけcommitし、失敗時はrollbackする。原因を保持する。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。
- PostgreSQLが必要です。[準備](../../docs/database.md)を済ませてください。

## Run

```bash
go test ./10-database/09-rollback
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
