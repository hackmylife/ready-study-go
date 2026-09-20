# Exercise: goroutine

## Goal

goroutineを起動し完了を待つ。

## Task

Startはworkを別goroutineで呼び、完了時に閉じるchannelを返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| 終了を待たされているworkを `Start` に渡す | `Start` 自体はchannelを返す。work実行中はそのchannelが閉じていない |
| workが完了する | 返したchannelが閉じ、呼び出し元が完了を知ることができる |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./06-concurrency/01-goroutine
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
