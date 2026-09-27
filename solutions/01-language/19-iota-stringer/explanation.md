# Why?

`iota`は`const`ブロックの行番号なので、`StatusPending, StatusActive, StatusClosed`は`0, 1, 2`になります。`String`を持つと`fmt.Sprint(StatusActive)`は`1`ではなく`"active"`を返し、ログやエラーメッセージで状態を読めます。`Status(7)`のような範囲外の値も`"Status(7)"`と表示され、空文字列で情報を失いません。

# Alternatives

状態名を`[...]string{"pending", "active", "closed"}`の配列に置き、範囲内なら添字で引く書き方もあります。値が増えたときは`go generate`と`stringer`コマンドで`String`を生成する方法もあります。

# Idiomatic Go

`String`の中で`fmt.Sprintf("%v", s)`のように自分自身を`%v`で表示すると、`String`が再帰して止まりません。`int(s)`に変換してから整形します。

# 振り返り

`StatusActive`と`StatusClosed`の間に新しい状態を追加すると、既存の定数の値がどう変わるか説明してください。値を保存・送信している場合に何が起きるかも考えてみましょう。
