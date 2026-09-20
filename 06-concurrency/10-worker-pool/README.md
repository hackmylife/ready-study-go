# Exercise: worker-pool

## Goal

固定数のworkerで仕事を配る。

## Task

Mapを実装する。入力順を保持し、最初のerrorで中止して全workerを待つ。workはcontextに協調して終了する。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./06-concurrency/10-worker-pool
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
