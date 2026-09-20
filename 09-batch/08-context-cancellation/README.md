# Exercise: context-cancellation

## Goal

レコード間でキャンセルを観測する。

## Task

Processは各レコードの前にctxを確認しhandleへ渡す。キャンセル後に次のレコードを始めない。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./09-batch/08-context-cancellation
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
