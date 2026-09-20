# Exercise: remove-constructor

## Goal

不要なconstructorを削る。

## Task

NewCounterを削除し、ゼロ値のCounterでAddとValueを使えるようにする。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `var c Counter`、`c.Add(2)`、`c.Value()` | `2` |
| 変更後のコード | `NewCounter` が存在せず、上の操作がそのまま使える |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./05-idiomatic-go/02-remove-constructor
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
