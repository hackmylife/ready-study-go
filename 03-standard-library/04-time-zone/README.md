# Exercise: time-zone

## Goal

瞬間を保って表示タイムゾーンを変える。

## Task

InTokyoは同じ瞬間をAsia/Tokyoで返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `2026-01-01 16:00 UTC` を `InTokyo` に渡す | `2026-01-02 01:00 +0900`、locationは `Asia/Tokyo`、errorは `nil`。表している瞬間は同じ |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./03-standard-library/04-time-zone
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
