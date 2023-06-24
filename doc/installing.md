---
discourse: 8178, 16551
relatedlinks: "[LXD のインストール](https://linuxcontainers.org/ja/lxd/getting-started-cli/)"
---

# LXDをインストールするには

LXDをインストールする最も簡単な方法は{ref}`提供されているパッケージのどれかをインストールする <installing-from-package>`ことですが、{ref}`ソースからLXDをインストールする<installing_from_source>`こともできます。

LXDのインストール後、`lxd`グループがシステム内に存在することを確認してください。
このグループ内のユーザがLXDを操作できます。
手順は{ref}`installing-manage-access`を参照してください。

## リリースを選択する

LXDはことrなるリリースブランチを並行して維持しています。

- 長期サポート(LTS)リリース：現在は LXD 5.0.x と LXD 4.0.x
- 機能リリース：LXD 5.x

本番環境にはLTSを推奨します。通常のバグフィクスとセキュリティアップデートの恩恵を受けられるからです。
しかし、長期リリースには新しい機能はやどんな種類の挙動の変更も追加されません。

LXDの最新の機能と毎月の更新を得るには、代わりに機能リリースを使ってください。

(installing-from-package)=
## LXDをパッケージからインストールする

LXDデーモンはLinuxでのみ稼働します。
クライアントツール(`lxc`)はほとんどのプラットフォームで利用できます。

### Linux

LXDをインストールする最も簡単な方法は{ref}`installing-snap-package`をインストールすることです。これはさまざまなLinuxディストリビューションで利用可能です。

この選択肢が使えない場合、{ref}`installing-other`を参照してください。

(installing-snap-package)=
#### Snapパッケージ

LXDはいくつかのLinuxディストリビューション(例えば、Ubuntu、Arch Linux、Debian、Fedora、そしてOpenSUSE)で動作する[snapパッケージ](https://snapcraft.io/lxd)を公開しテストしています。 

snapをインストールするには以下の手順を実行してください。

1. [提供されているディストリビューション一覧](https://jenkins.linuxcontainers.org/job/lxd-test-snap-latest-stable/)を見て、お使いのLinuxディストリビューションで利用可能かを確認してください。
   利用可能ではない場合、{ref}`installing-other`のいずれかで対応してください。

1. `snapd`をインストールします。
   Snapcraftドキュメント[インストール手順](https://snapcraft.io/docs/core/install)を参照してください。

1. snapパッケージをインストールします。
   最新の機能リリースをインストールするには以下のようにします。

        sudo snap install lxd

   LXD 5.0 LTS リリースの場合は以下のようにします。

        sudo snap install lxd --channel=5.0/stable

LXDのsnapパッケージについてより詳細な情報(上記以外のバージョン、更新の管理など)については[Managing the LXD snap](https://discuss.linuxcontainers.org/t/managing-the-lxd-snap/8178)を参照してください。

```{note}
Ubuntu 18.04では、もしLXDのdebパッケージを過去にインストールしていた場合、既存の全てのデータを以下のコマンドで移行できます。

        sudo lxd.migrate
```

(installing-other)=
#### 他のインストール方法

いくつかのLinuxディストリビューションではsnapパッケージ以外のインストール方法を提供しています。

````{tabs}

```{group-tab} Alpine Linux

Alpine Linuxで機能リリースのLXDをインストールするには、以下のようにします。

    apk add lxd
```

```{group-tab} Arch Linux

Arch Linuxで機能リリースのLXDをインストールするには、以下のようにします。

    pacman -S lxd
```

```{group-tab} Fedora

LXC/LXDのFedora RPMパッケージが[COPR レポジトリ](https://copr.fedorainfracloud.org/coprs/ganto/lxc4/)で利用可能です。

機能リリースのLXDパッケージをインストールするには、以下のようにします。

    dnf copr enable ganto/lxc4
    dnf install lxd

インストール手順のより詳細な情報については[インストールガイド](https://github.com/ganto/copr-lxc4/wiki)を参照してください。
```

```{group-tab} Gentoo

Gentooで機能リリースのLXDをインストールするには、以下のようにします。

    emerge --ask lxd
```

````

### 他のオペレーティングシステム

```{important}
他のオペレーティングシステム向けのビルドはクライアントのみを含み、サーバーは含みません。
```

````{tabs}

```{group-tab} macOS

LXDはmacOSのLXDクライアントのビルドを[Homebrew](https://brew.sh/)で公開しています。

機能リリースのLXDをインストールするには、以下のようにします。

    brew install lxc
```

```{group-tab} Windows

Windows版のLXDクライアントは[Chocolatey](https://community.chocolatey.org/packages/lxc)パッケージとして提供されています。
インストールするためには以下のようにします。

1. [インストール手順](https://docs.chocolatey.org/en-us/choco/setup)に従ってChocolateyをインストールします。
1. LXDクライアントをインストールします。

        choco install lxc
```

````

[GitHub](https://github.com/lxc/lxd/actions)にもLXDクライアントのネイティブビルドがあります。
特定のビルドをダウンロードするには以下のようにします。

1. GitHubアカウントにログインします。
1. 興味のあるブランチやタグ(例えば、最新のリリースタグあるいは`master`)でフィルタリングします。 <!-- wokeignore:rule=master -->
1. 最新のビルドを選択し、適切なアーティファクトをダウンロードします。

(installing_from_source)=
## LXDをソースからインストールする

LXDをソースコードからビルドとインストールしたい場合、以下の手順に従ってください。

LXDの開発には`liblxc`の最新バージョン(4.0.0以上が必要)を使用することをおすすめします。
さらにLXDが動作するためには Golang 1.18 以上が必要です。
Ubuntu では次のようにインストールできます:

```bash
sudo apt update
sudo apt install acl attr autoconf automake dnsmasq-base git golang libacl1-dev libcap-dev liblxc1 liblxc-dev libsqlite3-dev libtool libudev-dev liblz4-dev libuv1-dev make pkg-config rsync squashfs-tools tar tcl xz-utils ebtables
```

デフォルトのストレージドライバである`dir`ドライバに加えて、LXDではいくつかのストレージドライバが使えます。
これらのツールをインストールすると、initramfsへの追加が行われ、ホストのブートが少しだけ遅くなるかもしれませんが、特定のドライバを使いたい場合には必要です:

```bash
sudo apt install lvm2 thin-provisioning-tools
sudo apt install btrfs-progs
```

テストスイートを実行するには、次のパッケージも必要です:

```bash
sudo apt install curl gettext jq sqlite3 socat bind9-dnsutils
```

### ソースから最新版をビルドする

この方法は LXD の最新版をビルドしたい開発者やLinuxディストリビューションで提供されないLXDの特定のリリースをビルドするためのものです。
Linuxディストリビューションへ統合するためのソースからのビルドはここでは説明しません。
それは将来、別のドキュメントで取り扱うかもしれません。

```bash
git clone https://github.com/lxc/lxd
cd lxd
```

これでLXDの現在の開発ツリーをダウンロードしてソースツリー内に移動します。
その後下記の手順にしたがって実際にLXDをビルド、インストールしてください。

### ソースからリリース版をビルドする

LXD のリリースtarballは完全な依存ツリーと`libraft`とLXDのデータベースのセットアップに使用する`libdqlite`のローカルコピーをバンドルしています。

```bash
tar zxvf lxd-4.18.tar.gz
cd lxd-4.18
```

これでリリースtarballを解凍し、ソースツリー内に移動します。
その後下記の手順にしたがって実際にLXDをビルド、インストールしてください。

### ビルドを開始する

実際のビルドはMakefileの2回の別々の実行により行われます。
1つは`make deps`でこれはLXDに必要とされるライブラリーをビルドします。
もう1つは`make`でLXD自体をビルドします。
`make deps`の最後に`make`の実行に必要な環境変数を設定するための手順が表示されます。
新しいバージョンのLXDがリリースされたらこれらの環境変数の設定は変わるかもしれませんので、`make deps`の最後に表示された手順を使うようにしてください。
下記の手順(例示のために表示します)はあなたがビルドするLXDのバージョンのものとは一致しないかもしれません。

ビルドには最低2GBのRAMを搭載することを推奨します。

```{terminal}
:input: make deps

...
make[1]: Leaving directory '/root/go/deps/dqlite'
# environment

Please set the following in your environment (possibly ~/.bashrc)
#  export CGO_CFLAGS="${CGO_CFLAGS} -I$(go env GOPATH)/deps/dqlite/include/ -I$(go env GOPATH)/deps/raft/include/"
#  export CGO_LDFLAGS="${CGO_LDFLAGS} -L$(go env GOPATH)/deps/dqlite/.libs/ -L$(go env GOPATH)/deps/raft/.libs/"
#  export LD_LIBRARY_PATH="$(go env GOPATH)/deps/dqlite/.libs/:$(go env GOPATH)/deps/raft/.libs/:${LD_LIBRARY_PATH}"
#  export CGO_LDFLAGS_ALLOW="(-Wl,-wrap,pthread_create)|(-Wl,-z,now)"
:input: make
```

### ソースからのビルド結果のインストール

ビルドが完了したら、ソースツリーを維持したまま、あなたのお使いのシェルのパスに`$(go env GOPATH)/bin`を追加し、`LD_LIBRARY_PATH`環境変数を`make deps`で表示された値に設定すれば、 LXD が利用できます。
`~/.bashrc`ファイルの場合は以下のようになります。

```bash
export PATH="${PATH}:$(go env GOPATH)/bin"
export LD_LIBRARY_PATH="$(go env GOPATH)/deps/dqlite/.libs/:$(go env GOPATH)/deps/raft/.libs/:${LD_LIBRARY_PATH}"
```

これで`lxd`と`lxc`コマンドの実行ファイルが利用可能になりLXDをセットアップするのに使用できます。
`LD_LIBRARY_PATH`環境変数のおかげで実行ファイルは`$(go env GOPATH)/deps`にビルドされた依存ライブラリーを自動的に見つけて使用します。

### マシンセットアップ

LXDが非特権コンテナを作成できるように、rootユーザーに対するsub{u,g}idの設定が必要です。

```bash
echo "root:1000000:1000000000" | sudo tee -a /etc/subuid /etc/subgid
```

これでデーモンを実行できます(`sudo`グループに属する全員がLXDとやりとりできるように `--group sudo` を指定します。別に指定したいグループを作ることもできます)。

```bash
sudo -E PATH=${PATH} LD_LIBRARY_PATH=${LD_LIBRARY_PATH} $(go env GOPATH)/bin/lxd --group sudo
```

```{note}
`newuidmap/newgidmap`ツールがシステムに存在し、`/etc/subuid`、`/etc/subgid`が存在する場合は、rootユーザーに少なくとも10MのUID/GIDの連続した範囲を許可するように設定する必要があります。
```

(installing-manage-access)=
## LXDへのアクセスを管理する

LXDのアクセス制御はグループのメンバーシップに基づいています。
rootユーザと`lxd`グループの全てのメンバーはローカルデーモンとやりとりできます。
詳細は{ref}`security-daemon-access`を参照してください。

お使いのシステムに`lxd`グループが存在しない場合は、作成してLXDデーモンを再起動してください。
このグループに追加されたメンバーはLXDの完全な制御ができます。

グループのメンバーシップは通常ログイン時にのみ適用されますので、セッションを開き直すか、LXDとやりとりするシェル上で`newgrp lxd`コマンドを実行する必要があります。

````{important}
% Include content from [../README.md](../README.md)
```{include} ../README.md
    :start-after: <!-- Include start security note -->
    :end-before: <!-- Include end security note -->
```
````

(installing-upgrade)=
## LXDをアップグレードする

LXDを新しいバージョンにアップグレードした後、LXDはデータベースを新しいスキーマにアップデートする必要があるかもしれません。
このアップデートはLXDのアップグレードの後のデーモン起動時に自動的に実行されます。
アップデート前のデータベースのバックアップはアクティブなデータベースと同じ場所(例えばsnap の場合は`/var/snap/lxd/common/lxd/database`)に保存されます。

```{important}
スキーマのアップデート後は、古いバージョンのLXDはデータベースを無効とみなすかもしれません。
これはつまりLXDをダウングレードしてもあなたのLXDの環境は利用不可能と言われるかもしれないということです。

このようなダウングレードが必要な場合は、ダウングレードを行う前にデータベースのバックアップをリストアしてください。
```
