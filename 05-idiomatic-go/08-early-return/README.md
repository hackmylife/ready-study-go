# Exercise: early-return

## Goal

失敗条件を先に処理して通常経路を平坦にする。

## Task

CanShipを早期returnで書き直す。elseをなくし、動作を維持する。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `CanShip(true, 1)` | `nil` |
| `CanShip(false, 0)` | 表示が `"unpaid"` のerror |
| `CanShip(true, 0)` | 表示が `"out of stock"` のerror |
| 変更後のコード | 上の結果を保ち、elseが存在しない |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./05-idiomatic-go/08-early-return
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
