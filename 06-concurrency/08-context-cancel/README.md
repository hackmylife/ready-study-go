# Exercise: context-cancel

## Goal

送信待ちのgoroutineをキャンセルで終了させる。

## Task

Generateは0からの連番を送り、ctx終了時に出力を閉じる。受信者がいなくても終了できること。

## Examples

キャンセルと同時に送信可能になった値を受信する場合があります。キャンセル後の送信件数を固定せず、終了することを確認します。

| 入力・操作 | 期待する結果 |
|---|---|
| `Generate(ctx)` の出力を受信する | 先頭から `0, 1, 2, …` |
| ctxをキャンセルする | 生成処理が終了し、出力channelが閉じる |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./06-concurrency/08-context-cancel
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
