# Exercise: variables

## Goal

変数の宣言と代入を使う。

## Task

Swapを実装し、受け取った値を交換して返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `Swap(2, 9)` | 戻り値は `(9, 2)` |
| `Swap(-3, 0)` | 戻り値は `(0, -3)` |
| `Swap(4, 4)` | 戻り値は `(4, 4)` |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

リポジトリのルートから、演習ディレクトリを指定します。

```bash
go test ./01-language/01-variables
```

`exercise_test.go`だけを指定すると`exercise.go`が読み込まれず、`undefined: Swap`になります。このディレクトリ内で実行する場合は、引数なしの`go test`を使ってください。

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
