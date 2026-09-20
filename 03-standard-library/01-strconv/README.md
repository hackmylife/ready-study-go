# Exercise: strconv

## Goal

文字列を数値に変換しエラーを扱う。

## Task

ParsePortは1〜65535の整数を受け付ける。空白、非数値、範囲外を拒否する。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `ParsePort("8080")` | `(8080, nil)` |
| `ParsePort("65535")` | `(65535, nil)` |
| `ParsePort("0")` / `ParsePort("65536")` / `ParsePort(" 80")` | いずれもerrorが非nil |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./03-standard-library/01-strconv
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
