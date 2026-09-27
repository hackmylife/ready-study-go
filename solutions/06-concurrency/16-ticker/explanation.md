# Why?

`select`は`ctx.Done()`と`ticker.C`のどちらか先に準備できたほうを選びます。期限が`2500ms`なら、`1s`と`2s`のtickで`ready`を呼び、`2500ms`で`ctx.Done()`が閉じた瞬間に`context.DeadlineExceeded`を返します。`time.Sleep`で待つと、`3s`まで起きずに3回目の`ready`を呼んでしまいます。

# Alternatives

最初の確認を待たずに行いたい場合は、`for`の前に一度`ready()`を呼びます。一回だけ待つなら`time.NewTimer`や`time.After`を使います。

# Idiomatic Go

Tickerを作った関数が`defer ticker.Stop()`で止めます。goroutineやTickerは、作った側が終了条件を持ちます。テストでは`testing/synctest`の仮想時計を使うと、`3s`の待ち時間を実際に待たずに、経過時間まで正確に検証できます。

# 振り返り

`ready`の実行に`interval`より長い時間がかかったとき、次の`ready`がいつ呼ばれるか`time.Ticker`のドキュメントで確認してください。
