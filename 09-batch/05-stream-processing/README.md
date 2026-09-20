# Exercise: stream-processing

## Goal

入力を全保持せず出力へ流す。

## Task

Transformは一行ずつ大文字にし改行付きでWriterへ出力する。入力・出力のerrorを返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| 内容 `"go\n猫\n"` を `Transform` に渡す | 出力は `"GO\n猫\n"`、戻り値は `nil` |
| 出力先が `io.ErrClosedPipe` を返す | 処理を中止し、その原因errorを返す |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./09-batch/05-stream-processing
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
