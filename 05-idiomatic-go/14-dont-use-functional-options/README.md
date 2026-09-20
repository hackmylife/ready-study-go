# Exercise: dont-use-functional-options

## Goal

設定が一つのAPIから不要なoptionを削る。

## Task

OptionとWithPrefixを削除し、NewFormatter(prefix string) *Formatterに変更する。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./05-idiomatic-go/14-dont-use-functional-options
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
