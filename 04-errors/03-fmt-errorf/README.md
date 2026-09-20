# Exercise: fmt-errorf

## Goal

エラーに具体的な値を含める。

## Task

ValidateAgeは0〜150を受け付け、範囲外はage out of range: 値というerrorにする。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `ValidateAge(0)` / `ValidateAge(150)` | `nil` |
| `ValidateAge(151)` | `Error()` が `"age out of range: 151"` のerror |
| `ValidateAge(-1)` | `Error()` が `"age out of range: -1"` のerror |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./04-errors/03-fmt-errorf
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
