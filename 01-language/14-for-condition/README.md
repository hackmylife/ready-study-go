# Exercise: 14-for-condition

## Goal

条件だけのforで反復の終了を表す。

## Task

`DigitCount`は非負整数`n`の十進表記の桁数を返します。`0`も一桁です。負数は入力されません。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `DigitCount(0)` | `1` |
| `DigitCount(7)` | `1` |
| `DigitCount(10)` | `2` |
| `DigitCount(1200)` | `4` |

## Constraints

- `for 条件 { ... }`を使ってください。文字列への変換は使いません。
- テストと関数のsignatureを変更しないでください。
- テストは振る舞いを確認します。構文の使い分けは模範解答・解説でも確認してください。

## Run

```bash
go run ./cmd/koans check 01-language/14-for-condition
```

[段階的なヒント](HINTS.md) / [ループとコレクションの案内](../../docs/loops-and-collections.md)
