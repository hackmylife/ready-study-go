# Exercise: test-helper

## Goal

重複する検証をhelperにまとめる。

## Task

ParseBoolの成功・失敗テストを書く。helper内でt.Helperを呼び、失敗箇所が呼び出し元を示すようにする。

## Constraints

- exercise_test.goを編集する。exercise.goとmutants/は変更しない。
- 成功例だけでなく指定された境界値・失敗も検証する。
- go run ./cmd/koans check 07-testing/04-test-helper で欠陥実装を見抜けるか確認する。

## Run

```bash
go run ./cmd/koans check 07-testing/04-test-helper
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
