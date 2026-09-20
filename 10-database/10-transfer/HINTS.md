# Hints

<details>
<summary>Hint 1</summary>

SELECT FOR UPDATEで残高を確認する前に行をlockします。

</details>

<details>
<summary>Hint 2</summary>

逆方向の送金でも同じ順でlockします。

</details>

<details>
<summary>Hint 3</summary>

残高不足の判定と二つの更新を同じtransactionで行います。

</details>

