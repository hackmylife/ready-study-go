# Exercise: graceful-shutdown

## Goal

処理中のrequestを待ち、期限超過時は接続を閉じる。

## Task

Stopは新規受付を止め、timeoutまで待つ。期限超過時はCloseも呼び、元のerrorを返す。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./08-http/13-graceful-shutdown
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
