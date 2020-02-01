# クラスタリング
<!-- Clustering -->

<!--
LXD can be run in clustering mode, where any number of LXD servers
share the same distributed database and can be managed uniformly using
the lxc client or the REST API.
-->
LXD はクラスタリングモードで実行できます。クラスタリングモードでは複数台の LXD サーバが同じ分散データベースを共有し、REST API や lxc クライアントで統合管理できます。

<!--
Note that this feature was introduced as part of the API extension 
"clustering".
-->
この機能は API 拡張の "clustering" の一部として導入しました。

## クラスターの形成 <!-- Forming a cluster -->

<!--
First you need to choose a bootstrap LXD node. It can be an existing
LXD server or a brand new one. Then you need to initialize the
bootstrap node and join further nodes to the cluster. This can be done
interactively or with a preseed file.
-->
まず、ブートストラップノードを選択する必要があります。既存の LXD サーバでも新しいインスタンスでもブートストラップノードになれます。ブートストラップノードとなるサーバを決めた後は、ブートストラップノードを初期化し、それからクラスターへ追加ノードを参加させます。この処理はインタラクティブに行えますし、前もって定義ファイルを作成しても行えます。

<!--
Note that all further nodes joining the cluster must have identical
configuration to the bootstrap node, in terms of storage pools and
networks. The only configuration that can be node-specific are the
`source` and `size` keys for storage pools and the
`bridge.external_interfaces` key for networks.
-->
クラスターに追加するノードはすべて、ストーレージプールとネットワークについて、ブートストラップノードと同じ構成を持たなければなりません。ノード特有の設定として持てる唯一の設定は、ストレージプールに対する `source` と `size`、ネットワークに対する `bridge.external_interface` です。

<!--
It is strongly recommended that the number of nodes in the cluster be 
at least three, so the cluster can survive the loss of at least one node 
and still be able to establish quorum for its distributed state (which is
kept in a SQLite database replicated using the Raft algorithm). If the 
number of nodes is less than three, then only one node in the cluster
will store the SQLite database. When the third node joins the cluster,
both the second and third nodes will receive a replica of the database.
-->
クラスター内のノード数としては 3 以上を強く推奨します。これは少なくとも 1 ノードが落ちても分散状態のクオラムを確立できるからです（分散状態は Raft アルゴリズムを使ってレプリケーションされている SQLite データベースに保管されています）。ノード数が 3 より小さくなるとクラスタ内のただ 1 つのノードだけが SQLite データベースを保管します。第 3 のノードがクラスターに参加したときに、第 2 と第 3 のノードがデータベースの複製を受け取ります。

### インタラクティブに行う方法 <!-- Interactively -->

<!--
Run `lxd init` and answer `yes` to the very first question ("Would you
like to use LXD clustering?"). Then choose a name for identifying the
node, and an IP or DNS address that other nodes can use to connect to
it, and answer `no` to the question about whether you're joining an
existing cluster. Finally, optionally create a storage pool and a
network bridge. At this point your first cluster node should be up and
available on your network.
-->
`lxd init` を実行し、最初の質問（"Would you like to use LXD clustering?"）に `yes` と答えます。そして、そのノードを特定する名前、他のノードが接続するための IP もしくは DNS アドレスを選択します。そして、既存のクラスターに加わるかどうかの質問には `no` と答えます。最後に、オプショナルでストレージプールとネットワークブリッジを作成できます。これで、最初のクラスターノードが起動し、ネットワークが利用できるようになります。

<!--
You can now join further nodes to the cluster. Note however that these
nodes should be brand new LXD servers, or alternatively you should
clear their contents before joining, since any existing data on them
will be lost.
-->
更に追加のノードをクラスターに追加できます。しかし、追加ノードの既存データはすべて失われるため、追加のノードは完全に新しい LXD サーバであるか、追加前にすべての情報をクリアしたノードである必要があります。

<!--
To add an additional node, run `lxd init` and answer `yes` to the question
about whether to use clustering. Choose a node name that is different from
the one chosen for the bootstrap node or any other nodes you have joined so
far. Then pick an IP or DNS address for the node and answer `yes` to the
question about whether you're joining an existing cluster. Pick an address
of an existing node in the cluster and check the fingerprint that gets
printed.
-->
ノードを追加するために、`lxd init` を実行し、クラスタリングを使うかどうかの質問に `yes` と答えます。ブートストラップノード、それまでに参加したノードとは異なる名前を指定します。IP もしくは DNS アドレスを指定し、既存のクラスターに加わるかどうかの質問には `yes` と答えます。クラスター内の既存のノードのアドレスを指定し、表示されたフィンガープリントを確認します。

### 事前に定義して行う方法 <!-- Preseed -->

<!--
Create a preseed file for the bootstrap node with the configuration
you want, for example:
-->
事前にブートストラップノードの設定内容を書いた定義ファイルを作成できます。例えば:

```yaml
config:
  core.trust_password: sekret
  core.https_address: 10.55.60.171:8443
  images.auto_update_interval: 15
storage_pools:
- name: default
  driver: dir
networks:
- name: lxdbr0
  type: bridge
  config:
    ipv4.address: 192.168.100.14/24
    ipv6.address: none
profiles:
- name: default
  devices:
    root:
      path: /
      pool: default
      type: disk
    eth0:
      name: eth0
      nictype: bridged
      parent: lxdbr0
      type: nic
cluster:
  server_name: node1
  enabled: true
```

<!--
Then run `cat <preseed-file> | lxd init \-\-preseed` and your first node
should be bootstrapped.
-->
定義ファイルを作成したあと、`cat <preseed-file> | lxd init --preseed` を実行し、最初のノードを作成します。

<!--
Now create a bootstrap file for another node. You only need to fill in the
``cluster`` section with data and config values that are specific to the joining
node.
-->
次に、他のノードのブートストラップファイルを作成します。``cluster`` セクションに、追加するノード固有のデータと設定値を指定するだけです。

<!--
Be sure to include the address and certificate of the target bootstrap node. To
create a YAML-compatible entry for the ``cluster_certificate`` key you can use a
command like `sed ':a;N;$!ba;s/\n/\n\n/g' /var/lib/lxd/server.crt`, which you
have to run on the bootstrap node.
-->
ターゲットとなるブートストラップノードのアドレスと証明書を必ず含めてください。``cluster_certificate`` に対する YAML 互換のエントリを作成するには、`sed ':a;N;$!ba;s/\n/\n\n/g' /var/lib/lxd/server.crt` のようにコマンドを実行します。このコマンドはブートストラップノードで実行する必要があります。

<!--
For example:
-->
例えば:

```yaml
cluster:
  enabled: true
  server_name: node2
  server_address: 10.55.60.155:8443
  cluster_address: 10.55.60.171:8443
  cluster_certificate: "-----BEGIN CERTIFICATE-----

opyQ1VRpAg2sV2C4W8irbNqeUsTeZZxhLqp4vNOXXBBrSqUCdPu1JXADV0kavg1l

2sXYoMobyV3K+RaJgsr1OiHjacGiGCQT3YyNGGY/n5zgT/8xI0Dquvja0bNkaf6f

...

-----END CERTIFICATE-----
"
  cluster_password: sekret
  member_config:
  - entity: storage-pool
    name: default
    key: source
    value: ""
```

## クラスターの管理 <!-- Managing a cluster -->

<!--
Once your cluster is formed you can see a list of its nodes and their
status by running `lxc cluster list`. More detailed information about
an individual node is available with `lxc cluster show <node name>`.
-->
クラスタが形成されると、`lxc cluster list` を実行して、ノードのリストと状態を見ることができます。ノードそれぞれのもっと詳細な情報は `lxc cluster show <node name>` を実行して取得できます。


### ノードの削除 <!-- Deleting nodes -->

<!--
To cleanly delete a node from the cluster use `lxc cluster remove <node name>`.
-->
クラスタからノードをクリーンに削除するには、`lxc cluster remove <node name>` を使います。

### オフラインノードとフォールトトレランス <!-- Offline nodes and fault tolerance -->

<!--
At each time there will be an elected cluster leader that will monitor
the health of the other nodes. If a node is down for more than 20
seconds, its status will be marked as OFFLINE and no operation will be
possible on it, as well as operations that require a state change
across all nodes.
-->
都度、選出されたクラスタリーダーが存在し、そのリーダーが他のノードの健全性をモニタリングします。20 秒以上ノードがダウンした場合は、ステータスは OFFLINE とマークされ、そのノード上での操作はできなくなります。また、すべてのノードで状態の変更が必要な操作が可能です。

<!--
If the node that goes offline is the leader itself, the other nodes
will elect a new leader.
-->
リーダーがオフラインに移行した場合、他のノードが新しいリーダーに選出されます。

<!--
As soon as the offline node comes back online, operations will be
available again.
-->
オフラインノードがオンラインに戻るとすぐに、ふたたび操作できるようになります。

<!--
If you can't or don't want to bring the node back online, you can
delete it from the cluster using `lxc cluster remove \-\-force <node name>`.
-->
ノードをオンラインに戻せないとき、ノードをオンラインに戻したくないときは、`lxc cluster remove --force <node name>` を使ってクラスターからノードを削除できます。

### ノードのアップグレード <!-- Upgrading nodes -->

<!--
To upgrade a cluster you need to upgrade all of its nodes, making sure
that they all upgrade to the same version of LXD.
-->
クラスターをアップグレードするためには、すべてのノードをアップグレードし、すべてが確実に同じバージョンの LXD にする必要があります。

<!--
To upgrade a single node, simply upgrade the lxd/lxc binaries on the
host (via snap or other packaging systems) and restart the lxd daemon.
-->
単一のノードをアップグレードするには、単にホスト上で（snap や他のパッケージ管理システムを使って） lxd/lxc バイナリをアップグレードし、lxd デーモンを再起動します。

<!--
If the new version of the daemon has database schema or API changes,
the restarted node might transition into a Blocked state. That happens
if there are still nodes in the cluster that have not been upgraded
and that are running an older version. When a node is in the
Blocked state it will not serve any LXD API requests (in particular,
lxc commands on that node will not work, although any running
instance will continue to run).
-->
デーモンの新バージョンでデータベーススキーマや API が変更になった場合は、再起動したノードは Blocked 状態に移行する可能性があります。これは、クラスタ内にまだアップグレードされていないノードが存在し、その上で古いバージョンが動作している場合に起こります。ノードが Blocked 状態にあるとき、このノードは LXD API リクエストを一切受け付けません（詳しく言うと、実行中のインスタンスは動き続けますが、ノード上の lxc コマンドは動きません）。

<!--
You can see if some nodes are blocked by running `lxc cluster list` on
a node which is not blocked.
-->
ブロックされていないノード上で `lxc cluster list` を実行すると、ノードがブロックされているかどうかを確認できます。

<!--
As you proceed upgrading the rest of the nodes, they will all
transition to the Blocked state, until you upgrade the very last
one. At that point the blocked nodes will notice that there is no
out-of-date node left and will become operational again.
-->
残りのノードのアップグレードを進めると、最後のノードをアップグレードするまでは、ノードはすべて Blocked 状態に移行します。その時点で、Blocked ノードは古いノードが残っていないかを確認し、再度操作できるようになります。

### ディザスターリカバリー <!-- Disaster recovery -->

各 LXD クラスタはデータベースノードとして機能するメンバーを最大 3 つまで持つことができます。
恒久的にデータベースノードとして機能するクラスタメンバーの過半数を失った場合 (例えば 3 メンバーのクラスタで 2 メンバーを失った場合)、
クラスタは利用不可能になります。しかし、 1 つでもデータベースノードが生き残っていれば、クラスタをリカバーすることができます。
<!--
Every LXD cluster has up to 3 members that serve as database nodes. If you
permanently lose a majority of the cluster members that are serving as database
nodes (for example you have a 3-member cluster and you lose 2 members), the
cluster will become unavailable. However, if at least one database node has
survived, you will be able to recover the cluster.
-->

クラスタメンバーがデータベースノードとして設定されているかどうかをチェックするには、クラスタのいずれかの生き残っているメンバーにログオンして以下のコマンドを実行します。
<!--
In order to check which cluster members are configured as database nodes, log on
any survived member of your cluster and run the command:
-->

```
lxd cluster list-database
```

これは LXD デーモンが実行中でなくても実行できます。
<!--
This will work even if the LXD daemon is not running.
-->

一覧表示されたメンバーの中で、生き残っていてログインしたものを選びます (コマンドを実行したメンバーと異なる場合)。
<!--
Among the listed members, pick the one that has survived and log into it (if it
differs from the one you have run the command on).
-->

LXD デーモンが実行していないことを確認したうえで次のコマンドを実行します。
<!--
Now make sure the LXD daemon is not running and then issue the command:
-->

```
lxd cluster recover-from-quorum-loss
```

この時点で LXD デーモンを再起動できるようになり、データベースはオンラインに復帰するはずです。
<!--
At this point you can restart the LXD daemon and the database should be back
online.
-->

データベースからは何の情報も削除されていないことに注意してください。特に失われたクラスタメンバーに関する情報は、それらのインスタンスについてのメタデータも含めて、まだそこに残っています。
この情報は失われたインスタンスを再度作成する必要がある場合に、さらなるリカバーのステップで利用することができます。
<!--
Note that no information has been deleted from the database, in particular all
information about the cluster members that you have lost is still there,
including the metadata about their instances. This can help you with further
recovery steps in case you need to re-create the lost instances.
-->

失われたクラスタメンバーを恒久的に削除するためには、次のコマンドが利用できます。
<!--
In order to permanently delete the cluster members that you have lost, you can
run the command:
-->

```
lxc cluster remove <name> --force
```

ここでは ``lxd``` ではなく通常の ```lxc``` コマンドを使う必要があることに注意してください。
<!--
Note that this time you have to use the regular ```lxc``` command line tool, not
```lxd```.
-->

## インスタンス <!-- Instances -->

<!--
You can launch an instance on any node in the cluster from any node in
the cluster. For example, from node1:
-->
クラスタ上の任意のノード上でインスタンスを起動できます。例えば、node1 から:

```bash
lxc launch --target node2 ubuntu:18.04 bionic
```

<!--
will launch an Ubuntu 18.04 container on node2.
-->
のように実行すれば、node2 上で Ubuntu 18.04 コンテナが起動します。

<!--
When you launch an instance without defining a target, the instance will be 
launched on the server which has the lowest number of instances.
If all the servers have the same amount of instances, it will choose one 
at random.
-->
ターゲットを指定せずにインスタンスを起動したときは、インスタンスの数が一番少ないサーバ上でインスタンスが起動されます。全てのサーバが同じ数のインスタンスを持っている場合はランダムに選ばれます。

<!--
You can list all instances in the cluster with:
-->
以下のように実行すると、インスタンス上のすべてのコンテナをリストできます:

```bash
lxc list
```

<!--
The NODE column will indicate on which node they are running.
-->
NODE 列がコンテナが実行中のノードを示します。

<!--
After an instance is launched, you can operate it from any node. For
example, from node1:
-->
インスタンスが起動後、任意のノードからそのコンテナを操作できます。例えば、node1 から:

```bash
lxc exec bionic ls /
lxc stop bionic
lxc delete bionic
lxc pull file bionic/etc/hosts .
```

のように操作できます。

## イメージ <!-- Images -->

デフォルトではデータベースメンバを持っているのと同じ数のクラスタに
LXD はイメージを複製します。これは通常はクラスタ内で最大3つのコピーを
持つことを意味します。
<!--
By default, LXD will replicate images on as many cluster members as you
have database members. This typically means up to 3 copies within the cluster.
-->

耐障害性とイメージがローカルにある可能性を上げるためにこの数を増やす
ことができます。
<!--
That number can be increased to improve fault tolerance and likelihood
of the image being locally available.
-->

特別な値である "-1" は全てのノードにイメージをコピーするために使用できます。
<!--
The special value of "-1" may be used to have the image copied on all nodes.
-->


この数を 1 に設定することでイメージの複製を無効にできます。
<!--
You can disable the image replication in the cluster by setting the count down to 1:
-->

```bash
lxc config set cluster.images_minimal_replica 1
```

## ストレージプール <!-- Storage pools -->

<!--
As mentioned above, all nodes must have identical storage pools. The
only difference between pools on different nodes might be their
`source`, `size` or `zfs.pool\_name` configuration keys.
-->
先に述べたように、すべてのノードは同一のストレージプールを持たなければなりません。異なるノード上のプール間の違いは、設定項目、`source`、`size`、`zfs.pool\_name` のみです。

<!--
To create a new storage pool, you first have to define it across all
nodes, for example:
-->
新たにストレージプールを作成するためには、すべてのノードでストレージプールをを定義する必要があります。例えば:

```bash
lxc storage create --target node1 data zfs source=/dev/vdb1
lxc storage create --target node2 data zfs source=/dev/vdc1
```
のようにします。

<!--
Note that when defining a new storage pool on a node the only valid
configuration keys you can pass are the node-specific ones mentioned above.
-->
新しいストレージプールをノード上に定義する際、ノード固有で与えることのできる設定項目は上記設定のみです。

<!--
At this point the pool hasn't been actually created yet, but just
defined (it's state is marked as Pending if you run `lxc storage list`).
-->
この時点ではプールはまだ実際には作られていませんが、定義はされています（`lxc storage list` を実行すると、状態が Pending とマークされています）。

<!--
Now run:
-->
次のように実行しましょう:

```bash
lxc storage create data zfs
```

<!--
and the storage will be instantiated on all nodes. If you didn't
define it on a particular node, or a node is down, an error will be
returned.
-->
するとストレージがすべてのノードでインスタンス化されます。特定のノードで定義を行っていない場合、もしくはノードがダウンしている場合は、エラーが返ります。

<!--
You can pass to this final ``storage create`` command any configuration key
which is not node-specific (see above).
-->
この最後の ``storage create`` コマンドには、ノード固有ではない（上記参照）任意の設定項目を与えることができます。

## ストレージボリューム <!-- Storage volumes -->

<!--
Each volume lives on a specific node. The `lxc storage volume list`
includes a `NODE` column to indicate on which node a certain volume
resides.
-->
各ボリュームは特定のノード上に存在しています。`lxc storage volume list` は、特定のボリュームがどのノードにあるかを示す `NODE` 列を表示します。

<!--
Different volumes can have the same name as long as they live on
different nodes (for example image volumes). You can manage storage
volumes in the same way you do in non-clustered deployments, except
that you'll have to pass a `\-\-target <node name>` parameter to volume
commands if more than one node has a volume with the given name.
-->
異なるボリュームは、異なるノード（例えば image volumes）上に存在する限りは同じ名前を持てます。複数のノードが与えた名前のボリュームを持つ場合には、ボリュームコマンドに `--target <node name>` を与える必要がある点を除いて、ストレージボリュームはクラスタ化されていない場合と同じ方法で管理できます。

<!--
For example:
-->
例えば:

```bash
# Create a volume on the node this client is pointing at
lxc storage volume create default web

# Create a volume with the same node on another node
lxc storage volume create default web --target node2

# Show the two volumes defined
lxc storage volume show default web --target node1
lxc storage volume show default web --target node2
```

## ネットワーク <!-- Networks -->

<!--
As mentioned above, all nodes must have identical networks defined. The only
difference between networks on different nodes might be their
`bridge.external_interfaces` optional configuration key (see also documentation
about [network configuration](networks.md)).
-->
先に述べたように、すべてのノードは同じネットワークを定義しなければなりません。異なるノード間のネットワークで異なっても良い設定は、`bridge.external_interfaces` というオプショナルの設定項目です（[ネットワーク設定](networks.md)の文書を参照してください）

<!--
To create a new network, you first have to define it across all
nodes, for example:
-->
新しいネットワークを作成するには、最初にすべてのノードで以下のように定義を行う必要があります:

```bash
lxc network create --target node1 my-network
lxc network create --target node2 my-network
```

<!--
Note that when defining a new network on a node the only valid configuration
key you can pass is `bridge.external_interfaces`, as mentioned above.
-->
ノード上に新しいネットワークを定義する場合、先に述べたように `bridge.external_interfaces` のみ有効な設定として与えることができます。

<!--
At this point the network hasn't been actually created yet, but just
defined (it's state is marked as Pending if you run `lxc network list`).
-->
この時点では、ネットワークはまだ実際には作成されていません。しかし定義はされています（`lxc network list` を実行すると、状態が Pending とマークされています）。

<!--
Now run:
-->
次のように実行しましょう:

```bash
lxc network create my-network
```

<!--
and the network will be instantiated on all nodes. If you didn't
define it on a particular node, or a node is down, an error will be
returned.
-->
するとネットワークがすべてのノード上でインスタンス化されます。特定のノードで定義していない場合、もしくはノードがダウンしている場合は、エラーが返ります。

<!--
You can pass to this final ``network create`` command any configuration key
which is not node-specific (see above).
-->
この最後の ``network create`` コマンドには、ノード固有ではない（上記参照）任意の設定項目を与えることができます。

## 分離した REST API とクラスタネットワーク <!-- Separate REST API and clustering networks -->

クライアントの REST API エンドポイントとクラスタ内のノード間の内部的なトラフィック
（例えば REST API に DNS ラウンドロビンとともに仮想 IP アドレスを使うために）
で別のネットワークを設定できます。
<!--
You can configure different networks for the REST API endpoint of your clients
and for internal traffic between the nodes of your cluster (for example in order
to use a virtual address for your REST API, with DNS round robin).
-->

このためには、クラスタの最初のノードを ```cluster.https_address``` 設定キーを
使ってブートストラップする必要があります。例えば以下の定義ファイルを使うと
<!--
To do that, you need to bootstrap the first node of the cluster using the
```cluster.https_address``` config key. For example, when using preseed:
-->

```yaml
config:
  core.trust_password: sekret
  core.https_address: my.lxd.cluster:8443
  cluster.https_address: 10.55.60.171:8443
...
```

（YAML 定義ファイルの残りは上記と同じ）。
<!--
(the rest of the preseed YAML is the same as above).
-->

新しいノードを参加させるには、まず REST API のアドレスを設定します。
例えば ```lxc``` クライアントを使って以下のように実行し
<!--
To join a new node, first set its REST API address, for instance using the
```lxc``` client:
-->

```bash
lxc config set core.https_address my.lxd.cluster:8443
```

そして通常通り ```PUT /1.0/cluster``` API エンドポイントを使って、
```server_address``` フィールドで参加するノードのアドレスを設定します。
定義ファイルを使うなら YAML のペイロードは完全に上記のものと同じに
なるでしょう。
<!--
and then use the ```PUT /1.0/cluster``` API endpoint as usual, specifying the
address of the joining node with the ```server_address``` field. If you use
preseed, the YAML payload would be exactly like the one above.
-->
