[![LXD](https://linuxcontainers.org/static/img/containers.png)](https://linuxcontainers.org/lxd)
# LXD
<!--
LXD is a next generation system container and virtual machine manager.
It offers a unified user experience around full Linux systems running inside containers or virtual machines.
-->
LXD は次世代のシステムコンテナーおよび仮想マシンのマネージャーです。
コンテナーあるいは仮想マシンの内部で稼働する完全な Linux システムに対して統一されたユーザーエクスペリエンスを提供します。

<!--
It's image based with pre-made images available for a [wide number of Linux distributions](https://images.linuxcontainers.org)  
and is built around a very powerful, yet pretty simple, REST API.
-->
いろいろな [Linux ディストリビューション](https://images.linuxcontainers.org) のあらかじめビルドされたイメージを使ったイメージベースのマネージャーであり、非常に強力でありながら、非常にシンプルに、REST API を使って構築されます。

<!--
To get a better idea of what LXD is and what it does, you can [try it online](https://linuxcontainers.org/lxd/try-it/)!  
Then if you want to run it locally, take a look at our [getting started guide](https://linuxcontainers.org/lxd/getting-started-cli/).
-->
LXD がどういうものであり、何をするのかをよく理解するために、[オンラインで試用](https://linuxcontainers.org/lxd/try-it/) できます。  
そして、ローカルで実行してみたい場合は、[はじめに](https://linuxcontainers.org/lxd/getting-started-cli/) という文書をご覧ください。

<!--
- Release announcements: <https://linuxcontainers.org/lxd/news/>
- Release tarballs: <https://linuxcontainers.org/lxd/downloads/>
- Documentation: <https://linuxcontainers.org/lxd/docs/master/>
-->
- リリースアナウンス: <https://linuxcontainers.org/lxd/news/>
- リリース tarball: <https://linuxcontainers.org/lxd/downloads/>
- ドキュメント: <https://linuxcontainers.org/lxd/docs/master/>

## ステータス <!-- Status -->
Type                | Service               | Status
---                 | ---                   | ---
CI (client)         | GitHub                | [![Build Status](https://github.com/lxc/lxd/workflows/Client%20build%20and%20unit%20tests/badge.svg)](https://github.com/lxc/lxd/actions)
CI (server)         | Jenkins               | [![Build Status](https://jenkins.linuxcontainers.org/job/lxd-github-commit/badge/icon)](https://jenkins.linuxcontainers.org/job/lxd-github-commit/)
Go documentation    | Godoc                 | [![GoDoc](https://godoc.org/github.com/lxc/lxd/client?status.svg)](https://godoc.org/github.com/lxc/lxd/client)
Static analysis     | GoReport              | [![Go Report Card](https://goreportcard.com/badge/github.com/lxc/lxd)](https://goreportcard.com/report/github.com/lxc/lxd)
Translations        | Weblate               | [![Translation status](https://hosted.weblate.org/widgets/linux-containers/-/svg-badge.svg)](https://hosted.weblate.org/projects/linux-containers/lxd/)
Project status      | CII Best Practices    | [![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/1086/badge)](https://bestpractices.coreinfrastructure.org/projects/1086)

## LXD のパッケージからのインストール <!-- Installing LXD from packages -->
<!--
The LXD daemon only works on Linux but the client tool (`lxc`) is available on most platforms.
-->
LXD デーモンは Linux でしか動きませんが、クライアントツール (`lxc`) はほとんどのプラットフォームで動作します。

OS                  | 形式 <!-- Format -->                                            | コマンド <!-- Command -->
---                 | ---                                               | ---
Linux               | [Snap](https://snapcraft.io/lxd)                  | snap install lxd
Windows             | [Chocolatey](https://chocolatey.org/packages/lxc) | choco install lxc
MacOS               | [Homebrew](https://formulae.brew.sh/formula/lxc)  | brew install lxc

<!--
More instructions on installing LXD for a wide variety of Linux distributions and operating systems [can be found on our website](https://linuxcontainers.org/lxd/getting-started-cli/).
-->
さまざまな Linux ディストリビューションとオペレーティングシステムで LXD をインストールするためのより詳細な方法は、[公式サイト](https://linuxcontainers.org/lxd/getting-started-cli/) をご覧ください。

<!--
To install LXD from source, see [Installing LXD](installing.md) in the documentation.
-->
ソースから LXD をインストールするには、このドキュメント内の [LXD のインストール](installing.md) を参照してください。

## セキュリティ <!-- Security -->
<!--
LXD, similar to other container and VM managers provides a UNIX socket for local communication.
-->
LXD は他のコンテナーおよびVMの管理システムと同様にローカル通信用に UNIX ソケットを提供します。

<!--
**WARNING**: Anyone with access to that socket can fully control LXD, which includes
the ability to attach host devices and filesystems, this should
therefore only be given to users who would be trusted with root access
to the host.
-->
**警告**: このソケットにアクセスできる人は LXD を完全に制御できます。
これはホストのデバイスやファイルシステムにアタッチする能力も含みます。
ですので、ホストに root 権限でアクセスを許可するほどに信頼できる
ユーザーだけにこのソケットを与えるようにすべきです。

<!--
When listening on the network, the same API is available on a TLS socket
(HTTPS), specific access on the remote API can be restricted through
Canonical RBAC.
-->
ネットワークでリッスンする時、同じ API が TLS ソケット (HTTPS) 上で
利用可能です。リモート API の特定のアクセスは Canonical RBAC 経由で
制限することができます。

<!--
More details are [available here](security.md).
-->
より詳細は[こちらを参照してください](security.md).

## サポートとコミュニティ <!-- Support and community -->

LXD コミュニティと交流するには以下のチャンネルが利用できます。
<!--
The following channels are available for you to interact with the LXD community.
-->

### バグレポート <!-- Bug reports -->
バグ報告と機能リクエストはこちらから行えます: <https://github.com/lxc/lxd/issues/new>
<!--
You can file bug reports and feature requests at: <https://github.com/lxc/lxd/issues/new>
-->

### フォーラム <!-- Forum -->
ディスカッションフォーラムを使えます: <https://discuss.linuxcontainers.org>
<!--
A discussion forum is available at: <https://discuss.linuxcontainers.org>
-->

### メーリングリスト <!-- Mailing-lists -->
開発者向けとユーザー向けのディスカッションに LXC のメーリングリストを使っています。次の URL から見つけられますし、購読もできます: <https://lists.linuxcontainers.org>
<!--
We use the LXC mailing-lists for developer and user discussions, you can
find and subscribe to those at: <https://lists.linuxcontainers.org>
-->

### IRC
ライブのディスカッションがお好みなら、irc.libera.chat の [#lxc](https://kiwiirc.com/client/irc.libera.chat/#lxc) で私たちを見つけられます。 必要に応じて [Getting started with IRC](https://discuss.linuxcontainers.org/t/getting-started-with-irc/11920) を参照してください。
<!--
If you prefer live discussions, you can find us in [#lxc](https://kiwiirc.com/client/irc.libera.chat/#lxc) on irc.libera.chat. See [Getting started with IRC](https://discuss.linuxcontainers.org/t/getting-started-with-irc/11920) if needed.
-->

## コントリビュート <!-- Contributing -->
修正や新機能の追加は大歓迎です。最初に忘れずに [contributing guidelines](contributing.md) を読んでください！
<!--
Fixes and new features are greatly appreciated. Make sure to read our [contributing guidelines](contributing.md) first!
-->

<!--
```{toctree}
:hidden:
:titlesonly:

self
getting_started
configuration
images
operation
restapi_landing
internals
external_resources
```
-->
