# Exercise: race-condition

## Goal

共有カウンタの競合を修正する。

## Task

Countをrace detectorが報告しない実装に直す。n >= 0。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./06-concurrency/12-race-condition
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
