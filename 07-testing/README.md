# 07-testing

[全体の学習ガイド](../README.md)

この章ではexercise_test.goを編集します。koans checkで、書いたテストが欠陥実装を検出できるかも確認してください。

| 演習 | 学習目標 |
|---|---|
| [01-basic-test](01-basic-test/README.md) | 標準testingで結果を検証する。 |
| [02-table-driven-test](02-table-driven-test/README.md) | 入力と期待値をtableにまとめる。 |
| [03-subtest](03-subtest/README.md) | ケース名で失敗を特定する。 |
| [04-test-helper](04-test-helper/README.md) | 重複する検証をhelperにまとめる。 |
| [05-error-test](05-error-test/README.md) | errorの原因を文字列以外で検証する。 |
| [06-http-test](06-http-test/README.md) | httptestでHTTPを検証する。 |
| [07-fake](07-fake/README.md) | 小さなfakeで外部依存を検証する。 |
| [08-small-interface-for-testing](08-small-interface-for-testing/README.md) | 必要な操作だけのfakeで設計を確かめる。 |
| [09-fuzz](09-fuzz/README.md) | fuzz testで入力に依存しない性質を検証する。 |
| [10-benchmark](10-benchmark/README.md) | benchmarkで性能を測り、割り当て回数をテストで固定する。 |
| [11-synctest](11-synctest/README.md) | testing/synctestの仮想時計で、時間に依存する処理を実際に待たずに検証する。 |
