# Why?

`return checkAge(age)`は`*ValidationError`型のnilを`error`へ変換します。変換後のinterfaceは`(型: *ValidationError, 値: nil)`を持つため、`ValidateAge(20) != nil`が真になります。具体型のまま`err != nil`を判定し、成功時は型を持たない`nil`を返すと、`ValidateAge(20) == nil`になります。

# Alternatives

`checkAge`の戻り値を`error`にすれば、`return nil`が最初から型を持たないnilになり、この問題は起きません。自分で書く関数は、失敗を返す戻り値を具体型ではなく`error`にしてください。

# Idiomatic Go

errorを返す関数は、戻り値の型を`error`として宣言します。`*MyError`を返す関数の結果を`error`変数へ代入すると、同じ罠が呼び出し元で起きます。

# 振り返り

`var p *ValidationError; var err error = p`のとき、`p == nil`と`err == nil`がそれぞれ何になるか説明してください。
