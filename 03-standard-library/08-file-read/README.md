# Exercise: file-read

## Goal

ファイルを読み、OSのエラーを返す。

## Task

Readはファイル全体を文字列で返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| 内容が `"Go\n猫"` のファイルのパスを `Read` に渡す | 戻り値は `"Go\n猫"`、errorは `nil` |
| 存在しないファイルのパスを渡す | `errors.Is(err, os.ErrNotExist)` が `true` |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./03-standard-library/08-file-read
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
