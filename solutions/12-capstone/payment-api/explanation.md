# Why?

a=100,b=50から30を送るとa=70,b=80となり、送金履歴が一件残ります。同じIdempotency-Keyと内容で再送しても残高は変わりません。

口座行はID順にlockし、残高検査・両残高更新・送金履歴保存を一つのtransactionで行います。lock待機中に先行requestが完了することがあるため、lock後にもキーを調べます。残高不足でも再送結果を取得できるよう、既存送金の照合は残高検査より先です。

HTTPは入力を検証し、Storeへrequest contextを渡し、errorを公開用statusへ変換します。DBの診断情報は構造化ログに出し、responseには固定コードだけを返します。

# Alternatives

この規模ではStoreという具体型と標準ServeMuxで責務を追えます。必要になる前にcontroller/service/repository階層やDI frameworkを追加しません。別の永続化先が必要になった時は、その利用側に必要な小さなinterfaceを定義できます。

# Idiomatic Go

新規201・同じ再送200・キーの内容変更409を区別し、transaction完了後にHTTP応答を書きます。

# 振り返り

テストが示す具体例を一つ選び、入力から出力まで手で追ってください。境界値のテストを一つ自分で追加してから、もう一度解いてみましょう。
