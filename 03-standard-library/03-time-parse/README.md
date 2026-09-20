# Exercise: time-parse

## Goal

Goのlayoutで日時をパースする。

## Task

ParseDateはYYYY-MM-DDをUTCの00:00として読む。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `ParseDate("2024-02-29")` | `2024-02-29 00:00 UTC` と `nil` error |
| `ParseDate("2023-02-29")` | 存在しない日付なのでerror |
| `ParseDate("2024/02/29")` | 形式が違うのでerror |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./03-standard-library/03-time-parse
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
