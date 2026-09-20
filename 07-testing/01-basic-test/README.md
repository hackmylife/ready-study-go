# Exercise: basic-test

## Goal

標準testingで結果を検証する。

## Task

Absのテストを書く。正数・負数・0を検証する。最小intは対象外。

## Examples

以下は、あなたが書くテストで検証する対象関数の振る舞いです。テストコード自体を実装してください。

| 入力・操作 | 期待する結果 |
|---|---|
| `Abs(5)` / `Abs(-5)` | どちらも `5` |
| `Abs(0)` | `0` |

## Constraints

- exercise_test.goを編集する。exercise.goとmutants/は変更しない。
- 成功例だけでなく指定された境界値・失敗も検証する。
- go run ./cmd/koans check 07-testing/01-basic-test で欠陥実装を見抜けるか確認する。

## Run

```bash
go run ./cmd/koans check 07-testing/01-basic-test
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
