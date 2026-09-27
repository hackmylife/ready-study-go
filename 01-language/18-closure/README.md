# Exercise: 18-closure

## Goal

関数値が外側の変数を捕捉して状態を保つ。

## Task

`Counter`は呼び出すたびに次の番号を返す関数を返します。返された関数は最初の呼び出しで`1`、以降は`2`、`3`と1ずつ増えた値を返します。`Counter`を呼ぶたびに独立したカウンタを作り、別のカウンタの呼び出しは互いの値に影響しません。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `next := Counter(); next(); next(); next()` | `1`、`2`、`3`を順に返す |
| 上に続けて`other := Counter(); other()` | `1`（`next`の値を引き継がない） |
| 上に続けて`next()` | `4`（`other`の呼び出しで変わらない） |

## Constraints

- package変数を使わず、`Counter`の中の変数を返す関数から更新してください。
- テストと関数のsignatureを変更しないでください。

## Run

```bash
go run ./cmd/koans check 01-language/18-closure
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
