# Exercise: not-found

## Goal

DB固有の未登録を用途のerrorへ変換する。

## Task

Findは未登録だけをErrUserNotFoundへ変換する。他のDB errorは保持する。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| usersに `u1/Aki` があり、`Find(ctx, db, "u1")` | `("Aki", nil)` |
| `Find(ctx, db, "missing")` | errorは `ErrUserNotFound` |
| キャンセル済みctxで検索する | `context.Canceled` を原因に含むerror。未登録には変換しない |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。
- PostgreSQLが必要です。[準備](../../docs/database.md)を済ませてください。

## Run

```bash
go test ./10-database/07-not-found
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
