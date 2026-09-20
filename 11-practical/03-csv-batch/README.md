# Exercise: csv-batch

## Goal

CSVの入力から検証・変換・出力まで作る。

## Task

Runはname,quantityを読み、名前を正規化してpricesから金額を計算しname,quantity,totalを出力する。不正行で中止。先に出た行は残る。整数overflowとFlushのerrorも扱う。

## Examples

`\n` は改行です。単価表は関数の引数として渡します。

| 入力・操作 | 期待する結果 |
|---|---|
| 入力CSV `name,quantity\n TEA ,2\ncoffee,3\n`、単価はtea=100・coffee=150 | 出力CSVは `name,quantity,total\ntea,2,200\ncoffee,3,450\n`。戻り値は `nil` |
| 入力行のquantityが `0` | その行でerrorになり、以降を処理しない |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./11-practical/03-csv-batch
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
