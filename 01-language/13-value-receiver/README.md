# Exercise: value-receiver

## Goal

value receiverで元の値を保つ。

## Task

点の座標 `X` と `Y` にそれぞれ移動量 `dx` と `dy` を加えた、新しい `Point` を返す `Moved` を実装してください。元の点の座標は変更しません。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `p := Point{X: 2, Y: 5}` に対して `p.Moved(3, -2)` | 戻り値は `Point{X: 5, Y: 3}`。元の `p` は `Point{X: 2, Y: 5}` のまま |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./01-language/13-value-receiver
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
