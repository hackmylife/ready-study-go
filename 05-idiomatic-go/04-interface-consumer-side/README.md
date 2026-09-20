# Exercise: interface-consumer-side

## Goal

利用側で必要な依存を定義する。

## Task

Greeterが利用するNameLookupをこのpackageで定義し、Greetingを実装する。永続化方法は要求しない。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./05-idiomatic-go/04-interface-consumer-side
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
