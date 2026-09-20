# Exercise: select

## Goal

複数の通信をselectで待つ。

## Task

Receiveは入力の値かcontextの終了を待つ。閉じた入力はio.EOF。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./06-concurrency/06-select
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
