# Exercise: bufio

## Goal

Scannerで行を読み、走査エラーを確認する。

## Task

Linesは改行を除いた行を返す。空入力はnil。読み取りエラーは返す。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| 内容が `"a\n\nb\r\nlast"` のReaderを `Lines` に渡す | `[]string{"a", "", "b", "last"}` と `nil` error |
| 空のReader | `nil` sliceと `nil` error |
| 読み取りに失敗するReader | errorを返す |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./03-standard-library/10-bufio
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
