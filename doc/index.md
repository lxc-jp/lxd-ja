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
Release announcements can be found here: <https://linuxcontainers.org/lxd/news/>  
And the release tarballs here: <https://linuxcontainers.org/lxd/downloads/>
The documentation is here: <https://linuxcontainers.org/lxd/docs/master/>
-->
リリースアナウンスはこちらでご覧になれます: <https://linuxcontainers.org/lxd/news/>  
リリース tarball はこちらから取得できます: <https://linuxcontainers.org/lxd/downloads/>
ドキュメントはこちらでご覧になれます: <https://linuxcontainers.org/lxd/docs/master/>

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

## LXD のソースからのインストール <!-- Installing LXD from source -->
LXD の開発には liblxc の最新バージョン（3.0.0 以上が必要）を使用することをおすすめします。
さらに LXD が動作するためには Golang 1.13 以上が必要です。
Ubuntu では次のようにインストールできます:
<!--
We recommend having the latest versions of liblxc (>= 3.0.0 required)
available for LXD development. Additionally, LXD requires Golang 1.13 or
later to work. On ubuntu, you can get those with:
-->

```bash
sudo apt update
sudo apt install acl autoconf dnsmasq-base git golang libacl1-dev libcap-dev liblxc1 liblxc-dev libsqlite3-dev libtool libudev-dev liblz4-dev libuv1-dev make pkg-config rsync squashfs-tools tar tcl xz-utils ebtables
```

<!--
Note that when building LXC yourself, ensure to build it with the appropriate
security related libraries installed which our testsuite tests. Again, on
ubuntu, you can get those with:
-->
LXC を自分でビルドする場合は、テストスイートがテストする、関連する適切なセキュリティ関連のライブラリーがインストールされていることを確認してください。
Ubuntu であれば次のようにインストールできます:

```bash
sudo apt install libapparmor-dev libseccomp-dev libcap-dev
```

<!--
There are a few storage backends for LXD besides the default "directory" backend.
Installing these tools adds a bit to initramfs and may slow down your
host boot, but are needed if you'd like to use a particular backend:
-->
デフォルトのストレージバックエンドである "directory" に加えて、LXD ではいくつかのストレージバックエンドが使えます。
これらのツールをインストールすると、initramfs への追加が行われ、ホストのブートが少しだけ遅くなるかもしれませんが、特定のバックエンドを使いたい場合には必要です:

```bash
sudo apt install lvm2 thin-provisioning-tools
sudo apt install btrfs-progs
```

<!--
To run the testsuite, you'll also need:
-->
テストスイートを実行するには、次のパッケージも必要です:

```bash
sudo apt install curl gettext jq sqlite3 uuid-runtime socat
```

### ソースからの最新版のビルド <!-- From Source: Building the latest version -->
この方法は LXD の最新版をビルドしたい開発者や Linux ディストリビューションで提供されない LXD の特定のリリースをビルドするためのものです。 Linux ディストリビューションへ統合するためのソースからのビルドはここでは説明しません。それは将来別のドキュメントで取り扱うかもしれません。
<!--
These instructions for building from source are suitable for individual developers who want to build the latest version
of LXD, or build a specific release of LXD which may not be offered by their Linux distribution. Source builds for
integration into Linux distributions are not covered here and may be covered in detail in a separate document in the
future.
-->

ソースからビルドする際は、ビルド対象のソースコードを含む `GOPATH` をカスタムで設定します。ソースをビルドしたら `lxc` と `lxd` の実行ファイルが `$GOPATH/bin` に生成され、あとは `LD_LIBRARY_PATH` を設定（後述）するだけで、これらの実行ファイルがビルドされたソースツリーから直接実行できます。
<!--
When building from source, it is customary to configure a `GOPATH` which contains the to-be-built source code. When 
the sources are done building, the `lxc` and `lxd` binaries will be available at `$GOPATH/bin`, and with a little
`LD_LIBRARY_PATH` magic (described later), these binaries can be run directly from the built source tree. 
-->

以下に GitHub の LXD のソースの最新版に対して `GOPATH` を設定する手順を示します。
<!--
The following lines demonstrate how to configure a `GOPATH` with the most recent LXD sources from GitHub:
-->

```bash
mkdir -p ~/go
export GOPATH=~/go
go get -d -v github.com/lxc/lxd/lxd
cd $GOPATH/src/github.com/lxc/lxd
```

ビルドプロセスが開始したら、 Makefile は `go get` と `git clean` を使ってビルドに必要な全ての依存ライブラリーを取得します。
<!--
When the build process starts, the Makefile will use `go get` and `git clone` to grab all necessary dependencies 
needed for building.
-->

### ソースからのリリース版のビルド <!-- From Source: Building a Release -->

LXD の公式リリースをビルドするには、リリース版の tarball をダウンロード、解凍し、その中の `_dist` ディレクトリを指すように GOPATH を設定してください。 `_dist` ディレクトリは GOPATH として使用できるように構成されており必要なソース全てのスナップショットを含んでいます。 LXD は `live` なソースを `go get` と `git clone` で取得する代わりにこれらのスナップショットを使ってビルドします。リリース版の tarball をダウンロード、解凍したら以下のように `GOPATH` を設定してください。
<!--
To build an official release of LXD, download and extract a release tarball, and then set up GOPATH to point to the
`_dist` directory inside it, which is configured to be used as a GOPATH and contains snapshots of all necessary sources. LXD
will then build using these snapshots rather than grabbing 'live' sources using `go get` and `git clone`. Once the release
tarball is downloaded and extracted, set the `GOPATH` as follows:
-->

```bash
cd lxd-3.18
export GOPATH=$(pwd)/_dist
export GOBIN=$GOPATH/bin
```

### ビルドの開始 <!-- Starting the Build -->

`GOPATH` を設定したら、以下の手順で GitHub の最新版あるいは公式リリース版の LXD をビルド出来ます。
<!--
Once the `GOPATH` is configured, either to build the latest GitHub version or an official release, the following steps
can be used to build LXD.
-->

実際のビルドは Makefile の 2 回の別々の実行により行われます。 1 つは `make deps` でこれは LXD に必要とされるライブラリーをビルドします。もう 1 つは `make` で LXD 自体をビルドします。 `make deps` の最後に `make` の実行に必要な環境変数を設定するための手順が表示されます。新しいバージョンの LXD がリリースされたらこれらの環境変数の設定は変わるかもしれませんので、 `make deps` の最後に表示された手順を使うようにしてください。下記の手順（例示のために表示します）はあなたがビルドする LXD のバージョンのものとは一致しないかもしれません。
<!--
The actual building is done by two separate invocations of the Makefile: `make deps` - - which builds libraries required 
by LXD - - and `make`, which builds LXD itself. At the end of `make deps`, a message will be displayed which will specify environment variables that should be set prior to invoking `make`. As new versions of LXD are released, these environment
variable settings may change, so be sure to use the ones displayed at the end of the `make deps` process, as the ones
below (shown for example purposes) may not exactly match what your version of LXD requires:
-->


```bash
make deps
# `make deps` が出力した export のコマンド列を使ってください。下記はあくまで例です。
export CGO_CFLAGS="${CGO_CFLAGS} -I${GOPATH}/deps/dqlite/include/ -I${GOPATH}/deps/raft/include/"
export CGO_LDFLAGS="${CGO_LDFLAGS} -L${GOPATH}/deps/dqlite/.libs/ -L${GOPATH}/deps/raft/.libs/"
export LD_LIBRARY_PATH="${GOPATH}/deps/dqlite/.libs/:${GOPATH}/deps/raft/.libs/:${LD_LIBRARY_PATH}"
export CGO_LDFLAGS_ALLOW="(-Wl,-wrap,pthread_create)|(-Wl,-z,now)"
make
```

<!--
```bash
make deps
# Use the export statements printed in the output of 'make deps' - - these are examples: 
export CGO_CFLAGS="${CGO_CFLAGS} -I${GOPATH}/deps/dqlite/include/ -I${GOPATH}/deps/raft/include/"
export CGO_LDFLAGS="${CGO_LDFLAGS} -L${GOPATH}/deps/dqlite/.libs/ -L${GOPATH}/deps/raft/.libs/"
export LD_LIBRARY_PATH="${GOPATH}/deps/dqlite/.libs/:${GOPATH}/deps/raft/.libs/:${LD_LIBRARY_PATH}"
export CGO_LDFLAGS_ALLOW="(-Wl,-wrap,pthread_create)|(-Wl,-z,now)"
make
```
-->

### ソースからのビルド結果のインストール

ビルドが完了したら、ソースツリーを維持したまま、あなたのお使いのシェルのパスに `$GOPATH/bin` を追加し `LD_LIBRARY_PATH` 環境変数を `make deps` で表示された値に設定すれば、 LXD が利用できます。 `~/.bashrc` ファイルの場合は以下のようになります。
<!--
Once the build completes, you simply keep the source tree, add the directory referenced by `$GOPATH/bin` to 
your shell path, and set the `LD_LIBRARY_PATH` variable printed by `make deps` to your environment. This might look
something like this for a `~/.bashrc` file:
-->

```bash
# GOPATH は export する必要はありません。
GOPATH=~/go
# しかしこの2つは export する必要があります。
export PATH="$PATH:$GOPATH/bin"
export LD_LIBRARY_PATH="${GOPATH}/deps/dqlite/.libs/:${GOPATH}/deps/raft/.libs/:${LD_LIBRARY_PATH}"
```

<!--
```bash
# No need to export GOPATH:
GOPATH=~/go
# But we need to export these:
export PATH="$PATH:$GOPATH/bin"
export LD_LIBRARY_PATH="${GOPATH}/deps/dqlite/.libs/:${GOPATH}/deps/raft/.libs/:${LD_LIBRARY_PATH}"
```
-->

これで `lxd` と `lxc` コマンドの実行ファイルが利用可能になり LXD をセットアップするのに使用できます。 `LD_LIBRARY_PATH` 環境変数のおかげで実行ファイルは `$GOPATH/deps` にビルドされた依存ライブラリーを自動的に見つけて使用します。
<!--
Now, the `lxd` and `lxc` binaries will be available to you and can be used to set up LXD. The binaries will automatically find and use the dependencies built in `$GOPATH/deps` thanks to the `LD_LIBRARY_PATH` environment variable.
-->

### マシンセットアップ <!-- Machine Setup -->
<!--
You'll need sub{u,g}ids for root, so that LXD can create the unprivileged containers:
-->
LXD が非特権コンテナーを作成できるように、root ユーザーに対する sub{u,g}id の設定が必要です:

```bash
echo "root:1000000:65536" | sudo tee -a /etc/subuid /etc/subgid
```

<!--
Now you can run the daemon (the `\-\-group sudo` bit allows everyone in the `sudo`
group to talk to LXD; you can create your own group if you want):
-->
これでデーモンを実行できます（`sudo` グループに属する全員が LXD とやりとりできるように `--group sudo` を指定します。別に指定したいグループを作ることもできます）:

```bash
sudo -E PATH=$PATH LD_LIBRARY_PATH=$LD_LIBRARY_PATH $GOPATH/bin/lxd --group sudo
```

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

## LXD を使い始める <!-- Getting started with LXD -->
<!--
Now that you have LXD running on your system you can read the [getting started guide](https://linuxcontainers.org/lxd/getting-started-cli/) or go through more examples and configurations in [our documentation](https://github.com/lxc/lxd/tree/master/doc).
-->
ここまでで、システム上で LXD が実行されているでしょうから、[はじめに](https://linuxcontainers.org/lxd/getting-started-cli/) という文書を読んだり、[ドキュメント](https://linuxcontainers.org/lxd/docs/master/) （[日本語訳](https://lxd-ja.readthedocs.io/)）の例や設定を見たりできます。

## バグレポート <!-- Bug reports -->
<!--
Bug reports and Feature requests can be filed at: <https://github.com/lxc/lxd/issues/new>
-->
バグ報告と機能リクエストはこちらから行えます: <https://github.com/lxc/lxd/issues/new>

## コントリビュート <!-- Contributing -->
<!--
Fixes and new features are greatly appreciated but please read our [contributing guidelines](contributing.md) first.
-->
修正や新機能の追加は歓迎です。最初に [contributing guidelines](contributing.md) を読んでください。

## サポートとディスカッション <!-- Support and discussions -->
### フォーラム <!-- Forum -->
<!--
A discussion forum is available at: <https://discuss.linuxcontainers.org>
-->
ディスカッションフォーラムを使えます: <https://discuss.linuxcontainers.org>

### メーリングリスト <!-- Mailing-lists -->
<!--
We use the LXC mailing-lists for developer and user discussions, you can
find and subscribe to those at: <https://lists.linuxcontainers.org>
-->
開発者向けとユーザー向けのディスカッションに LXC のメーリングリストを使っています。次の URL から見つけられますし、購読もできます: <https://lists.linuxcontainers.org>

### IRC
<!--
If you prefer live discussions, you can find us in [#lxc](https://kiwiirc.com/client/irc.libera.chat/#lxc) on irc.libera.chat.
-->
ライブのディスカッションがお好みなら、irc.libera.chat の [#lxc](https://kiwiirc.com/client/irc.libera.chat/#lxc) で私たちを見つけられます。

## FAQ
#### LXD サーバにリモートからアクセスできるようにするには? <!-- How to enable LXD server for remote access? -->
<!--
By default LXD server is not accessible from the networks as it only listens
on a local unix socket. You can make LXD available from the network by specifying
additional addresses to listen to. This is done with the `core.https_address`
config variable.
-->
デフォルトでは、LXD サーバーはネットワークからのアクセスを許可せず、ローカルの unix ソケットのみで待ち受けます。
待ち受ける追加のアドレスを指定して、ネットワーク経由で利用できるように設定できます。
これには `core.https_address` を使用します。

<!--
To see the current server configuration, run:
-->
現在のサーバーの設定を確認するには次のように実行します:

```bash
lxc config show
```

<!--
To set the address to listen to, find out what addresses are available and use
the `config set` command on the server:
-->
待ち受けるアドレスを設定するには、どのアドレスが使用できるかを調べ、`config set` コマンドをサーバー上で実行します:

```bash
ip addr
lxc config set core.https_address 192.168.1.15
```

#### https 越しに `lxc remote add` を実行したとき、パスワードを聞かれますか? <!-- When I do a `lxc remote add` over https, it asks for a password? -->
<!--
By default, LXD has no password for security reasons, so you can't do a remote
add this way. In order to set a password, do:
-->
デフォルトではセキュリティー上の理由から、LXD にはパスワードはありません。
ですのでこの方法でリモートから追加できません。パスワードを設定するには、LXD を実行中のホスト上で次のコマンドを実行します:

```bash
lxc config set core.trust_password SECRET
```

<!--
on the host LXD is running on. This will set the remote password that you can
then use to do `lxc remote add`.
-->
これで、`lxc remote add` を使えるように、リモートのパスワードを設定します。

<!--
You can also access the server without setting a password by copying the client
certificate from `.config/lxc/client.crt` to the server and adding it with:
-->
クライアント上の `.config/lxc/client.crt` にある証明書を次のコマンドでサーバにコピーすることで、パスワードを設定しなくてもサーバにアクセスできます。

```bash
lxc config trust add client.crt
```

#### どのように LXD のストレージを設定するのですか? <!-- How do I configure LXD storage? -->
<!--
LXD supports btrfs, ceph, directory, lvm and zfs based storage.
-->
LXD は btrfs、ceph、ディレクトリ、lvm、zfs を使ったストーレジをサポートします。

<!--
First make sure you have the relevant tools for your filesystem of
choice installed on the machine (btrfs-progs, lvm2 or zfsutils-linux).
-->
まず、選択したファイルシステムを扱うツールをマシンにインストールしてください（btrfs-progs, lvm2, zfsutils-linux）。

<!--
By default, LXD comes with no configured network or storage.
You can get a basic configuration done with:
-->
デフォルトでは、LXD ではネットワークとストレージが設定されていません。
基本的な設定は次のコマンドで設定できます:

```bash
    lxd init
```

<!--
`lxd init` supports both directory based storage and ZFS.
If you want something else, you'll need to use the `lxc storage` command:
-->
`lxd init` はディレクトリと ZFS ベースのストレージの両方をサポートします。
他のファイルシステムを使いたい場合は、`lxc storage` コマンドを使う必要があります:

```bash
lxc storage create default BACKEND [OPTIONS...]
lxc profile device add default root disk path=/ pool=default
```

<!--
BACKEND is one of `btrfs`, `ceph`, `dir`, `lvm` or `zfs`.
-->
`BACKEND` は `btrfs`、`ceph`、`dir`、`lvm`、`zfs` のどれかです。

<!--
Unless specified otherwise, LXD will setup loop based storage with a sane default size.
-->
特に指定しないと、LXD はデフォルトサイズの loop ベースのストレージをセットアップします。

<!--
For production environments, you should be using block backed storage
instead both for performance and reliability reasons.
-->
プロダクション環境では、パフォーマンスと信頼性を確保するために、loop ベースではなく、ブロックストレージを使うべきです。

#### LXD を使ってコンテナーのライブマイグレーションはできますか? <!-- How can I live migrate a container using LXD? -->
<!--
Live migration requires a tool installed on both hosts called
[CRIU](https://criu.org), which is available in Ubuntu via:
-->
ライブマイグレーションには、送受信それぞれのホスト上に [CRIU](http://criu.org) というツールが必要です。
Ubuntu では次のようにインストールできます:

```bash
sudo apt install criu
```

<!--
Then, launch your container with the following,
-->
そして、次のようにコンテナーを起動します。

```bash
lxc launch ubuntu $somename
sleep 5s # let the container get to an interesting state
lxc move host1:$somename host2:$somename
```

<!--
And with luck you'll have migrated the container :). Migration is still in
experimental stages and may not work for all workloads. Please report bugs on
lxc-devel, and we can escalate to CRIU lists as necessary.
-->
運が良ければ、コンテナーがマイグレーションされるでしょう :)
マイグレーションはまだ実験段階のステージで、すべてのケースで動作しないかもしれません。
そういう場合は lxc-devel にバグレポートをしてください。必要であれば CRIU にもエスカレーションします。

#### 私のホームディレクトリをコンテナー内にバインドマウントできますか? <!-- Can I bind mount my home directory in a container? -->
はい。ディスクデバイスを使用して以下のように出来ます。
<!--
Yes. This can be done using a disk device:
-->

```bash
lxc config device add container-name home disk source=/home/$USER path=/home/ubuntu
```

非特権コンテナーの場合は、さらに以下のいずれかが必要です。
<!--
For unprivileged containers, you will also need one of:
-->

 - `lxc config device add` の実行時に `shift=true` を指定する。これは `shiftfs` がサポートされている場合にのみ使えます（ `lxc info` 参照）。 <!-- Pass `shift=true` to the `lxc config device add` call. This depends on `shiftfs` being supported (see `lxc info`) -->
 - raw.idmap エントリー（[ユーザー名前空間 (user namespace) 用の ID のマッピング](userns-idmap.md) 参照） <!-- raw.idmap entry (see [Idmaps for user namespace](userns-idmap.md)) -->
 - ホームディレクトリーに配置した再帰的な POSIX ACL <!-- Recursive POSIX ACLs placed on your home directory -->

これらのいずれかにより、コンテナー内のユーザーが実際に効果のある read/write パーミッションを持てるようになります。
これらの 1 つも設定しないときは、すべてが uid/gid がオーバーフロー (65536:65536) ように見えて、 world リーダブルでないものへのアクセスは全て失敗します。
<!--
Either of those can be used to allow the user in the container to have working read/write permissions.
When not setting one of those, everything will show up as the overflow uid/gid (65536:65536)
and access to anything that's not world readable will fail.
-->


特権コンテナーではコンテナー内の全ての uid/gid が外部と同じためこの問題はありません。
しかし、このことは特権コンテナーに関するセキュリティーの問題の主な原因でもあります。
<!--
Privileged containers do not have this issue as all uid/gid inthe container are the same outside.
But that's also the cause of most of the security issues with such privileged containers.
-->

#### LXD コンテナー内で docker を実行できますか? <!-- How can I run docker inside a LXD container? -->
<!--
In order to run Docker inside a LXD container the `security.nesting` property of the container should be set to `true`. 
-->
LXD コンテナー内で Docker を実行するには、コンテナーの `security.nesting` プロパティを `true` に設定します。

```bash
lxc config set <container> security.nesting true
```

<!--
Note that LXD containers cannot load kernel modules, so depending on your
Docker configuration you may need to have the needed extra kernel modules
loaded by the host.
-->
LXD コンテナー内ではカーネルモジュールはロードできませんので、Docker の設定に従って、ホスト側で必要なカーネルモジュールをロードしておく必要があることに注意してください。

<!--
You can do so by setting a comma separate list of kernel modules that your container needs with:
-->
コンテナーで必要なカーネルモジュールをカンマ区切りのリストで次のように設定しておけます:

```bash
lxc config set <container> linux.kernel_modules <modules>
```

<!--
We have also received some reports that creating a `/.dockerenv` file in your
container can help Docker ignore some errors it's getting due to running in a
nested environment.
-->
コンテナー内に `/.dockerenv` ファイルを作ることで、ネストされた環境内で実行することによりおこるエラーのいくつかを Docker に無視させることができるというレポートをいくつか受け取っています。

## LXD のハック <!-- Hacking on LXD -->
### 直接 REST API を使って <!-- Directly using the REST API -->
<!--
The LXD REST API can be used locally via unauthenticated Unix socket or remotely via SSL encapsulated TCP.
-->
LXD の REST API は、認証不要なローカルの Unix ソケット経由でも、SSL で暗号化された TCP 経由でも使えます。

#### UNIX ソケット経由 <!-- Via Unix socket -->

```bash
curl --unix-socket /var/lib/lxd/unix.socket \
    -H "Content-Type: application/json" \
    -X POST \
    -d @hello-ubuntu.json \
    lxd/1.0/containers
```

あるいは snap ユーザーの場合は

```bash
curl --unix-socket /var/snap/lxd/common/lxd/unix.socket \
    -H "Content-Type: application/json" \
    -X POST \
    -d @hello-ubuntu.json \
    lxd/1.0/containers
```

#### TCP 経由 <!-- Via TCP -->
<!--
TCP requires some additional configuration and is not enabled by default.
-->
TCP 経由では、デフォルトでは有効ではない追加の設定が必要です。

```bash
lxc config set core.https_address "[::]:8443"
```

```bash
curl -k -L \
    --cert ~/.config/lxc/client.crt \
    --key ~/.config/lxc/client.key \
    -H "Content-Type: application/json" \
    -X POST \
    -d @hello-ubuntu.json \
    "https://127.0.0.1:8443/1.0/containers"
```

#### 事前に用意する JSON ファイル <!-- JSON payload -->
<!--
The `hello-ubuntu.json` file referenced above could contain something like:
-->
上記の `hello-ubuntu.json` ファイルは以下のような内容です。

```json
{
    "name":"some-ubuntu",
    "ephemeral":true,
    "config":{
        "limits.cpu":"2"
    },
    "source": {
        "type":"image",
        "mode":"pull",
        "protocol":"simplestreams",
        "server":"https://cloud-images.ubuntu.com/releases",
        "alias":"18.04"
    }
}
```
