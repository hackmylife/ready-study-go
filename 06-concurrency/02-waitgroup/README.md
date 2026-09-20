# Exercise: waitgroup

## Goal

複数のgoroutineを待ち合わせる。

## Task

RunAllは全jobを並行実行し、全て終わってから戻る。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./06-concurrency/02-waitgroup
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
