# サーバ設定 <!-- Server configuration -->
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

 - `core` コア・デーモン設定 <!-- (core daemon configuration) -->
 - `images` イメージ設定 <!-- (image configuration) -->
 - `maas` MAAS 統合 <!-- (MAAS integration) -->

キー <!-- Key -->                             | 型 <!-- Type -->      | デフォルト値 <!-- Default -->   | API 拡張 <!-- API extension -->            | 説明 <!-- Description -->
:--                             | :---      | :------   | :------------            | :----------
backups.compression\_algorithm  | string    | gzip      | backup\_compression      | 新規のイメージに用いる圧縮アルゴリズム (bzip2, gzip, lzma, xz, none のいずれか) <!-- Compression algorithm to use for new images (bzip2, gzip, lzma, xz or none) -->
candid.api.key                  | string    | -         | candid\_config\_key      | Candid サーバーの公開鍵（HTTPのみのサーバーで必要） <!-- Public key of the candid server (required for HTTP-only servers) -->
candid.api.url                  | string    | -         | candid\_authentication   | Candid を使用する外部認証エンドポイントの URL <!-- URL of the the external authentication endpoint using Candid -->
candid.expiry                   | integer   | 3600      | candid\_config           | Canded macaroon の有効期間 (秒で指定) <!-- Candid macaroon expiry in seconds -->
candid.domains                  | string    | -         | candid\_config           | 許可される Candid ドメインのカンマ区切りリスト (空文字は全てのドメインが有効という意味になります) <!-- Comma-separated list of allowed Candid domains (empty string means all domains are valid) -->
cluster.offline\_threshold      | integer   | 20        | clustering               | 無反応なノードをオフラインとみなす秒数 <!-- Number of seconds after which an unresponsive node is considered offline -->
core.debug\_address             | string    | -         | pprof\_http              | pprof デバッグサーバがバインドするアドレス (HTTP) <!-- Address to bind the pprof debug server to (HTTP) -->
core.https\_address             | string    | -         | -                        | リモート API がバインドするアドレス (HTTPs) <!-- Address to bind for the remote API (HTTPs) -->
core.https\_allowed\_credentials| boolean   | -         | -                        | Access-Control-Allow-Credentials HTTP ヘッダの値を "true" にするかどうか <!-- Whether to set Access-Control-Allow-Credentials http header value to "true" -->
core.https\_allowed\_headers    | string    | -         | -                        | Access-Control-Allow-Headers HTTP ヘッダの値 <!-- Access-Control-Allow-Headers http header value -->
core.https\_allowed\_methods    | string    | -         | -                        | Access-Control-Allow-Methods HTTP ヘッダの値 <!-- Access-Control-Allow-Methods http header value -->
core.https\_allowed\_origin     | string    | -         | -                        | Access-Control-Allow-Origin HTTP ヘッダの値 <!-- Access-Control-Allow-Origin http header value -->
core.proxy\_https               | string    | -         | -                        | HTTPS プロキシを使用する場合はその URL (未指定の場合は HTTPS\_PROXY 環境変数を参照) <!-- https proxy to use, if any (falls back to HTTPS\_PROXY environment variable) -->
core.proxy\_http                | string    | -         | -                        | HTTP プロキシを使用する場合はその URL (未指定の場合は HTTP\_PROXY 環境変数を参照) <!-- http proxy to use, if any (falls back to HTTP\_PROXY environment variable) -->
core.proxy\_ignore\_hosts       | string    | -         | -                        | プロキシが不要なホスト (NO\_PROXY と同様な形式、例えば 1.2.3.4,1.2.3.5, を指定。未指定の場合は NO\_PROXY 環境変数を参照) <!-- hosts which don't need the proxy for use (similar format to NO\_PROXY, e.g. 1.2.3.4,1.2.3.5, falls back to NO\_PROXY environment variable) -->
core.trust\_password            | string    | -         | -                        | 信頼を確立するためにクライアントに要求するパスワード <!-- Password to be provided by clients to setup a trust -->
images.auto\_update\_cached     | boolean   | true      | -                        | LXD がキャッシュしているイメージを自動的に更新するかどうか <!-- Whether to automatically update any image that LXD caches -->
images.auto\_update\_interval   | integer   | 6         | -                        | キャッシュされているイメージが更新されているかチェックする間隔を時間単位で指定 <!-- Interval in hours at which to look for update to cached images (0 disables it) -->
images.compression\_algorithm   | string    | gzip      | -                        | 新しいイメージに使用する圧縮アルゴリズム (bzip2, gzip, lzma, xz あるいは none) <!-- Compression algorithm to use for new images (bzip2, gzip, lzma, xz or none) -->
images.remote\_cache\_expiry    | integer   | 10        | -                        | キャッシュされたが未使用のイメージを破棄するまでの日数 <!-- Number of days after which an unused cached remote image will be flushed -->
maas.api.key                    | string    | -         | maas\_network            | MAAS を管理するための API キー <!-- API key to manage MAAS -->
maas.api.url                    | string    | -         | maas\_network            | MAAS サーバの URL <!-- URL of the MAAS server -->
maas.machine                    | string    | hostname  | maas\_network            | この LXD ホストの MAAS での名前 <!-- Name of this LXD host in MAAS -->

これらのキーは lxc コマンドで次のように設定します。 <!-- Those keys can be set using the lxc tool with: -->

```bash
lxc config set <key> <value>
```
