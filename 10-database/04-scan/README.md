# Exercise: scan

## Goal

SQL NULLとゼロ値を区別する。

## Task

Emailはnullableなemailを値と存在可否で返す。空文字は存在する値。未登録行はerror。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| 指定したcontactのemailがSQLの `NULL` | `("", false, nil)` |
| emailが空文字 | `("", true, nil)` |
| emailが `"a@example.test"` | `("a@example.test", true, nil)` |
| 指定したcontactがない | errorは `sql.ErrNoRows` |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。
- PostgreSQLが必要です。[準備](../../docs/database.md)を済ませてください。

## Run

```bash
go test ./10-database/04-scan
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
