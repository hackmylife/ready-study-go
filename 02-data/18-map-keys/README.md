# Exercise: 18-map-keys

## Goal

mapsとslicesでmapのキーを順序付きで取り出す。

## Task

`SortedKeys`はmapのキーを文字列の昇順で返します。値の大小は順序に影響しません。nil・空mapにはnilを返し、入力は変更しません。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `SortedKeys(map[string]int{"pear": 1, "apple": 9})` | `[]string{"apple", "pear"}` |
| `SortedKeys(nil)` | `nil` |

## Constraints

- 標準パッケージ`maps`と`slices`を使ってください。必要なimportも自分で追加します。
- テストと関数のsignatureを変更しないでください。
- テストは振る舞いを確認します。構文の使い分けは模範解答・解説でも確認してください。

## Run

```bash
go run ./cmd/koans check 02-data/18-map-keys
```

[段階的なヒント](HINTS.md) / [ループとコレクションの案内](../../docs/loops-and-collections.md)
