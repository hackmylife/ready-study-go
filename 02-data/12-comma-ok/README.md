# Exercise: comma-ok

## Goal

登録済みゼロ値と未登録を区別する。

## Task

Lookupは値とキーの存在を返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `Lookup(map[string]int{"zero": 0}, "zero")` | `(0, true)`（登録されている） |
| `Lookup(map[string]int{"zero": 0}, "missing")` | `(0, false)`（登録されていない） |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./02-data/12-comma-ok
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
