# Exercise: error-response

## Goal

内部errorを公開用のHTTP応答へ写す。

## Task

WriteErrorはErrNotFoundを404/not_found、それ以外を500/internal_errorとしてJSONで返す。内部error文字列は公開しない。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./08-http/08-error-response
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
