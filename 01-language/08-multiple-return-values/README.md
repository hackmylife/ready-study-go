# Exercise: multiple-return-values

## Goal

結果と成功可否を複数の戻り値で返す。

## Task

Divideは整数の商と除算可能かを返す。ゼロ除算では(0, false)。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `Divide(7, 2)` | `(3, true)`（整数の商） |
| `Divide(0, 2)` | `(0, true)`（成功した結果が0） |
| `Divide(7, 0)` | `(0, false)`（除算できない） |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./01-language/08-multiple-return-values
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
