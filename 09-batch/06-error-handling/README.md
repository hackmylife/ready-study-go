# Exercise: error-handling

## Goal

行番号をerrorへ追加し原因を保持する。

## Task

Processは一行ずつhandleへ渡し、失敗時にline N: 原因でwrapする。行番号は1から。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| 内容 `"ok\nbad\nnever"` で、`bad` の処理が `"invalid quantity"` のerrorを返す | 返るerrorの表示は `"line 2: invalid quantity"`。元の原因を保持し、`never` は処理しない |
| 空のReader | handleを呼ばず、`nil` を返す |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./09-batch/06-error-handling
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
