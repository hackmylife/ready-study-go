# Hints

<details>
<summary>Hint 1</summary>

`func FuzzReverse(f *testing.F)`の中で`f.Add`にseedを渡し、`f.Fuzz(func(t *testing.T, s string) { ... })`で性質を検証します。

</details>

<details>
<summary>Hint 2</summary>

`unicode/utf8`の`ValidString`・`DecodeRuneInString`・`DecodeLastRuneInString`で、UTF-8の妥当性と先頭・末尾のruneを調べられます。

</details>

<details>
<summary>Hint 3</summary>

何もしない`Reverse`でも「二回反転すると戻る」は成り立ちます。先頭と末尾のruneを比べる性質も必要です。

</details>
