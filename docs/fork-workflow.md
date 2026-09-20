# 自分のforkで学ぶ

教材本体をupstream、自分のforkをoriginとして扱います。解答は自分のforkに保存し、教材への改善提案だけをupstreamへ送る運用を想定しています。

```bash
git remote add upstream <教材本体のURL>
git switch -c study
```

演習を解いた後は、対象ファイルと学習ノートを選んで自分のcommitに保存します。通常は`exercise.go`を編集し、Testing章では`exercise_test.go`を編集します。「削る」演習では指定された型・関数・ファイルも削除します。

## 上流の教材更新を取り込む

作業を保存してから、上流の差分を確認してmergeします。以下のmainは教材本体のdefault branch名に合わせてください。

```bash
git fetch upstream
git log --oneline HEAD..upstream/main
git merge upstream/main
```

解答済みファイルに上流の変更があればconflictが起きることがあります。自分の解答と新しい要求・テストを比較して解消し、対象の`koans check`を再実行します。解答を消すresetや履歴の書き換えは必要ありません。

## forkのCI

Actionsタブでworkflowを有効にするとCurriculumが教材を検証します。Exercise workflowでは、`01-language/01-variables`のように演習パスを指定して自分の解答をチェックできます。branchを選ぶ欄では解答を保存したbranchを選んでください。

workflowに個人のDB接続情報を登録する必要はありません。CIはジョブ専用のPostgreSQL serviceを使用します。

模範解答はforkにも含まれます。隠しファイルや暗号化ではなく、`solutions/`を意識して開くまで見ないという学習上の区切りです。進捗は自己学習のためのもので、不正防止の試験システムではありません。
