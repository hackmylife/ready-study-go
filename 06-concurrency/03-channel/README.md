# Exercise: channel

## Goal

channelで値を受け渡す。

## Task

Sumはchannelが閉じるまで受信して合計を返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| 入力channelに `2, 0, 5` を送り、閉じてから `Sum` に渡す | 戻り値は `7` |
| 値を送らずに閉じたchannelを渡す | `0` |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./06-concurrency/03-channel
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
