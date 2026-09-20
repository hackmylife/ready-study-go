# Exercise: graceful-shutdown

## Goal

処理中のrequestを待ち、期限超過時は接続を閉じる。

## Task

Stopは新規受付を止め、timeoutまで待つ。期限超過時はCloseも呼び、元のerrorを返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| 処理中requestがないserverを `Stop` する | 新規受付を止めて `nil` を返す |
| 処理中requestがtimeout内に完了する | その応答が完了するまで待ち、`nil` を返す |
| 処理中requestが残った状態でtimeoutを `0` にする | 接続を閉じ、`context.DeadlineExceeded` を原因に含むerrorを返す |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./08-http/13-graceful-shutdown
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
