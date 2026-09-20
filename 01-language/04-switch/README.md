# Exercise: switch

## Goal

switchで離散的な値を分類する。

## Task

DayKindは土日をweekend、平日をweekday、それ以外をinvalidとして返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `DayKind("Mon")` | `"weekday"` |
| `DayKind("Sun")` | `"weekend"` |
| `DayKind("sun")` | `"invalid"`（大文字小文字を区別） |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./01-language/04-switch
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
