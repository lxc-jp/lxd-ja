# 環境変数

以下の環境変数を設定することで、LXDのクライアントとデーモンをユーザーの環境に適合させることができ、いくつかの高度な機能を有効または無効にすることができます。

```{note}
snap版のLXDをお使いの場合はこれらの環境変数は利用できません。
```

## クライアントとサーバー共通の環境変数

名前          | 説明
:---          | :----
`LXD_DIR`     | LXDのデータディレクトリ
`PATH`        | 実行ファイルの検索対象のパスのリスト
`http_proxy`  | HTTP用のプロキシサーバーのURL
`https_proxy` | HTTPS用のプロキシサーバーのURL
`no_proxy`    | プロキシが不要なドメイン、IPアドレスあるいはCIDRレンジのリスト

## クライアントの環境変数

名前              | 説明
:---              | :----
`EDITOR`          | 使用するテキストエディタ
`VISUAL`          | (`EDITOR` が設定されてないときに)使用するテキストエディタ
`LXD_CONF`        | LXC設定ディレクトリーのパス
`LXD_GLOBAL_CONF` | LXCグローバル設定ディレクトリーのパス
`LXC_REMOTE`      | 使用するリモートの名前（設定されたデフォルトのリモートよりも優先されます）

## サーバーの環境変数

名前                            | 説明
:---                            | :----
`LXD_EXEC_PATH`                 | (サブコマンド実行時に使用される)LXD実行ファイルのフルパス
`LXD_LXC_TEMPLATE_CONFIG`       | LXCテンプレート設定ディレクトリ
`LXD_SECURITY_APPARMOR`         | `false`に設定するとAppArmorを無効にします
`LXD_UNPRIVILEGED_ONLY`         | `true`に設定すると非特権コンテナしか作れなくなるように強制します。LXD_UNPRIVILEGED_ONLYを設定する前に作られた特権コンテナは引き続き特権を持つことに注意してください。このオプションをLXDデーモンを最初にセットアップするときに設定するのが実用的です。
`LXD_OVMF_PATH`                 | `OVMF_CODE.fd`と`OVMF_VARS.ms.fd`を含むOVMFビルドへのパス
`LXD_SHIFTFS_DISABLE`           | `shiftfs`のサポートを無効にする(従来のUIDシフトを試す際に有用です)
`LXD_IDMAPPED_MOUNTS_DISABLE`   | idmapを使ったマウントを無効にする(従来のUIDシフトを試す際に有用です)
`LXD_DEVMONITOR_DIR`            | デバイスモニターでモニターするパス。主にテスト用。
