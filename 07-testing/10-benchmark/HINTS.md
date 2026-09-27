# Hints

<details>
<summary>Hint 1</summary>

`func BenchmarkJoin(b *testing.B)`の中で`for b.Loop() { ... }`と書くと、測定に必要な回数だけ繰り返されます。`b.ReportAllocs()`で割り当て回数も表示します。

</details>

<details>
<summary>Hint 2</summary>

`testing.AllocsPerRun(100, func() { ... })`は関数を100回実行し、1回あたりの平均割り当て回数を返します。

</details>

<details>
<summary>Hint 3</summary>

`+=`で連結する実装や、`Grow`で容量を確保しない実装も同じ文字列を返します。結果の比較だけでは見抜けないので、割り当て回数を検証します。

</details>
