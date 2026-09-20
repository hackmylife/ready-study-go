# Exercise: context

## Goal

HTTP requestのcontextを下流へ伝える。

## Task

Handlerはloadにr.Contextを渡す。成功は本文200、失敗は503。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| loadが `("ok", nil)` を返す | 応答はstatus `200`、本文 `ok` |
| requestのctxがキャンセル済みで、loadがctxのerrorを返す | loadに同じctxが届き、応答は `503` |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./08-http/12-context
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
