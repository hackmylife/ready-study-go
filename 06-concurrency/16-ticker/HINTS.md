# Hints

<details>
<summary>Hint 1</summary>

`time.Sleep(interval)`で待つと、待っている間に`ctx`が終了しても気付けません。期限が`2500ms`のテストでは`3s`まで戻らなくなります。

</details>

<details>
<summary>Hint 2</summary>

`ticker := time.NewTicker(interval)`は`interval`ごとに`ticker.C`へ時刻を送ります。`defer ticker.Stop()`で止めます。

</details>

<details>
<summary>Hint 3</summary>

`for`の中の`select`で`<-ctx.Done()`と`<-ticker.C`を同時に待ちます。

</details>
