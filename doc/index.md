[![LXD](https://linuxcontainers.org/static/img/containers.png)](https://linuxcontainers.org/lxd)
# LXD
<!--
LXD is a next generation system container manager.
It offers a user experience similar to virtual machines but using Linux containers instead.
-->
LXD は次世代のシステムコンテナマネージャーです。
仮想マシンと同様のユーザーエクスペリエンスを提供しますが、仮想マシンの代わりに Linux コンテナを使用します。

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
-->
リリースアナウンスはこちらでご覧になれます: <https://linuxcontainers.org/lxd/news/>  
リリース tarball はこちらから取得できます: <https://linuxcontainers.org/lxd/downloads/>

## ステータス <!-- Status -->
Type            | Service               | Status
---             | ---                   | ---
CI (Linux)      | Jenkins               | [![Build Status](https://jenkins.linuxcontainers.org/job/lxd-github-commit/badge/icon)](https://jenkins.linuxcontainers.org/job/lxd-github-commit/)
CI (macOS)      | Travis                | [![Build Status](https://travis-ci.org/lxc/lxd.svg?branch=master)](https://travis-ci.org/lxc/lxd/)
CI (Windows)    | AppVeyor              | [![Build Status](https://ci.appveyor.com/api/projects/status/rb4141dsi2xm3n0x/branch/master?svg=true)](https://ci.appveyor.com/project/lxc/lxd/)
Documentation   | Godoc                 | [![GoDoc](https://godoc.org/github.com/lxc/lxd/client?status.svg)](https://godoc.org/github.com/lxc/lxd/client)
Static analysis | GoReport              | [![Go Report Card](https://goreportcard.com/badge/github.com/lxc/lxd)](https://goreportcard.com/report/github.com/lxc/lxd)
Translations    | Weblate               | [![Translation status](https://hosted.weblate.org/widgets/linux-containers/-/svg-badge.svg)](https://hosted.weblate.org/projects/linux-containers/lxd/)
Project status  | CII Best Practices    | [![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/1086/badge)](https://bestpractices.coreinfrastructure.org/projects/1086)

## LXD のパッケージからのインストール <!-- Installing LXD from packages -->
<!--
Instructions on installing LXD for a wide variety of Linux distributions and operating systems [can be found on our website](https://linuxcontainers.org/lxd/getting-started-cli/).
-->
さまざまな Linux ディストリビューションとオペレーティングシステムで LXD をインストールする方法は、[公式サイト](https://linuxcontainers.org/lxd/getting-started-cli/) をご覧ください。

## LXD のソースからのインストール <!-- Installing LXD from source -->
<!--
We recommend having the latest versions of liblxc (>= 2.0.0 required)
available for LXD development. Additionally, LXD requires Golang 1.9 or
later to work. On ubuntu, you can get those with:
-->
LXD の開発には liblxc の最新バージョン（2.0.0 以上が必要）を使用することをおすすめします。さらに Golang 1.9 以上が動作する必要があります。
Ubuntu では次のようにインストールできます:

```bash
sudo apt update
sudo apt install acl dnsmasq-base git golang liblxc1 lxc-dev libacl1-dev make pkg-config rsync squashfs-tools tar xz-utils
```

<!--
Note that when building LXC yourself, ensure to build it with the appropriate
security related libraries installed which our testsuite tests. Again, on
ubuntu, you can get those with:
-->
LXC を自分でビルドする場合は、テストスイートがテストする、関連する適切なセキュリティ関連のライブラリがインストールされていることを確認してください。
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
sudo apt install btrfs-tools
```

<!--
To run the testsuite, you'll also need:
-->
テストスイートを実行するには、次のパッケージも必要です:

```bash
sudo apt install curl gettext jq sqlite3 uuid-runtime bzr
```


### ツールのビルド <!-- Building the tools -->
<!--
LXD consists of two binaries, a client called `lxc` and a server called `lxd`.
These live in the source tree in the `lxc/` and `lxd/` dirs, respectively.
To get the code, set up your go environment:
-->
LXD にはふたつのバイナリが含まれます。クライアントである `lxc` と、サーバである `lxd` です。
これらはソースツリーのそれぞれ `lxc/`、`lxd/` に含まれます。
コードを取得するには、Go 環境をセットアップします:

```bash
mkdir -p ~/go
export GOPATH=~/go
```

<!--
And then download it as usual:
-->
そして次のようにダウンロードしてください:


```bash
go get -d -v github.com/lxc/lxd/lxd
cd $GOPATH/src/github.com/lxc/lxd
make
```

<!--
...which will give you two binaries in `$GOPATH/bin`, `lxd` the daemon binary,
and `lxc` a command line client to that daemon.
-->
すると、ふたつのバイナリは `$GOPATH/bin` から取得できます。`lxd` はデーモンバイナリであり、このデーモンに接続するためのコマンドラインクライアントは `lxc` です。

### マシンセットアップ <!-- Machine Setup -->
<!--
You'll need sub{u,g}ids for root, so that LXD can create the unprivileged
containers:
-->
LXD が非特権コンテナを作成できるように、root ユーザに対する sub{u,g}id の設定が必要です:

```bash
echo "root:1000000:65536" | sudo tee -a /etc/subuid /etc/subgid
```

<!--
Now you can run the daemon (the `--group` sudo bit allows everyone in the sudo
group to talk to LXD; you can create your own group if you want):
-->
これでデーモンを実行できます（sudo グループに属する全員が LXD とやりとりできるように `--group sudo` を指定します。別に指定したいグループを作ることもできます）:

```bash
sudo -E $GOPATH/bin/lxd --group sudo
```

## LXD を使い始める <!-- Getting started with LXD -->
<!--
Now that you have LXD running on your system you can read the [getting started guide](https://linuxcontainers.org/lxd/getting-started-cli/) or go through more examples and configurations in [our documentation](https://github.com/lxc/lxd/tree/master/doc).
-->
ここまでで、システム上で LXD が実行されているでしょうから、[はじめに](https://linuxcontainers.org/lxd/getting-started-cli/) という文書を読んだり、[ドキュメント](https://github.com/lxc/lxd/tree/master/doc) の例や設定を見たりできます。

## バグレポート <!-- Bug reports -->
<!--
Bug reports can be filed at: <https://github.com/lxc/lxd/issues/new>
-->
バグ報告はこちらから行えます: <https://github.com/lxc/lxd/issues/new>

## コントリビュート <!-- Contributing -->
<!--
Fixes and new features are greatly appreciated but please read our [contributing guidelines](contributing.md) first.
-->
修正や新機能の追加は歓迎です。最初に [contributing guidelines](contributing.md) を読んでください。

## サポートとディスカッション <!-- Support and discussions -->
### Forum
A discussion forum is available at: <https://discuss.linuxcontainers.org>

### Mailing-lists
We use the LXC mailing-lists for developer and user discussions, you can
find and subscribe to those at: <https://lists.linuxcontainers.org>

### IRC
If you prefer live discussions, some of us also hang out in
[#lxcontainers](http://webchat.freenode.net/?channels=#lxcontainers) on irc.freenode.net.

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

#### How do I configure LXD storage?
LXD supports btrfs, ceph, directory, lvm and zfs based storage.

First make sure you have the relevant tools for your filesystem of
choice installed on the machine (btrfs-progs, lvm2 or zfsutils-linux).

By default, LXD comes with no configured network or storage.
You can get a basic configuration done with:

```bash
    lxd init
```

`lxd init` supports both directory based storage and ZFS.
If you want something else, you'll need to use the `lxc storage` command:

```bash
lxc storage create default BACKEND [OPTIONS...]
lxc profile device add default root disk path=/ pool=default
```

BACKEND is one of `btrfs`, `ceph`, `dir`, `lvm` or `zfs`.

Unless specified otherwise, LXD will setup loop based storage with a sane default size.

For production environments, you should be using block backed storage
instead both for performance and reliability reasons.

#### How can I live migrate a container using LXD?
Live migration requires a tool installed on both hosts called
[CRIU](http://criu.org), which is available in Ubuntu via:

```bash
sudo apt-get install criu
```

Then, launch your container with the following,

```bash
lxc launch ubuntu $somename
sleep 5s # let the container get to an interesting state
lxc move host1:$somename host2:$somename
```

And with luck you'll have migrated the container :). Migration is still in
experimental stages and may not work for all workloads. Please report bugs on
lxc-devel, and we can escalate to CRIU lists as necessary.

#### Can I bind mount my home directory in a container?
Yes. The easiest way to do that is using a privileged container to avoid file ownership issues:

1.a) create a container.

```bash
lxc launch ubuntu privilegedContainerName -c security.privileged=true
```

1.b) or, if your container already exists.

```bash
lxc config set privilegedContainerName security.privileged true
```

2) then.

```bash
lxc config device add privilegedContainerName shareName disk source=/home/$USER path=/home/ubuntu
```

#### How can I run docker inside a LXD container?
In order to run Docker inside a LXD container the `security.nesting` property of the container should be set to `true`. 

```bash
lxc config set <container> security.nesting true
```

Note that LXD containers cannot load kernel modules, so depending on your
Docker configuration you may need to have the needed extra kernel modules
loaded by the host.

You can do so by setting a comma separate list of kernel modules that your container needs with:

```bash
lxc config set <container> linux.kernel_modules <modules>
```

We have also received some reports that creating a `/.dockerenv` file in your
container can help Docker ignore some errors it's getting due to running in a
nested environment.

## Hacking on LXD
### Directly using the REST API
The LXD REST API can be used locally via unauthenticated Unix socket or remotely via SSL encapsulated TCP.

#### Via Unix socket

```bash
curl --unix-socket /var/lib/lxd/unix.socket \
    -H "Content-Type: application/json" \
    -X POST \
    -d @hello-ubuntu.json \
    lxd/1.0/containers
```

#### Via TCP
TCP requires some additional configuration and is not enabled by default.

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

#### JSON payload
The `hello-ubuntu.json` file referenced above could contain something like:

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
        "alias":"14.04"
    }
}
```
