# Exercise: not-found

## Goal

DB固有の未登録を用途のerrorへ変換する。

## Task

Findは未登録だけをErrUserNotFoundへ変換する。他のDB errorは保持する。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。
- PostgreSQLが必要です。[準備](../../docs/database.md)を済ませてください。

## Run

```bash
go test ./10-database/07-not-found
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
