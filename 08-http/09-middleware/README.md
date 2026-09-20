# Exercise: middleware

## Goal

handlerを包んで共通処理を合成する。

## Task

RequestIDはX-Request-IDを応答へ設定してnextを呼ぶ。未指定ならunknown。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| requestに `X-Request-ID: r-1` があり、nextは204を返す | nextを一度呼び、応答は204と `X-Request-ID: r-1` |
| 同じnextで、requestにX-Request-IDがない | 応答は204と `X-Request-ID: unknown` |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./08-http/09-middleware
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
