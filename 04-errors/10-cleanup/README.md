# Exercise: cleanup

## Goal

本処理とcloseの両方のerrorを保持する。

## Task

ReadAndCloseはReaderを読み、Closeを必ず呼び、両方のerrorを保持する。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./04-errors/10-cleanup
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
