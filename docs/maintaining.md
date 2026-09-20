# 教材を検証・追加する

## 未着手と完成の区別

未着手の演習はコンパイルできますがテストが失敗します。`go vet ./...`も未着手の状態で通ります。race修正の演習には意図したdata raceがあり、`go test -race`で失敗します。

模範解答は`solutions/<演習>/solution.go`にあり、`//go:build ignore`で通常のbuildから除外しています。Testing章では`solution_test.go`が模範テストです。`koans verify`はリポジトリを一時コピーし、解答で対象ファイルを置換してから検証します。学習者の解答は上書きしません。

```bash
make fmt-check
go vet ./...
go run ./cmd/koans verify
```

`verify`では模範解答に対してvet、staticcheck、race付きテストを実行します。Testing章はさらに欠陥実装でテストが失敗することを確認します。DBなしの実行はDBテストをSKIPするため、完全な検証には[DB設定](database.md)が必要です。

教材本体のリリース前には、まだ解答していないcheckoutで次も実行します。

```bash
go run ./cmd/koans verify --starters
```

これは各演習の未着手テストが実際に失敗することを確認します。学習者のforkでは解答済み演習が成功するので、通常のCurriculum CIには`--starters`を付けません。

## 演習を追加する

1. `spec.md`の粒度に従い、具体的な入力・出力・境界値を決める。
2. `exercise.go`、`exercise_test.go`、`README.md`、`HINTS.md`を作る。
3. `solutions/`に対応する解答と`explanation.md`を置く。
4. `curriculum.json`へパス・学習目標・DB要否・kindを追加する。
5. 解答でテストが通ることを確かめる。
6. 対象の振る舞いを壊し、テストのassertionが失敗することを確かめる。
7. 章のREADMEと全演習一覧を更新する。

各演習は独立したGo packageです。入力型を共有するだけのために隣の演習へ依存させません。DBのテスト準備だけは`internal/dbtest`を共通利用します。共有helperは学習者が解くコードへ業務上の設計を押し込みません。

「削る」演習の構造チェックは課題で指定した削除を検証します。通常の実装問題ではコードの表記を固定せず、振る舞いを検証してください。

## Visible・追加・非公開テスト

リポジトリ内のテストはforkした人から読めます。公開リポジトリだけで本当にHiddenなテストを提供することはできません。この教材ではVisible Testsに境界値も含め、Testing章では`mutants/*.go.txt`を公開して、何を検出すべきかを確認できるようにしています。

運営者が非公開の追加テストを使う場合は、別管理の採点環境で受講者のcheckoutへ`*_test.go`を追加し、対象packageのテストを実行します。公開workflowへ秘密のテストを取得するtokenを渡す方式にはしません。受講者のコードとworkflowは受講者が変更できるため、公開forkのCIを試験の信頼境界にはできません。

テストを読めることは学習の一部です。テストだけでは読みやすさ、不要な抽象化、適切な依存の置き方を全て判定できません。解答の解説と振り返りも修了条件に含めてください。

## バージョン

Goは`go.mod`と`.tool-versions`、staticcheckは`go.mod`のtool dependencyに固定します。更新時は[Goの公式リリース](https://go.dev/dl/?mode=json)を確認し、全検証を再実行します。CIのGo versionは`go.mod`から読みます。
