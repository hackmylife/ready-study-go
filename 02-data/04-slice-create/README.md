# Exercise: slice-create

## Goal

長さを持つsliceを作る。

## Task

Zerosは長さnのゼロで初期化されたsliceを返す。n >= 0。n=0でもnilにはしない。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `Zeros(3)` | `[]int{0, 0, 0}` |
| `Zeros(0)` | `[]int{}`（nilではない空slice） |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./02-data/04-slice-create
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
