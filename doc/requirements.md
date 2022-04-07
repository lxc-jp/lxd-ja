# 動作環境
LXD は Go 1.18 以上を必要とし、 golang のコンパイラのみでテストされています。
(訳注: 以前は gccgo もサポートされていましたが golang のみになりました)

ビルドには最低 2GB の RAM を推奨します。

## 必要なカーネルバージョン
サポートされる最小のカーネルバージョンは 5.4 です。

LXD には以下の機能をサポートするカーネルが必要です。

 * Namespaces (pid, net, uts, ipc と mount)
 * Seccomp

以下のオプションの機能はさらなるカーネルオプションを必要とします。

 * Namespaces (user と cgroup)
 * AppArmor (mount mediation に対する Ubuntu パッチを含む)
 * Control Groups (blkio, cpuset, devices, memory, pids と net\_prio)
 * CRIU (正確な詳細は CRIU のアップストリームを参照のこと)

さらに使用している LXC のバージョンで必要とされる他のカーネルの機能も
必要です。

## LXC
LXD は以下のビルドオプションでビルドされた LXC 4.0.0 以上を必要とします。

 * apparmor (もし LXD の apparmor サポートを使用するのであれば)
 * seccomp

Ubuntu を含む、さまざまなディストリビューションの最近のバージョンを
動かすためには、 LXCFS もインストールする必要があります。

## QEMU
仮想マシンを利用するには QEMU 6.0 以降が必要です。

## 追加のライブラリー(と開発用のヘッダ)
LXD はデータベースとして `dqlite` を使用しています。
ビルドしセットアップするためには `make deps` を実行してください。

LXD は他にもいくつかの (たいていはパッケージ化されている) C ライブラリーを使用しています。

 - libacl1
 - libcap2
 - liblz4 (`dqlite` で使用)
 - libuv1 (`dqlite` で使用)
 - libsqlite3 >= 3.25.0 (`dqlite` で使用)

ライブラリーそのものとライブラリーの開発用ヘッダ (-dev パッケージ)の全てを
インストールしたことを確認してください。
