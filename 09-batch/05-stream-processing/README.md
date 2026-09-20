# Exercise: stream-processing

## Goal

入力を全保持せず出力へ流す。

## Task

Transformは一行ずつ大文字にし改行付きでWriterへ出力する。入力・出力のerrorを返す。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./09-batch/05-stream-processing
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
