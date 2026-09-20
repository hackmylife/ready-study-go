# Exercise: package-design

## Goal

責務に沿って公開APIを絞る。

## Task

税計算の内部関数を非公開にし、InvoiceTotalだけを公開する。整数の金額で切り捨て計算する。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./05-idiomatic-go/09-package-design
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
