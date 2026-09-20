# 06-concurrency

[全体の学習ガイド](../README.md)

go test -raceで検証してください。goroutineの終了条件とchannelの所有者を説明できることも学習目標です。

| 演習 | 学習目標 |
|---|---|
| [01-goroutine](01-goroutine/README.md) | goroutineを起動し完了を待つ。 |
| [02-waitgroup](02-waitgroup/README.md) | 複数のgoroutineを待ち合わせる。 |
| [03-channel](03-channel/README.md) | channelで値を受け渡す。 |
| [04-buffered-channel](04-buffered-channel/README.md) | bufferを使って即時受信可能なchannelを作る。 |
| [05-close-channel](05-close-channel/README.md) | 送信側がchannelの終了を所有する。 |
| [06-select](06-select/README.md) | 複数の通信をselectで待つ。 |
| [07-context](07-context/README.md) | contextを下流へ伝播する。 |
| [08-context-cancel](08-context-cancel/README.md) | 送信待ちのgoroutineをキャンセルで終了させる。 |
| [09-context-timeout](09-context-timeout/README.md) | 子contextの期限と後始末を管理する。 |
| [10-worker-pool](10-worker-pool/README.md) | 固定数のworkerで仕事を配る。 |
| [11-bounded-concurrency](11-bounded-concurrency/README.md) | 同時実行数の上限を守る。 |
| [12-race-condition](12-race-condition/README.md) | 共有カウンタの競合を修正する。 |
| [13-mutex](13-mutex/README.md) | 複数操作をmutexで保護する。 |
| [14-remove-goroutine](14-remove-goroutine/README.md) | 即座に待つだけのgoroutineを削る。 |
