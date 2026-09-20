# Exercise: early-return

## Goal

失敗条件を先に処理して通常経路を平坦にする。

## Task

CanShipを早期returnで書き直す。elseをなくし、動作を維持する。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./05-idiomatic-go/08-early-return
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
