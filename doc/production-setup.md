# プロダクション環境のセットアップ
<!-- Production setup -->
## イントロダクション <!-- Introduction -->
あなたは [LXD live online](https://linuxcontainers.org/lxd/try-it/) か、
なんらかのサーバで LXD を試してみました。結果に満足して、今度は LXD で
本格的な作業を試してみたいと思います。
<!--
So you've made it past trying out [LXD live online](https://linuxcontainers.org/lxd/try-it/),
or on a server scavenged from random parts. You like what you see,
and now you want to try doing some serious work with LXD.
-->

大多数の Linux ディストリビューションでは大量のコンテナーを稼働させるのに最適化されたカーネルの設定はされていません。
このドキュメントでの指示はコンテナーを稼働させる際にひっかかりがちな制限値のほとんどとお勧めの変更後の値をカバーしています。
<!--
The vast majority of Linux distributions do not come with optimized
kernel settings suitable for the operation of a large number of
containers. The instructions in this document cover the most common
limits that you're likely to hit when running containers and suggested
updated values.
-->


### よく遭遇するエラー <!-- Common errors that may be encountered -->

`Failed to allocate directory watch: Too many open files`

`<Error> <Error>: Too many open files`

`failed to open stream: Too many open files in...`

`neighbour: ndisc_cache: neighbor table overflow!`

## サーバの変更 <!-- Server Changes -->
### /etc/security/limits.conf

ドメイン <!-- Domain -->  | 種別 <!-- Type -->  | 項目 <!-- Item -->    | 値 <!-- Value -->    | デフォルト <!-- Default -->  | 説明 <!-- Description -->
:-----  | :---  | :----   | :-------- | :-------- | :----------
\*      | soft  | nofile  | 1048576   | 未設定 <!-- unset -->     | オープンするファイルの最大数 <!-- maximum number of open files -->
\*      | hard  | nofile  | 1048576   | 未設定 <!-- unset -->     | オープンするファイルの最大数 <!-- maximum number of open files -->
root    | soft  | nofile  | 1048576   | 未設定 <!-- unset -->     | オープンするファイルの最大数 <!-- maximum number of open files -->
root    | hard  | nofile  | 1048576   | 未設定 <!-- unset -->     | オープンするファイルの最大数 <!-- maximum number of open files -->
\*      | soft  | memlock | unlimited | 未設定 <!-- unset -->     | ロックされたメモリ内の最大のアドレス空間 (KB) <!-- maximum locked-in-memory address space (KB) -->
\*      | hard  | memlock | unlimited | 未設定 <!-- unset -->     | ロックされたメモリ内の最大のアドレス空間 (KB) <!-- maximum locked-in-memory address space (KB) -->
root    | soft  | memlock | unlimited | 未設定 <!-- unset -->     | ロックされたメモリ内の最大のアドレス空間 (KB) (`bpf` システムコール監視にのみ必要) <!-- maximum locked-in-memory address space (KB) (Only need with `bpf` syscall supervision) -->
root    | hard  | memlock | unlimited | 未設定 <!-- unset -->     | ロックされたメモリ内の最大のアドレス空間 (KB) (`bpf` システムコール監視にのみ必要) <!-- maximum locked-in-memory address space (KB) (Only need with `bpf` syscall supervision) -->


### /etc/sysctl.conf

パラメータ <!-- Parameter -->       | 値 <!-- Value -->     | デフォルト <!-- Default --> | 説明 <!-- Description -->
:-----                              | :---       | :---      | :---
fs.inotify.max\_queued\_events      | 1048576    | 16384     | これは対応する inotify のインスタンスにキューイングされるイベント数の上限を指定します。 <!-- This specifies an upper limit on the number of events that can be queued to the corresponding inotify instance. --> [1]
fs.inotify.max\_user\_instances     | 1048576    | 128       | これは実ユーザー ID ごとに作成可能な inotify のインスタンス数の上限を指定します。 <!-- This specifies an upper limit on the number of inotify instances that can be created per real user ID. --> [1]
fs.inotify.max\_user\_watches       | 1048576    | 8192      | これは実ユーザー ID ごとに作成可能な watch 数の上限を指定します。 <!-- This specifies an upper limit on the number of watches that can be created per real user ID. --> [1]
vm.max\_map\_count                  | 262144     | 65530     | このファイルはプロセスが持つメモリマップ領域の最大数を含みます。malloc の呼び出しの副作用として、 直接的にはmmap と mprotect によって、また、共有ライブラリーをロードすることによって、メモリマップ領域を使います。  <!-- This file contains the maximum number of memory map areas a process may have. Memory map areas are used as a side-effect of calling malloc, directly by mmap and mprotect, and also when loading shared libraries. -->
kernel.dmesg\_restrict              | 1          | 0         | この設定を有効にするとコンテナーがカーネルのリングバッファー内のメッセージにアクセスするのを拒否します。この設定はホスト・システム上の非 root ユーザーへのアクセスも拒否することに注意してください。 <!-- This denies container access to the messages in the kernel ring buffer. Please note that this also will deny access to non-root users on the host system. -->
net.ipv4.neigh.default.gc\_thresh3  | 8192       | 1024      | これは ARP テーブル (IPv4) 内のエントリーの最大数です。1024 個を超えるコンテナーを作成するなら増やすべきです。増やさなければ ARP テーブルがフルになったときに `neighbour: ndisc_cache: neighbor table overflow!` というエラーが発生し、コンテナーがネットワーク設定を取得できなくなります。 [2] <!-- This is the maximum number of entries in ARP table (IPv4). You should increase this if you create over 1024 containers. Otherwise, you will get the error `neighbour: ndisc_cache: neighbor table overflow!` when the ARP table gets full and those containers will not be able to get a network configuration. [2] -->
net.ipv6.neigh.default.gc\_thresh3  | 8192       | 1024      | これは ARP テーブル (IPv6) 内のエントリーの最大数です。1024 個を超えるコンテナーを作成するなら増やすべきです。増やさなければ ARP テーブルがフルになったときに `neighbour: ndisc_cache: neighbor table overflow!` というエラーが発生し、コンテナーがネットワーク設定を取得できなくなります。 [2] <!-- This is the maximum number of entries in ARP table (IPv6). You should increase this if you plan to create over 1024 containers. Otherwise, you will get the error `neighbour: ndisc_cache: neighbor table overflow!` when the ARP table gets full and those containers will not be able to get a network configuration. [2] -->
net.core.bpf\_jit\_limit            | 3000000000 | 264241152 | eBPF JIT アロケーションのサイズの上限値で、通常は PAGE_SIZE * 40000 に設定されます。カーネルが `CONFIG_BPF_JIT_ALWAYS_ON=y` の設定でコンパイルされている場合は `/proc/sys/net/core/bpf_jit_enable` が `1` に設定され変更できません。そのようなカーネルでは eBPF JIT コンパイラーは `seccomp` のような bpf のプログラムを JIT コンパイルする際の失敗を他のカーネルでは続行可能なエラーとして扱う場合でも致命的なエラーとして扱います。そのようなカーネルでは eBPF の JIT コンパイルされたプログラムのサイズの上限値は大幅に増やす必要があります。 <!-- This is a limit on the size of eBPF JIT allocations which is usually set to PAGE_SIZE * 40000. When your kernel is compiled with `CONFIG_BPF_JIT_ALWAYS_ON=y` then `/proc/sys/net/core/bpf_jit_enable` is set to `1` and can't be changed. On such kernels the eBPF JIT compiler will treat failure to JIT compile a bpf program such as a `seccomp` filter as fatal when it would continue on another kernel. On such kernels the limit for eBPF jitted programs needs to be increased siginficantly. -->
kernel.keys.maxkeys                 | 2000       | 200       | 非 root ユーザーが使用できるキーの最大数で、コンテナー数より大きくなければなりません <!-- This is the maximum number of keys a non-root user can use, should be higher than the number of containers -->
kernel.keys.maxbytes                | 2000000    | 20000     | 非 root ユーザーが使用できる keyring の最大サイズ <!-- This is the maximum size of the keyring non-root users can use -->
fs.aio-max-nr                       | 524288     | 65536     | これは並行に実行される非同期 I/O 操作の最大数です。 AIO サブシステムを使うワークロードが大量にある場合（例: MySQL ）、これを増やす必要があるかもしれません。 <!-- This is the maximum number of concurrent async I/O operations. You might need to increase it further if you have a lot of workloads that use the AIO subsystem (e.g. MySQL) -->

設定後、サーバの再起動が必要です。
<!--
Then, reboot the server.
-->

[1]: http://man7.org/linux/man-pages/man7/inotify.7.html
[2]: https://www.kernel.org/doc/Documentation/networking/ip-sysctl.txt

### ネットワーク帯域の調整 <!-- Network Bandwidth Tweaking -->
大量の (コンテナー・コンテナー間、あるいはホスト・コンテナー間の) ローカル・アクティビティを持つ
LXD ホスト上に 1GbE 以上の NIC をお持ちか、 LXD ホストに 1GbE 以上のインターネット接続を
お持ちでしたら、 txqueuelen を調整する価値があります。これらの設定は 10GbE NIC ではさらに
よく機能します。
<!--
If you have at least 1GbE NIC on your lxd host with a lot of local
activity (container - container connections, or host - container
connections), or you have 1GbE or better internet connection on your lxd
host it worth play with txqueuelen. These settings work even better with
10GbE NIC.
-->

#### サーバの変更 <!-- Server Changes -->

##### txqueuelen 
(あなたにとっての最適な値はわかりませんが) あなたの実 NIC の `txqueuelen` を
10000 に変える必要がある場合、 lxdbr0 インタフェースの `txqueuelen` も 10000 に
変更してください。
<!--
You need to change `txqueuelen` of your real NIC to 10000 (not sure
about the best possible value for you), and change and change lxdbr0
interface `txqueuelen` to 10000.  
-->

Debian ベースのディストリビューションでは `/etc/network/interfaces` 内で
`txqueuelen` を恒久的に変更できます。
例えば `up ip link set eth0 txqueuelen 10000` という設定を加えることで
起動時にインタフェースの txqueuelen の値を設定できます。
<!--
In Debian-based distros you can change `txqueuelen` permanently in `/etc/network/interfaces`  
You can add for ex.: `up ip link set eth0 txqueuelen 10000` to your interface configuration to set txqueuelen value on boot.  
You could set it txqueuelen temporary (for test purpose) with `ifconfig <interface> txqueuelen 10000`
-->

##### /etc/sysctl.conf

`net.core.netdev_max_backlog` の値も増やす必要があります。
`/etc/sysctl.conf` に `net.core.netdev_max_backlog = 182757` という設定を加えれば
(再起動後に) 恒久的に設定できます。
(テストの目的で) `netdev_max_backlog` を一時的に設定するには
`echo 182757 > /proc/sys/net/core/netdev_max_backlog` と実行します。
注意: この値が大きすぎると思うかもしれません。多くの人は
`netdev_max_backlog` = `net.ipv4.tcp_mem` の最小値と設定することを好んでいます。
例えば私は `net.ipv4.tcp_mem = 182757 243679 365514` という値を使用しています。
<!--
You also need to increase `net.core.netdev_max_backlog` value.  
You can add `net.core.netdev_max_backlog = 182757` to `/etc/sysctl.conf` to set it permanently (after reboot)
You set `netdev_max_backlog` temporary (for test purpose) with `echo 182757 > /proc/sys/net/core/netdev_max_backlog`
Note: You can find this value too high, most people prefer set `netdev_max_backlog` = `net.ipv4.tcp_mem` min. value.
For example I use this values `net.ipv4.tcp_mem = 182757 243679 365514`
-->

#### コンテナーの変更 <!-- Containers changes -->

コンテナー内のイーサネット・インタフェース全ての txqueuelen の値を変更する必要も
あります。
Debian ベースのディストリビューションでは `/etc/network/interfaces` 内で恒久的に
txqueuelen を変更できます。
例えば `up ip link set eth0 txqueuelen 10000` という設定を加えることで
起動時にインタフェースの txqueuelen の値を設定できます。
<!--
You also need to change txqueuelen value for all you ethernet interfaces in containers.  
In Debian-based distros you can change txqueuelen permanently in `/etc/network/interfaces`  
You can add for ex.: `up ip link set eth0 txqueuelen 10000` to your interface configuration to set txqueuelen value on boot.
-->

#### この変更についての注意 <!-- Notes regarding this change -->

10000 という txqueuelen の値は 10GbE NIC ではよく使われます。基本的には、 小さな
txqueuelen の値は高レイテンシで低速なデバイスと低レイテンシで高速なデバイスで
使われます。個人的にはこれらの設定で (ホスト・コンテナー間、コンテナー・コンテナー間の)
ローカル通信とインターネット接続が 3〜5% 改善しています。
txqueuelen の値の調整の良いところは、使用するコンテナー数が増えれば増えるほど、この
調整の恩恵を受けられることです。そして、この値はいつでも一時的に変更することができ、
あなたの環境で LXD ホストの再起動無しに変更の結果を確認することができます。
<!--
10000 txqueuelen value commonly used with 10GbE NICs. Basically small
txqueuelen values used with slow devices with a high latency, and higher
with devices with low latency. I personally have like 3-5% improvement
with these settings for local (host with container, container vs
container) and internet connections. Good thing about txqueuelen value
tweak, the more containers you use, the more you can be can benefit from
this tweak. And you can always temporary set this values and check this
tweak in your environment without lxd host reboot.
-->
