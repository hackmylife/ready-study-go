# Exercise: read-csv

## Goal

CSVを業務上のレコードへ読む。

## Task

ReadOrdersはheaderがname,quantityのCSVをOrderへ読む。数量は整数とし、値域検証は後の演習に任せる。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./09-batch/02-read-csv
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
