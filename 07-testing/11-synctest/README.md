# Exercise: synctest

## Goal

testing/synctestの仮想時計で、時間に依存する処理を実際に待たずに検証する。

## Task

`Retry`のテストを書いてください。`Retry`は`fn`が`nil`を返すまで最大`attempts`回呼び出し、失敗するたびに`delay`だけ待ってから次を呼びます。待っている間に`ctx`が終了したら、すぐに`ctx.Err()`を返します。全て失敗したら最後のerrorを返します。`attempts`は1以上とします。

## Examples

以下は、あなたが書くテストで検証する対象関数の振る舞いです。テストコード自体を実装してください。時刻は`Retry`を呼んだ時点からの経過時間です。

| 入力・操作 | 期待する結果 |
|---|---|
| `attempts=5`、`delay=1s`、`fn`が3回目で成功 | `fn`を3回呼び、`2s`で`nil`を返す |
| `attempts=3`、`fn`が`first`・`second`・`third`を順に返す | `fn`を3回呼び、`third`を返す |
| `attempts=5`、`delay=1s`、`ctx`の期限が`1500ms`、`fn`は常に失敗 | `fn`を2回呼び、`1500ms`で`context.DeadlineExceeded`を返す |

## Constraints

- exercise_test.goを編集する。exercise.goとmutants/は変更しない。
- `synctest.Test`の中で`Retry`を呼び、呼び出し回数に加えて経過時間も検証する。
- go run ./cmd/koans check 07-testing/11-synctest で欠陥実装を見抜けるか確認する。

## Run

```bash
go run ./cmd/koans check 07-testing/11-synctest
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
