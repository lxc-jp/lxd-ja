[![LXD](https://linuxcontainers.org/static/img/containers.png)](https://linuxcontainers.org/lxd)

<!-- Include start LXD intro -->
# LXD
LXDは、次世代のシステムコンテナおよび仮想マシンマネージャです。
コンテナや仮想マシンの中で動作する完全なLinuxシステムに統一されたユーザーエクスペリエンスを提供します。

LXDはイメージベースで、[多くのLinuxディストリビューション](https://images.linuxcontainers.org)に対応しています。
そして、非常にパワフルでありながら、非常にシンプルなREST APIを中心に構築されています。

LXDとは何か、何ができるのか、より良いアイデアを得るためには、[オンラインで試す](https://linuxcontainers.org/lxd/try-it/)ことができます!
また、ローカルで動作させたい場合は、[Getting Started Guide](https://linuxcontainers.org/lxd/getting-started-cli/)をご覧ください。

- リリースのアナウンス: <https://linuxcontainers.org/lxd/news/>
- リリースのtarball: <https://linuxcontainers.org/lxd/downloads/>
- ドキュメント: <https://linuxcontainers.org/lxd/docs/master/>

<!-- Include end LXD intro -->

## ステータス
タイプ             | サービス           | ステータス
---                | ---                | ---
CI（クライアント） | GitHub             | [![Build Status](https://github.com/lxc/lxd/workflows/Client%20build%20and%20unit%20tests/badge.svg)](https://github.com/lxc/lxd/actions)
CI（サーバ）       | Jenkins            | [![Build Status](https://jenkins.linuxcontainers.org/job/lxd-github-commit/badge/icon)](https://jenkins.linuxcontainers.org/job/lxd-github-commit/)
Goドキュメント     | Godoc              | [![GoDoc](https://godoc.org/github.com/lxc/lxd/client?status.svg)](https://godoc.org/github.com/lxc/lxd/client)
静的解析           | GoReport           | [![Go Report Card](https://goreportcard.com/badge/github.com/lxc/lxd)](https://goreportcard.com/report/github.com/lxc/lxd)
翻訳               | Weblate            | [![翻訳状況](https://hosted.weblate.org/widgets/linux-containers/-/svg-badge.svg)](https://hosted.weblate.org/projects/linux-containers/lxd/)
プロジェクトの状況 | CII Best Practices | [![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/1086/badge)](https://bestpractices.coreinfrastructure.org/projects/1086)

<!-- Include start installing -->

## パッケージからのLXDのインストール
LXDのデーモンはLinuxでしか動作しませんが、クライアントツール(`lxc`)はほとんどのプラットフォームで利用可能です。

OS      | フォーマット                                      |コマンド
---     | ---                                               | ---
Linux   | [Snap](https://snapcraft.io/lxd)                  | snap install lxd
Windows | [Chocolatey](https://chocolatey.org/packages/lxc) | choco install lxc
MacOS   | [Homebrew](https://formulae.brew.sh/formula/lxc)  | brew install lxc

様々なLinuxディストリビューションやOSへのLXDのインストールについては、[私たちのウェブサイト](https://linuxcontainers.org/lxd/getting-started-cli/)に詳しい説明があります。
<!-- Include end installing -->

LXDをソースからインストールするには、ドキュメントの[Installing LXD](doc/installing.md)を参照してください。

## セキュリティ

<!-- Include start security -->

LXDのインストールが安全であることを保証するために、以下の点を考慮してください。

- オペレーティングシステムを最新に保ち、利用可能なすべてのセキュリティパッチをインストールする。
- サポートされているLXDのバージョン（LTSリリースまたは月例機能リリース）のみを使用する。
- LXDデーモンとリモートAPIへのアクセスを制限すること。
- 必要とされない限り、特権コンテナを使わないこと。特権的なコンテナを使う場合は、適切なセキュリティ対策をしてください。詳細は[LXCセキュリティページ](https://linuxcontainers.org/lxc/security/)を参照してください。
- ネットワークインターフェイスを安全に設定してください。
<!--Include end security -->

詳しい情報は[Security](doc/security.md)を参照してください。

**重要：**。
<!-- Include start security note -->
UNIXソケットを介したLXDへのローカルアクセスは、常にLXDへのフルアクセスを許可します。
これは、任意のインスタンス上のセキュリティ機能を微調整するのと同様に、任意のインスタンスにファイルシステムパスやデバイスをアタッチする能力を含みます。

したがって、あなたのシステムへのルートアクセスを信頼できるユーザーにのみ、このようなアクセスを与えるべきです。
<!-- Include end security note -->
<!-- Include start support -->

## サポートとコミュニティ

LXDコミュニティと交流するために以下のチャンネルが用意されています。

### バグレポート
バグレポートや機能要求は以下の場所で受け付けています。<https://github.com/lxc/lxd/issues/new>

### フォーラム
フォーラムは以下の場所にあります。<https://discuss.linuxcontainers.org>

### メーリングリスト
開発者やユーザーの議論にはLXCのメーリングリストを利用しています。
メーリングリストは以下の場所にあります。<https://lists.linuxcontainers.org>

### IRC
ライブの議論がお好みならば、irc.libera.chatの[#lxc](https://kiwiirc.com/client/irc.libera.chat/#lxc)で私たちを見つけることができます。必要であれば [Getting started with IRC](https://discuss.linuxcontainers.org/t/getting-started-with-irc/11920) を参照してください。

### 商用サポート

LXDの商用サポートは、[Canonical Ltd](https://www.canonical.com)を通じて受けることができます。

## ドキュメント
公式ドキュメントは https://lxd-ja.readthedocs.io/ja/latest/ (原文は https://linuxcontainers.org/lxd/docs/master/) で入手できます。

その他の資料は、[website](https://linuxcontainers.org/lxd/articles)、[YouTube](https://www.youtube.com/channel/UCuP6xPt0WTeZu32CkQPpbvA)、フォーラムの[Tutorials section](https://discuss.linuxcontainers.org/c/tutorials/)にあります。

<!-- Include end support -->

## コントリビュート
修正や新機能の提供は大歓迎です。まずは、[コントリビュートガイド](CONTRIBUTING.md)をお読みください!
