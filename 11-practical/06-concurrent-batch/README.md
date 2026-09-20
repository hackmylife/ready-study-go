# Exercise: concurrent-batch

## Goal

I/O・変換・並行実行・errorを組み合わせる。

## Task

Runは整数行を全て読み、上限付きでworkを実行し、入力順に改行付きで出力する。読み取り・変換失敗時は出力しない。出力途中の失敗では書いた分が残る。入力全体がメモリに収まる前提。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./11-practical/06-concurrent-batch
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
