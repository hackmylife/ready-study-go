# Exercise: fuzz

## Goal

fuzz testで入力に依存しない性質を検証する。

## Task

`Reverse`のfuzz testを`FuzzReverse`として書いてください。個々の入力と期待値を並べる代わりに、どの入力でも成り立つ性質を検証します。`Reverse`の入力は正しいUTF-8の文字列に限ります。

- 結果も正しいUTF-8である。
- 二回反転すると元の文字列に戻る。
- 結果の最初のruneは、元の文字列の最後のruneである。

## Examples

以下は、あなたが書くテストで検証する対象関数の振る舞いです。テストコード自体を実装してください。

| 入力・操作 | 期待する結果 |
|---|---|
| `Reverse("Go言語")` | `"語言oG"` |
| `Reverse(Reverse("hello, 世界"))` | `"hello, 世界"` |
| `Reverse("")` | `""` |

## Constraints

- exercise_test.goを編集する。exercise.goとmutants/は変更しない。
- `f.Add`で複数バイトの文字を含むseedを登録する。`koans check`はseedだけを実行します。
- 正しいUTF-8でない入力は`t.Skip`で除外する。
- go run ./cmd/koans check 07-testing/09-fuzz で欠陥実装を見抜けるか確認する。

## Run

```bash
go run ./cmd/koans check 07-testing/09-fuzz
# ランダムな入力で10秒間探索する
go test -fuzz=FuzzReverse -fuzztime=10s ./07-testing/09-fuzz
```

探索で失敗が見つかると、入力が`testdata/fuzz/FuzzReverse/`に保存され、以降の`go test`でも毎回実行されます。

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
