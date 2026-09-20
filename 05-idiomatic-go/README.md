# 05-idiomatic-go

[全体の学習ガイド](../README.md)

| 演習 | 学習目標 |
|---|---|
| [01-zero-value](01-zero-value/README.md) | 初期化しなくても使える型を設計する。 |
| [02-remove-constructor](02-remove-constructor/README.md) | 不要なconstructorを削る。 |
| [03-small-interface](03-small-interface/README.md) | 利用側が必要とするmethodだけを要求する。 |
| [04-interface-consumer-side](04-interface-consumer-side/README.md) | 利用側で必要な依存を定義する。 |
| [05-remove-interface](05-remove-interface/README.md) | 一つの具体型にしか役割がないinterfaceを削る。 |
| [06-accept-interface-return-struct](06-accept-interface-return-struct/README.md) | 入力はinterface、出力は具体型にする。 |
| [07-composition](07-composition/README.md) | 部品を組み合わせて振る舞いを加える。 |
| [08-early-return](08-early-return/README.md) | 失敗条件を先に処理して通常経路を平坦にする。 |
| [09-package-design](09-package-design/README.md) | 責務に沿って公開APIを絞る。 |
| [10-remove-utils](10-remove-utils/README.md) | 用途の分からないutilsから責務のある場所へ処理を移す。 |
| [11-explicit-dependency](11-explicit-dependency/README.md) | 時計を明示的な依存として渡す。 |
| [12-simple-api](12-simple-api/README.md) | boolフラグの意味が伝わるAPIにする。 |
| [13-functional-options](13-functional-options/README.md) | 省略可能な設定にfunctional optionsを使う。 |
| [14-dont-use-functional-options](14-dont-use-functional-options/README.md) | 設定が一つのAPIから不要なoptionを削る。 |
