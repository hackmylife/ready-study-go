# Exercise: structs

## Goal

関連する値をstructにまとめる。

## Task

NewUserでNameとAgeを設定したUserを返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `NewUser("Aki", 21)` | `User{Name: "Aki", Age: 21}` |
| `NewUser("", 0)` | `User{Name: "", Age: 0}` |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./01-language/10-structs
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
