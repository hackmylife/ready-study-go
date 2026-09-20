# Exercise: http-client

## Goal

HTTP Clientのstatusとbody寿命を扱う。

## Task

FetchはGETで200の本文を返す。他statusはerror。response bodyは必ずCloseする。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./08-http/10-http-client
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
