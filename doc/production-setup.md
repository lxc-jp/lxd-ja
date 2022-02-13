# プロダクション環境のセットアップ
あなたは [LXD live online](https://linuxcontainers.org/lxd/try-it/) か、
なんらかのサーバで LXD を試してみました。結果に満足して、今度は LXD で
本格的な作業を試してみたいと思います。

大多数の Linux ディストリビューションでは大量のコンテナーを稼働させるのに最適化されたカーネルの設定はされていません。
このドキュメントでの指示はコンテナーを稼働させる際にひっかかりがちな制限値のほとんどとお勧めの変更後の値をカバーしています。


### よく遭遇するエラー

`Failed to allocate directory watch: Too many open files`

`<Error> <Error>: Too many open files`

`failed to open stream: Too many open files in...`

`neighbour: ndisc_cache: neighbor table overflow!`

## サーバの変更
### /etc/security/limits.conf

ドメイン | 種別 | 項目    | 値        | デフォルト | 説明
:-----   | :--- | :----   | :-------- | :--------  | :----------
\*       | soft | nofile  | 1048576   | 未設定     | オープンするファイルの最大数
\*       | hard | nofile  | 1048576   | 未設定     | オープンするファイルの最大数
root     | soft | nofile  | 1048576   | 未設定     | オープンするファイルの最大数
root     | hard | nofile  | 1048576   | 未設定     | オープンするファイルの最大数
\*       | soft | memlock | unlimited | 未設定     | ロックされたメモリ内の最大のアドレス空間 (KB)
\*       | hard | memlock | unlimited | 未設定     | ロックされたメモリ内の最大のアドレス空間 (KB)
root     | soft | memlock | unlimited | 未設定     | ロックされたメモリ内の最大のアドレス空間 (KB) (`bpf` システムコール監視にのみ必要)
root     | hard | memlock | unlimited | 未設定     | ロックされたメモリ内の最大のアドレス空間 (KB) (`bpf` システムコール監視にのみ必要)


注意: snap ユーザーの場合はこれらの制限は snap/LXD によって自動的に上げられます。

### /etc/sysctl.conf

パラメータ                         | 値         | デフォルト | 説明
:-----                             | :---       | :---       | :---
fs.aio-max-nr                      | 524288     | 65536      | これは並行に実行される非同期 I/O 操作の最大数です。 AIO サブシステムを使うワークロードが大量にある場合（例: MySQL ）、これを増やす必要があるかもしれません。
fs.inotify.max\_queued\_events     | 1048576    | 16384      | これは対応する inotify のインスタンスにキューイングされるイベント数の上限を指定します。 [1]
fs.inotify.max\_user\_instances    | 1048576    | 128        | これは実ユーザー ID ごとに作成可能な inotify のインスタンス数の上限を指定します。 [1]
fs.inotify.max\_user\_watches      | 1048576    | 8192       | これは実ユーザー ID ごとに作成可能な watch 数の上限を指定します。 [1]
kernel.dmesg\_restrict             | 1          | 0          | この設定を有効にするとコンテナーがカーネルのリングバッファー内のメッセージにアクセスするのを拒否します。この設定はホスト・システム上の非 root ユーザーへのアクセスも拒否することに注意してください。
kernel.keys.maxbytes               | 2000000    | 20000      | 非 root ユーザーが使用できる keyring の最大サイズ
kernel.keys.maxkeys                | 2000       | 200        | 非 root ユーザーが使用できるキーの最大数で、コンテナー数より大きくなければなりません
net.core.bpf\_jit\_limit           | 3000000000 | 264241152  | eBPF JIT アロケーションのサイズの上限値で、通常は PAGE_SIZE * 40000 に設定されます。カーネルが `CONFIG_BPF_JIT_ALWAYS_ON=y` の設定でコンパイルされている場合は `/proc/sys/net/core/bpf_jit_enable` が `1` に設定され変更できません。そのようなカーネルでは eBPF JIT コンパイラーは `seccomp` のような bpf のプログラムを JIT コンパイルする際の失敗を他のカーネルでは続行可能なエラーとして扱う場合でも致命的なエラーとして扱います。そのようなカーネルでは eBPF の JIT コンパイルされたプログラムのサイズの上限値は大幅に増やす必要があります。
net.ipv4.neigh.default.gc\_thresh3 | 8192       | 1024       | これは ARP テーブル (IPv4) 内のエントリーの最大数です。1024 個を超えるコンテナーを作成するなら増やすべきです。増やさなければ ARP テーブルがフルになったときに `neighbour: ndisc_cache: neighbor table overflow!` というエラーが発生し、コンテナーがネットワーク設定を取得できなくなります。 [2]
net.ipv6.neigh.default.gc\_thresh3 | 8192       | 1024       | これは ARP テーブル (IPv6) 内のエントリーの最大数です。1024 個を超えるコンテナーを作成するなら増やすべきです。増やさなければ ARP テーブルがフルになったときに `neighbour: ndisc_cache: neighbor table overflow!` というエラーが発生し、コンテナーがネットワーク設定を取得できなくなります。 [2]
vm.max\_map\_count                 | 262144     | 65530      | このファイルはプロセスが持つメモリマップ領域の最大数を含みます。malloc の呼び出しの副作用として、 直接的にはmmap と mprotect によって、また、共有ライブラリーをロードすることによって、メモリマップ領域を使います。

設定後、サーバの再起動が必要です。

[1]: http://man7.org/linux/man-pages/man7/inotify.7.html
[2]: https://www.kernel.org/doc/Documentation/networking/ip-sysctl.txt

### コンテナー名の漏洩防止
/sys/kernel/slab と /proc/sched\_debug はともにシステム上の全ての cgroup の一覧を表示し、拡張を使えばコンテナー一覧を表示するのを容易にします。

一覧が見られるのを防ぐためには、コンテナーを開始する前に以下のコマンドを忘れずに実行してください。

 - chmod 400 /proc/sched\_debug
 - chmod 700 /sys/kernel/slab/

### ネットワーク帯域の調整
大量の (コンテナー・コンテナー間、あるいはホスト・コンテナー間の) ローカル・アクティビティを持つ
LXD ホスト上に 1GbE 以上の NIC をお持ちか、 LXD ホストに 1GbE 以上のインターネット接続を
お持ちでしたら、 txqueuelen を調整する価値があります。これらの設定は 10GbE NIC ではさらに
よく機能します。

#### サーバの変更

##### txqueuelen
(あなたにとっての最適な値はわかりませんが) あなたの実 NIC の `txqueuelen` を
10000 に変える必要がある場合、 lxdbr0 インタフェースの `txqueuelen` も 10000 に
変更してください。

Debian ベースのディストリビューションでは `/etc/network/interfaces` 内で
`txqueuelen` を恒久的に変更できます。
例えば `up ip link set eth0 txqueuelen 10000` という設定を加えることで
起動時にインタフェースの txqueuelen の値を設定できます。

##### /etc/sysctl.conf

`net.core.netdev_max_backlog` の値も増やす必要があります。
`/etc/sysctl.conf` に `net.core.netdev_max_backlog = 182757` という設定を加えれば
(再起動後に) 恒久的に設定できます。
(テストの目的で) `netdev_max_backlog` を一時的に設定するには
`echo 182757 > /proc/sys/net/core/netdev_max_backlog` と実行します。
注意: この値が大きすぎると思うかもしれません。多くの人は
`netdev_max_backlog` = `net.ipv4.tcp_mem` の最小値と設定することを好んでいます。
例えば私は `net.ipv4.tcp_mem = 182757 243679 365514` という値を使用しています。

#### コンテナーの変更

コンテナー内のイーサネット・インタフェース全ての txqueuelen の値を変更する必要も
あります。
Debian ベースのディストリビューションでは `/etc/network/interfaces` 内で恒久的に
txqueuelen を変更できます。
例えば `up ip link set eth0 txqueuelen 10000` という設定を加えることで
起動時にインタフェースの txqueuelen の値を設定できます。

#### この変更についての注意

10000 という txqueuelen の値は 10GbE NIC ではよく使われます。基本的には、 小さな
txqueuelen の値は高レイテンシで低速なデバイスと低レイテンシで高速なデバイスで
使われます。個人的にはこれらの設定で (ホスト・コンテナー間、コンテナー・コンテナー間の)
ローカル通信とインターネット接続が 3〜5% 改善しています。
txqueuelen の値の調整の良いところは、使用するコンテナー数が増えれば増えるほど、この
調整の恩恵を受けられることです。そして、この値はいつでも一時的に変更することができ、
あなたの環境で LXD ホストの再起動無しに変更の結果を確認することができます。
