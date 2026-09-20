# Exercise: context-cancellation

## Goal

レコード間でキャンセルを観測する。

## Task

Processは各レコードの前にctxを確認しhandleへ渡す。キャンセル後に次のレコードを始めない。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `[]string{"a", "b"}` の処理中に、aのhandleがctxをキャンセルする | handleへ渡るのは `"a"` だけ。戻り値は `context.Canceled` |
| キャンセルされていないctxで空のrowsを渡す | 何も処理せず、`nil` を返す |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./09-batch/08-context-cancellation
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
