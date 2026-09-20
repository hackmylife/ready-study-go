# Why?

NewClient()はTimeout=5秒、WithTimeout(2秒)なら2秒です。optionは具体型の設定を順番に適用する関数です。

# Alternatives

設定が少なければConfig structや通常の引数で十分です。この演習の規模では次の演習のように削る判断も合理的です。

# Idiomatic Go

optionalな設定の組み合わせが増えた場合にfunctional optionsを検討します。

# 振り返り

テストが示す具体例を一つ選び、入力から出力まで手で追ってください。境界値のテストを一つ自分で追加してから、もう一度解いてみましょう。
