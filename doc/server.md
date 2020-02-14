# サーバ設定
<!-- Server configuration -->
サーバ設定は単純なキーと値の組です。
<!--
The server configuration is a simple set of key and values.
-->

key/value 設定は現在サポートされている以下のネームスペースによって
名前空間が分けられています。
<!--
The key/value configuration is namespaced with the following namespaces
currently supported:
-->

 - `backups` バックアップ設定 <!-- (backups configuration) -->
 - `candid` Candid 認証の統合 <!-- (Candid authentication integration) -->
 - `cluster` クラスタ設定 <!-- (cluster configuration) -->
 - `core` コア・デーモン設定 <!-- (core daemon configuration) -->
 - `images` イメージ設定 <!-- (image configuration) -->
 - `maas` MAAS 統合 <!-- (MAAS integration) -->
 - `rbac`  ロールベースのアクセス制御の統合 <!-- (Role Based Access Control integration) -->

キー <!-- Key -->                   | 型 <!-- Type --> | スコープ <!-- Scope --> | デフォルト値 <!-- Default -->   | API 拡張 <!-- API extension -->            | 説明 <!-- Description -->
:--                                 | :---      | :----     | :------   | :------------            | :----------
backups.compression\_algorithm      | string    | global    | gzip      | backup\_compression      | 新規のイメージに用いる圧縮アルゴリズム (bzip2, gzip, lzma, xz, none のいずれか) <!-- Compression algorithm to use for new images (bzip2, gzip, lzma, xz or none) -->
candid.api.key                      | string    | global    | -         | candid\_config\_key      | Candid サーバーの公開鍵（HTTPのみのサーバーで必要） <!-- Public key of the candid server (required for HTTP-only servers) -->
candid.api.url                      | string    | global    | -         | candid\_authentication   | Candid を使用する外部認証エンドポイントの URL <!-- URL of the the external authentication endpoint using Candid -->
candid.expiry                       | integer   | global    | 3600      | candid\_config           | Canded macaroon の有効期間 (秒で指定) <!-- Candid macaroon expiry in seconds -->
candid.domains                      | string    | global    | -         | candid\_config           | 許可される Candid ドメインのカンマ区切りリスト (空文字は全てのドメインが有効という意味になります) <!-- Comma-separated list of allowed Candid domains (empty string means all domains are valid) -->
cluster.https\_address              | string    | local     | -         | clustering\_server\_address       | クラスタのトラフィックに使用すべきサーバのアドレス <!-- Address the server should using for clustering traffic -->
cluster.offline\_threshold          | integer   | global    | 20        | clustering               | 無反応なノードをオフラインとみなす秒数 <!-- Number of seconds after which an unresponsive node is considered offline -->
cluster.images\_minimal\_replica    | integer   | global    | 3         | clustering\_image\_replication    | 特定のイメージのコピーを持つべきクラスタメンバの最小数 (リプリケーションなしは 1 を、全メンバにコピーは -1 を設定) <!-- Minimal numbers of cluster members with a copy of a particular image (set 1 for no replication, -1 for all members) -->
cluster.max\_voters                 | integer   | global    | 3         | clustering\_sizing                | データベースの投票者の役割を割り当てられるクラスターメンバーの最大数 <!-- Maximum number of cluster members that will be assigned the database voter role -->
cluster.max\_standby                | integer   | global    | 2         | clustering\_sizing                | データベースのスタンバイの役割を割り当てられるクラスターメンバーの最大数 <!-- Maximum number of cluster members that will be assigned the database stand-by role -->
core.debug\_address                 | string    | local     | -         | pprof\_http              | pprof デバッグサーバがバインドするアドレス (HTTP) <!-- Address to bind the pprof debug server to (HTTP) -->
core.https\_address                 | string    | local     | -         | -                        | リモート API がバインドするアドレス (HTTPS) <!-- Address to bind for the remote API (HTTPS) -->
core.https\_allowed\_credentials    | boolean   | global    | -         | -                        | Access-Control-Allow-Credentials HTTP ヘッダの値を "true" にするかどうか <!-- Whether to set Access-Control-Allow-Credentials http header value to "true" -->
core.https\_allowed\_headers        | string    | global    | -         | -                        | Access-Control-Allow-Headers HTTP ヘッダの値 <!-- Access-Control-Allow-Headers http header value -->
core.https\_allowed\_methods        | string    | global    | -         | -                        | Access-Control-Allow-Methods HTTP ヘッダの値 <!-- Access-Control-Allow-Methods http header value -->
core.https\_allowed\_origin         | string    | global    | -         | -                        | Access-Control-Allow-Origin HTTP ヘッダの値 <!-- Access-Control-Allow-Origin http header value -->
core.proxy\_https                   | string    | global    | -         | -                        | HTTPS プロキシを使用する場合はその URL (未指定の場合は HTTPS\_PROXY 環境変数を参照) <!-- https proxy to use, if any (falls back to HTTPS\_PROXY environment variable) -->
core.proxy\_http                    | string    | global    | -         | -                        | HTTP プロキシを使用する場合はその URL (未指定の場合は HTTP\_PROXY 環境変数を参照) <!-- http proxy to use, if any (falls back to HTTP\_PROXY environment variable) -->
core.proxy\_ignore\_hosts           | string    | global    | -         | -                        | プロキシが不要なホスト (NO\_PROXY と同様な形式、例えば 1.2.3.4,1.2.3.5, を指定。未指定の場合は NO\_PROXY 環境変数を参照) <!-- hosts which don't need the proxy for use (similar format to NO\_PROXY, e.g. 1.2.3.4,1.2.3.5, falls back to NO\_PROXY environment variable) -->
core.trust\_password                | string    | global    | -         | -                        | 信頼を確立するためにクライアントに要求するパスワード <!-- Password to be provided by clients to setup a trust -->
images.auto\_update\_cached         | boolean   | global    | true      | -                        | LXD がキャッシュしているイメージを自動的に更新するかどうか <!-- Whether to automatically update any image that LXD caches -->
images.auto\_update\_interval       | integer   | global    | 6         | -                        | キャッシュされているイメージが更新されているかチェックする間隔を時間単位で指定 <!-- Interval in hours at which to look for update to cached images (0 disables it) -->
images.compression\_algorithm       | string    | global    | gzip      | -                        | 新しいイメージに使用する圧縮アルゴリズム (bzip2, gzip, lzma, xz あるいは none) <!-- Compression algorithm to use for new images (bzip2, gzip, lzma, xz or none) -->
images.remote\_cache\_expiry        | integer   | global    | 10        | -                        | キャッシュされたが未使用のイメージを破棄するまでの日数 <!-- Number of days after which an unused cached remote image will be flushed -->
maas.api.key                        | string    | global    | -         | maas\_network            | MAAS を管理するための API キー <!-- API key to manage MAAS -->
maas.api.url                        | string    | global    | -         | maas\_network            | MAAS サーバの URL <!-- URL of the MAAS server -->
maas.machine                        | string    | local     | hostname  | maas\_network            | この LXD ホストの MAAS での名前 <!-- Name of this LXD host in MAAS -->
rbac.agent.url                      | string    | global    | -         | rbac                              | RBAC 登録中に提供される Candid エージェントの URL <!-- The Candid agent url as provided during RBAC registration -->
rbac.agent.username                 | string    | global    | -         | rbac                              | RBAC 登録中に提供される Candid エージェントのユーザ名 <!-- The Candid agent username as provided during RBAC registration -->
rbac.agent.public\_key              | string    | global    | -         | rbac                              | RBAC 登録中に提供される Candid エージェントの公開鍵 <!-- The Candid agent public key as provided during RBAC registration -->
rbac.agent.private\_key             | string    | global    | -         | rbac                              | RBAC 登録中に提供される Candid エージェントの秘密鍵 <!-- The Candid agent private key as provided during RBAC registration -->
rbac.api.expiry                     | integer   | global    | -         | rbac                              | RBAC の macaroon の有効期限 (秒) <!-- RBAC macaroon expiry in seconds -->
rbac.api.key                        | string    | global    | -         | rbac                              | RBAC サーバの公開鍵 (HTTP のみ有効なサーバで必要) <!-- Public key of the RBAC server (required for HTTP-only servers) -->
rbac.api.url                        | string    | global    | -         | rbac                              | 外部の RBAC サーバの URL <!-- URL of the external RBAC server -->
storage.backups\_volume             | string    | local     | -         | daemon\_storage                   | バックアップの tarball を保管するのに使用するボリューム (POOL/VOLUME 形式で指定) <!-- Volume to use to store the backup tarballs (syntax is POOL/VOLUME) -->
storage.images\_volume              | string    | local     | -         | daemon\_storage                   | イメージの tarball を保管するのに使用するボリューム (POOL/VOLUME 形式で指定) <!-- Volume to use to store the image tarballs (syntax is POOL/VOLUME) -->

これらのキーは lxc コマンドで次のように設定します。 <!-- Those keys can be set using the lxc tool with: -->

```bash
lxc config set <key> <value>
```

クラスタの一部として動作するときは、上記の表でスコープが `global` のキーは全てのクラスタメンバーに即座に反映されます。スコープが `local` のキーはコマンドラインツールの `\-\-target` オプションを使ってメンバーごとに設定する必要があります。
<!--
When operating as part of a cluster, the keys marked with a `global`
scope will immediately be applied to all the cluster members. Those keys
with a `local` scope must be set on a per member basis using the
`\-\-target` option of the command line tool.
-->
