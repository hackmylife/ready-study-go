# Why?

GETの結果はContent-Type: application/jsonと{"name":"Aki"}です。headerは本文を書き始める前に設定します。

# Alternatives

Encodeの書き込みエラー後にhttp.Errorを追加しても、すでに送ったstatusや本文は戻せません。実務ではログ記録や接続終了の扱いを検討します。

# Idiomatic Go

書き込み後に応答statusを変更することはできません。

# 振り返り

テストが示す具体例を一つ選び、入力から出力まで手で追ってください。境界値のテストを一つ自分で追加してから、もう一度解いてみましょう。
