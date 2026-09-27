# Exercise: 15-rwmutex

## Goal

読み取りの多い共有データをRWMutexで保護する。

## Task

名前ごとの得点を保持する`Scores`を実装してください。ゼロ値のまま使え、複数のgoroutineから同時に呼び出されます。

- `Set(name, score)`は得点を登録し、同じ名前なら上書きします。
- `Get(name)`は得点と登録済みかどうかを返します。
- `View(fn)`は現在の全得点を`fn`へ渡します。`fn`は内容を読むだけで変更しません。`View`の実行中に別の`View`や`Get`を呼んでも、先の`View`の終了を待たずに進みます。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `var s Scores; s.Get("go")` | `(0, false)` |
| `s.Set("go", 90); s.Set("go", 95); s.Get("go")` | `(95, true)` |
| `View`の`fn`が戻る前に、別のgoroutineで`View`を呼ぶ | 二つ目の`fn`がすぐに呼ばれ、`scores["go"]`は`95` |
| 100個のgoroutineが`Set`・`Get`・`View`を同時に呼ぶ | race detectorが競合を報告しない |

## Constraints

- テストと関数のsignatureを変更しないでください。
- `go test -race`で検証します。

## Run

```bash
go run ./cmd/koans check 06-concurrency/15-rwmutex
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
