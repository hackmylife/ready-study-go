# Exercise: error-handling

## Goal

行番号をerrorへ追加し原因を保持する。

## Task

Processは一行ずつhandleへ渡し、失敗時にline N: 原因でwrapする。行番号は1から。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./09-batch/06-error-handling
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
