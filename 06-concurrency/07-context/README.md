# Exercise: context

## Goal

contextを下流へ伝播する。

## Task

Fetchは受け取ったcontextをloadにそのまま渡す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| キャンセル済みctxを `Fetch` に渡し、loadが受け取ったctxのerrorを返す | loadに同じctxが届き、`Fetch` のerrorは `context.Canceled` |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./06-concurrency/07-context
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
