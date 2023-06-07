---
discourse: 16647
relatedlinks: https://www.youtube.com/watch?v=pM0EgUqj2a0
---

(remote-image-servers)=
# リモートイメージサーバー

`lxc` CLIコマンドは下記のデフォルトリモートイメージサーバーが初期設定されています。

`images:`
: このサーバーはさまざまなLinuxディストリビューションの非公式イメージを提供します。
  イメージはLXDチームによりメンテナンスされ、コンパクトで最小限にビルドされています。

  利用可能なイメージの概要については[`images.linuxcontainers.org`](https://images.linuxcontainers.org)を参照してください。

`ubuntu:`
: このサーバーは公式の安定版のUbuntuイメージを提供します。
  全てのイメージはcloudイメージです。これは`cloud-init`と`lxd-agent`の両方を含んでいることを意味します。

  利用可能なイメージの概要については[`cloud-images.ubuntu.com/releases`](https://cloud-images.ubuntu.com/releases/)を参照してください。

`ubuntu-daily:`
: このサーバーは公式のデイリービルド版のUbuntuイメージを提供します。
  全てのイメージはcloudイメージです。これは`cloud-init`と`lxd-agent`の両方を含んでいることを意味します。

  利用可能なイメージの概要については[`cloud-images.ubuntu.com/daily`](https://cloud-images.ubuntu.com/daily/)を参照してください。

(remote-image-server-types)=
## リモートサーバータイプ

LXDは下記のタイプのリモートイメージサーバーをサポートします。

simple streamsサーバー
: [simple streams形式](https://git.launchpad.net/simplestreams/tree/)を使う純粋なイメージサーバー。
  デフォルトのイメージサーバーはsimple streamsサーバーです。

公開LXDサーバー
: イメージを配布するためだけに稼働し、このサーバー自身ではインスタンスを稼働しないLXDサーバー。

  LXDサーバーをポート8443で公開で利用可能にするには、[`core.https_address`](server-options-core)設定オプションを`:8443`に設定し、認証方法をなにも設定しないようにします(詳細は{ref}`server-expose`参照)。
  そして共有したいイメージを`public`にセットします。

LXDサーバー
: ネットワーク越しに管理できる通常のLXDサーバー、イメージサーバーとしても利用可能。

  セキュリティ上の理由により、リモートAPIへのアクセスを制限し、アクセス制御のための認証方法を設定するほうが良いです。
  詳細な情報は{ref}`server-expose`と{ref}`authentication`を参照してください。
