# Exercise: http-client

## Goal

HTTP Clientのstatusとbody寿命を扱う。

## Task

FetchはGETで200の本文を返す。他statusはerror。response bodyは必ずCloseする。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| FetchのGETに対し、サーバーが200と本文 `Go` を返す | 戻り値は `("Go", nil)`。response bodyをCloseする |
| サーバーが503を返す | errorを返す。この場合もbodyをCloseする |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./08-http/10-http-client
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
