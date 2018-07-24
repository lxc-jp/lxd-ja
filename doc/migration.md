# LXD でのライブマイグレーション <!-- Live Migration in LXD -->

## 概要 <!-- Overview -->

<!--
Migration has two pieces, a "source", that is, the host that already has the
container, and a "sink", the host that's getting the container. Currently,
in the `pull` mode, the source sets up an operation, and the sink connects
to the source and pulls the container.
-->
マイグレーションには 2 つの要素があります。1 つは「ソース」、つまり
既にコンテナを保持しているホストです。もう 1 つは「シンク」、コンテナを
受け取るホストです。現在、 `pull` モードでは、ソースが操作をセットアップし、
シンクがソースに接続してコンテナを pull します。

<!--
There are three websockets (channels) used in migration:

  1. the control stream
  2. the criu images stream
  3. the filesystem stream
-->
マイグレーションでは以下の 3 つの websocket (チャンネル) を使用します。

  1. コントロール・ストリーム
  2. criu イメージ・ストリーム
  3. ファイルシステム・ストリーム

<!--
When a migration is initiated, information about the container, its
configuration, etc. are sent over the control channel (a full
description of this process is below), the criu images and container
filesystem are synced over their respective channels, and the result of
the restore operation is sent from the sink to the source over the
control channel.
-->
マイグレーションが開始されると、コンテナに関する情報、コンテナの設定などが
コントロール・チャンネル上を流れます (このプロセスの完全な説明は後述します)。
criu イメージとコンテナのファイルシステムはそれぞれ個別のチャンネルを
使って同期され、リストア操作の結果はシンクからソースにコントロール・チャンネル
上で送られます。

<!--
In particular, the protocol that is spoken over the criu channel and filesystem
channel can vary, depending on what is negotiated over the control socket. For
example, both the source and the sink's LXD directory is on btrfs, the
filesystem socket can speak btrfs-send/receive. Additionally, although we do a
"stop the world" type migration right now, support for criu's p.haul protocol
will happen over the criu socket at some later time.
-->
特に、 criu チャンネルとファイルシステム・チャンネルの上で話されるプロトコルは
コントロール・ソケット上で交渉されたものによって異なる場合があります。例えば、
ソースとシンクの両方の LXD ディレクトリが btrfs 上にある場合、ファイルシステム・
ソケットは btrfs の send/receive を話せます。さらに、現時点では我々は
「ストップ・ザ・ワールド」タイプのマイグレーションを実行しますが、 criu の
p.haul プロトコルはいつか criu ソケット上で実現されるでしょう。

## コントロール・ソケット <!-- Control Socket -->

<!--
Once all three websockets are connected between the two endpoints, the
source sends a MigrationHeader (protobuf description found in
`/lxd/migration/migrate.proto`). This header contains the container
configuration which will be added to the new container.
-->
2 つのエンドポイント間で 3 つの websocket が全て接続されたら、ソースは
MigrationHeader (protobuf の記述が `/lxd/migration/migrate.proto` にあります)
を送ります。このヘッダはコンテナの設定を含んでおり、それは新しいコンテナに
追加されます。

<!--
There are also two fields indicating the filesystem and criu protocol to speak.
For example, if a server is hosted on a btrfs filesystem, it can indicate that it
wants to do a `btrfs send` instead of a simple rsync (similarly, it could
indicate that it wants to speak the p.haul protocol, instead of just rsyncing
the images over slowly).
-->
話す予定のファイルシステムと criu のプロトコルを示す 2 つのフィールドもあります。
例えば、サーバが btrfs ファイルシステム上にホストされている場合、単純な rsync の
代わりに `btrfs send` を使いたいと示すことができます (同様に単にイメージを rsync
で低速に転送する代わりに p.haul プロトコルを話したいと示すこともできるかもしれません)。

<!--
The sink then examines this message and responds with whatever it
supports. Continuing our example, if the sink is not on a btrfs
filesystem, it responds with the lowest common denominator (rsync, in
this case), and the source is to send the root filesystem using rsync.
Similarly with the criu connection; if the sink doesn't have support for
the p.haul protocol (or whatever), we fall back to rsync.
-->
次にシンクはこのメッセージを調べてシンクがサポートするもので応答します。
上の例を続けると、シンクが btrfs ファイルシステム上にない場合、最小公倍数
(この場合は rsync) で応答し、ソースはルート・ファイルシステムを rsync で
送ることになります。同様に criu コネクションの例でシンクが p.haul プロトコル
(や他の何か) をサポートしない場合は、 rsync にフォールバックします。
