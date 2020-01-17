# コンテナ
<!-- Containers -->

## はじめに <!-- Introduction -->
コンテナは LXD のデフォルトタイプであり、現時点では一番機能を持っており、LXD インスタンスの完全な実装です。
<!--
Containers are the default type for LXD and currently the most
featureful and complete implementation of LXD instances.
-->

<!--
They are implemented through the use of `liblxc` (LXC).
-->
これは `liblxc` (LXC) を使って実装しています。

## 設定 <!-- Configuration -->
<!--
See [instance configuration](instances.md) for valid configuration options.
-->
設定オプションについては [インスタンスの設定](instances.md) をご覧ください。

## ライブマイグレーション <!-- Live migration -->
<!--
LXD supports live migration of containers using [CRIU](http://criu.org). In
order to optimize the memory transfer for a container LXD can be instructed to
make use of CRIU's pre-copy features by setting the
`migration.incremental.memory` property to `true`. This means LXD will request
CRIU to perform a series of memory dumps for the container. After each dump LXD
will send the memory dump to the specified remote. In an ideal scenario each
memory dump will decrease the delta to the previous memory dump thereby
increasing the percentage of memory that is already synced. When the percentage
of synced memory is equal to or greater than the threshold specified via
`migration.incremental.memory.goal` LXD will request CRIU to perform a final
memory dump and transfer it. If the threshold is not reached after the maximum
number of allowed iterations specified via
`migration.incremental.memory.iterations` LXD will request a final memory dump
from CRIU and migrate the container.
-->
LXD では、[CRIU](http://criu.org) を使ったコンテナのライブマイグレーションが使えます。
コンテナのメモリ転送を最適化するために、`migration.incremental.memory` プロパティを `true` に設定することで、CRIU の pre-copy 機能を使うように LXD を設定できます。
つまり、LXD は CRIU にコンテナの一連のメモリダンプを実行するように要求します。
ダンプが終わると、LXD はメモリダンプを指定したリモートホストに送ります。
理想的なシナリオでは、各メモリダンプを前のメモリダンプとの差分にまで減らし、それによりすでに同期されたメモリの割合を増やします。
同期されたメモリの割合が `migration.incremental.memory.goal` で設定したしきい値と等しいか超えた場合、LXD は CRIU に最終的なメモリダンプを実行し、転送するように要求します。
`migration.incremental.memory.iterations` で指定したメモリダンプの最大許容回数に達した後、まだしきい値に達していない場合は、LXD は最終的なメモリダンプを CRIU に要求し、コンテナを移行します。
