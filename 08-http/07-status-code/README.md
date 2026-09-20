# Exercise: status-code

## Goal

statusとLocationでリソース作成を伝える。

## Task

Createdは201、Location: /users/u1、空本文を返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| Createdへ `POST /users` を渡す | statusは `201`、Locationは `/users/u1`、本文は空 |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./08-http/07-status-code
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
