# Why?

待機中のrequestがある状態でtimeout=0ならDeadlineExceededを返します。Shutdownで待った後、期限を超えた接続をCloseします。終了用にはキャンセル済みのrequest contextを使いません。

# Alternatives

別の書き方も、同じ入力・出力とエラー契約を満たせます。模範解答と自分のコードを比較し、読み手が追う状態や分岐が増えていないか説明してください。

# Idiomatic Go

graceful shutdownには待機上限と、超過後の処理を決めます。

# 振り返り

テストが示す具体例を一つ選び、入力から出力まで手で追ってください。境界値のテストを一つ自分で追加してから、もう一度解いてみましょう。
