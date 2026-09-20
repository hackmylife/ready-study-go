# Exercise: remove-interface

## Goal

一つの具体型にしか役割がないinterfaceを削る。

## Task

Repository interfaceを削除しNewRepositoryは*MemoryRepositoryを返す。Findの動作を保つ。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./05-idiomatic-go/05-remove-interface
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
