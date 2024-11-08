# miseとは

miseは、開発環境のバージョン管理とタスクランナーの機能を統合したモダンなツールです。asdfの後継として開発され、Rustで実装されているため高速な動作が特徴です。元々、RTXという名前でしたがmiseに変更されました。direnvやasdfを使ったことがあれば、高速な動作性能が気に入るはずです。

Home | mise-en-place https://mise.jdx.dev/

## miseの特徴とメリット

miseには以下のような特徴があり、開発環境の管理を効率化できます

- 統合的な環境管理機能
  - Node.js、Python等の異なる言語を一元管理
  - 環境変数(.env)の管理
  - asdfとの互換性を維持

- 使いやすい設定システム
  - TOMLフォーマットによる読みやすい設定
  - 環境変数とツールのバージョンを同一ファイルで管理

- 強力なタスクランナー機能
  - 依存関係を持つタスクの定義が可能
  - シェルスクリプトとの連携が容易

- 高いパフォーマンス
  - Rustで実装されたネイティブバイナリ
  - asdf（シェル実装）と比べて低オーバーヘッドで高速

# インストール方法

```bash
# Homebrewでインストール
brew install mise

# または直接インストール
curl https://mise.run | sh

# シェルの設定（zshの場合）
echo 'eval "$(~/.local/bin/mise activate zsh)"' >> ~/.zshrc
```

## Node.jsのインストール方法

よく使用するコマンドの例を以下に示します。

```bash
# インストール可能なバージョンの一覧表示
mise ls-remote node

# インストール済みのバージョン一覧
mise ls node

# 最新のLTSバージョンをインストール
mise use node@lts

# 特定のバージョンをインストール
mise use node@20.14.0

# グローバルで使用する場合
mise use -g node@20.14.0

```

# 基本的な設定例

nodejsを使うだけならこれコピペで大丈夫です。僕はこれ以外にもJava, Rubyを管理しています。

```toml:.mise.toml
[env]
_.file = [".env"]        # プロジェクトの環境変数ファイルを読み込み
_.path = ['./node_modules/.bin']  # ローカルのnode_modulesにパスを通す

[tools]
node = '20'  # メジャーバージョンのみ指定（20.x.xの最新版を自動選択）
```

# タスク管理

僕はおまけ程度でしか使ってませんが、プロジェクトのクリーンナップ用のシェルが書きたくなったらmiseで書いたりします。
nodeの世界以外のタスクを管理するイメージです。docker操作、成果物の一括削除等
具体的な使用例として、React NativeのプロジェクトでRailsとJava, xcode, Androidの世界線で操作する必要がありmiseにまとめたりしてました。

## 設定ファイル

```toml
[env]
_.file = [".env"]
_.path = ['./node_modules/.bin']

[tools]
node = '20.11.0'

[tasks.clean]
run = '''
#!/usr/bin/env bash
rm -rf node_modules .next
yarn
'''

[tasks.format]
run = "yarn eslint"

[tasks."build:node"]
run = "yarn build"
depends = ["format"]

[tasks.ci]
depends = ["build:node", "build:docker"]
```

## タスク実行例

```bash
# クリーンアップ
mise run clean

# ビルド
mise run build:docker

# CI用のタスク実行
mise run ci
```
