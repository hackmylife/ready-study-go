# Exercise: composition

## Goal

部品を組み合わせて振る舞いを加える。

## Task

CountingWriterは渡されたWriterへ書き、実際に書けたバイト数を累積する。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| CountingWriterへ `"Go"`、続けて `"!"` を書く | 出力先の内容は `"Go!"`、累積 `Bytes` は `3` |
| 3バイトを渡したが出力先が1バイトだけ書いてerrorを返す | 戻り値の件数も増える `Bytes` も `1`。原因errorも返る |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./05-idiomatic-go/07-composition
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
