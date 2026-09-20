# Exercise: context-cancel

## Goal

送信待ちのgoroutineをキャンセルで終了させる。

## Task

Generateは0からの連番を送り、ctx終了時に出力を閉じる。受信者がいなくても終了できること。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./06-concurrency/08-context-cancel
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
