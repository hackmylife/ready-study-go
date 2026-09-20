# Exercise: mutex

## Goal

複数操作をmutexで保護する。

## Task

CounterのAddとValueを並行に呼べるようにする。Counterをコピーしない。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./06-concurrency/13-mutex
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
