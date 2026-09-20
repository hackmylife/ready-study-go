# Exercise: open

## Goal

database/sqlの接続poolを開き接続を確認する。

## Task

Connectはpgx driverでDBを開きPingContextする。失敗時はpoolを閉じる。成功時のCloseは呼び出し元が担当する。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。
- PostgreSQLが必要です。[準備](../../docs/database.md)を済ませてください。

## Run

```bash
go test ./10-database/01-open
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
