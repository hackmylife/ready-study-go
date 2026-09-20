# Exercise: json-response

## Goal

JSONのheaderと本文を返す。

## Task

UserHandlerは200で{"name":"Aki"}を返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| UserHandlerへGET requestを渡す | statusは `200`、Content-Typeは `application/json`、JSON本文は `{"name":"Aki"}` |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./08-http/05-json-response
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
