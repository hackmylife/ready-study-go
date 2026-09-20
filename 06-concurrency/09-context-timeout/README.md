# Exercise: context-timeout

## Goal

子contextの期限と後始末を管理する。

## Task

WithTimeoutは期限付きcontextでworkを実行し、戻る時にcancelする。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| timeoutを `0` にしてworkへctxを渡す | workから見えるctxのerrorは `context.DeadlineExceeded` |
| timeoutを `time.Hour` にしてworkがすぐ成功する | 戻り値は `nil`。関数から戻った時点で子ctxはキャンセル済み |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./06-concurrency/09-context-timeout
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
