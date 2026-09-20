# Exercise: time

## Goal

time.Durationで経過時間を扱う。

## Task

ExpiresAtは起点にTTLを加える。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| 起点が `2026-01-01 23:50 UTC`、TTLが `20*time.Minute` | `2026-01-02 00:10 UTC` |
| 同じ起点でTTLが `0` | 起点と同じ時刻 |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./03-standard-library/02-time
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
