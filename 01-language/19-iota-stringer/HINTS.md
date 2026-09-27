# Hints

<details>
<summary>Hint 1</summary>

`const`ブロックの中で`iota`は`0`から始まり、行ごとに1ずつ増えます。

</details>

<details>
<summary>Hint 2</summary>

最初の行に`StatusPending Status = iota`と書けば、続く行は型と式を省略できます。

</details>

<details>
<summary>Hint 3</summary>

`fmt`は`String() string`を持つ値を表示するときにそのmethodを呼びます。未定義の値は`fmt.Sprintf("Status(%d)", int(s))`で表せます。

</details>
