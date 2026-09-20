# Hints

<details>
<summary>Hint 1</summary>

DoのerrorとStatusCodeを分けて確認します。

</details>

<details>
<summary>Hint 2</summary>

Do成功直後にdefer Body.Closeを置きます。

</details>

<details>
<summary>Hint 3</summary>

HTTP応答の本文を所有する側が必ずcloseします。

</details>

