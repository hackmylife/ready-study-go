# Exercise: package-design

## Goal

責務に沿って公開APIを絞る。

## Task

税計算の内部関数を非公開にし、InvoiceTotalだけを公開する。整数の金額で切り捨て計算する。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `InvoiceTotal(101, 10)`（金額101、税率10%） | `111`（小数部分を切り捨てた税込金額） |
| `InvoiceTotal(0, 10)` | `0` |
| 変更後の公開関数 | `InvoiceTotal` だけ。税計算の内部関数は非公開 |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./05-idiomatic-go/09-package-design
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
