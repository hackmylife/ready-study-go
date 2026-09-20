# Exercise: 15-for-break-continue

## Goal

breakとcontinueで反復を制御する。

## Task

`SumEvenBeforeNegative`は先頭から値を読み、最初の負数の直前までにある偶数だけを合計します。奇数は飛ばします。負数以降は合計に含めません。空入力の結果は0です。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `SumEvenBeforeNegative([]int{2, 3, 4, -1, 10})` | `6` |
| `SumEvenBeforeNegative([]int{-2, 8})` | `0` |
| `SumEvenBeforeNegative(nil)` | `0` |

## Constraints

- `for { ... }`を使い、末尾到達と負数で`break`、奇数で`continue`してください。
- テストと関数のsignatureを変更しないでください。
- テストは振る舞いを確認します。構文の使い分けは模範解答・解説でも確認してください。

## Run

```bash
go run ./cmd/koans check 01-language/15-for-break-continue
```

[段階的なヒント](HINTS.md) / [ループとコレクションの案内](../../docs/loops-and-collections.md)
