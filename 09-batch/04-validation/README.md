# Exercise: validation

## Goal

レコードの妥当性を検証する。

## Task

Validateは空白だけのNameと0以下のQuantityを拒否する。入力自体は変更しない。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `Validate(Order{Name: "tea", Quantity: 1})` | `nil` |
| 名前が `" "`、数量は `1` | error |
| 名前が `"tea"`、数量は `0` または `-2` | error |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./09-batch/04-validation
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
