# Exercise: transform

## Goal

値を変更せず業務変換する。

## Task

NormalizeはNameの前後空白を除き小文字にし、Quantityを保って返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `Normalize(Order{Name: "  TEA ", Quantity: 3})` | `Order{Name: "tea", Quantity: 3}`。元のOrderは変わらない |
| `Normalize(Order{})` | `Order{}` |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./09-batch/03-transform
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
