# LXD のインストール <!-- Installing LXD -->

LXD をインストールする最も簡単な方法は [Getting started guide](https://linuxcontainers.org/lxd/getting-started-cli/#installing-a-package) で説明されているパッケージのどれかをインストールすることですが、ソースから LXD をインストールすることもできます。
<!--
The easiest way to install LXD is to install one of the available packages as described in the [Getting started guide](https://linuxcontainers.org/lxd/getting-started-cli/#installing-a-package), but you can also install LXD from the sources.
-->

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
sudo apt install acl attr autoconf dnsmasq-base git golang libacl1-dev libcap-dev liblxc1 liblxc-dev libsqlite3-dev libtool libudev-dev liblz4-dev libuv1-dev make pkg-config rsync squashfs-tools tar tcl xz-utils ebtables
```

デフォルトのストレージバックエンドである "directory" に加えて、LXD ではいくつかのストレージバックエンドが使えます。
これらのツールをインストールすると、initramfs への追加が行われ、ホストのブートが少しだけ遅くなるかもしれませんが、特定のバックエンドを使いたい場合には必要です:
<!--
There are a few storage backends for LXD besides the default "directory" backend.
Installing these tools adds a bit to initramfs and may slow down your
host boot, but are needed if you'd like to use a particular backend:
-->

```bash
sudo apt install lvm2 thin-provisioning-tools
sudo apt install btrfs-progs
```

テストスイートを実行するには、次のパッケージも必要です:
<!--
To run the testsuite, you'll also need:
-->

```bash
sudo apt install curl gettext jq sqlite3 uuid-runtime socat bind9-dnsutils
```

### ソースからの最新版のビルド <!-- From Source: Building the latest version -->
この方法は LXD の最新版をビルドしたい開発者や Linux ディストリビューションで提供されない LXD の特定のリリースをビルドするためのものです。 Linux ディストリビューションへ統合するためのソースからのビルドはここでは説明しません。それは将来別のドキュメントで取り扱うかもしれません。
<!--
These instructions for building from source are suitable for individual developers who want to build the latest version
of LXD, or build a specific release of LXD which may not be offered by their Linux distribution. Source builds for
integration into Linux distributions are not covered here and may be covered in detail in a separate document in the
future.
-->

```bash
git clone https://github.com/lxc/lxd
cd lxd
```

これで LXD の現在の開発ツリーをダウンロードしてソースツリー内に移動します。
その後下記の手順にしたがって実際に LXD をビルド、インストールしてください。
<!--
This will download the current development tree of LXD and place you in the source tree.
Then proceed to the instructions below to actually build and install LXD.
-->

### ソースからのリリース版のビルド <!-- From Source: Building a Release -->

LXD のリリース tarball は完全な依存ツリーと libraft と LXD のデータベースのセットアップに使用する libdqlite のローカルコピーをバンドルしています。
<!--
The LXD release tarballs bundle a complete dependency tree as well as a
local copy of libraft and libdqlite for LXD's database setup.
-->

```bash
tar zxvf lxd-4.18.tar.gz
cd lxd-4.18
```

これでリリース tarball を解凍し、ソースツリー内に移動します。
その後下記の手順にしたがって実際に LXD をビルド、インストールしてください。
<!--
This will unpack the release tarball and place you inside of the source tree.
Then proceed to the instructions below to actually build and install LXD.
-->

### ビルドの開始 <!-- Starting the Build -->

実際のビルドは Makefile の 2 回の別々の実行により行われます。 1 つは `make deps` でこれは LXD に必要とされるライブラリーをビルドします。もう 1 つは `make` で LXD 自体をビルドします。 `make deps` の最後に `make` の実行に必要な環境変数を設定するための手順が表示されます。新しいバージョンの LXD がリリースされたらこれらの環境変数の設定は変わるかもしれませんので、 `make deps` の最後に表示された手順を使うようにしてください。下記の手順（例示のために表示します）はあなたがビルドする LXD のバージョンのものとは一致しないかもしれません。
<!--
The actual building is done by two separate invocations of the Makefile: `make deps` -\- which builds libraries required
by LXD -\- and `make`, which builds LXD itself. At the end of `make deps`, a message will be displayed which will specify environment variables that should be set prior to invoking `make`. As new versions of LXD are released, these environment
variable settings may change, so be sure to use the ones displayed at the end of the `make deps` process, as the ones
below (shown for example purposes) may not exactly match what your version of LXD requires:
-->

ビルドには最低 2GB の RAM を推奨します。
<!--
We recommend having at least 2GB of RAM to allow the build to complete.
-->

```bash
make deps
# `make deps` が出力した export のコマンド列を使って環境変数を設定してください。
# 例:
#  export CGO_CFLAGS="${CGO_CFLAGS} -I$(go env GOPATH)/deps/dqlite/include/ -I$(go env GOPATH)/deps/raft/include/"
#  export CGO_LDFLAGS="${CGO_LDFLAGS} -L$(go env GOPATH)/deps/dqlite/.libs/ -L$(go env GOPATH)/deps/raft/.libs/"
#  export LD_LIBRARY_PATH="$(go env GOPATH)/deps/dqlite/.libs/:$(go env GOPATH)/deps/raft/.libs/:${LD_LIBRARY_PATH}"
#  export CGO_LDFLAGS_ALLOW="(-Wl,-wrap,pthread_create)|(-Wl,-z,now)"
make
```
<!--
```bash
make deps
# Follow the instructions from `make deps` to export the required environment variables.
# For example:
#  export CGO_CFLAGS="${CGO_CFLAGS} -I$(go env GOPATH)/deps/dqlite/include/ -I$(go env GOPATH)/deps/raft/include/"
#  export CGO_LDFLAGS="${CGO_LDFLAGS} -L$(go env GOPATH)/deps/dqlite/.libs/ -L$(go env GOPATH)/deps/raft/.libs/"
#  export LD_LIBRARY_PATH="$(go env GOPATH)/deps/dqlite/.libs/:$(go env GOPATH)/deps/raft/.libs/:${LD_LIBRARY_PATH}"
#  export CGO_LDFLAGS_ALLOW="(-Wl,-wrap,pthread_create)|(-Wl,-z,now)"
make
```
-->

### ソースからのビルド結果のインストール

ビルドが完了したら、ソースツリーを維持したまま、あなたのお使いのシェルのパスに `$(go env GOPATH)/bin` を追加し `LD_LIBRARY_PATH` 環境変数を `make deps` で表示された値に設定すれば、 LXD が利用できます。 `~/.bashrc` ファイルの場合は以下のようになります。
<!--
Once the build completes, you simply keep the source tree, add the directory referenced by `$(go env GOPATH)/bin` to
your shell path, and set the `LD_LIBRARY_PATH` variable printed by `make deps` to your environment. This might look
something like this for a `~/.bashrc` file:
-->

```bash
export PATH="${PATH}:$(go env GOPATH)/bin"
export LD_LIBRARY_PATH="$(go env GOPATH)/deps/dqlite/.libs/:$(go env GOPATH)/deps/raft/.libs/:${LD_LIBRARY_PATH}"
```

これで `lxd` と `lxc` コマンドの実行ファイルが利用可能になり LXD をセットアップするのに使用できます。 `LD_LIBRARY_PATH` 環境変数のおかげで実行ファイルは `$(go env GOPATH)/deps` にビルドされた依存ライブラリーを自動的に見つけて使用します。
<!--
Now, the `lxd` and `lxc` binaries will be available to you and can be used to set up LXD. The binaries will automatically find and use the dependencies built in `$(go env GOPATH)/deps` thanks to the `LD_LIBRARY_PATH` environment variable.
-->

### マシンセットアップ <!-- Machine Setup -->
LXD が非特権コンテナーを作成できるように、root ユーザーに対する sub{u,g}id の設定が必要です:
<!--
You'll need sub{u,g}ids for root, so that LXD can create the unprivileged containers:
-->

```bash
echo "root:1000000:1000000000" | sudo tee -a /etc/subuid /etc/subgid
```

これでデーモンを実行できます（`sudo` グループに属する全員が LXD とやりとりできるように `--group sudo` を指定します。別に指定したいグループを作ることもできます）:
<!--
Now you can run the daemon (the `-\-group sudo` bit allows everyone in the `sudo`
group to talk to LXD; you can create your own group if you want):
-->

```bash
sudo -E PATH=${PATH} LD_LIBRARY_PATH=${LD_LIBRARY_PATH} $(go env GOPATH)/bin/lxd --group sudo
```
