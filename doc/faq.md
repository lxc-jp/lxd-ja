# よく聞かれる質問 (FAQ)

### LXD サーバーをリモートからアクセス可能にするには？
デフォルトでは LXD サーバーはローカルの unix ソケットのみをリッスンしているためネットワークからはアクセス可能ではありません。
リッスンする追加のアドレスを指定することでネットワークから LXD を利用可能にできます。
これは `core.https_address` 設定で実現できます。

現状のサーバー設定を確認するには、以下のコマンドを実行します。

```bash
lxc config show
```

リッスンするアドレスを設定するには、利用可能なアドレスを調べた上でサーバーで `config set` コマンドを実行します。

```bash
ip addr
lxc config set core.https_address 192.168.1.15
```

### https 経由で `lxc remote add` を実行したらパスワードを聞かれたがどうすればよいか？
デフォルトではセキュリティー上の理由から LXD はパスワードを設定していないため、 `lxc remote add` でリモートは追加できません。
パスワードを設定するには LXD が実行中のホスト上で以下のコマンドを実行します。

```bash
lxc config set core.trust_password SECRET
```

これでリモートパスワードが設定されるので、 `lxc remote add` 実行時にこのパスワードを使用できます。

あるいはクライアント証明書を `.config/lxc/client.crt` からサーバーにコピーして以下のコマンドで追加すれば、パスワードを設定しなくてもサーバーにアクセスできます。

```bash
lxc config trust add client.crt
```

### LXD のストレージを設定するには？
LXD は btrfs, ceph, directory, lvm と zfs ベースのストレージをサポートします。

まず、あなたが選択したファイルシステムに関連するツール（btrfs-progs, lvm2 あるいは zfsutils-linux）をマシーン上にインストールしてください。

（訳注：LXD をインストールしただけの）デフォルトの状態では LXD はネットワークやストレージが設定されていません。
以下のコマンドにより基本の設定を実行できます。

```bash
lxd init
```

`lxd init` はディレクトリーベースのストレージと ZFS の両方をサポートします。
それ以外のストレージを使いたい場合は `lxc storage` コマンドを使う必要があります。

```bash
lxc storage create default BACKEND [OPTIONS...]
lxc profile device add default root disk path=/ pool=default
```

BACKEND は `btrfs`, `ceph`, `dir`, `lvm`, `zfs` のいずれかです。

明示的に指定しない場合、 LXD は妥当なデフォルトサイズでループデバイスをベースにしたストレージをセットアップします。

本番環境ではパフォーマンスと信頼性の両方の理由でブロックデバイスをベースにしたストレージを使うべきです。

### LXD を使ってコンテナーをマイグレートするには？
ライブマイグレーションには [CRIU](https://criu.org) と呼ばれるツールを両方のホストにインストールする必要があります。
Ubuntu では以下のコマンドでインストールできます。

```bash
sudo apt install criu
```

次に以下のコマンドでコンテナーを起動します。

```bash
lxc launch ubuntu SOME-NAME
sleep 5s # コンテナーの起動が完了するのをを待ちます。
lxc move host1:SOME-NAME host2:SOME-NAME
```


これで運が良ければコンテナーがマイグレートされます :)。
マイグレーションはいまだ実験的な段階にあり環境によっては動かないかもしれません。
動かない場合は lxc-devel メーリングリストに報告してください。
そうすれば私たちが必要に応じて CRIU メーリングリストに報告します。

### 自分のホームディレクトリをコンテナー内にバインドマウントできますか？
はい。ディスクデバイスを使って以下のようにすればできます。

```bash
lxc config device add container-name home disk source=/home/${USER} path=/home/ubuntu
```

非特権コンテナーの場合は、以下のいずれかも必要です。

 - `lxc config device add` の実行に `shift=true` を指定する。これは `shiftfs` がサポートされているかに依存します（`lxc info` 参照）。
 - raw.idmap エントリーを使用する（[ユーザー名前空間 (user namespace) 用の ID のマッピング](userns-idmap.md) 参照）。
 - マウントしたいホームディレクトリに再帰的な POSIX ACL を設定する。

上記のいずれかを実行すればコンテナー内のユーザーは read/write パーミッションに沿ってアクセス可能です。
上記のいずれも設定しない場合、アクセスしようとすると uid/gid (65536:65536) のオーバーフローが発生し、全ユーザーで読み取り可能 (world readable) 以外のファイルへのアクセスは失敗します。

特権コンテナーではコンテナー内の uid/gid が外部と同じなためこの問題はありません。
しかしこれは特権コンテナーのセキュリティーの問題のほとんどの原因でもあります。

### LXD コンテナー内で docker を動かすには？
LXD のコンテナー内で Docker を動かすにはコンテナーの `security.nesting` プロパティーを `true` にする必要があります。

```bash
lxc config set <container> security.nesting true
```

LXD コンテナーはカーネルモジュールをロードすることはできないので、お使いの Docker の設定によっては、ホストで追加のカーネルモジュールをロードする必要があることに注意してください。

コンテナーが必要とするカーネルモジュールのカンマ区切りリストを以下のコマンドで指定すればホストでそれらのモジュールをロードできます。

```bash
lxc config set <container> linux.kernel_modules <modules>
```

コンテナー内に `/.dockerenv` ファイルを作成するとネストした環境内で実行しているために発生するエラーを Docker が無視するようにできるという報告もあります。

## コンテナーの起動に関する問題
もしコンテナーが起動しない場合や、期待通りの動きをしない場合に最初にすべきことは、コンテナーが生成したコンソールログを見ることです。
これには `lxc console --show-log CONTAINERNAME` コマンドを使います。

次の例では、`systemd` が起動しない RHEL 7 システムを調べています。

    # lxc console --show-log systemd
    Console log:

    Failed to insert module 'autofs4'
    Failed to insert module 'unix'
    Failed to mount sysfs at /sys: Operation not permitted
    Failed to mount proc at /proc: Operation not permitted
    [!!!!!!] Failed to mount API filesystems, freezing.

ここでのエラーは、/sys と /proc がマウントできないというエラーです。これは非特権コンテナーでは正しい動きです。
しかし、LXD は _可能であれば_ 自動的にこれらのファイルシステムをマウントします。

[コンテナーの要件](container-environment.md) では、コンテナーには `/sbin/init` が存在するだけでなく、空の `/dev`、`/proc`、`/sys` フォルダーが存在していなければならないと定められています。
もしこれらのフォルダーが存在しなければ、LXD はこれらをマウントできません。そして、systemd がこれらをマウントしようとします。
非特権コンテナーでは、systemd はこれを行う権限はなく、フリーズしてしまいます。

何かが変更される前に環境を見ることはできます。`raw.lxc` 設定パラメーターを使って、明示的にコンテナー内の init を変更できます。
これは Linux カーネルコマンドラインに `init=/bin/bash` を設定するのと同じです。

    lxc config set systemd raw.lxc 'lxc.init.cmd = /bin/bash'

次のようになります:

    root@lxc-01:~# lxc config set systemd raw.lxc 'lxc.init.cmd = /bin/bash'
    root@lxc-01:~# lxc start systemd
    root@lxc-01:~# lxc console --show-log systemd

    Console log:

    [root@systemd /]#
    root@lxc-01:~#

コンテナーが起動しましたので、コンテナー内で期待通りに動いていないことを確認できます。

    root@lxc-01:~# lxc exec systemd bash
    [root@systemd ~]# ls
    [root@systemd ~]# mount
    mount: failed to read mtab: No such file or directory
    [root@systemd ~]# cd /
    [root@systemd /]# ls /proc/
    sys
    [root@systemd /]# exit

LXD は自動修復を試みますので、起動時に作成されたフォルダもあります。コンテナーをシャットダウンして再起動すると問題は解決されます。
しかし問題の根源は依然として存在しています。**テンプレートに必要なファイルが含まれていないという問題です**。

## ネットワークの問題

大規模な[プロダクション環境](production-setup.md)では、複数の VLAN を持ち、LXD クライアントを直接それらの VLAN に接続するのが一般的です。
netplan と systemd-networkd を使っている場合、いくつかの最悪の問題を引き起こす可能性があるバグに遭遇するでしょう。

### VLAN ベースのブリッジでは netplan で systemd-networkd が使えない

執筆時点（2019-03-05）では、netplan は VLAN にアタッチされたブリッジにランダムな MAC アドレスを割り当てられません。
常に同じ MAC アドレスを選択するため、同じネットワークセグメントに複数のマシンが存在する場合、レイヤー 2 の問題が発生します。
複数のブリッジを作成することも困難です。代わりに `network-manager` を使ってください。
設定例は次のようになります。管理アドレスが 10.61.0.25 で、VLAN102 をクライアントのトラフィックに使います。

    network:
      version: 2
      renderer: NetworkManager
      ethernets:
        eth0:
          dhcp4: no
          accept-ra: no
          # This is the 'Management Address'
          addresses: [ 10.61.0.25/24 ]
          gateway4: 10.61.0.1
          nameservers:
            addresses: [ 1.1.1.1, 8.8.8.8 ]
        eth1:
          dhcp4: no
          accept-ra: no
          # A bogus IP address is required to ensure the link state is up
          addresses: [ 10.254.254.25/32 ]

      vlans:
        vlan102:
          accept-ra: no
          dhcp4: no
          id: 102
          link: eth1

      bridges:
        br102:
          accept-ra: no
          dhcp4: no
          interfaces: [ "vlan102" ]
          # A bogus IP address is required to ensure the link state is up
          addresses: [ 10.254.102.25/32 ]
          parameters:
            stp: false

#### 注意事項

* eth0 はデフォルトゲートウェイの指定がある管理インターフェースです
* vlan102 は eth1 を使います
* br102 は vlan102 を使います。そして __bogus な /32 の IP アドレスが割り当てられています__

他に重要なこととして、`stp: false` を設定することがあります。そうしなければ、ブリッジは最大で 10 秒間 `learning` 状態となります。これはほとんどの DHCP リクエストが投げられる期間よりも長いです。
クロスコネクトされてループを引き起こす可能性はありませんので、このように設定しても安全です。

### 'port security' に気をつける

スイッチは MAC アドレスの変更を許さず、不正な MAC アドレスのトラフィックをドロップするか、ポートを完全に無効にするものが多いです。
ホストから LXD インスタンスに ping できたとしても、_異なった_ ホストから ping できない場合は、これが原因の可能性があります。
この原因を突き止める方法は、アップリンク（この場合は eth1）で tcpdump を実行することです。
すると、応答は送るが ACK を取得できない 'ARP Who has xx.xx.xx.xx tell yy.yy.yy.yy'、もしくは ICMP パケットが行き来しているものの、決して他のホストで受け取られないのが見えるでしょう。

### 不必要に特権コンテナーを実行しない

特権コンテナーはホスト全体に影響する処理を行うことができます。例えば、ネットワークカードをリセットするために、/sys 内のものを使えます。
これは **ホスト全体** に対してリセットを行い、ネットワークの切断を引き起こします。
ほぼすべてのことが非特権コンテナーで実行できます。コンテナー内から NFS マウントしたいというような、通常とは異なる特権が必要なケースでは、バインドマウントを使う必要があるかもしれません。
