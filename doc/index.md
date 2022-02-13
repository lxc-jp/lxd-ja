[![LXD](https://linuxcontainers.org/static/img/containers.png)](https://linuxcontainers.org/lxd)
# LXD
LXD は次世代のシステムコンテナーおよび仮想マシンのマネージャーです。
コンテナーあるいは仮想マシンの内部で稼働する完全な Linux システムに対して統一されたユーザーエクスペリエンスを提供します。

いろいろな [Linux ディストリビューション](https://images.linuxcontainers.org) のあらかじめビルドされたイメージを使ったイメージベースのマネージャーであり、非常に強力でありながら、非常にシンプルに、REST API を使って構築されます。

LXD がどういうものであり、何をするのかをよく理解するために、[オンラインで試用](https://linuxcontainers.org/lxd/try-it/) できます。
そして、ローカルで実行してみたい場合は、[はじめに](https://linuxcontainers.org/lxd/getting-started-cli/) という文書をご覧ください。

- リリースアナウンス: <https://linuxcontainers.org/lxd/news/>
- リリース tarball: <https://linuxcontainers.org/lxd/downloads/>
- ドキュメント: <https://linuxcontainers.org/lxd/docs/master/>

## ステータス
Type                | Service               | Status
---                 | ---                   | ---
CI (client)         | GitHub                | [![Build Status](https://github.com/lxc/lxd/workflows/Client%20build%20and%20unit%20tests/badge.svg)](https://github.com/lxc/lxd/actions)
CI (server)         | Jenkins               | [![Build Status](https://jenkins.linuxcontainers.org/job/lxd-github-commit/badge/icon)](https://jenkins.linuxcontainers.org/job/lxd-github-commit/)
Go documentation    | Godoc                 | [![GoDoc](https://godoc.org/github.com/lxc/lxd/client?status.svg)](https://godoc.org/github.com/lxc/lxd/client)
Static analysis     | GoReport              | [![Go Report Card](https://goreportcard.com/badge/github.com/lxc/lxd)](https://goreportcard.com/report/github.com/lxc/lxd)
Translations        | Weblate               | [![Translation status](https://hosted.weblate.org/widgets/linux-containers/-/svg-badge.svg)](https://hosted.weblate.org/projects/linux-containers/lxd/)
Project status      | CII Best Practices    | [![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/1086/badge)](https://bestpractices.coreinfrastructure.org/projects/1086)

## LXD のパッケージからのインストール
LXD デーモンは Linux でしか動きませんが、クライアントツール (`lxc`) はほとんどのプラットフォームで動作します。

OS                  | 形式                                            | コマンド
---                 | ---                                               | ---
Linux               | [Snap](https://snapcraft.io/lxd)                  | snap install lxd
Windows             | [Chocolatey](https://chocolatey.org/packages/lxc) | choco install lxc
MacOS               | [Homebrew](https://formulae.brew.sh/formula/lxc)  | brew install lxc

さまざまな Linux ディストリビューションとオペレーティングシステムで LXD をインストールするためのより詳細な方法は、[公式サイト](https://linuxcontainers.org/lxd/getting-started-cli/) をご覧ください。

ソースから LXD をインストールするには、このドキュメント内の [LXD のインストール](installing.md) を参照してください。

## セキュリティ
LXD は他のコンテナーおよびVMの管理システムと同様にローカル通信用に UNIX ソケットを提供します。

**警告**: このソケットにアクセスできる人は LXD を完全に制御できます。
これはホストのデバイスやファイルシステムにアタッチする能力も含みます。
ですので、ホストに root 権限でアクセスを許可するほどに信頼できる
ユーザーだけにこのソケットを与えるようにすべきです。

ネットワークでリッスンする時、同じ API が TLS ソケット (HTTPS) 上で
利用可能です。リモート API の特定のアクセスは Canonical RBAC 経由で
制限することができます。

より詳細は[こちらを参照してください](security.md).

## サポートとコミュニティ

LXD コミュニティと交流するには以下のチャンネルが利用できます。

### バグレポート
バグ報告と機能リクエストはこちらから行えます: <https://github.com/lxc/lxd/issues/new>

### フォーラム
ディスカッションフォーラムを使えます: <https://discuss.linuxcontainers.org>

### メーリングリスト
開発者向けとユーザー向けのディスカッションに LXC のメーリングリストを使っています。次の URL から見つけられますし、購読もできます: <https://lists.linuxcontainers.org>

### IRC
ライブのディスカッションがお好みなら、irc.libera.chat の [#lxc](https://kiwiirc.com/client/irc.libera.chat/#lxc) で私たちを見つけられます。 必要に応じて [Getting started with IRC](https://discuss.linuxcontainers.org/t/getting-started-with-irc/11920) を参照してください。

## コントリビュート
修正や新機能の追加は大歓迎です。最初に忘れずに [contributing guidelines](contributing.md) を読んでください！

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
