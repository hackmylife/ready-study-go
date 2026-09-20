# Exercise: waitgroup

## Goal

複数のgoroutineを待ち合わせる。

## Task

RunAllは全jobを並行実行し、全て終わってから戻る。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| カウンタに1を加えるjobを二つ渡す（カウンタ自体は競合しないもの） | `RunAll` から戻った時にカウンタは `2` |
| 両方とも外部からの解除を待つjobを渡す | 両方が開始でき、両方が終わるまで `RunAll` は戻らない |
| `RunAll(nil)` | 何も実行せず戻る |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./06-concurrency/02-waitgroup
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
