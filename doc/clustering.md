# クラスタリング <!-- Clustering -->

<!--
LXD can be run in clustering mode, where any number of LXD instances
share the same distributed database and can be managed uniformly using
the lxc client or the REST API.
-->
LXD はクラスタリングモードで実行できます。クラスタリングモードでは複数台の LXD インスタンスが同じ分散データベースを共有し、REST API や lxc クライアントで統合管理できます。

<!--
Note that this feature was introduced as part of the API extension 
"clustering".
-->
この機能は API 拡張の "clustering" の一部として導入しました。

## クラスターの形成 <!-- Forming a cluster -->

<!--
First you need to choose a bootstrap LXD node. It can be an existing
LXD instance or a brand new one. Then you need to initialize the
bootstrap node and join further nodes to the cluster. This can be done
interactively or with a preseed file.
-->
まず、ブートストラップノードを選択する必要があります。既存の LXD インスタンスでも新しいインスタンスでもブートストラップノードになれます。ブートストラップノードとなるインスタンスを決めた後は、ブートストラップノードを初期化し、それからクラスターへ追加ノードを参加させます。この処理はインタラクティブに行えますし、前もって定義ファイルを作成しても行えます。

<!--
Note that all further nodes joining the cluster must have identical
configuration to the bootstrap node, in terms of storage pools and
networks. The only configuration that can be node-specific are the
`source` and `size` keys for storage pools and the
`bridge.external_interfaces` key for networks.
-->
クラスターに追加するノードはすべて、ストーレージプールとネットワークについて、ブートストラップノードと同じ構成を持たなければなりません。ノード特有の設定として持てる唯一の設定は、ストレージプールに対する `source` と `size`、ネットワークに対する `bridge.external_interface` です。

<!--
It is recommended that the number of nodes in the cluster be at least
three, so the cluster can survive the loss of at least one node and
still be able to establish quorum for its distributed state (which is
kept in a SQLite database replicated using the Raft algorithm).
-->
クラスター内のノード数としては 3 以上を推奨します。これは少なくとも 1 ノードが落ちてもクラスターが生存でき、分散状態で必要な数の生存を確立できるからです（これは Raft アルゴリズムを使った SQLite データベースのレプリケーションを維持できるということです）

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
nodes should be brand new LXD instances, or alternatively you should
clear their contents before joining, since any existing data on them
will be lost.
-->
更に追加のノードをクラスターに追加できます。しかし、追加ノードの既存データはすべて失われるため、追加のノードは完全に新しい LXD インスタンスであるか、追加前にすべての情報をクリアしたノードである必要があります。

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
Then run `cat <preseed-file> | lxd init --preseed` and your first node
should be bootstrapped.
-->
定義ファイルを作成したあと、`cat <preseed-file> | lxd init --preseed` を実行し、最初のノードを作成します。

<!--
Now create a bootstrap file for another node. Be sure to specify the
address and certificate of the target bootstrap node. To create a
YAML-compatible entry for the `<cert>` key you can use a command like
`sed ':a;N;$!ba;s/\n/\n\n/g' /var/lib/lxd/server.crt`, which you have to
run on the bootstrap node.
-->
次に、他のノードのブートストラップファイルを作成します。追加するクラスターのブートストラップノードのアドレスと証明書を必ず指定してください。`<cert>` キーと YAML 互換であるエントリを作成するには、`sed ':a;N;$!ba;s/\n/\n\n/g' /var/lib/lxd/server.crt` のようにコマンドを実行します。このコマンドはブートストラップノード上で実行します。

<!--
For example:
-->
例えば:

```yaml
config:
  core.https_address: 10.55.60.155:8443
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
  server_name: node2
  enabled: true
  cluster_address: 10.55.60.171:8443
  cluster_certificate: "-----BEGIN CERTIFICATE-----

opyQ1VRpAg2sV2C4W8irbNqeUsTeZZxhLqp4vNOXXBBrSqUCdPu1JXADV0kavg1l

2sXYoMobyV3K+RaJgsr1OiHjacGiGCQT3YyNGGY/n5zgT/8xI0Dquvja0bNkaf6f

...

-----END CERTIFICATE-----
"
  cluster_password: sekret
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
delete it from the cluster using `lxc cluster remove --force <node name>`.
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
container will continue to run).
-->
デーモンの新バージョンでデータベーススキーマや API が変更になった場合は、再起動したノードは Blocked 状態に移行する可能性があります。これは、クラスタ内にまだアップグレードされていないノードが存在し、その上で古いバージョンが動作している場合に起こります。ノードが Blocked 状態にあるとき、このノードは LXD API リクエストを一切受け付けません（詳しく言うと、実行中のコンテナは動き続けますが、ノード上の lxc コマンドは動きません）。

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


## コンテナ <!-- Containers -->

<!--
You can launch a container on any node in the cluster from any node in
the cluster. For example, from node1:
-->
クラスタ上の任意のノード上でコンテナを起動できます。例えば、node1 から:

```bash
lxc launch --target node2 ubuntu:16.04 xenial
```

<!--
will launch an Ubuntu 16.04 container on node2.
-->
のように実行すれば、node2 上で Ubuntu 16.04 コンテナが起動します。

<!--
You can list all containers in the cluster with:
-->
以下のように実行すると、クラスタ上のすべてのコンテナをリストできます:

```bash
lxc list
```

<!--
The NODE column will indicate on which node they are running.
-->
NODE 列がコンテナが実行中のノードを示します。

<!--
After a container is launched, you can operate it from any node. For
example, from node1:
-->
コンテナが起動後、任意のノードからそのコンテナを操作できます。例えば、node1 から:

```bash
lxc exec xenial ls /
lxc stop xenial
lxc delete xenial
lxc pull file xenial/etc/hosts .
```

のように操作できます。

## ストレージプール <!-- Storage pools -->

<!--
As mentioned above, all nodes must have identical storage pools. The
only difference between pools on different nodes might be their
`source`, `size` or `zfs.pool_name` configuration keys.
-->
先に述べたように、すべてのノードは同一のストレージプールを持たなければなりません。異なるノード上のプール間の違いは、設定項目、`source`、`size`、`zfs.pool_name` のみです。

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
that you'll have to pass a `--target <node name>` parameter to volume
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
