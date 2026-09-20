# Exercise: csv

## Goal

CSVの引用符と区切りを正しく扱う。

## Task

ReadCSVはCSVをレコードとして読む。列数は最初の行に合わせる。

## Examples

入力欄の `\n` は改行を表します。

| 入力・操作 | 期待する結果 |
|---|---|
| CSV `name,note\nAki,"hello, Go"\n` を `ReadCSV` に渡す | `[][]string{{"name", "note"}, {"Aki", "hello, Go"}}` と `nil` error |
| CSV `a,b\nonly-one\n` | 列数が違うのでerror |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./03-standard-library/11-csv
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
