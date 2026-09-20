# Exercise: Payment API

## Goal

口座と送金を持つ小さなPayment APIを完成させる。

## Task

CreateAccount・Account・TransferByID・Send・Routesを実装する。[API契約](API.md)とテストを読み、作成→送金→参照を通す。DBのschema、server起動、終了処理は用意済み。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。
- PostgreSQLが必要です。[準備](../../docs/database.md)を済ませてください。

## Run

```bash
go test -race ./12-capstone/payment-api
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
