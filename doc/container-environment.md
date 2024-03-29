(container-runtime-environment)=
# コンテナ実行環境

LXD は実行する全てのコンテナに一貫性のある環境を提供しようとします。

正確な環境はカーネルの機能やユーザーの設定によって若干異なりますが、それ以外は全てのコンテナに対して同一です。

## ファイルシステム

LXDは使用するどのイメージから生成する新規のコンテナは少なくとも以下のルートレベルのディレクトリが存在することを前提とします。

 - `/dev` (空のディレクトリ)
 - `/proc` (空のディレクトリ)
 - `/sbin/init` (実行ファイル)
 - `/sys` (空のディレクトリ)

## デバイス

LXD のコンテナは`tmpfs`ファイルシステムをベースとする最低限で一時的な`/dev`を持ちます。
これは`tmpfs`であって`devtmpfs`ファイルシステムではないので、デバイスノードは手動で作成されたときのみ現れます。

デバイスノードの標準セットでは以下のデバイスが自動的にセットアップされます。

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

標準セットのデバイスに加えて、以下のデバイスも利便性のためにセットアップされます。

 - `/dev/fuse`
 - `/dev/net/tun`
 - `/dev/mqueue`

## ネットワーク

LXDコンテナはネットワークデバイスをいくつでもアタッチできます。
これらの名前はユーザーにオーバーライドされない限りは`ethX`で`X`は連番です。

## コンテナからホストへのコミュニケーション

LXDは`/dev/lxd/sock`にソケットをセットアップし、コンテナ内のrootユーザーはこれを使ってホストのLXDとコミュニケーションできます。

APIドキュメントは{doc}`dev-lxd`を参照してください。

## マウント

以下のマウントがデフォルトでセットアップされます。

 - `/proc` (`proc`)
 - `/sys` (`sysfs`)
 - `/sys/fs/cgroup/*` (`cgroupfs`) (cgroup namespace サポートを欠くカーネルの場合のみ)

以下のパスがホスト上に存在する場合は自動的にマウントされます。

 - `/proc/sys/fs/binfmt_misc`
 - `/sys/firmware/efi/efivars`
 - `/sys/fs/fuse/connections`
 - `/sys/fs/pstore`
 - `/sys/kernel/debug`
 - `/sys/kernel/security`

これらのパスを引き渡す理由は、これらがマウントされているか、コンテナ内でマウント可能であることが必要とされているレガシーなinitシステムのためです。

これらのパスほとんどは非特権コンテナ内からは書き込み可能ではなく(あるいは読み取り可能ですらなく)、特権コンテナ内ではLXDのAppArmorポリシーによってブロックされます。

## LXCFS

ホストに LXCFS がある場合は、コンテナ用に自動的にセットアップされます。

これは通常いくつかの`/proc`ファイルになり、それらは bind mount を通してオーバーライドされます。
古いカーネルでは`/sys/fs/cgroup`の仮想バージョンもLXCFSによりセットアップされるかもしれません。

## PID1

LXDは何であれ`/sbin/init`に置かれているものをコンテナの初期プロセス(PID 1)として起動します。
このバイナリは親が変更されたプロセス(訳注: ゾンビプロセスなど)の処理を含めて適切なinitシステムとして振る舞う必要があります。

LXDがコンテナのPID1とコミュニケーションするのは以下の2つのシグナルだけです。

 - `SIGINT` コンテナのリブートをトリガーする
 - `SIGPWR` (かあるいは `SIGRTMIN`+3) コンテナのクリーンなシャットダウンをトリガーする

PID1の初期環境は`container=lxc`以外は空です。
initシステムは`container=lxc`をランタイムの検出(訳注: lxcで動いていることを知る)に使用できます。

デフォルトの3個(訳注: stdin, stdout, stderr)より上の全てのファイルディスクリプタはPID1が起動される前に閉じられます。
