# Exercise: http-test

## Goal

httptestでHTTPを検証する。

## Task

Healthのstatus・Content-Type・JSON本文を検証する。外部通信は不要。

## Constraints

- exercise_test.goを編集する。exercise.goとmutants/は変更しない。
- 成功例だけでなく指定された境界値・失敗も検証する。
- go run ./cmd/koans check 07-testing/06-http-test で欠陥実装を見抜けるか確認する。

## Run

```bash
go run ./cmd/koans check 07-testing/06-http-test
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
