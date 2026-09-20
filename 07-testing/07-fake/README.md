# Exercise: fake

## Goal

小さなfakeで外部依存を検証する。

## Task

Sendが宛先と本文をNotifierへ渡し、通知のerrorを返すことを検証する。

## Examples

以下は、あなたが書くテストで検証する対象関数の振る舞いです。テストコード自体を実装してください。

| 入力・操作 | 期待する結果 |
|---|---|
| `Send(fake, "aki@example.test")` | fakeの `Notify` に宛先 `"aki@example.test"` と本文 `"Welcome"` が渡る |
| fakeが `Notify` でerrorを返す | `Send` の戻り値にその原因errorが保持される |

## Constraints

- exercise_test.goを編集する。exercise.goとmutants/は変更しない。
- 成功例だけでなく指定された境界値・失敗も検証する。
- go run ./cmd/koans check 07-testing/07-fake で欠陥実装を見抜けるか確認する。

## Run

```bash
go run ./cmd/koans check 07-testing/07-fake
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
