# プロジェクト名

このプロジェクトは Go言語のアプリケーションと PostgreSQL データベースを Docker Compose を使ってローカルで実行します。

## 前提条件

*   [Docker](https://docs.docker.com/get-docker/) がインストールされていること。
*   [Docker Compose](https://docs.docker.com/compose/install/) (Docker Desktop には通常含まれています) がインストールされていること。

## セットアップ

1.  **リポジトリのクローン:**
    ```bash
    git clone <リポジトリのURL>
    cd <プロジェクトディレクトリ>
    ```

2.  **環境変数ファイル (.env) の作成:**
    プロジェクトルートにある `.env.example` ファイルをコピーして `.env` ファイルを作成し、必要な環境変数を設定してください。
    ```bash
    cp .env.example .env
    ```
    次に、`.env` ファイルをエディタで開き、データベース接続情報などを自分のローカル環境に合わせて編集します。
    **注意:** `.env` ファイルは機密情報を含むため、Git リポジトリにはコミットしないでください (`.gitignore` に追加されています)。

## 開発環境の起動

以下のコマンドを実行すると、Docker イメージがビルドされ (初回または変更時)、Go アプリケーションと PostgreSQL のコンテナがバックグラウンドで起動します。

```bash
docker compose up -d --build
```

# dbコンテナの中に入る
```
docker compose exec db bash
```
# コンテナ内でpsqlコマンドを実行 (ユーザー名、DB名は.envの値)
```
psql -U ${DB_USER} -d ${DB_NAME}
```
# psqlプロンプトが表示されたら、SQLコマンドを実行できる (例: \dt でテーブル一覧)
# 終了は \q

# コンテナから抜ける
```
exit
```
## Dockerを終了する

```
docker compose down
```
