# Exercise: functional-options

## Goal

省略可能な設定にfunctional optionsを使う。

## Task

NewClientとWithTimeoutを実装する。既定値は5秒、正のtimeoutだけを許可し、後の指定を優先する。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `NewClient()` | Timeoutが `5*time.Second` のClientと `nil` error |
| `NewClient(WithTimeout(time.Second), WithTimeout(2*time.Second))` | Timeoutが `2*time.Second` のClientと `nil` error |
| `NewClient(WithTimeout(0))` | error |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./05-idiomatic-go/13-functional-options
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
