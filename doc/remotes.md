# リモート

## イントロダクション

リモートは LXD のコマンドラインクライアント内の概念でありさまざまな LXD サーバやクラスタを参照するのに使用します。
リモートは実質的には特定の LXD サーバをサーバにログインしたり認証するのに必要な認証情報も含めて指定する URL を指す名前です。
LXD には次の 4 種類のリモートがあります。

- Static
- Default
- Global (システムごと)
- Local (ユーザーごと)

### Static

Static リモートは
- `local` (default)
- `ubuntu`
- `ubuntu-daily`

これらはハードコードされておりユーザーが変更できません。

### Default

初回の使用時に自動的に追加されます。

### Global (システムごと)

デフォルトでは global の設定ファイルは `/etc/lxd/config.yml`、 または snap の場合は `/var/snap/lxd/common/global-conf/`、または `LXD_GLOBAL_CONF` 環境変数が定義されていればそのパスが使用されます。
この設定ファイルを手動で編集して global リモートを追加できます。
これらのリモートの証明書は `servercerts` ディレクトリ (例: `/etc/lxd/servercerts/`) に置き、リモートの名前にマッチ (例: `foo.crt`) させます。

設定例を以下に示します。

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

### Local (ユーザーごと)

local レベルのリモートは CLI (`lxc`) で次のように管理します。
`lxc remote [command]`

デフォルトでは local の設定ファイルは `~/.config/lxc/config.yml` に置かれます。
`LXD_CONF` 環境変数でパスを変更できます。
ユーザーはシステムのリモートを (例: `lxc remote name` や `lxc remote set-url` を実行することで) オーバーライドすることができます。
その場合リモートの設定は関連する証明書と共にコピーされます。
