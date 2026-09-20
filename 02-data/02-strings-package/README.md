# Exercise: strings-package

## Goal

stringsで空白を正規化する。

## Task

Normalizeは前後・連続する空白を除き、語の間を半角スペース一つにする。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `Normalize("  Go\t is\n fun  ")` | `"Go is fun"` |
| `Normalize("　猫　犬")`（全角スペース） | `"猫 犬"` |
| `Normalize(" ")` | `""` |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./02-data/02-strings-package
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
