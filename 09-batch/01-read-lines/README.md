# Exercise: read-lines

## Goal

行入力を逐次処理する。

## Task

EachLineは一行ずつvisitへ渡す。visitのerrorで中止し、読み取りerrorも返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| 内容 `"a\n\nb"` のReaderを `EachLine` に渡す | visitには順に `"a"`、`""`、`"b"` が渡る。全て成功なら戻り値は `nil` |
| 内容 `"a\nb"` に対し、最初のvisitがerrorを返す | `"b"` は処理せず、そのerrorを返す |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./09-batch/01-read-lines
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
