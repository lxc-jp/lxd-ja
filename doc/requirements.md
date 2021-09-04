# 動作環境
<!-- Requirements -->
## Go

<!--
LXD requires Go 1.13 or higher and is only tested with the golang compiler.
-->
LXD は Go 1.13 以上を必要とし、 golang のコンパイラのみでテストされています。
(訳注: 以前は gccgo もサポートされていましたが golang のみになりました)

ビルドには最低 2GB の RAM を推奨します。
<!--
We recommend having at least 2GB of RAM to allow the build to complete.
-->

## 必要なカーネルバージョン <!-- Kernel requirements -->
<!--
The minimum supported kernel version is 3.13.
-->
サポートされる最小のカーネルバージョンは 3.13 です。

<!--
LXD requires a kernel with support for:
-->
LXD には以下の機能をサポートするカーネルが必要です。

<!--
 * Namespaces (pid, net, uts, ipc and mount)
 * Seccomp
-->
 * Namespaces (pid, net, uts, ipc と mount)
 * Seccomp

<!--
The following optional features also require extra kernel options:
-->
以下のオプションの機能はさらなるカーネルオプションを必要とします。

<!--
 * Namespaces (user and cgroup)
 * AppArmor (including Ubuntu patch for mount mediation)
 * Control Groups (blkio, cpuset, devices, memory, pids and net\_prio)
 * CRIU (exact details to be found with CRIU upstream)
-->
 * Namespaces (user と cgroup)
 * AppArmor (mount mediation に対する Ubuntu パッチを含む)
 * Control Groups (blkio, cpuset, devices, memory, pids と net\_prio)
 * CRIU (正確な詳細は CRIU のアップストリームを参照のこと)

<!--
As well as any other kernel feature required by the LXC version in use.
-->
さらに使用している LXC のバージョンで必要とされる他のカーネルの機能も
必要です。

## LXC
<!--
LXD requires LXC 3.0.0 or higher with the following build options:
-->
LXD は以下のビルドオプションでビルドされた LXC 3.0.0 以上を必要とします。

<!--
 * apparmor (if using LXD's apparmor support)
 * seccomp
-->
 * apparmor (もし LXD の apparmor サポートを使用するのであれば)
 * seccomp

<!--
To run recent version of various distributions, including Ubuntu, LXCFS
should also be installed.
-->
Ubuntu を含む、さまざまなディストリビューションの最近のバージョンを
動かすためには、 LXCFS もインストールする必要があります。

## QEMU
仮想マシンを利用するには QEMU 4.2 以降が望ましいです。
それより古いバージョンは QEMU 2.11 までは動作報告がありますが、古いバージョンのサポートは将来の LXD のリリースで誤ってリグレッションが起きる可能性があります。
<!--
For virtual machines, QEMU 4.2 or higher is preferred.
Older versions, as far back as QEMU 2.11 have been reported to work
properly, but support for those may accidentally regress in future LXD
releases.
-->

## 追加のライブラリー(と開発用のヘッダ) <!-- Additional libraries (and development headers) -->
<!--
LXD uses `dqlite` for its database, to build and setup it, you can
run `make deps`.
-->
LXD はデータベースとして `dqlite` を使用しています。
ビルドしセットアップするためには `make deps` を実行してください。

<!--
LXD itself also uses a number of (usually packaged) C libraries:
-->
LXD は他にもいくつかの (たいていはパッケージ化されている) C ライブラリーを使用しています。

 - libacl1
 - libcap2
 - libuv1 (`dqlite` で使用) <!-- (for `dqlite`) -->
 - libsqlite3 >= 3.25.0 (`dqlite` で使用) <!-- (for `dqlite`) -->

<!--
Make sure you have all these libraries themselves and their development
headers (-dev packages) installed.
-->
ライブラリーそのものとライブラリーの開発用ヘッダ (-dev パッケージ)の全てを
インストールしたことを確認してください。
