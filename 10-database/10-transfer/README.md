# Exercise: transfer

## Goal

残高移動の原子性と並行実行を守る。

## Task

Transferは整数の金額を送金する。同一口座・0以下・残高不足・未登録を拒否し、両口座を一定順でlockする。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。
- PostgreSQLが必要です。[準備](../../docs/database.md)を済ませてください。

## Run

```bash
go test ./10-database/10-transfer
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
