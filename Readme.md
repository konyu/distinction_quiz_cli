# Distinction Quiz CLI
Distinction Quiz CLIは、エクセルファイルからクイズを生成し、コマンドラインで実行するGo製のアプリです。
Distinctionとついている通り、https://distinction.atsueigo.com/ の単語の勉強をスキマ時間で行うためのものです。


https://github.com/konyu/distinction_quiz_cli/assets/1217706/18a80ce2-b6d0-4b32-a0b8-6142cb1d1978



## 特徴

- エクセルファイルからクイズデータを読み込む
- クイズの問題数を指定可能
- 乱数のシード値を指定可能

## 単語帳の元データについて
著作権の関係上Distinctionのデータは共有できません。
エクセルファイルを自分で作成する必要があります。
Distinctionまたは、自身で勉強したい単語帳などのデータを作成してください。

プロジェクトのルート直下にある `spreadsheet.xlsx` を置き換えてください。

### エクセルファイルのフォーマット
- \# 番号:
- Word :単語
- Translation: 訳

| #   | Word                | Translation                  |
| --- | ------------------- | ---------------------------- |
| 1   | low key             | 控えめな、秘密な             |
| 2   | drive sb up the wall| 〜をイライラさせる           |
| 3   | black and white     | 明らかな、白黒はっきりしている |

エクセルファイルはN枚のシートに対応しています。上記のフォーマットで2枚目3枚目のシートを作成してもクイズを生成することができます。

## 使用方法

以下のコマンドでプロジェクトをビルドします：

```sh
go build -o bin/distinction_quiz_cli
```

## 以下のコマンドでクイズを実行します：

```
./bin/distinction_quiz_cli
```

## オプション
- --seed : 乱数の生成するための文字列を指定します。文字列を指定することで同じ問題が出題されます。
- --xlsx : 読み込むエクセルファイルのパスを指定します。デフォルトはspreadsheet.xlsxです。
- --num : 生成するクイズの問題数を指定します。デフォルトは10です。

## プロジェクト構成
- main.go：アプリケーションのエントリーポイントです。
- quiz/quiz.go：クイズの生成と実行を担当します。
- sheets/sheets.go：エクセルファイルのデータを取得します。
- utils/utils.go：ユーティリティ関数を提供します。

## ライセンス
MIT
