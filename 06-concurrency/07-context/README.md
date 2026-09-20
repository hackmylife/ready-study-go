# Exercise: context

## Goal

contextを下流へ伝播する。

## Task

Fetchは受け取ったcontextをloadにそのまま渡す。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./06-concurrency/07-context
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
