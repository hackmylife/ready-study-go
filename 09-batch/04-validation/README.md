# Exercise: validation

## Goal

レコードの妥当性を検証する。

## Task

Validateは空白だけのNameと0以下のQuantityを拒否する。入力自体は変更しない。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./09-batch/04-validation
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
