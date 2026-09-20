# Exercise: for

## Goal

forで反復と累積を書く。

## Task

SumToは1からnまでの和を返す。n <= 0なら0を返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `SumTo(4)` | `10`（1から4までの合計） |
| `SumTo(1)` | `1` |
| `SumTo(0)` または `SumTo(-2)` | `0` |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./01-language/05-for
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
