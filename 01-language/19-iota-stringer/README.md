# Exercise: 19-iota-stringer

## Goal

iotaで列挙値を定義し、Stringerで表示名を持たせる。

## Task

注文の状態を表す`Status`型に、`StatusPending`・`StatusActive`・`StatusClosed`の三つの定数を`0`・`1`・`2`の順で定義してください。また`Status.String`を実装し、`fmt`で表示したときに状態名になるようにします。定義していない値は`Status(値)`と表示します。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `int(StatusPending)` / `int(StatusActive)` / `int(StatusClosed)` | `0` / `1` / `2` |
| `fmt.Sprint(StatusActive)` | `"active"` |
| `fmt.Sprint(StatusClosed)` | `"closed"` |
| `fmt.Sprint(Status(7))` | `"Status(7)"` |

## Constraints

- 定数は`iota`で定義し、値を一つずつ書かないでください。
- テストと関数のsignatureを変更しないでください。

## Run

```bash
go run ./cmd/koans check 01-language/19-iota-stringer
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
