# Exercise: 17-range-index

## Goal

rangeの添字を使って元のsliceを更新する。

## Task

`DoubleInPlace`は受け取ったsliceの各要素を二倍にします。戻り値はなく、呼び出し元のsliceが変わります。nil・空sliceには何もしません。入力は二倍してintに収まる値です。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `values := []int{2, -3, 0}; DoubleInPlace(values)` | `valuesが[]int{4, -6, 0}になる` |
| `DoubleInPlace(nil)` | `何もせず終了` |

## Constraints

- `for i := range values`で元の要素を更新してください。
- テストと関数のsignatureを変更しないでください。
- テストは振る舞いを確認します。構文の使い分けは模範解答・解説でも確認してください。

## Run

```bash
go run ./cmd/koans check 01-language/17-range-index
```

[段階的なヒント](HINTS.md) / [ループとコレクションの案内](../../docs/loops-and-collections.md)
