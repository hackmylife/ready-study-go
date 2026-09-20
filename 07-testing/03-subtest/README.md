# Exercise: subtest

## Goal

ケース名で失敗を特定する。

## Task

Slugのテストをt.Runで分ける。前後空白と大文字を検証する。

## Examples

以下は、あなたが書くテストで検証する対象関数の振る舞いです。テストコード自体を実装してください。

| 入力・操作 | 期待する結果 |
|---|---|
| `Slug(" go ")` | `"go"` |
| `Slug("GO")` | `"go"` |
| `Slug("")` | `""` |

## Constraints

- exercise_test.goを編集する。exercise.goとmutants/は変更しない。
- 成功例だけでなく指定された境界値・失敗も検証する。
- go run ./cmd/koans check 07-testing/03-subtest で欠陥実装を見抜けるか確認する。

## Run

```bash
go run ./cmd/koans check 07-testing/03-subtest
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
