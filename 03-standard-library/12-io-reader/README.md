# Exercise: io-reader

## Goal

具体的な入力元に依存せず読む。

## Task

ReadAllはReaderを最後まで読み、途中のエラーも返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| 内容が `"猫"` のReaderを `ReadAll` に渡す | バイト列の内容は `"猫"`、errorは `nil` |
| 最後のReadで `"Go"` のデータと `io.EOF` を同時に返すReader | バイト列の内容は `"Go"`、errorは `nil`。最後のデータも含む |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./03-standard-library/12-io-reader
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
