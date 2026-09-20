# Exercise: accept-interface-return-struct

## Goal

入力はinterface、出力は具体型にする。

## Task

NewDecoderはio.Readerを受け取り*Decoderを返す。Decoder.Nextは一行を読む。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./05-idiomatic-go/06-accept-interface-return-struct
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
