# Hints

<details>
<summary>Hint 1</summary>

先に口座作成と取得を通し、HTTPを介さないStoreの送金テストへ進みます。

</details>

<details>
<summary>Hint 2</summary>

送金ではORDER BY id FOR UPDATE、同じtxによる更新、uniqueなIdempotency-Keyを組み合わせます。

</details>

<details>
<summary>Hint 3</summary>

新規201・同じ再送200・キーの内容変更409を区別し、transaction完了後にHTTP応答を書きます。

</details>

