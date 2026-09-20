# Exercise: goroutine

## Goal

goroutineを起動し完了を待つ。

## Task

Startはworkを別goroutineで呼び、完了時に閉じるchannelを返す。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./06-concurrency/01-goroutine
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
