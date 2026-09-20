# Exercise: pointers

## Goal

ポインタを介して呼び出し元の値を更新する。

## Task

Incrementは値を増やしtrueを返す。nilならfalseを返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `n := 4` として `Increment(&n)` を呼ぶ | 戻り値は `true`、呼び出し後の `n` は `5` |
| `Increment(nil)` | `false`。panicしない |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./01-language/09-pointers
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
