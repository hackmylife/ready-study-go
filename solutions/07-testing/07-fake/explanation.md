# Why?

Send(fake,"aki@example.test")はNotify("aki@example.test","Welcome")を呼びます。fakeが受け取った引数を保持すると、呼び出し契約を直接検証できます。

# Alternatives

別の書き方も、同じ入力・出力とエラー契約を満たせます。模範解答と自分のコードを比較し、読み手が追う状態や分岐が増えていないか説明してください。

# Idiomatic Go

mock frameworkを使わず小さなfakeで必要な振る舞いを表せます。

# 振り返り

テストが示す具体例を一つ選び、入力から出力まで手で追ってください。境界値のテストを一つ自分で追加してから、もう一度解いてみましょう。
