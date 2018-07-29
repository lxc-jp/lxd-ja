# イントロダクション <!-- Introduction -->
<!--
LXD runs safe containers. This is achieved mostly through the use of
user namespaces which make it possible to run containers unprivileged,
greatly limiting the attack surface.
-->
LXD は安全なコンテナを実行します。これは主にユーザー・ネームスペースの使用
によって実現されています。ユーザー・ネームスペースはコンテナを非特権で実行
することを可能にし、攻撃対象を大幅に限定します。

<!--
User namespaces work by mapping a set of uids and gids on the host to a
set of uids and gids in the container.
-->
ユーザー・ネームスペースはコンテナの uid と gid の組をホストの uid と
gid の組にマッピングすることで機能します。


<!--
For example, we can define that the host uids and gids from 100000 to
165535 may be used by LXD and should be mapped to uid/gid 0 through
65535 in the container.
-->
例えば、 100000 から 165535 までのホストの uid と gid を LXD が使用できる
ようにし、コンテナで 0 から 65535 までの uid/gid にマッッピングするように
設定できます。

<!--
As a result a process running as uid 0 in the container will actually be
running as uid 100000.
-->
この結果、コンテナで 0 の uid で動くプロセスが実際には uid 100000 で動く
ことになります。

<!--
Allocations should always be of at least 65536 uids and gids to cover
the POSIX range including root (0) and nobody (65534).
-->
root (0) と nobody (65534) の POSIX の範囲をカバーするため、割当は必ず
最低 65535 個の uid と gid であるべきです。

# カーネルのサポート <!-- Kernel support -->
<!--
User namespaces require a kernel >= 3.12, LXD will start even on older
kernels but will refuse to start containers.
-->
ユーザー・ネームスペースの使用にはカーネル 3.12 以上が必要です。 LXD は
古いカーネルでも起動しますが、コンテナを起動するのは拒否します。

# 使用可能な範囲 <!-- Allowed ranges -->
<!--
On most hosts, LXD will check `/etc/subuid` and `/etc/subgid` for
allocations for the "lxd" user and on first start, set the default
profile to use the first 65536 uids and gids from that range.
-->
ほとんどのホストでは、 LXD は初回起動時に "lxd" ユーザの割当のために
`/etc/subuid` と `/etc/subgid` をチェックし、そこで指定されている範囲の
最初の 65536 個の uid と gid をデフォルト・プロファイルで使用するように
設定します。

<!--
If the range is shorter than 65536 (which includes no range at all),
then LXD will fail to create or start any container until this is corrected.
-->
範囲が 65536 より小さい場合 (範囲が全く無い場合を含む)、これが修正される
まで LXD はコンテナの作成と起動に失敗します。

<!--
If some but not all of `/etc/subuid`, `/etc/subgid`, `newuidmap` (path lookup)
and `newgidmap` (path lookup) can be found on the system, LXD will fail
the startup of any container until this is corrected as this shows a
broken shadow setup.
-->
`/etc/subuid` 、 `/etc/subgid` 、 `newuidmap` (パスを検索)、 `newgidmap`
(パスを検索) のいくつか (ただし全部ではない) がシステムに存在する場合、
これはシャドーの設定が間違っていることを示しているので、これが修正されるまで
LXD はコンテナの起動に失敗します。

<!--
If none of those 4 files can be found, then LXD will assume it's running
on a host using an old version of shadow. In this mode, LXD will assume
it can use any uids and gids above 65535 and will take the first 65536
as its default map.
-->
これらの 4 つのファイルが 1 つも無い場合、 LXD はホストが古いバージョンの
シャドーの上で稼働していると想定します。このモードでは LXD は 65535 より
上のどんな uid と gid も使用可能と想定し、最初の 65536 個をデフォルトの
マッピングに使用します。

# ホスト間で異なる範囲の使用 <!-- Varying ranges between hosts -->
<!--
The source map is sent when moving containers between hosts so that they
can be remapped on the receiving host.
-->
ホスト間でコンテナを移動する時、送信側のマッピングが送られるので、受信側の
ホストで異なる範囲にマッピング可能です。

# コンテナ毎に異なる ID マッピング <!-- Different idmaps per container -->
<!--
LXD supports using different idmaps per container, to further isolate
containers from each other. This is controlled with two per-container
configuration keys, `security.idmap.isolated` and `security.idmap.size`.
-->
コンテナを他のコンテナからより一層隔離するために、 LXD はコンテナ毎に
異なる ID マッピングを使用することをサポートしています。これはコンテナ毎に
`security.idmap.isolated` と `security.idmap.size` という 2 つの設定項目で
制御できます。

<!--
Containers with `security.idmap.isolated` will have a unique id range computed
for them among the other containers with `security.idmap.isolated` set (if none
is available, setting this key will simply fail).
-->
`security.idmap.isolated` が設定されたコンテナは
`security.idmap.isolated` が設定された他のコンテナと衝突しないユニークな
ID の範囲を持つように設定されます (もしそのようなコンテナが 1 つも存在しない場合、
このキーを設定しようとしても失敗します)。

<!--
Containers with `security.idmap.size` set will have their id range set to this
size. Isolated containers without this property set default to a id range of
size 65536; this allows for POSIX compliance and a "nobody" user inside the
container.
-->
`security.idmap.size` が設定されたコンテナはこのサイズに ID の範囲が設定
されます。このプロパティが設定されていない隔離されたコンテナは ID の範囲が
デフォルトのサイズ 65536 に設定されます。これにより POSIX に準拠し、コンテナ内で
"nobody" ユーザが使用できます。

<!--
To select a specific map, the `security.idmap.base` key will let you
override the auto-detection mechanism and tell LXD what host uid/gid you
want to use as the base for the container.
-->
特定のマッピングを選択するには `security.idmap.base` を設定すると
自動検出機構をオーバーライドし、コンテナでベースとして使用したい
ホストの uid/gid を LXD に伝えることができます。

<!--
These properties require a container reboot to take effect.
-->
これらのプロパティを反映するにはコンテナの再起動が必要です。

# カスタムの ID マッピング <!-- Custom idmaps -->
<!--
LXD also supports customizing bits of the idmap, e.g. to allow users to bind
mount parts of the host's filesystem into a container without the need for any
uid-shifting filesystem. The per-container configuration key for this is
`raw.idmap`, and looks like:
-->
さらに LXD は ID マッピングの一部をカスタマイズすることをサポートします。例えば、
uid を変更するファイルシステムを必要とせずに、ホストのファイルシステムの一部を
コンテナに bind mount することをユーザに許可できます。このためのコンテナ毎の
設定項目は `raw.idmap` で、設定例は以下のようになります。

    both 1000 1000
    uid 50-60 500-510
    gid 100000-110000 10000-20000

<!--
The first line configures both the uid and gid 1000 on the host to map to uid
1000 inside the container (this can be used for example to bind mount a user's
home directory into a container).
-->
1 行目は、ホストの uid と gid 1000 の両方をコンテナ内の uid 1000 にマッピング
する設定です (これは例えばユーザのホームディレクトリをコンテナ内に bind mount
するのに使用できます)。

<!--
The second and third lines map only the uid or gid ranges into the container,
respectively. The second entry per line is the source id, i.e. the id on the
host, and the third entry is the range inside the container. These ranges must
be the same size.
-->
2 行目と 3 行目は uid または gid のどちらかだけをコンテナ内にマッピングする設定
です。行の中の 2 番目のエントリはソース ID 、 つまりホスト上の ID で、 3 番目の
エントリはコンテナ内部での範囲です。これらの範囲は同じサイズでなければなりません。

<!--
This property requires a container reboot to take effect.
-->
このプロパティを反映するにはコンテナの再起動が必要です。
