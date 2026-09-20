# Exercise: worker-pool

## Goal

固定数のworkerで仕事を配る。

## Task

Mapを実装する。入力順を保持し、最初のerrorで中止して全workerを待つ。workはcontextに協調して終了する。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| 入力 `[]int{3, 1, 2}`、並行実行の上限 `2`、workは値を二倍にする | `[]int{6, 2, 4}` と `nil` error。出力は入力順、同時実行は上限以内 |
| 途中のworkがerrorを返す | その原因errorを返し、全workerが終了してから戻る |
| 並行実行の上限が `0` | error |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./06-concurrency/10-worker-pool
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
