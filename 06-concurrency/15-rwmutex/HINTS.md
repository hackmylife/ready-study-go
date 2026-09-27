# Hints

<details>
<summary>Hint 1</summary>

`sync.Mutex`は読み取り同士も一つずつしか進めません。二つ目の`View`が待たされるテストはこれを検出しています。

</details>

<details>
<summary>Hint 2</summary>

`sync.RWMutex`は`RLock`同士なら同時に保持できます。書き込みには`Lock`を使います。

</details>

<details>
<summary>Hint 3</summary>

ゼロ値で使うため、mapは`Set`の中でnilなら作ります。nilのmapからの読み取りは安全です。

</details>
