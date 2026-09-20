# Go Koans — Implementation Specification

## 1. 概要

Goを「知っている」状態ではなく、**AIに頼らず、手で自然にGoコードを書ける状態**を目指す演習型教材を作成する。

教材形式は Kotlin Koans を参考にする。

https://github.com/Kotlin/kotlin-koans-edu

基本的な学習サイクルは以下とする。

```text
失敗するテストを読む
    ↓
TODO部分を自分で実装する
    ↓
go test
    ↓
失敗
    ↓
修正
    ↓
テスト成功
    ↓
模範解答と比較
    ↓
「なぜGoではこの書き方をするのか」を理解する
```

単なるGo文法教材ではなく、

**Syntax → Standard Library → Go Idioms → Design Judgment → Practical Application**

まで段階的に習得できる教材とする。

---

# 2. 学習目標

修了時に以下をAIなしで実装できることを目標とする。

## 基本文法

* variable / constant
* if / switch
* for / range
* function
* multiple return values
* pointer
* struct
* method
* slice
* map
* string / rune

## 標準ライブラリ

少なくとも以下を自然に利用できること。

* strings
* strconv
* time
* errors
* fmt
* io
* os
* bufio
* encoding/json
* encoding/csv
* context
* net/http
* sync

## Go固有の考え方

以下をコードとして理解する。

* Error is a value
* useful zero value
* early return
* small interfaces
* interfaceは利用側で定義する
* accept interfaces, return structs
* composition
* explicit dependencies
* defer
* context cancellation
* bounded concurrency
* goroutine lifecycle
* channel ownership
* simple code over clever code

## 実務能力

以下のようなコードを自力で書けること。

* HTTP Server
* HTTP Client
* JSON API
* Batch
* File processing
* Database access
* Transaction
* Concurrent processing
* Graceful shutdown

---

# 3. 教材設計原則

## 3.1 Koans形式

基本的に各exerciseには未完成コードとテストを配置する。

例：

```text
03-slices/
  01-append/
    exercise.go
    exercise_test.go
    README.md
```

`exercise.go`：

```go
func AppendValue(values []string, value string) []string {
    // TODO: implement
    return nil
}
```

`exercise_test.go`：

```go
func TestAppendValue(t *testing.T) {
    got := AppendValue([]string{"a", "b"}, "c")
    want := []string{"a", "b", "c"}

    if !reflect.DeepEqual(got, want) {
        t.Fatalf("want %v, got %v", want, got)
    }
}
```

受講者はTODOを実装し、

```bash
go test ./...
```

を成功させる。

---

# 4. Exerciseの粒度

原則として、

**1 Exercise = 1 Concept**

とする。

例えば「slice」という巨大なExerciseにはしない。

```text
slices/
  01-create
  02-append
  03-range
  04-filter
  05-copy
  06-nil-slice
```

のように分割する。

基本Exerciseは5〜15分程度で解けるサイズとする。

実務Exerciseは30〜60分程度でもよい。

Capstoneのみ1〜3時間程度を想定する。

---

# 5. Exerciseの構成

各Exerciseは原則として以下を持つ。

```text
exercise.go
exercise_test.go
README.md
```

必要に応じてfixture等を追加してよい。

READMEには以下だけを書く。

```markdown
# Exercise: Error Wrapping

## Goal

error wrappingを理解する。

## Task

LoadUserを実装する。

## Constraints

- panic禁止
- errorを握りつぶさない

## Run

go test
```

READMEに答えを書きすぎないこと。

**テストと関数signatureから要求を読み取ること自体も学習対象とする。**

---

# 6. Solution

各Exerciseには模範解答を用意する。

ただし通常のexerciseから直接見えにくい構造にする。

例：

```text
solutions/
  03-slices/
    01-append/
      solution.go
      explanation.md
```

solutionにはコードだけでなく、

```markdown
# Why?

この実装ではappendを利用している。

Goのsliceは可変長配列そのものではなく、
underlying arrayへのdescriptorである。

...

# Alternatives

別の実装方法も存在する。

...

# Idiomatic Go

このケースでは明示的なループよりappendを利用する方が
意図が明確になる。
```

のような説明を付ける。

---

# 7. Goらしさ

この教材では、

**「テストが通る = 正解」だけにはしない。**

以下を学習対象とする。

## Clear is better than clever

短さより読みやすさを優先する。

## Error is a value

panic/exception的な思考ではなく、errorを通常の値として扱う。

## Useful zero value

不要なconstructorを作らない。

## Small interfaces

巨大interfaceを避ける。

## Interface belongs to the consumer

interfaceを実装側ではなく利用側で定義する。

## Explicit dependencies

依存関係を隠さない。

## Goroutine lifecycle

作ったgoroutineがいつ終了するのか説明できるようにする。

## Context propagation

context.Contextを適切に伝播する。

## Don't over-engineer

以下のような不要な抽象化を避ける。

```text
utils
helpers
common
AbstractFactory
過剰なRepository interface
不要なDI framework
不要なgoroutine
```

---

# 8. 「削る」Exercise

Goらしさを学習するため、

**コードを書くExerciseだけでなく、不要なコードを削るExercise**

を用意する。

例：

### Remove Unnecessary Interface

```go
type UserRepository interface {
    Find(string) (*User, error)
}

type PostgresUserRepository struct{}
```

要件：

> 不要なinterfaceを削除し、よりシンプルな設計に変更する。

---

### Remove Constructor

```go
type Counter struct {
    value int
}

func NewCounter() *Counter {
    return &Counter{}
}
```

zero valueで利用可能ならconstructorを削除する。

---

### Remove Goroutine

不要に

```go
go func() {
    ...
}()
```

しているコードを同期処理へ戻す。

---

### Remove Utils Package

```text
utils/
  string.go
  time.go
```

を責務のあるpackageへ移動する。

---

# 9. Curriculum

## Level 1 — Language Basics

目的：

**Goの基本構文を考えずに書けるようにする。**

```text
01-language/

01-variables
02-constants
03-if
04-switch
05-for
06-range
07-functions
08-multiple-return-values
09-pointers
10-structs
11-methods
12-pointer-receiver
13-value-receiver
```

---

# 10. Level 2 — Data Structures

```text
02-data/

01-string
02-strings-package
03-rune
04-slice-create
05-slice-append
06-slice-range
07-slice-filter
08-slice-copy
09-nil-slice
10-map-create
11-map-lookup
12-comma-ok
13-map-count
14-map-grouping
```

ここでは特に

```go
var values []string
```

と

```go
values := []string{}
```

の違いも扱う。

---

# 11. Level 3 — Standard Library

```text
03-standard-library/

01-strconv
02-time
03-time-parse
04-time-zone
05-json-marshal
06-json-unmarshal
07-json-decoder
08-file-read
09-file-write
10-bufio
11-csv
12-io-reader
13-io-writer
```

特に`io.Reader` / `io.Writer`は重要Conceptとして扱う。

---

# 12. Level 4 — Error Handling

```text
04-errors/

01-error
02-errors-new
03-fmt-errorf
04-error-wrapping
05-errors-is
06-errors-as
07-sentinel-error
08-custom-error
09-defer
10-cleanup
```

重要原則：

```text
error is a value
```

通常処理でpanicを使用しないことを理解させる。

---

# 13. Level 5 — Idiomatic Go

このLevelを教材の重要部分とする。

```text
05-idiomatic-go/

01-zero-value
02-remove-constructor
03-small-interface
04-interface-consumer-side
05-remove-interface
06-accept-interface-return-struct
07-composition
08-early-return
09-package-design
10-remove-utils
11-explicit-dependency
12-simple-api
13-functional-options
14-dont-use-functional-options
```

重要なのは、

**パターンを覚えるのではなく、いつ使わないかも学習すること。**

---

# 14. Level 6 — Concurrency

```text
06-concurrency/

01-goroutine
02-waitgroup
03-channel
04-buffered-channel
05-close-channel
06-select
07-context
08-context-cancel
09-context-timeout
10-worker-pool
11-bounded-concurrency
12-race-condition
13-mutex
14-remove-goroutine
```

Concurrency Exerciseでは必要に応じて、

```bash
go test -race ./...
```

を利用する。

重要Concept：

```text
goroutineには終了条件を持たせる

channelを誰がcloseするか明確にする

無制限にgoroutineを生成しない

concurrencyは目的ではなく手段
```

---

# 15. Level 7 — Testing

```text
07-testing/

01-basic-test
02-table-driven-test
03-subtest
04-test-helper
05-error-test
06-http-test
07-fake
08-small-interface-for-testing
```

mock frameworkへの依存は避ける。

Go標準のテスト手法を優先する。

---

# 16. Level 8 — HTTP

最初はframeworkを使用しない。

`net/http`を利用する。

```text
08-http/

01-handler
02-servemux
03-path-parameter
04-query-parameter
05-json-response
06-json-request
07-status-code
08-error-response
09-middleware
10-http-client
11-client-timeout
12-context
13-graceful-shutdown
```

Gin / Echo等は教材の中心にしない。

まず標準ライブラリを理解する。

---

# 17. Level 9 — Batch

```text
09-batch/

01-read-lines
02-read-csv
03-transform
04-validation
05-stream-processing
06-error-handling
07-concurrent-processing
08-context-cancellation
```

最終的には、

```text
CSV
 ↓
parse
 ↓
validate
 ↓
transform
 ↓
output
```

という小さなbatchを完成させる。

---

# 18. Level 10 — Database

PostgreSQLを想定する。

最初は`database/sql`を利用する。

```text
10-database/

01-open
02-query-row
03-query
04-scan
05-exec
06-context
07-not-found
08-transaction
09-rollback
10-transfer
```

ORMに依存しない。

DBアクセスの基本を理解した後であれば、

```text
pgx
sqlx
scany
```

等を発展課題として追加してもよい。

---

# 19. Practical Exercises

ここからは1 Conceptではなく、複数Conceptを組み合わせる。

```text
11-practical/

01-http-server
02-http-client
03-csv-batch
04-database-crud
05-transaction
06-concurrent-batch
```

---

# 20. Capstone

最終課題として小さなPayment APIを作る。

```text
12-capstone/
  payment-api/
```

API：

```text
POST /accounts
GET  /accounts/{id}

POST /transfers
GET  /transfers/{id}
```

必要Concept：

* net/http
* JSON
* PostgreSQL
* transaction
* validation
* error handling
* context
* structured logging
* graceful shutdown

---

# 21. Capstoneの設計方針

最初から以下のようなArchitectureを強制しない。

```text
controller
service
repository
domain
usecase
```

必要な責務が発生したときにpackageを分割する。

最終的には例えば、

```text
cmd/
  server/

internal/
  account/
  transfer/
  postgres/
```

程度のシンプルな構成を目指す。

---

# 22. Test Strategy

テストは3種類に分ける。

## Visible Tests

受講者が読めるテスト。

要求理解にも利用する。

## Hidden Tests

単純なhard codingや境界値漏れを防止する。

CI上で実行することを想定する。

## Quality Checks

以下を実行する。

```bash
gofmt
go vet
go test
go test -race
staticcheck
```

`-race`は必要なExerciseのみでもよい。

---

# 23. Progression

Exerciseは順番に解くことを推奨するが、

各Exerciseは可能な限り独立させる。

```text
Language
   ↓
Data
   ↓
Standard Library
   ↓
Errors
   ↓
Idiomatic Go
   ↓
Concurrency
   ↓
Testing
   ↓
HTTP / Batch / DB
   ↓
Practical
   ↓
Capstone
```

---

# 24. Hint System

AIを使わず学習できるよう、段階的なHintを用意する。

各Exerciseについて可能なら、

```text
Hint 1
Hint 2
Hint 3
Solution
```

を用意する。

例：

### Hint 1

> strings packageを調べてみよう。

### Hint 2

> strings.TrimSpaceが利用できる。

### Hint 3

> strings.Fieldsを利用すると連続した空白を扱える。

Hintは答えそのものを最初から提示しない。

---

# 25. AIに依存しない学習体験

この教材の重要な目的は、

> AIにコードを書かせる能力

ではなく、

> AIなしでもGoコードを書ける能力

を身につけることである。

そのためExerciseは、

```text
問題を見る
↓
自分で考える
↓
コンパイルする
↓
エラーを読む
↓
修正する
↓
テストする
```

というフィードバックループを短くする。

受講者が困った場合も、

```text
AI
```

ではなく、

```text
compiler
go doc
standard library documentation
hint
```

の順番で解決することを推奨する。

---

# 26. Repository Structure

最終的なrepositoryは概ね以下とする。

```text
go-koans/

README.md
go.mod
Makefile

01-language/
02-data/
03-standard-library/
04-errors/
05-idiomatic-go/
06-concurrency/
07-testing/
08-http/
09-batch/
10-database/
11-practical/
12-capstone/

solutions/
```

各directory以下にExerciseを配置する。

---

# 27. Root README

Root READMEには以下を記載する。

## What is Go Koans?

Goを手で書くための反復練習教材。

## Philosophy

```text
Don't ask AI first.

Read.
Think.
Write.
Compile.
Test.
Understand.
```

## Getting Started

```bash
git clone ...
cd go-koans

go test ./01-language/01-variables/...
```

## Rules

1. まず自分で書く
2. compiler errorを読む
3. test failureを読む
4. `go doc`を使う
5. Hintを見る
6. それでも分からなければSolutionを見る

---

# 28. Implementation Requirements

Codexによる実装では以下を守ること。

* 最新のstable Goを対象とする
* `gofmt`済みであること
* `go vet ./...`が成功すること
* 通常Exerciseでは外部dependencyを極力使用しない
* Go標準ライブラリを優先する
* Exercise間の不要なdependencyを作らない
* TODOは明確にする
* テストから要求を理解できるようにする
* Solutionを用意する
* Solutionには`explanation.md`を付ける
* 過度な抽象化をしない
* 教材コード自体もIdiomatic Goで書く

---

# 29. Exercise作成時の重要ルール

CodexはExerciseを大量生成することを優先してはならない。

優先順位は、

```text
学習目的が明確
    ↓
Exerciseが小さい
    ↓
Testが分かりやすい
    ↓
自分で考える余地がある
    ↓
SolutionがIdiomatic
    ↓
解説で理由が理解できる
```

とする。

単なる穴埋め問題集にしない。

特にIdiomatic GoのExerciseでは、

**「どう書くか」だけではなく「なぜそう書くか」「なぜ別の書き方をしないか」**

を学べるようにする。

---

# 30. Definition of Done

教材全体として以下を満たした状態を完成とする。

* 基本文法から実務コードまで段階的に進める
* 各Exerciseは基本5〜15分で完了できる
* 1 Exercise = 1 Conceptを基本とする
* 全Exerciseにテストが存在する
* TODO状態では対象テストが失敗する
* 正しく実装するとテストが成功する
* Goらしさを学ぶExerciseが存在する
* 「不要なものを削る」Exerciseが存在する
* concurrencyをrace detectorで検証できる
* HTTP / Batch / DBの実務Exerciseが存在する
* 最終的に小規模なPayment APIを完成できる
* Solutionと解説が存在する
* AIなしで学習を進められる
* `gofmt`, `go vet`, `go test`が正常に動作する

最終的な目的は、

> **Goの正解を知っている人ではなく、エディタを開いたら自然にGoを書き始められる人を育てること。**

である。

