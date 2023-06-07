(images-remote)=
# リモートイメージを使用するには

`lxc` CLIコマンドはいくつかのリモートイメージサーバーを初期設定されています。
概要は{ref}`remote-image-servers`を参照してください。

## 設定されたリモートを一覧表示する

設定された全てのリモートサーバーを見るには、以下のコマンドを入力します。

    lxc remote list

[simple streams形式](https://git.launchpad.net/simplestreams/tree/)を使用するリモートサーバーは純粋なイメージサーバーです。
`lxd`形式を使用するサーバーはLXDサーバーであり、イメージサーバーだけとして稼働しているか、通常のLXDサーバーとして稼働するのに加えて追加のイメージを提供しているかのどちらかです。
詳細は{ref}`remote-image-server-types`を参照してください。

## リモート上の利用可能なイメージを一覧表示する

サーバー上の全てのリモートイメージを一覧表示するには、以下のコマンドを入力します。

    lxc image list <remote>:

結果をフィルタできます。
手順は{ref}`images-manage-filter`を参照してください。

## リモートサーバーを追加する

どのようにリモートを追加するかはサーバーが使用しているプロトコルに依存します。

### simple streamsサーバーを追加する

simple streamsサーバーをリモートとして追加するには、以下のコマンドを入力します。

    lxc remote add <remote_name> <URL> --protocol=simplestreams

URLはHTTPSでなければなりません。

### リモートのLXDサーバーを追加する

LXDサーバーをリモートして追加するには、以下のコマンドを入力します。

    lxc remote add <remote_name> <IP|FQDN|URL> [flags]

認証方法によっては固有のフラグが必要です(例えば、Candid認証では`lxc remote add <remote_name> <IP|FQDN|URL> --auth-type=candid`を使います)。
詳細は{ref}`authentication`を参照してください。

例えば、IPアドレスを指定してリモートを追加するには以下のコマンドを入力します。

    lxc remote add my-remote 192.0.2.10

リモートサーバーのフィンガープリントを確認するプロンプトが表示され、リモートで使用している認証方法によってパスワードまたはトークンの入力を求められます。

## イメージを参照する

イメージを参照するには、リモートとイメージのエイリアスまたはフィンガープリントをコロンで区切って指定します。
例:

    images:ubuntu/22.04
    ubuntu:22.04
    local:ed7509d7e83f

(images-remote-default)=
## デフォルトのリモートを選択する

リモート名前を指定せずにイメージ名だけ指定すると、デフォルトのイメージサーバーが使用されます。

どのサーバーがデフォルトのイメージサーバーと設定されているか表示するには、以下のコマンドを入力します。

    lxc remote get-default

別のリモートをデフォルトのイメージサーバーに選択するには、以下のコマンドを入力します。

    lxc remote switch <remote_name>
