# Exercise: rollback

## Goal

callback失敗時にtransactionを取り消す。

## Task

WithTxはfnが成功した時だけcommitし、失敗時はrollbackする。原因を保持する。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| 初期残高100の口座をcallback内で0に更新した後、callbackが `cause` を返す | `WithTx` は原因 `cause` を保持したerrorを返す。DBの残高は `100` のまま |
| 初期残高100をcallback内で80に更新し、callbackが成功する | 戻り値は `nil`、DBの残高は `80` |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。
- PostgreSQLが必要です。[準備](../../docs/database.md)を済ませてください。

## Run

```bash
go test ./10-database/09-rollback
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
