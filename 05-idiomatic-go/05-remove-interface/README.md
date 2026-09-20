# Exercise: remove-interface

## Goal

一つの具体型にしか役割がないinterfaceを削る。

## Task

Repository interfaceを削除しNewRepositoryは*MemoryRepositoryを返す。Findの動作を保つ。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `NewRepository()` で取得した値の型 | `*MemoryRepository` |
| その値で `Find("u1")` / `Find("missing")` | それぞれ `("Aki", true)` / `("", false)` |
| 変更後のコード | `Repository` interfaceが存在しない |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./05-idiomatic-go/05-remove-interface
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
