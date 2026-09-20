# Exercise: defer

## Goal

成功・失敗のどちらでも後処理する。

## Task

Useはworkの後にcloseを必ず一度呼ぶ。workのerrorを返す。closeは失敗しない。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./04-errors/09-defer
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
