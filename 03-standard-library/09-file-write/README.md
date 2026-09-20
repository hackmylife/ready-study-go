# Exercise: file-write

## Goal

ファイルの上書きと権限を扱う。

## Task

Writeは内容を上書きする。新規ファイルの権限は0600。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `Write(path, "long content")` の後に `Write(path, "Go")` | どちらも `nil` error。ファイル内容は `"Go"` だけになる |
| まだ存在しないファイルに書く | ファイルを作成し、他のユーザーへのアクセス権を付けない（作成権限 `0600`） |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./03-standard-library/09-file-write
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
