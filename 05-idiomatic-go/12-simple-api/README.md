# Exercise: simple-api

## Goal

boolフラグの意味が伝わるAPIにする。

## Task

既存のFormat(name, loud)を削除し、GreetingとLoudGreetingという二つの関数へ分ける。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `Greeting("Aki")` | `"Hello, Aki"` |
| `LoudGreeting("Aki")` | `"HELLO, AKI"` |
| 変更後のコード | 旧APIの `Format(name, loud)` が存在しない |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./05-idiomatic-go/12-simple-api
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
