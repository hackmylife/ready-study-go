# Exercise: context-timeout

## Goal

子contextの期限と後始末を管理する。

## Task

WithTimeoutは期限付きcontextでworkを実行し、戻る時にcancelする。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./06-concurrency/09-context-timeout
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
