# Exercise: map-create

## Goal

書き込み可能なmapを作る。

## Task

Indexは名前をキー、最後に現れた添字を値にする。空入力でも書き込めるmapを返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `Index([]string{"a", "b", "a"})` | `map[string]int{"a": 2, "b": 1}` |
| `Index(nil)` | 空のmap。返されたmapに新しいキーを書き込める |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./02-data/10-map-create
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
