# Hints

<details>
<summary>Hint 1</summary>

`synctest.Test(t, func(t *testing.T) { ... })`の中では時計が仮想化されます。中の全てのgoroutineが待ち状態になると、時計は次のtimerの時刻まで進みます。

</details>

<details>
<summary>Hint 2</summary>

`start := time.Now()`と`time.Since(start)`で、`Retry`が戻るまでの仮想時間を`2*time.Second`のように`==`で比較できます。

</details>

<details>
<summary>Hint 3</summary>

待たない実装、1回多く呼ぶ実装、`time.Sleep`で待ち`ctx`を無視する実装を見抜く必要があります。経過時間・呼び出し回数・返すerrorのそれぞれを検証します。

</details>
