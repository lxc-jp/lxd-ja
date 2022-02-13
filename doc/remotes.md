# Remotes
## イントロダクション <!-- Introduction -->
リモートは LXD のコマンドラインクライアント内の概念でありさまざまな LXD サーバーやクラスターを参照するのに使用します。
リモートは実質的には特定の LXD サーバーをサーバーにログインしたり認証するのに必要な認証情報も含めて指定する URL を指す名前です。
LXD には次の 4 種類のリモートがあります。
<!--
Remotes are a concept in the LXD command line client which are used to refer to various LXD servers or clusters.
A remote is effectively a name pointing to the URL of a particular LXD server as well as needed credentials to login and authenticate the server.
LXD has four types of remotes:
-->

- Static
- Default
- Global (システムごと) <!-- (per-system) -->
- Local (ユーザーごと) <!-- (per-user) -->

### Static
Static リモートは
<!--
Static remotes are:
-->
- local (default)
- ubuntu
- ubuntu-daily

これらはハードコードされておりユーザーが変更できません。
<!--
They are hardcoded and can't be modified by the user.
-->

### Default
初回の使用時に自動的に追加されます。
<!--
Automatically added on first use.
-->

### Global (システムごと) <!-- (per-system) -->
デフォルトでは global の設定ファイルは `/etc/lxc/config.yml` に置かれます。
`LXD_GLOBAL_CONF` 環境変数でパスを変更できます。
この設定ファイルを手動で編集して global リモートを追加できます。
これらのリモートの証明書は `servercerts` ディレクトリ (例: /etc/lxc/servercerts/) に置き、リモートの名前にマッチ (例: `foo.crt`) させます。
<!--
By default the global configuration file is kept in `/etc/lxc/config.yml` or in `LXD_GLOBAL_CONF` if defined.
The configuration file can be manually edited to add global remotes. Certificates for those remotes should be stored inside the `servercerts` directory (e.g. /etc/lxc/servercerts/) and match the remote name (e.g. `foo.crt`).
-->

設定例を以下に示します。
<!--
An example config is below:
-->
```
remotes:
  foo:
    addr: https://10.0.2.4:8443
    auth_type: tls
    project: default
    protocol: lxd
    public: false
  bar:
    addr: https://10.0.2.5:8443
    auth_type: tls
    project: default
    protocol: lxd
    public: false
```

### Local (ユーザーごと) <!-- (per-user) -->
local レベルのリモートは CLI (`lxc`) で次のように管理します。
<!--
Local level remotes are managed from the CLI (`lxc`) with:
-->
`lxc remote [command]`

デフォルトでは local の設定ファイルは `~/.config/lxc/config.yml` に置かれます。
`LXD_CONF` 環境変数でパスを変更できます。
ユーザーはシステムのリモートを (例: `lxc remote name` や `lxc remote set-url` を実行することで) オーバーライドすることができます。
その場合リモートの設定は関連する証明書と共にコピーされます。
<!--
By default the configuration file is kept in `~/.config/lxc/config.yml` or in `LXD_CONF` if defined.
Users have the possibility to override system remotes (e.g. by running `lxc remote rename` or `lxc remote set-url`)
which results in the remote being copied to their own config, including any associated certificates.
-->
