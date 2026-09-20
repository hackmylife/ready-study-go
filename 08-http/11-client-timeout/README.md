# Exercise: client-timeout

## Goal

Client全体のtimeoutを設定する。

## Task

NewClientは正のtimeoutを設定した新しいClientを返す。既定の共有Clientは変更しない。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./08-http/11-client-timeout
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
