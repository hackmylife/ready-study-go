# Exercise: fake

## Goal

小さなfakeで外部依存を検証する。

## Task

Sendが宛先と本文をNotifierへ渡し、通知のerrorを返すことを検証する。

## Constraints

- exercise_test.goを編集する。exercise.goとmutants/は変更しない。
- 成功例だけでなく指定された境界値・失敗も検証する。
- go run ./cmd/koans check 07-testing/07-fake で欠陥実装を見抜けるか確認する。

## Run

```bash
go run ./cmd/koans check 07-testing/07-fake
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
