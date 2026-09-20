# Exercise: explicit-dependency

## Goal

時計を明示的な依存として渡す。

## Task

IssuedAtは引数nowを一度呼びRFC3339で返す。現在時刻を直接読まない。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./05-idiomatic-go/11-explicit-dependency
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
