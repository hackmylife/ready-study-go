# Exercise: errors-is

## Goal

wrapされたsentinel errorを判定する。

## Task

Missingはos.ErrNotExistを原因に含むかを返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `Missing(os.ErrNotExist)` | `true` |
| `os.ErrNotExist` をwrapしたerrorを `Missing` に渡す | `true` |
| 同じ文面で別に作ったerror、または `nil` | `false` |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./04-errors/05-errors-is
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
