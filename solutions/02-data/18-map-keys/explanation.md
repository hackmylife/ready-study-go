# Why?

`maps.Keys`はキーのiteratorを返し、`slices.Sorted`はそれをsliceに集めて整列します。`{"pear":1,"apple":9}`なら`["apple","pear"]`です。通常のrangeでキーを集めてslices.Sortしても同じ契約を満たせます。標準のmapsパッケージは辞書の操作用であり、要素変換のMap関数とは別です。

## 振り返り

入力例を一つ選び、各反復で変化する値と終了条件を説明してください。
