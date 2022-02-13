# 環境変数
以下の環境変数を設定することで、 LXD のクライアントとデーモンを
ユーザーの環境に適合させることができ、いくつかの高度な機能を有効または
無効にすることができます。

## クライアントとサーバ共通の環境変数
名前                           | 説明
:---                            | :----
`LXD_DIR`                       | LXD のデータディレクトリ
`PATH`                          | 実行ファイルの検索対象のパスのリスト
`http_proxy`                    | HTTP 用のプロキシサーバの URL
`https_proxy`                   | HTTPS 用のプロキシサーバの URL
`no_proxy`                      | プロキシが不要なドメイン、IPアドレスあるいは CIDR レンジのリスト

## クライアントの環境変数
名前                           | 説明
:---                            | :----
`EDITOR`                        | 使用するテキストエディタ
`VISUAL`                        | (`EDITOR` が設定されてないときに) 使用するテキストエディタ
`LXD_CONF`                      | LXC 設定ディレクトリーのパス
`LXD_GLOBAL_CONF`               | LXC グローバル設定ディレクトリーのパス

## サーバの環境変数
名前                           | 説明
:---                            | :----
`LXD_EXEC_PATH`                 | (サブコマンド実行時に使用される) LXD 実行ファイルのフルパス
`LXD_LXC_TEMPLATE_CONFIG`       | LXC テンプレート設定ディレクトリ
`LXD_SECURITY_APPARMOR`         | `false` に設定すると AppArmor を無効にします
`LXD_UNPRIVILEGED_ONLY`         | `true` に設定すると非特権コンテナーしか作れなくなるように強制します。 LXD_UNPRIVILEGED_ONLY を設定する前に作られた特権コンテナーだけが引き続き特権を持つことに注意してください。このオプションを LXD デーモンを最初にセットアップするときに設定するのが実用的です。
`LXD_OVMF_PATH`                 | `OVMF_CODE.fd` と `OVMF_VARS.ms.fd` を含む OVMF ビルドへのパス
`LXD_SHIFTFS_DISABLE`           | shiftfs のサポートを無効にする（従来の UID シフトを試す際に有用です）
`LXD_DEVMONITOR_DIR`            | デバイスモニターでモニターするパス。主にテスト用。
