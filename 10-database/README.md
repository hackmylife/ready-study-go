# 10-database

[全体の学習ガイド](../README.md)

| 演習 | 学習目標 |
|---|---|
| [01-open](01-open/README.md) | database/sqlの接続poolを開き接続を確認する。 |
| [02-query-row](02-query-row/README.md) | 単一行の値を読む。 |
| [03-query](03-query/README.md) | 複数行を読み、Rowsを閉じる。 |
| [04-scan](04-scan/README.md) | SQL NULLとゼロ値を区別する。 |
| [05-exec](05-exec/README.md) | 更新件数を確認する。 |
| [06-context](06-context/README.md) | DB操作にキャンセルを伝える。 |
| [07-not-found](07-not-found/README.md) | DB固有の未登録を用途のerrorへ変換する。 |
| [08-transaction](08-transaction/README.md) | 複数更新を一つのtransactionにする。 |
| [09-rollback](09-rollback/README.md) | callback失敗時にtransactionを取り消す。 |
| [10-transfer](10-transfer/README.md) | 残高移動の原子性と並行実行を守る。 |
