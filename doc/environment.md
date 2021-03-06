# 環境変数
<!-- Environment variables -->
## イントロダクション <!-- Introduction -->
以下の環境変数を設定することで、 LXD のクライアントとデーモンを
ユーザーの環境に適合させることができ、いくつかの高度な機能を有効または
無効にすることができます。
<!--
The LXD client and daemon respect some environment variables to adapt to
the user's environment and to turn some advanced features on and off.
-->

## クライアントとサーバ共通の環境変数 <!-- Common -->
名前 <!-- Name -->                           | 説明 <!-- Description -->
:---                            | :----
`LXD_DIR`                       | LXD のデータディレクトリ <!-- The LXD data directory -->
`PATH`                          | 実行ファイルの検索対象のパスのリスト <!-- List of paths to look into when resolving binaries -->
`http_proxy`                    | HTTP 用のプロキシサーバの URL <!-- Proxy server URL for HTTP -->
`https_proxy`                   | HTTPS 用のプロキシサーバの URL <!-- Proxy server URL for HTTPS -->
`no_proxy`                      | プロキシが不要なドメイン、IPアドレスあるいは CIDR レンジのリスト <!-- List of domains, IP addresses or CIDR ranges that don't require the use of a proxy -->

## クライアントの環境変数 <!-- Client environment variable -->
名前 <!-- Name -->                           | 説明 <!-- Description -->
:---                            | :----
`EDITOR`                        | 使用するテキストエディタ <!-- What text editor to use -->
`VISUAL`                        | (`EDITOR` が設定されてないときに) 使用するテキストエディタ <!-- What text editor to use (if `EDITOR` isn't set) -->
`LXD_CONF`                      | LXC 設定ディレクトリーのパス <!-- Path to the LXC configuration directory -->
`LXD_GLOBAL_CONF`               | LXC グローバル設定ディレクトリーのパス <!-- Path to the global LXC configuration directory -->

## サーバの環境変数 <!-- Server environment variable -->
名前 <!-- Name -->                           | 説明 <!-- Description -->
:---                            | :----
`LXD_EXEC_PATH`                 | (サブコマンド実行時に使用される) LXD 実行ファイルのフルパス <!-- Full path to the LXD binary (used when forking subcommands) -->
`LXD_LXC_TEMPLATE_CONFIG`       | LXC テンプレート設定ディレクトリ <!-- Path to the LXC template configuration directory -->
`LXD_SECURITY_APPARMOR`         | `false` に設定すると AppArmor を無効にします <!-- If set to `false`, forces AppArmor off -->
`LXD_UNPRIVILEGED_ONLY`         | `true` に設定すると非特権コンテナーしか作れなくなるように強制します。 LXD_UNPRIVILEGED_ONLY を設定する前に作られた特権コンテナーだけが引き続き特権を持つことに注意してください。このオプションを LXD デーモンを最初にセットアップするときに設定するのが実用的です。 <!-- If set to `true`, enforces that only unprivileged containers can be created. Note that any privileged containers that have been created before setting LXD_UNPRIVILEGED_ONLY will continue to be privileged. To use this option effectively it should be set when the LXD daemon is first setup. -->
`LXD_OVMF_PATH`                 | `OVMF_CODE.fd` と `OVMF_VARS.ms.fd` を含む OVMF ビルドへのパス <!-- Path to an OVMF build including `OVMF_CODE.fd` and `OVMF_VARS.ms.fd` -->
`LXD_SHIFTFS_DISABLE`           | shiftfs のサポートを無効にする（従来の UID シフトを試す際に有用です） <!-- Disable shiftfs support (useful when testing traditional UID shifting) -->
