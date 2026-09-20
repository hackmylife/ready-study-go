# Hints

<details>
<summary>Hint 1</summary>

callbackには*sql.Txを渡します。

</details>

<details>
<summary>Hint 2</summary>

callbackのerrorを確認してからCommitします。

</details>

<details>
<summary>Hint 3</summary>

commit後のdefer Rollbackが返すsql.ErrTxDoneは追加の失敗として扱いません。

</details>

