# Exercise: io-writer

## Goal

出力先をio.Writerとして受け取る。

## Task

WriteGreetingはHello, 名前と改行を書き、書き込みエラーを返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| 空のbufferと名前 `"猫"` を `WriteGreeting` に渡す | bufferの内容は `"Hello, 猫\n"`、errorは `nil` |
| 書き込みで `errWrite` を返すWriterに書く | 返るerrorに `errWrite` が保持される |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./03-standard-library/13-io-writer
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
