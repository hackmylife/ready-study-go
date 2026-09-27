# Why?

`RWMutex`の`RLock`は他の`RLock`と同時に保持できるため、`View`の`fn`が戻る前でも二つ目の`View`は`scores["go"] == 95`を読めます。`Set`は`Lock`を取り、全ての読み取りが終わるまで待ってからmapを書き換えます。mapをゼロ値のnilのままにし、最初の`Set`で作ると、`var s Scores`をそのまま使えます。

# Alternatives

読み取りが短く、書き込みと同程度の頻度なら`sync.Mutex`で十分です。`RWMutex`は読み取り中の処理が長い、または読み取りが大半を占める場合に効果があります。どちらが速いかは`go test -bench`で測ってから選びます。

# Idiomatic Go

`View`のようにロック中にcallbackを呼ぶ設計では、callbackの中で同じ`Scores`の`Set`を呼ぶと、`RLock`を保持したまま`Lock`を待つためデッドロックします。callbackに許す操作をAPIの契約として決めてください。

# 振り返り

`View`の`fn`が実行中に別のgoroutineが`Set`を呼んだとき、`Set`がいつ完了するか説明してください。
