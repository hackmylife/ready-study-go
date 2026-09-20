# Exercise: custom-error

## Goal

入力に結びつく情報をerror型で返す。

## Task

FieldError.ErrorとValidateを実装する。空のemailをFieldErrorで拒否する。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `Validate("")` | `*FieldError`。`Field` は `"email"`、`Reason` は `"required"`、表示は `"email: required"` |
| `Validate("a@example.test")` | `nil` |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./04-errors/08-custom-error
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
