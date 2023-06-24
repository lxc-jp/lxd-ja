# リモートサーバーを追加するには

リモートサーバーはLXDコマンドラインクライアント内の概念です。
デフォルトでは、コマンドラインクライアントはローカルのLXDデーモンとやりとりしますが、他のサーバーやクラスタを追加できます。

リモートサーバーの用途の1つはローカルサーバーでインスタンスを作成するのに使えるイメージを配布することです。
詳細は{ref}`remote-image-servers`を参照してください。

完全なLXDサーバーをお使いのクライアントにリモートサーバーとして追加することもできます。
この場合、ローカルのデーモンと同様にリモートサーバーとやりとりできます。
例えば、リモートサーバー上のインスタンスを管理したりサーバー設定を更新できます。

## 認証

LXDサーバーをリモートサーバーとして追加できるようにするには、サーバーのAPIが公開されている必要があります。
それはつまり、[`core.https_address`](server-options-core)サーバー設定オプションが設定されている必要があることを意味します。

サーバーを追加する際は、{ref}`authentication`の方法で認証する必要があります。

詳細は{ref}`server-expose`を参照してください。

## 追加されたリモートを一覧表示する

% Include parts of the content from file [howto/images_remote.md](howto/images_remote.md)
```{include} howto/images_remote.md
   :start-after: <!-- Include start list remotes -->
   :end-before: <!-- Include end list remotes -->
```

## リモートのLXDサーバーを追加する

% Include parts of the content from file [howto/images_remote.md](howto/images_remote.md)
```{include} howto/images_remote.md
   :start-after: <!-- Include start add remotes -->
   :end-before: <!-- Include end add remotes -->
```

## デフォルトのリモートを選択する

LXDコマンドラインクライアントは`local`リモート、つまりローカルのLXDデーモン、に接続する用に初期設定されています。

別のリモートをデフォルトのリモートとして選択するには、以下のように入力します。

    lxc remote switch <remote_name>

どのサーバーがデフォルトのリモートとして設定されているか確認するには、以下のように入力します。

    lxc remote get-default

## グローバルのリモートを設定する

グローバルなシステム毎の設定としてリモートを設定できます。
これらのリモートは、設定を追加したLXDサーバーの全てのユーザーで利用できます。

ユーザーはこれらのシステムで設定されたリモートを(例えば`lxc remote rename`または`lxc remote set-url`を実行することで)オーバーライドできます。
その結果、リモートと対応する証明書がユーザー設定にコピーされます。

グローバルリモートを設定するには、以下のいずれかのディレクトリに置かれた`config.yml`ファイルを編集します。

- (定義されていれば)`LXD_GLOBAL_CONF`で指定されるディレクトリ
- `/var/snap/lxd/common/global-conf/` (snapをお使いの場合)
- `/etc/lxd/` (snap以外の場合)

リモートへの接続用の証明書は同じ場所の`servercerts`ディレクトリ(例えば、`/etc/lxd/servercerts/`)に保管する必要があります。
証明書はリモート名に対応する(例えば、`foo.crt`)必要があります。

以下の設定例を参照してください。

```
remotes:
  foo:
    addr: https://192.0.2.4:8443
    auth_type: tls
    project: default
    protocol: lxd
    public: false
  bar:
    addr: https://192.0.2.5:8443
    auth_type: tls
    project: default
    protocol: lxd
    public: false
```
