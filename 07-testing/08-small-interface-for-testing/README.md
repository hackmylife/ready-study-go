# Exercise: small-interface-for-testing

## Goal

必要な操作だけのfakeで設計を確かめる。

## Task

CopyFirstはReaderから先頭nバイトを読む。短い入力と読み取りerrorを含めて検証する。n >= 0。

## Examples

以下は、あなたが書くテストで検証する対象関数の振る舞いです。テストコード自体を実装してください。

| 入力・操作 | 期待する結果 |
|---|---|
| 内容 `"abcdef"` のReaderと `n=3` を `CopyFirst` に渡す | バイト列の内容は `"abc"`、errorは `nil` |
| 内容 `"a"` のReaderと `n=3` | バイト列の内容は `"a"`、errorは `nil` |
| 読み取りerrorを返すReaderと `n=3` | 原因errorを返す |

## Constraints

- exercise_test.goを編集する。exercise.goとmutants/は変更しない。
- 成功例だけでなく指定された境界値・失敗も検証する。
- go run ./cmd/koans check 07-testing/08-small-interface-for-testing で欠陥実装を見抜けるか確認する。

## Run

```bash
go run ./cmd/koans check 07-testing/08-small-interface-for-testing
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
