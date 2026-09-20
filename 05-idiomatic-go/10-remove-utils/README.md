# Exercise: remove-utils

## Goal

用途の分からないutilsから責務のある場所へ処理を移す。

## Task

NormalizeCodeをこのpackageへ移動し、utils packageのNormalizeCodeを削除する。余計な公開関数を残さない。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `NormalizeCode("  go-42 ")` | `"GO-42"` |
| `NormalizeCode("")` | `""` |
| 変更後のファイル構成 | 上の結果を保ち、この演習内の `utils/` がなくなる |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./05-idiomatic-go/10-remove-utils
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
