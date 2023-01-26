---
discourse: 8178
relatedlinks: "[LXD のインストール](https://linuxcontainers.org/ja/lxd/getting-started-cli/)"
---

# LXDをインストールするには

LXD をインストールする最も簡単な方法は提供されているパッケージのどれかをインストールすることですが、ソースから LXD をインストールすることもできます。

% Include some content from [../README.md](../README.md)
```{include} ../README.md
    :start-after: <!-- Include start installing -->
    :end-before: <!-- Include end installing -->
```

(installing_from_source)=
## LXDをソースからインストールする

LXD の開発には `liblxc` の最新バージョン（4.0.0 以上が必要）を使用することをおすすめします。
さらに LXD が動作するためには Golang 1.18 以上が必要です。
Ubuntu では次のようにインストールできます:

```bash
sudo apt update
sudo apt install acl attr autoconf automake dnsmasq-base git golang libacl1-dev libcap-dev liblxc1 liblxc-dev libsqlite3-dev libtool libudev-dev liblz4-dev libuv1-dev make pkg-config rsync squashfs-tools tar tcl xz-utils ebtables
```

デフォルトのストレージドライバである `dir` ドライバに加えて、LXD ではいくつかのストレージドライバが使えます。
これらのツールをインストールすると、initramfs への追加が行われ、ホストのブートが少しだけ遅くなるかもしれませんが、特定のドライバを使いたい場合には必要です:

```bash
sudo apt install lvm2 thin-provisioning-tools
sudo apt install btrfs-progs
```

テストスイートを実行するには、次のパッケージも必要です:

```bash
sudo apt install curl gettext jq sqlite3 socat bind9-dnsutils
```

### ソースから最新版をビルドする

この方法は LXD の最新版をビルドしたい開発者や Linux ディストリビューションで提供されない LXD の特定のリリースをビルドするためのものです。 Linux ディストリビューションへ統合するためのソースからのビルドはここでは説明しません。それは将来、別のドキュメントで取り扱うかもしれません。

```bash
git clone https://github.com/lxc/lxd
cd lxd
```

これで LXD の現在の開発ツリーをダウンロードしてソースツリー内に移動します。
その後下記の手順にしたがって実際に LXD をビルド、インストールしてください。

### ソースからリリース版をビルドする

LXD のリリース tarball は完全な依存ツリーと `libraft` と LXD のデータベースのセットアップに使用する `libdqlite` のローカルコピーをバンドルしています。

```bash
tar zxvf lxd-4.18.tar.gz
cd lxd-4.18
```

これでリリース tarball を解凍し、ソースツリー内に移動します。
その後下記の手順にしたがって実際に LXD をビルド、インストールしてください。

### ビルドを開始する

実際のビルドは Makefile の 2 回の別々の実行により行われます。 1 つは `make deps` でこれは LXD に必要とされるライブラリーをビルドします。もう 1 つは `make` で LXD 自体をビルドします。 `make deps` の最後に `make` の実行に必要な環境変数を設定するための手順が表示されます。新しいバージョンの LXD がリリースされたらこれらの環境変数の設定は変わるかもしれませんので、 `make deps` の最後に表示された手順を使うようにしてください。下記の手順（例示のために表示します）はあなたがビルドする LXD のバージョンのものとは一致しないかもしれません。

ビルドには最低 2GB の RAM を搭載することを推奨します。

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

### ソースからのビルド結果のインストール

ビルドが完了したら、ソースツリーを維持したまま、あなたのお使いのシェルのパスに `$(go env GOPATH)/bin` を追加し `LD_LIBRARY_PATH` 環境変数を `make deps` で表示された値に設定すれば、 LXD が利用できます。 `~/.bashrc` ファイルの場合は以下のようになります。

```bash
export PATH="${PATH}:$(go env GOPATH)/bin"
export LD_LIBRARY_PATH="$(go env GOPATH)/deps/dqlite/.libs/:$(go env GOPATH)/deps/raft/.libs/:${LD_LIBRARY_PATH}"
```

これで `lxd` と `lxc` コマンドの実行ファイルが利用可能になり LXD をセットアップするのに使用できます。 `LD_LIBRARY_PATH` 環境変数のおかげで実行ファイルは `$(go env GOPATH)/deps` にビルドされた依存ライブラリーを自動的に見つけて使用します。

### マシンセットアップ

LXD が非特権コンテナを作成できるように、root ユーザーに対する sub{u,g}id の設定が必要です:

```bash
echo "root:1000000:1000000000" | sudo tee -a /etc/subuid /etc/subgid
```

これでデーモンを実行できます（`sudo` グループに属する全員が LXD とやりとりできるように `--group sudo` を指定します。別に指定したいグループを作ることもできます）:

```bash
sudo -E PATH=${PATH} LD_LIBRARY_PATH=${LD_LIBRARY_PATH} $(go env GOPATH)/bin/lxd --group sudo
```

```{note}
`newuidmap/newgidmap`ツールがシステムに存在し、`/etc/subuid`、`/etc/subgid`が存在する場合は、rootユーザーに少なくとも10MのUID/GIDの連続した範囲を許可するように設定する必要があります。
```

## LXDをアップグレードする

LXD を新しいバージョンにアップグレードした後、 LXD はデータベースを新しいスキーマにアップデートする必要があるかもしれません。
このアップデートは LXD のアップグレードの後のデーモン起動時に自動的に実行されます。
アップデート前のデータベースのバックアップはアクティブなデータベースと同じ場所 (例えば snap の場合は `/var/snap/lxd/common/lxd/database`) に保存されます。

```{important}
スキーマのアップデート後は、古いバージョンの LXD はデータベースを無効とみなすかもしれません。
これはつまり LXD をダウングレードしてもあなたの LXD の環境は利用不可能と言われるかもしれないということです。

このようなダウングレードが必要な場合は、ダウングレードを行う前にデータベースのバックアップをリストアしてください。
```
