# Exercise: concurrent-processing

## Goal

バッチ処理に上限付き並行実行を組み込む。

## Task

Mapでレコードの数値変換を並行に行う。順序・キャンセル・error・worker上限を守る。

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
go test -race ./09-batch/07-concurrent-processing
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
