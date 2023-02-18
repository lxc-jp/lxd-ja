(images-remote)=
# リモートイメージを使用するには

`lxc` CLIコマンドはいくつかのリモートイメージサーバを初期設定されています。
概要は{ref}`remote-image-servers`を参照してください。

## 設定されたリモートを一覧表示する

設定された全てのリモートサーバを見るには、以下のコマンドを入力します。

    lxc remote list

[simple streams形式](https://git.launchpad.net/simplestreams/tree/)を使用するリモートサーバは純粋なイメージサーバです。
`lxd`形式を使用するサーバはLXDサーバであり、イメージサーバだけとして稼働しているか、通常のLXDサーバとして稼働するのに加えて追加のイメージを提供しているかのどちらかです。
詳細は{ref}`remote-image-server-types`を参照してください。

## リモート上の利用可能なイメージを一覧表示する

サーバ上の全てのリモートイメージを一覧表示するには、以下のコマンドを入力します。

    lxc image list <remote>:

結果をフィルタできます。
手順は{ref}`images-manage-filter`を参照してください。

## リモートサーバを追加する

どのようにリモートを追加するかはサーバが使用しているプロトコルに依存します。

### simple streamsサーバを追加する

simple streamsサーバをリモートとして追加するには、以下のコマンドを入力します。

    lxc remote add <remote_name> <URL> --protocol=simplestreams

URLはHTTPSでなければなりません。

### リモートのLXDサーバを追加する

LXDサーバをリモートして追加するには、以下のコマンドを入力します。

    lxc remote add <remote_name> <IP|FQDN|URL> [flags]

認証方法によっては固有のフラグが必要です(例えば、Candid認証では`lxc remote add <remote_name> <IP|FQDN|URL> --auth-type=candid`を使います)。
詳細は{ref}`authentication`を参照してください。

例えば、IPアドレスを指定してリモートを追加するには以下のコマンドを入力します。

    lxc remote add my-remote 192.0.2.10

リモートサーバのフィンガープリントを確認するプロンプトが表示され、リモートで使用している認証方法によってパスワードまたはトークンの入力を求められます。

## イメージを参照する

イメージを参照するには、リモートとイメージのエイリアスまたはフィンガープリントをコロンで区切って指定します。
例:

    images:ubuntu/22.04
    ubuntu:22.04
    local:ed7509d7e83f

(images-remote-default)=
## デフォルトのリモートを選択する

リモート名前を指定せずにイメージ名だけ指定すると、デフォルトのイメージサーバが使用されます。

どのサーバがデフォルトのイメージサーバと設定されているか表示するには、以下のコマンドを入力します。

    lxc remote get-default

別のリモートをデフォルトのイメージサーバに選択するには、以下のコマンドを入力します。

    lxc remote switch <remote_name>
