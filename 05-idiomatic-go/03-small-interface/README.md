# Exercise: small-interface

## Goal

利用側が必要とするmethodだけを要求する。

## Task

Sendの引数をio.Writerに絞り、文字列を書く。呼び出し元にReadやCloseを要求しない。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./05-idiomatic-go/03-small-interface
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
