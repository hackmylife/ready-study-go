# Exercise: errors-as

## Goal

errorから型付き情報を取り出す。

## Task

RetryDelayはwrapされたRetryErrorから待ち時間と存在可否を返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `RetryError{Delay: 3*time.Second}` へのポインタをwrapしたerrorを渡す | `(3*time.Second, true)` |
| 無関係のerror、または `nil` を渡す | `(0, false)` |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./04-errors/06-errors-as
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
