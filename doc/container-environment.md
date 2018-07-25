# コンテナ実行環境 <!-- Container runtime environment -->
<!--
LXD attempts to present a consistent environment to the container it runs.
-->
LXD は実行するコンテナに一貫性のある環境を提供しようとします。

<!--
The exact environment will differ slightly based on kernel features and user
configuration but will otherwise be identical for all containers.
-->
正確な環境はカーネルの機能やユーザの設定によって若干異なりますが、それ以外は
全てのコンテナに対して同一です。

## PID1
<!--
LXD spawns whatever is located at `/sbin/init` as the initial process of the container (PID 1).
This binary should act as a proper init system, including handling re-parented processes.
-->
LXD は何であれ `/sbin/init` に置かれているものをコンテナの初期プロセス (PID 1) として起動します。
このバイナリは親が変更されたプロセス (訳注: ゾンビプロセスなど) の処理を含めて適切な init
システムとして振る舞う必要があります。

<!--
LXD's communication with PID1 in the container is limited to two signals:
 - `SIGINT` to trigger a reboot of the container
 - `SIGPWR` (or alternatively `SIGRTMIN`+3) to trigger a clean shutdown of the container
-->
LXD がコンテナの PID1 とコミュニケーションするのは以下の2つのシグナルだけです。
 - `SIGINT` コンテナのリブートをトリガーする
 - `SIGPWR` (かあるいは `SIGRTMIN`+3) コンテナのクリーンなシャットダウンをトリガーする

<!--
The initial environment of PID1 is blank except for `container=lxc` which can
be used by the init system to detect the runtime.
-->
PID1 の初期環境は `container=lxc` 以外は空です。 init システムは `container=lxc`
をランタイムを検出する (訳注: lxc で動いていることを知る) ために使用できます。

<!--
All file descriptors above the default 3 are closed prior to PID1 being spawned.
-->
デフォルトの 3 個 (訳注: stdin, stdout, stderr) より上の全てのファイルディスクリプタは
PID1 が起動される前に閉じられます。

## ファイルシステム <!-- Filesystem -->
<!--
LXD assumes that any image it uses to create a new container from will come with at least:
-->
LXD は使用するどのイメージから生成する新規のコンテナは少なくとも以下のファイルシステムを
含むことを前提とします。

 - `/dev` (空のディレクトリ) <!-- (empty) -->
 - `/proc` (空のディレクトリ) <!-- (empty) -->
 - `/sbin/init` (実行ファイル) <!-- (executable) -->
 - `/sys` (空のディレクトリ) <!-- (empty) -->

## デバイス <!-- Devices -->
<!--
LXD containers have a minimal and ephemeral `/dev` based on a tmpfs filesystem.
Since this is a tmpfs and not a devtmpfs, device nodes will only appear if manually created.
-->
LXD のコンテナは tmpfs ファイルシステムをベースとする最低限で一時的な `/dev` を
持ちます。これは tmpfs であって devtmpfs ではないので、デバイスノードは手動で作成
されたときのみ現れます。

<!--
The standard set of device nodes will be setup:
-->
デバイスノードの標準セットでは以下のデバイスがセットアップされます。

 - `/dev/console`
 - `/dev/fd`
 - `/dev/full`
 - `/dev/log`
 - `/dev/null`
 - `/dev/ptmx`
 - `/dev/random`
 - `/dev/stdin`
 - `/dev/stderr`
 - `/dev/stdout`
 - `/dev/tty`
 - `/dev/urandom`
 - `/dev/zero`

<!--
On top of the standard set of devices, the following are also setup for convenience:
-->
標準セットのデバイスに加えて、以下のデバイスも利便性のためにセットアップされます。

 - `/dev/fuse`
 - `/dev/net/tun`
 - `/dev/mqueue`

## マウント <!-- Mounts -->
<!--
The following mounts are setup by default under LXD:
-->
LXD では以下のマウントがデフォルトでセットアップされます。

 - `/proc` (proc)
 - `/sys` (sysfs)
 - `/sys/fs/cgroup/*` (cgroupfs) (cgroup namespace サポートを欠くカーネルの場合のみ) <!-- (only on kernels lacking cgroup namespace support) -->

<!--
The following paths will also be automatically mounted if present on the host:
-->
以下のパスがホスト上に存在する場合は自動的にマウントされます。

 - `/proc/sys/fs/binfmt_misc`
 - `/sys/firmware/efi/efivars`
 - `/sys/fs/fuse/connections`
 - `/sys/fs/pstore`
 - `/sys/kernel/debug`
 - `/sys/kernel/security`

<!--
The reason for passing all of those is legacy init systems which require
those to be mounted or be mountabled inside the container.
-->
これらを引き渡す理由は、これらがマウントされているか、コンテナ内でマウント
できるようになっているかが必要とされているレガシーな init システムのためです。

<!--
The majority of those will not be writable (or even readable) from inside an
unprivileged container and will be blocked by our AppArmor policy inside
privileged containers.
-->
これらのほとんどは非特権コンテナ内からは書き込み可能ではなく (あるいは読み取り可能
ですらなく)、特権コンテナ内では LXD の AppArmor ポリシーによってブロックされます。

## ネットワーク <!-- Network -->
<!--
LXD containers may have any number of network devices attached to them.
The naming for those unless overridden by the user is ethX where X is an incrementing number.
-->
LXD コンテナはネットワークデバイスをいくつでもアタッチできます。
これらの名前はユーザにオーバーライドされない限りは ethX で X は
連番です。

## コンテナからホストへのコミュニケーション <!-- Container to host communication -->
<!--
LXD sets up a socket at `/dev/lxd/sock` which root in the container can use to communicate with LXD on the host.
-->
LXD は `/dev/lxd/sock` にソケットをセットアップし、コンテナ内の root ユーザはこれを使ってホストの
LXD とコミュニケーションできます。

<!--
The API is [documented here](dev-lxd.md).
-->
API は [ここにドキュメント化されています](dev-lxd.md).

## LXCFS
<!--
If LXCFS is present on the host, it will automatically be setup for the container.
-->
ホストに LXCFS がある場合は、コンテナ用に自動的にセットアップされます。

<!--
This normally results in a number of `/proc` files being overridden through bind-mounts.
On older kernels a virtual version of `/sys/fs/cgroup` may also be setup by LXCFS.
-->
これは通常いくつかの `/proc` ファイルになり、それらは bind mount を通してオーバーライド
されます。古いカーネルでは `/sys/fs/cgroup` の仮想バージョンも LXCFS によりセットアップ
されるかもしれません。
