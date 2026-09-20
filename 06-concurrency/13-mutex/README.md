# Exercise: mutex

## Goal

複数操作をmutexで保護する。

## Task

CounterのAddとValueを並行に呼べるようにする。Counterをコピーしない。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| ゼロ値のCounterへ、並行に合計10000回 `Add(1)` を呼ぶ | 全て完了した後の `Value()` は `10000` |
| 更新中に別goroutineから `Value()` を呼ぶ | race detectorが競合を報告しない |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./06-concurrency/13-mutex
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
