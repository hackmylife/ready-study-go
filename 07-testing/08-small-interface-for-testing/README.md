# Exercise: small-interface-for-testing

## Goal

必要な操作だけのfakeで設計を確かめる。

## Task

CopyFirstはReaderから先頭nバイトを読む。短い入力と読み取りerrorを含めて検証する。n >= 0。

## Constraints

- exercise_test.goを編集する。exercise.goとmutants/は変更しない。
- 成功例だけでなく指定された境界値・失敗も検証する。
- go run ./cmd/koans check 07-testing/08-small-interface-for-testing で欠陥実装を見抜けるか確認する。

## Run

```bash
go run ./cmd/koans check 07-testing/08-small-interface-for-testing
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
