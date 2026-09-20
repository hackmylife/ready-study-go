# Exercise: Payment API

## Goal

口座と送金を持つ小さなPayment APIを完成させる。

## Task

CreateAccount・Account・TransferByID・Send・Routesを実装する。[API契約](API.md)とテストを読み、作成→送金→参照を通す。DBのschema、server起動、終了処理は用意済み。

## Examples

IDは説明用の例です。実際には作成応答で返ったIDを使います。各操作の詳しい契約は [API.md](API.md) を参照してください。

| 入力・操作 | 期待する結果 |
|---|---|
| `POST /accounts` に `{"name":"Aki","balance":100}` | `201`。例: `{"id":1,"name":"Aki","balance":100}`、Locationは `/accounts/1` |
| 口座1の残高100、口座2の残高50で、キー `order-1` と本文 `{"from_id":1,"to_id":2,"amount":30}` を `POST /transfers` へ送る | `201`。残高は口座1が `70`、口座2が `80`。送金履歴が作られる |
| 同じキーと同じ送金内容を再送する | `200` で同じ送金を返す。残高は `70` と `80` のまま |
| 同じキーでamountを31に変えて再送する | `409` と `{"error":"idempotency_conflict"}` |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。
- PostgreSQLが必要です。[準備](../../docs/database.md)を済ませてください。

## Run

```bash
go test -race ./12-capstone/payment-api
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
