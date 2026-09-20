# Exercise: map-grouping

## Goal

mapとsliceを組み合わせて分類する。

## Task

GroupはTeamごとに名前を入力順でまとめる。

## Examples

mapの表示順は問いません。各チーム内の名前の順序は入力順です。

| 入力・操作 | 期待する結果 |
|---|---|
| `Group([]Member{{Name: "A", Team: "red"}, {Name: "B", Team: "blue"}, {Name: "C", Team: "red"}})` | `red` の値は `[]string{"A", "C"}`、`blue` の値は `[]string{"B"}` |
| `Group(nil)` | 空のmap |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./02-data/14-map-grouping
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
