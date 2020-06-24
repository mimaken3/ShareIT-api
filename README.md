# ShareIT-api

## IT に関する情報共有サイト

ニュースや記事などで得た IT に関する知見、プログラミングでインプットした内容やエラーの解決法などをメモとして、ツイート感覚で投稿し共有することが出来るサイトです。<br>
主にちょっとしたメモ程度で、投稿する敷居が低いものなので、アウトプットの場としてぜひ使ってみて下さい！<br>
[ShareIT](https://shareit.fun)

![GIF](https://user-images.githubusercontent.com/29462808/85257058-282c9980-b4a0-11ea-8266-0d6eba9a2a6e.gif)

## アーキテクチャ

![ShareIT-Architecture](https://user-images.githubusercontent.com/29462808/85223849-f57f9400-b400-11ea-8ff0-27418d34ccd1.png)

# 主な機能

- 記事とそれに関するトピックのラベルを付けて投稿 ✍️
- 投稿された記事に対してのコメント 📝、いいね ❤️
- 記事に付けられたコメントといいねの通知(push 通知は未実装 🙅‍♂️)
- ユーザとトピックを指定した検索機能 🔍

# 実際にローカルで動かしてみる場合

Git , Docker, Docker compose がインストールされてる前提で進めていきます。

## 手順

1. GitHub のリポジトリから git clone する
2. Amazon S3 でユーザのアイコンを保存するための設定
3. 環境設定ファイルへの記述
4. Docker compose で起動

#### 1. GitHub のリポジトリから git clone する(本リポジトリのではなく[別のリポジトリ](https://github.com/mimaken3/ShareIT-docker-dev)のを clone)

```bash
$ git clone https://github.com/mimaken3/ShareIT-docker-dev
```

clone したディレクトリに移動

```bash
$ cd ShareIT-docker-dev
```

#### 2. Amazon S3 でユーザのアイコンを保存するための設定

1.  まずは S3 へ[ログインまたはサインイン](https://aws.amazon.com/jp/s3/getting-started/)して下さい。

2.  次に S3 のコンソールページへ移動して、`バケットを作成する` ボタンをクリックすると<br>
    `名前とリージョンタブ`が開かれるので以下のように入力

    - `バケット名`はユニークな名前を入力
    - `リージョン`はアジアパシフィック（東京）
    - `既存のバケットから設定をコピー`は何もしない
    - `次ヘ`ボタンをクリック

3.  次に`オプションの設定`タブが開かれますが、こちらは特に何も入力、設定せずに`次へ`ボタンを押す
4.  `アクセス許可の設定`タブが開かれるので以下のように設定

    - `パブリックアクセスをすべてブロック`のチェックを外し、<br>
      `現在の設定により、このバケットと中のオブジェクトがパブリックになる可能性があることを了承します。`にチェックを入れる。<br>
      ※注意 オブジェクトが公開される状態になるので、ShareIT のローカル動作が終わったらバケットを削除するか、<br>
      `パブリックアクセスをすべてブロック`に変更するようにして下さい。
    - `次ヘ`ボタンをクリック

5.  `確認`タブが開かれるので`バケットを作成`ボタンをクリック

6.  バケット一覧に先程作成したバケットが表示されるので、そのバケットをダブルクリックして移動。

7.  `フォルダを作成`ボタンをクリックし`users-icons`という名前でフォルダ名を入力し`保存`ボタンをクリック。

8.  フォルダが作成されたので、ダブルクリックしてそのフォルダに移動。

9.  ユーザの初期アイコンを以下のようにしてアップロード

    - Google 検索で初期アイコンで使えそうな画像を探してを`default.png`という名前でローカル（自分の PC）に保存
    - `users-icons`フォルダで`アップロード`ボタンをクリックして`ファイルを追加`、保存した画像を選択し、`アップロード`ボタンをクリック。

10. [CORS](https://developer.mozilla.org/ja/docs/Web/HTTP/CORS)の設定を行います。

    - `アクセス権限`タブにある`CORSの設定`をクリック
    - CORS 構成エディタが開かれるので下記を記述

```
    <?xml version="1.0" encoding="UTF-8"?>
    <CORSConfiguration xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
    <CORSRule>
        <AllowedOrigin>*</AllowedOrigin>
        <AllowedMethod>GET</AllowedMethod>
        <AllowedMethod>PUT</AllowedMethod>
        <AllowedMethod>POST</AllowedMethod>
        <AllowedMethod>DELETE</AllowedMethod>
        <MaxAgeSeconds>3000</MaxAgeSeconds>
        <AllowedHeader>*</AllowedHeader>
    </CORSRule>
    </CORSConfiguration>
```

11. 次にアクセスキーを取得します。[こちら](https://tech-blog.s-yoshiki.com/entry/135)を参考にして`アクセスキーID`と`シークレットアクセスキー`を取得。

#### 3. 環境設定ファイルへの記述

S3 の情報を Golang(API サーバ用), React.js(AP サーバ用)の環境設定ファイルにそれぞれ記述します。

- `ShareIT-docker-dev/docker/react-front/.env/`にある`dev.js`ファイルを開きます。

先程と同様、バケット名とアクセスキーを下記のダブルクオート内にそれぞれ記述。

```
AWS_S3_ACCESS_KEY_ID: "",
AWS_S3_SECRET_ACCESS_KEY: "",
AWS_S3_BUCKET_NAME: "",
```

- `ShareIT-docker-dev/docker/golang/`にある`.env`ファイルを開きます。

手順 3 のステップ 2 で作成したバケット名、ステップ 11 で取得したアクセスキーを下記に**ダブルクオートなし**でそれぞれ記述。また、SECRET_KEY には任意の文字列を入力。

```
AWS_S3_ACCESS_KEY_ID =
AWS_S3_SECRET_ACCESS_KEY =
DEV_BUCKET_NAME =
SECRET_KEY =
```

#### 4. Docker compose で起動

Docker compose を使って Go 、React 、MySQL コンテナを立ち上げます。<br>
(docker-compose は docker-compose.yml と同じ階層、つまり git clone した ShareIT-docker-dev ディレクトリ配下で行います。)

```bash
$ docker-compose build
```

※ 初回ビルドは PC のスペックやネットワーク環境にもよりますが 20〜30 分ほどかかります。

ビルドが完了したら起動します。

```bash
$ docker-compose up
```

Go, React はそれぞれ以下のように表示されたら起動が完了したことになります。<br>
(MySQL のプロセスが起動したあとに Go のプロセスが起動するようになっているため、確認するのは Go と React のみ)<br>

Go<br>
<img width="708" alt="Screen Shot 2020-06-18 at 18 47 43" src="https://user-images.githubusercontent.com/29462808/85006094-a6352b80-b194-11ea-8c41-3acf0465747e.png">

React
<img width="1427" alt="Screen Shot 2020-06-18 at 18 48 15" src="https://user-images.githubusercontent.com/29462808/85006099-a8978580-b194-11ea-9584-e0e715b1d660.png">

これで起動出来ました 🎉 <br>
http://localhost:8088 にアクセスすると以下の様なログイン画面に飛ぶはずです。

<img width="1436" alt="Screen Shot 2020-06-19 at 17 15 37" src="https://user-images.githubusercontent.com/29462808/85112414-90853c00-b250-11ea-81b5-2d57924ee013.png">

ちなみに、MySQL コンテナの起動時に、180 個ほどのトピックと admin ユーザのデータがインサートされています。<br>
なので管理者としてログインしたい場合は、<br>
**ユーザ名: admin**<br>
**パスワード: password**<br>
でログイン出来ます。<br>

コードを変更する場合、 Go, React 共に[fresh](https://github.com/gravityblast/fresh)と[webpack](https://webpack.js.org/)でホットリロードが効いているので、<br>
ファイルを保存したら自動でブラウザに反映されます。

使い終わったら Ctrl+C で止めます。

# 削除

今回インストールしたコンテナ、イメージ、ボリューム、ネットワークを削除する場合は、<br>

```bash
$ docker-compose down --rmi all --volumes
```

で一括で全削除出来ます。
以下参考<br>
https://qiita.com/suin/items/19d65e191b96a0079417

# 最後に

まだまだ足りない実装や課題点などありますので、少しずつ改善していくつもりです。<br>
また、フロント側(React)でのバグや質問、追加実装などございましたら[こちら](https://github.com/mimaken3/ShareIT-front)のリポジトリに、<br>
サーバ側(Go)の場合[こちら](https://github.com/mimaken3/ShareIT-api)のリポジトリに issue を立てていただければ幸いです 🙇‍♂️

では！

