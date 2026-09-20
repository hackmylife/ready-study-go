# Exercise: concurrent-processing

## Goal

バッチ処理に上限付き並行実行を組み込む。

## Task

Mapでレコードの数値変換を並行に行う。順序・キャンセル・error・worker上限を守る。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./09-batch/07-concurrent-processing
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
