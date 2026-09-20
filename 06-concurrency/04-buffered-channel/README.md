# Exercise: buffered-channel

## Goal

bufferを使って即時受信可能なchannelを作る。

## Task

Queueは全入力をbufferに詰めてcloseしてから返す。goroutineは不要。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./06-concurrency/04-buffered-channel
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
