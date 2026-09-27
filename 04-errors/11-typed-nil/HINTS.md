# Hints

<details>
<summary>Hint 1</summary>

失敗したテストの`%#v`の出力を読みます。`(*koan.ValidationError)(nil)`は「型は`*ValidationError`、値はnil」という意味です。

</details>

<details>
<summary>Hint 2</summary>

interfaceの値は型と値の組です。型を持つinterfaceは、中の値がnilでも`nil`と等しくなりません。

</details>

<details>
<summary>Hint 3</summary>

`checkAge`の結果を`*ValidationError`のまま`nil`と比較し、nilなら`return nil`と書きます。

</details>
