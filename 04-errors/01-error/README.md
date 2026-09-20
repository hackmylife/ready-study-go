# Exercise: error

## Goal

成功と失敗をerrorで表す。

## Task

Sqrtは非負の平方根を返し、負数ではerrorを返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `Sqrt(9)` | `(3, nil)` |
| `Sqrt(0)` | `(0, nil)` |
| `Sqrt(-1)` | errorが非nil |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./04-errors/01-error
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
