# Exercise: transfer

## Goal

残高移動の原子性と並行実行を守る。

## Task

Transferは整数の金額を送金する。同一口座・0以下・残高不足・未登録を拒否し、両口座を一定順でlockする。

## Examples

各行は独立した例です。金額と残高は整数で扱います。

| 入力・操作 | 期待する結果 |
|---|---|
| 初期残高a=100、b=50で `Transfer(ctx, db, "a", "b", 30)` | `nil`。残高はa=`70`、b=`80` |
| 初期残高a=100、b=50で、aからbへ80ずつ同時に送金する | 片方だけ成功、片方は `ErrFunds`。最終残高はa=`20`、b=`130` |
| 金額0、または送金元と送金先が同じ | `ErrInvalid`。残高は変わらない |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。
- PostgreSQLが必要です。[準備](../../docs/database.md)を済ませてください。

## Run

```bash
go test ./10-database/10-transfer
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
