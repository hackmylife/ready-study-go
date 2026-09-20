# Exercise: error-test

## Goal

errorの原因を文字列以外で検証する。

## Task

Load(false)がwrapされたErrMissingを返すこと、Load(true)が成功することを検証する。

## Constraints

- exercise_test.goを編集する。exercise.goとmutants/は変更しない。
- 成功例だけでなく指定された境界値・失敗も検証する。
- go run ./cmd/koans check 07-testing/05-error-test で欠陥実装を見抜けるか確認する。

## Run

```bash
go run ./cmd/koans check 07-testing/05-error-test
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
