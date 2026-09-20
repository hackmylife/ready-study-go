# Exercise: if

## Goal

ifで境界を含む条件分岐を書く。

## Task

数値 `value` が下限 `min` から上限 `max` の範囲に収まるようにする `Clamp` を実装してください。範囲より小さければ下限、大きければ上限、範囲内なら元の値を返します。`min <= max` を前提とします。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `Clamp(-1, 0, 10)` | `0`（下限） |
| `Clamp(4, 0, 10)` | `4`（元の値） |
| `Clamp(11, 0, 10)` | `10`（上限） |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./01-language/03-if
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
