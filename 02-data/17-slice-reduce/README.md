# Exercise: 17-slice-reduce

## Goal

初期値から順番に値を畳み込む。

## Task

`Reduce`はinitialを累積値にして、先頭からf(累積値, 要素)を呼び、最後の累積値を返します。各要素に一度ずつ適用します。空入力はinitialをそのまま返し、fを呼びません。

Tは要素型、Uは累積値の型です。fはnilではなく、渡された値の参照先を変更しないものとします。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `Reduce([]int{1, 2, 3}, 10, add)（addは加算する関数）` | `16` |
| `Reduce([]string{"a", "b"}, "start", join)（joinは間に/を入れる関数）` | `"start/a/b"` |
| `Reduce([]int(nil), 10, add)` | `10` |

## Constraints

- 入力順を守ってください。加算以外の操作でも使える実装にします。
- テストと関数のsignatureを変更しないでください。
- テストは振る舞いを確認します。構文の使い分けは模範解答・解説でも確認してください。

## Run

```bash
go run ./cmd/koans check 02-data/17-slice-reduce
```

[段階的なヒント](HINTS.md) / [ループとコレクションの案内](../../docs/loops-and-collections.md)
