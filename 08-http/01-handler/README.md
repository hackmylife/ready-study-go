# Exercise: handler

## Goal

http.HandlerFuncで応答を書く。

## Task

HelloはHello, Goと改行を返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| Helloへ `GET /` を渡す | statusは `200`、本文は `"Hello, Go\n"` |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./08-http/01-handler
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
