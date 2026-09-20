# Exercise: map-count

## Goal

mapのゼロ値を集計に使う。

## Task

Countsは各単語の出現数を返す。大文字小文字は区別する。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `Counts([]string{"Go", "go", "Go"})` | `map[string]int{"Go": 2, "go": 1}` |
| `Counts(nil)` | 空のmap |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./02-data/13-map-count
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
