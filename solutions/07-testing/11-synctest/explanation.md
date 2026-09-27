# Why?

`synctest.Test`の中では`time.After(1s)`を待つ間に仮想時計が進むため、`2s`待つテストも実時間では待ちません。時間が決定的になるので、`time.Since(start) == 2*time.Second`と正確に比べられます。待たない実装は`0s`で戻り、`time.Sleep`で待つ実装は`1500ms`の期限を無視して5回呼ぶため、それぞれ経過時間と呼び出し回数で検出できます。

# Alternatives

実時間で待つテストでも`delay`を`10ms`程度にすれば検証できますが、経過時間を`==`で比較できず、負荷の高いCIで不安定になります。時計を引数で渡す設計は05-idiomatic-go/11-explicit-dependencyで扱います。

# Idiomatic Go

`synctest`は実装に手を加えずに`time`パッケージの時計を差し替えます。bubbleの中のgoroutineが全て止まったまま進めなくなると、テストはデッドロックとして失敗します。

# 振り返り

`TestRetryStopsWhenContextIsDone`で、仮想時計が`0s → 1s → 1500ms`と進む間に、どのgoroutineが何を待っているか説明してください。
