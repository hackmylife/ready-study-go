# Exercise: range

## Goal

rangeで値を列挙する。

## Task

CountPositiveは正の整数の個数を返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `CountPositive([]int{-1, 0, 3, 4})` | `2`（3と4が正の数） |
| `CountPositive([]int{-3})` | `0` |
| `CountPositive(nil)` | `0` |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./01-language/06-range
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
