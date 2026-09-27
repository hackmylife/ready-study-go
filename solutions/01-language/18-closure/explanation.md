# Why?

返した関数は`Counter`の中の`count`を捕捉し、呼び出しの間も同じ変数を更新します。`next()`は`1 → 2 → 3`と進み、新しく作った`other()`は別の`count`を持つので`1`から始まります。package変数に置くと、`other()`が`4`を返してテストが失敗します。

# Alternatives

状態をstructのフィールドに置き、methodで更新する書き方もあります。状態が一つで操作も一つなら関数値で足り、操作が増えるならstructとmethodのほうが読みやすくなります。

# Idiomatic Go

関数値は状態を持てます。`http.HandlerFunc`や`slices.SortFunc`の比較関数のように、小さな振る舞いを引数として渡す場面で使います。

# 振り返り

`first := Counter()`と`second := Counter()`を作ったとき、`count`が何個存在するか説明してください。
