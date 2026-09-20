# PostgreSQLの準備

Database章、PracticalのDB課題、Capstoneで使用します。普通の演習はGo標準ライブラリを使い、DB接続にだけpgxの`database/sql` driverを使います。

```bash
docker compose -p ready-study-go-koans up -d --wait
export KOANS_DATABASE_URL='postgres://koans:koans@127.0.0.1:55432/koans?sslmode=disable'
go run ./cmd/koans check 10-database/01-open
```

Docker ComposeはPostgreSQL 18を127.0.0.1:55432で起動します。この接続先は教材専用です。コンテナのデータはtmpfsにあり、停止・再作成後の保持を前提にしません。残したい学習記録はforkへ保存してください。

55432が使用中なら起動時に変更できます。

```bash
KOANS_DB_PORT=55433 docker compose -p ready-study-go-koans up -d --wait
export KOANS_DATABASE_URL='postgres://koans:koans@127.0.0.1:55433/koans?sslmode=disable'
```

テストは`internal/dbtest`でランダムな名前のschemaを作り、そのschemaだけを最後に削除します。他の演習のテストとテーブル名が同じでも干渉しません。接続ユーザーにはschema作成権限が必要です。独自のPostgreSQLを使う場合も専用の学習用databaseを用意してください。

`KOANS_DATABASE_URL`未設定時、直接の`go test`ではDBテストはSKIPします。`koans check`は未設定をエラーとして案内し、`progress`は`NEEDS_DB`と表示します。全教材のDB込み検証は次のコマンドです。

```bash
go run ./cmd/koans verify
```

学習を終えたら教材のコンテナを停止できます。

```bash
docker compose -p ready-study-go-koans down
```

Capstoneのserverは同じ接続設定を使い、接続先に`schema.sql`を適用します。テストの一時schemaとは別に、接続先のsearch_path上へテーブルを作ります。
