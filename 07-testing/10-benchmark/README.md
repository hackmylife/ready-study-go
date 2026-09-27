# Exercise: benchmark

## Goal

benchmarkで性能を測り、割り当て回数をテストで固定する。

## Task

`Join`のテストとbenchmarkを書いてください。`Join`は`strings.Join`と同じ結果を返し、一回の呼び出しでのメモリ割り当てを1回以下に抑えます。

- 結果の文字列をテストで検証する。
- 割り当て回数を`testing.AllocsPerRun`で測り、1回を超えたら失敗させる。
- `BenchmarkJoin`を書き、`go test -bench`で実行時間と割り当て回数を表示する。

## Examples

以下は、あなたが書くテストで検証する対象関数の振る舞いです。テストコード自体を実装してください。

| 入力・操作 | 期待する結果 |
|---|---|
| `Join([]string{"a", "b", "c"}, ", ")` | `"a, b, c"` |
| `Join([]string{"a"}, ", ")` | `"a"` |
| `Join(nil, ", ")` | `""` |
| 8個の単語を`", "`で`Join`する1回の呼び出し | メモリ割り当ては1回以下 |

## Constraints

- exercise_test.goを編集する。exercise.goとmutants/は変更しない。
- 割り当て回数は8個程度の要素で測る。要素が1〜2個では欠陥実装と回数が変わらないことがあります。
- `koans check`はbenchmarkを実行しません。下の`go test -bench`で実行して結果を読んでください。
- go run ./cmd/koans check 07-testing/10-benchmark で欠陥実装を見抜けるか確認する。

## Run

```bash
go run ./cmd/koans check 07-testing/10-benchmark
go test -run '^$' -bench=Join ./07-testing/10-benchmark
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
