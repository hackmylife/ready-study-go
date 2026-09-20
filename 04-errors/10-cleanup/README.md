# Exercise: cleanup

## Goal

本処理とcloseの両方のerrorを保持する。

## Task

ReadAndCloseはReaderを読み、Closeを必ず呼び、両方のerrorを保持する。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| 読み取りが `readErr`、Closeが `closeErr` で失敗するReaderを渡す | Closeを一度呼ぶ。返るerrorに対する `errors.Is` は両方の原因で `true` |
| 読み取りもCloseも成功するReader | 読み取ったデータと `nil` errorを返す |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./04-errors/10-cleanup
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
