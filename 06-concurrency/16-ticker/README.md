# Exercise: 16-ticker

## Goal

Tickerとselectで、キャンセル可能な定期処理を書く。

## Task

`Poll`は`interval`ごとに`ready`を呼び、`ready`が`true`を返したら`nil`を返します。最初の呼び出しは開始から`interval`経過した時点です。`ctx`が先に終了したら、それ以上`ready`を呼ばず、待っている途中でもすぐに`ctx.Err()`を返します。`interval`は正の値とします。

## Examples

テストは`testing/synctest`の仮想時計で実行するため、実際には待ちません。時刻は`Poll`を呼んだ時点からの経過時間です。

| 入力・操作 | 期待する結果 |
|---|---|
| `interval`が`1s`、`ready`が3回目に`true` | `1s`・`2s`・`3s`に`ready`を呼び、`3s`で`nil`を返す |
| `ctx`の期限が`2500ms`、`ready`は常に`false`、`interval`が`1s` | `1s`・`2s`に`ready`を呼び、`2500ms`で`context.DeadlineExceeded`を返す |

## Constraints

- `time.NewTicker`と`select`を使い、終了時にTickerを止めてください。
- テストと関数のsignatureを変更しないでください。

## Run

```bash
go run ./cmd/koans check 06-concurrency/16-ticker
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
