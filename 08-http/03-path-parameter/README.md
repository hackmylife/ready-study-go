# Exercise: path-parameter

## Goal

PathValueでpathの値を読む。

## Task

RoutesはGET /users/{id}に対してidを返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `GET /users/u1` | statusは `200`、本文は `u1` |
| `GET /users/alice` | statusは `200`、本文は `alice` |
| `GET /users/a/extra` | `404` |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./08-http/03-path-parameter
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
