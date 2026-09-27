# Exercise: 11-typed-nil

## Goal

nilのポインタを入れたerrorがnilと等しくならないことを理解する。

## Task

`ValidateAge`は年齢が0以上なら`nil`、負なら`*ValidationError{Field: "age"}`を返します。現在の実装は成功時にも`err != nil`が真になります。検証には既存の`checkAge`を使い、呼び出し元が`if err != nil`で正しく判定できるように修正してください。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `err := ValidateAge(20); err == nil` | `true` |
| `ValidateAge(0)` | `nil` |
| `ValidateAge(-1)` | `errors.As`で`*ValidationError`を取り出せ、`Field`は`"age"` |

## Constraints

- `checkAge`と`ValidationError`は変更しないでください。
- テストと関数のsignatureを変更しないでください。

## Run

```bash
go run ./cmd/koans check 04-errors/11-typed-nil
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
