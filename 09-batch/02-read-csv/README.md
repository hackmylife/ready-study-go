# Exercise: read-csv

## Goal

CSVを業務上のレコードへ読む。

## Task

ReadOrdersはheaderがname,quantityのCSVをOrderへ読む。数量は整数とし、値域検証は後の演習に任せる。

## Examples

`\n` は改行です。この演習では負の数量も読み込み、値域の検証は行いません。

| 入力・操作 | 期待する結果 |
|---|---|
| CSV `name,quantity\n"tea, green",2\ncoffee,-1\n` | `[]Order{{Name: "tea, green", Quantity: 2}, {Name: "coffee", Quantity: -1}}` と `nil` error |
| CSV `name,quantity\ntea,x\n` | 数量を整数にできないのでerror |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./09-batch/02-read-csv
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
