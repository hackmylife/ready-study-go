# Exercise: json-decoder

## Goal

Readerから厳密に一つのJSON値を読む。

## Task

Decodeは未知フィールドと後続のJSON値を拒否する。空入力もエラー。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| 内容が `{"name":"Aki"}` のReaderを `Decode` に渡す | `Request{Name: "Aki"}` と `nil` error |
| 内容が `{"name":"x"} {}` のReader | JSON値が続くのでerror |
| 内容が `{"other":1}` のReader、または空のReader | error |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./03-standard-library/07-json-decoder
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
