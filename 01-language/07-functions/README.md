# Exercise: functions

## Goal

引数と戻り値を持つ関数を書く。

## Task

ApplyTwiceはfを二度適用する。fはnilでないものとする。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| 値 `3` と「受け取った値に2を足す関数」を `ApplyTwice` に渡す | `7`（3 → 5 → 7） |
| 値 `2` と「受け取った値を3倍にする関数」を渡す | `18`（2 → 6 → 18）。関数は二度呼ばれる |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./01-language/07-functions
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
