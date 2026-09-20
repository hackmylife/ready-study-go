# Exercise: bounded-concurrency

## Goal

同時実行数の上限を守る。

## Task

Mapを実装し、workの同時実行数をlimit以下にする。順序維持と終了条件はworker-pool演習と同じ。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./06-concurrency/11-bounded-concurrency
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
