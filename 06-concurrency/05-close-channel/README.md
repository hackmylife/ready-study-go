# Exercise: close-channel

## Goal

送信側がchannelの終了を所有する。

## Task

Forwardは入力を順に転送し、入力終了時に出力を閉じる。入力は閉じない。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `8, 9` を送り、閉じた入力channelを `Forward` に渡す | 出力channelで順に `8, 9` を受信し、その後に出力が閉じる |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./06-concurrency/05-close-channel
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
