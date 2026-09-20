# Exercise: errors-new

## Goal

固定メッセージのerrorを作る。

## Task

ValidateNameは空文字だけをエラーにする。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `ValidateName("Aki")` | `nil` |
| `ValidateName("")` | `Error()` が `"name is required"` のerror |
| `ValidateName(" ")` | `nil`（この課題では空文字だけを拒否） |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./04-errors/02-errors-new
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
