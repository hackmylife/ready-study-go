# Exercise: explicit-dependency

## Goal

時計を明示的な依存として渡す。

## Task

IssuedAtは引数nowを一度呼びRFC3339で返す。現在時刻を直接読まない。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `2026-01-02 03:04:05 UTC` を返す時計を `IssuedAt` に渡す | `"2026-01-02T03:04:05Z"`。渡された時計を一度呼ぶ |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./05-idiomatic-go/11-explicit-dependency
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
