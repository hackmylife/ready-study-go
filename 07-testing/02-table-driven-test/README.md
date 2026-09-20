# Exercise: table-driven-test

## Goal

入力と期待値をtableにまとめる。

## Task

Gradeのテーブルテストを書く。境界59/60/79/80と通常値を検証する。

## Examples

以下は、あなたが書くテストで検証する対象関数の振る舞いです。テストコード自体を実装してください。

| 入力・操作 | 期待する結果 |
|---|---|
| `Grade(59)` / `Grade(60)` | それぞれ `"C"` / `"B"` |
| `Grade(79)` / `Grade(80)` | それぞれ `"B"` / `"A"` |

## Constraints

- exercise_test.goを編集する。exercise.goとmutants/は変更しない。
- 成功例だけでなく指定された境界値・失敗も検証する。
- go run ./cmd/koans check 07-testing/02-table-driven-test で欠陥実装を見抜けるか確認する。

## Run

```bash
go run ./cmd/koans check 07-testing/02-table-driven-test
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
