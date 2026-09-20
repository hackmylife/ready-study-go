# Exercise: transaction

## Goal

残高と履歴の永続化を一つの単位にする。

## Task

Transferは送金とtransfer_eventsへの記録を同一transactionで行う。記録が失敗したら残高も戻す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| 初期残高a=100、b=50で、aからbへ10を送金する | 残高はa=`90`、b=`60`。履歴に `from_id="a", to_id="b", amount=10` が保存される |
| 履歴の保存がDB制約で失敗する | error。残高更新も取り消され、送金前の残高を保つ |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。
- PostgreSQLが必要です。[準備](../../docs/database.md)を済ませてください。

## Run

```bash
go test -race ./11-practical/05-transaction
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
