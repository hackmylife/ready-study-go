# Exercise: remove-goroutine

## Goal

即座に待つだけのgoroutineを削る。

## Task

Upperからgoroutineとchannelを削除し、同じ結果を同期的に返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `Upper("Go")` | `"GO"` |
| `Upper("猫")` | `"猫"` |
| 変更後のコード | 同じ結果を保ち、goroutineとchannelが存在しない |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./06-concurrency/14-remove-goroutine
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
