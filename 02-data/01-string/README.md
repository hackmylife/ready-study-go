# Exercise: string

## Goal

stringの長さがバイト数であることを理解する。

## Task

文字列 `s` が何バイトで表されるかを返す `ByteLength` を実装してください。見た目の文字数とは異なる場合があります。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `ByteLength("Go")` | `2` バイト |
| `ByteLength("猫")` | `3` バイト |
| `ByteLength("")` | `0` バイト |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./02-data/01-string
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
