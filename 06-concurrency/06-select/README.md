# Exercise: select

## Goal

複数の通信をselectで待つ。

## Task

Receiveは入力の値かcontextの終了を待つ。閉じた入力はio.EOF。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| 入力channelから `7` を受信できる状態で `Receive` を呼ぶ | `(7, nil)` |
| 値が残っていない閉じたchannelを渡す | errorは `io.EOF` |
| 空のchannelと、キャンセル済みのctxを渡す | errorは `context.Canceled` |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./06-concurrency/06-select
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
