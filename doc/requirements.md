# 動作環境 <!-- Requirements -->
## Go

<!--
LXD requires Go 1.9 or higher.
Both the golang and gccgo compilers are supported.
-->
LXD は Go 1.9 以上を必要とします。
golang と gccgo の両方のコンパイラがサポートされます。

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
LXD requires LXC 2.0.0 or higher with the following build options:
-->
LXD は以下のビルドオプションでビルドされた LXC 2.0.0 以上を必要とします。

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
