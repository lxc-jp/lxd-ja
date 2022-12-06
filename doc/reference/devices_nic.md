(devices-nic)=
# タイプ: `nic`

LXD では、様々な種類のネットワークデバイス（ネットワークインターフェースコントローラーや NIC と呼びます）が使えます:

インスタンスにネットワークデバイスを追加する際には、追加したいデバイスのタイプを選択するのに 2 つの方法があります。
`nictype` プロパティを指定するか `network` プロパティを使うかです。

## `network` プロパティを使って NIC を指定する

`network` プロパティを指定する場合、 NIC は既存の管理されたネットワークにリンクされ、 `nictype` はネットワークのタイプに応じて自動的に検出されます。

NIC の設定の一部は個々の NIC で変更可能ではなくネットワークから継承されます。

これらの詳細は下記の NIC 固有のセクションの "Managed" カラムに記載します。

## 利用可能な NIC

NIC ごとにどのプロパティが設定可能かの詳細については下記を参照してください。

次の NIC は `nictype` か `network` プロパティを使って選択できます。

- [`bridged`](#nic-bridged): ホスト上に存在するブリッジを使います。ホストのブリッジとインスタンスを接続する仮想デバイスペアを作成します。
- [`macvlan`](#nic-macvlan): 既存のネットワークデバイスをベースに MAC が異なる新しいネットワークデバイスを作成します。
- [`sriov`](#nic-sriov): SR-IOV が有効な物理ネットワークデバイスの仮想ファンクション（virtual function）をインスタンスに与えます。

次の NIC は `network` プロパティのみを使って選択できます。

- [`ovn`](#nic-ovn): 既存の OVN ネットワークを使用し、インスタンスが接続する仮想デバイスペアを作成します。

 の NIC は `nictype` プロパティのみを使って選択できます。

- [`physical`](#nic-physical): ホストの物理デバイスを直接使います。対象のデバイスはホスト上では見えなくなり、インスタンス内に出現します。
- [`ipvlan`](#nic-ipvlan): 既存のネットワークデバイスをベースに MAC アドレスは同じですが IP アドレスが異なる新しいネットワークデバイスを作成します。
- [`p2p`](#nic-p2p): 仮想デバイスペアを作成し、片方をインスタンス内に置き、残りの片方をホスト上に残します。
- [`routed`](#nic-routed): 仮想デバイスペアを作成し、ホストからインスタンスに繋いで静的ルートをセットアップし ARP/NDP エントリーをプロキシします。これにより指定された親インタフェースのネットワークにインスタンスが参加できるようになります。

(instance_device_type_nic_bridged)=
## `nic`: `bridged`

サポートされるインスタンスタイプ: コンテナ, VM

この NIC の指定に使えるプロパティ: `nictype`, `network`

ホストの既存のブリッジを使用し、ホストのブリッジをインスタンスに接続するための仮想デバイスのペアを作成します。

デバイス設定プロパティは以下の通りです。

キー                      | 型      | デフォルト値       | 必須 | 管理 | 説明
:--                       | :--     | :--                | :--  | :--  | :--
`parent`                  | string  | -                  | yes  | yes  | ホストデバイスの名前
`network`                 | string  | -                  | yes  | no   | （parent の代わりに）デバイスをリンクする先の LXD ネットワーク
`name`                    | string  | カーネルが割り当て | no   | no   | インスタンス内でのインタフェースの名前
`mtu`                     | integer | 親の MTU           | no   | yes  | 新しいインタフェースの MTU
`hwaddr`                  | string  | ランダムに割り当て | no   | no   | 新しいインタフェースの MAC アドレス
`host_name`               | string  | ランダムに割り当て | no   | no   | ホスト内でのインタフェースの名前
`limits.ingress`          | string  | -                  | no   | no   | 入力トラフィックの I/O 制限値（さまざまな単位が使用可能、 {ref}`instances-limit-units` 参照）
`limits.egress`           | string  | -                  | no   | no   | 出力トラフィックの I/O 制限値（さまざまな単位が使用可能、 {ref}`instances-limit-units` 参照）
`limits.max`              | string  | -                  | no   | no   | `limits.ingress` と `limits.egress` の両方を同じ値に変更する
`ipv4.address`            | string  | -                  | no   | no   | DHCP でインスタンスに割り当てる IPv4 アドレス（`security.ipv4_filtering` 設定時に全ての IPv4 トラフィックを制限するには `none` と設定可能）
`ipv6.address`            | string  | -                  | no   | no   | DHCP でインスタンスに割り当てる IPv6 アドレス（`security.ipv6_filtering` 設定時に全ての IPv6 トラフィックを制限するには `none` と設定可能）
`ipv4.routes`             | string  | -                  | no   | no   | ホスト上で NIC に追加する IPv4 静的ルートのカンマ区切りリスト
`ipv6.routes`             | string  | -                  | no   | no   | ホスト上で NIC に追加する IPv6 静的ルートのカンマ区切りリスト
`ipv4.routes.external`    | string  | -                  | no   | no   | NIC にルーティングしアップリンクのネットワーク (BGP) で公開する IPv4 静的ルートのカンマ区切りリスト
`ipv6.routes.external`    | string  | -                  | no   | no   | NIC にルーティングしアップリンクのネットワーク (BGP) で公開する IPv6 静的ルートのカンマ区切りリスト
`security.mac_filtering`  | bool    | `false`            | no   | no   | インスタンスが他のインスタンスの MAC アドレスになりすますのを防ぐ
`security.ipv4_filtering` | bool    | `false`            | no   | no   | インスタンスが他のインスタンスの IPv4 アドレスになりすますのを防ぐ (これを設定すると `mac_filtering` も有効になります）
`security.ipv6_filtering` | bool    | `false`            | no   | no   | インスタンスが他のインスタンスの IPv6 アドレスになりすますのを防ぐ (これを設定すると `mac_filtering` も有効になります）
`maas.subnet.ipv4`        | string  | -                  | no   | yes  | インスタンスを登録する MAAS IPv4 サブネット
`maas.subnet.ipv6`        | string  | -                  | no   | yes  | インスタンスを登録する MAAS IPv6 サブネット
`boot.priority`           | integer | -                  | no   | no   | VM のブート優先度 (高いほうが先にブート)
`vlan`                    | integer | -                  | no   | no   | タグなしのトラフィックに使用する VLAN ID （デフォルトの VLAN からポートを削除するには `none` を指定）
`vlan.tagged`             | integer | -                  | no   | no   | タグありのトラフィックに参加する VLAN ID または VLAN の範囲のカンマ区切りリスト
`security.port_isolation` | bool    | `false`            | no   | no   | NIC がポート隔離を有効にしたネットワーク内の他の NIC と通信するのを防ぐ

## `nic`: `macvlan`

サポートされるインスタンスタイプ: コンテナ, VM

この NIC の指定に使えるプロパティ: `nictype`, `network`

既存のネットワークデバイスを元に新しいネットワークデバイスをセットアップしますが、異なる MAC アドレスを用います。

デバイス設定プロパティは以下の通りです。

キー               | 型      | デフォルト値       | 必須 | 管理 | 説明
:--                | :--     | :--                | :--  | :--  | :--
`parent`           | string  | -                  | yes  | yes  | ホストデバイスの名前
`network`          | string  | -                  | yes  | no   | （parent の代わりに）デバイスをリンクする先の LXD ネットワーク
`name`             | string  | カーネルが割り当て | no   | no   | インスタンス内部でのインタフェース名
`mtu`              | integer | 親の MTU           | no   | yes  | 新しいインタフェースの MTU
`hwaddr`           | string  | ランダムに割り当て | no   | no   | 新しいインタフェースの MAC アドレス
`vlan`             | integer | -                  | no   | no   | アタッチ先の VLAN ID
`gvrp`             | bool    | `false`            | no   | no   | GARP VLAN Registration Protocol を使って VLAN を登録する
`maas.subnet.ipv4` | string  | -                  | no   | yes  | インスタンスを登録する MAAS IPv4 サブネット
`maas.subnet.ipv6` | string  | -                  | no   | yes  | インスタンスを登録する MAAS IPv6 サブネット
`boot.priority`    | integer | -                  | no   | no   | VM のブート優先度 (高いほうが先にブート)

## `nic`: `sriov`

サポートされるインスタンスタイプ: コンテナ, VM

この NIC の指定に使えるプロパティ: `nictype`, `network`

SR-IOV を有効にした物理ネットワークデバイスの仮想ファンクションをインスタンスに渡します。

デバイス設定プロパティは以下の通りです。

キー                     | 型      | デフォルト値       | 必須 | 管理 | 説明
:--                      | :--     | :--                | :--  | :--  | :--
`parent`                 | string  | -                  | yes  | yes  | ホストデバイスの名前
`network`                | string  | -                  | yes  | no   | （parent の代わりに）デバイスをリンクする先の LXD ネットワーク
`name`                   | string  | カーネルが割り当て | no   | no   | インスタンス内部でのインタフェース名
`mtu`                    | integer | カーネルが割り当て | no   | yes  | 新しいインタフェースの MTU
`hwaddr`                 | string  | ランダムに割り当て | no   | no   | 新しいインタフェースの MAC アドレス
`security.mac_filtering` | bool    | `false`            | no   | no   | インスタンスが他のインスタンスの MAC アドレスになりすますのを防ぐ
`vlan`                   | integer | -                  | no   | no   | アタッチ先の VLAN ID
`maas.subnet.ipv4`       | string  | -                  | no   | yes  | インスタンスを登録する MAAS IPv4 サブネット
`maas.subnet.ipv6`       | string  | -                  | no   | yes  | インスタンスを登録する MAAS IPv6 サブネット
`boot.priority`          | integer | -                  | no   | no   | VM のブート優先度 (高いほうが先にブート)

(instance_device_type_nic_ovn)=
## `nic`: `ovn`

サポートされるインスタンスタイプ: コンテナ, VM

この NIC の指定に使えるプロパティ: `network`

既存の OVN ネットワークを使用し、インスタンスが接続する仮想デバイスペアを作成します。

デバイス設定プロパティは以下の通りです。

キー                                   | 型      | デフォルト値       | 必須 | 管理 | 説明
:--                                    | :--     | :--                | :--  | :--  | :--
`network`                              | string  | -                  | yes  | yes  | デバイスの接続先の LXD ネットワーク
`acceleration`                         | string  | `none`             | no   | no   | ハードウェアオフローディングを有効にする。 `none` か `sriov` (下記の SR-IOV ハードウェアアクセラレーション参照)
`name`                                 | string  | カーネルが割り当て | no   | no   | インスタンス内部でのインタフェース名
`host_name`                            | string  | ランダムに割り当て | no   | no   | ホスト内部でのインタフェース名
`hwaddr`                               | string  | ランダムに割り当て | no   | no   | 新しいインターフェースの MAC アドレス
`ipv4.address`                         | string  | -                  | no   | no   | DHCP でインスタンスに割り当てる IPv4 アドレス
`ipv6.address`                         | string  | -                  | no   | no   | DHCP でインスタンスに割り当てる IPv6 アドレス
`ipv4.routes`                          | string  | -                  | no   | no   | ホスト上で NIC に追加する IPv4 静的ルートのカンマ区切りリスト
`ipv6.routes`                          | string  | -                  | no   | no   | ホスト上で NIC に追加する IPv6 静的ルートのカンマ区切りリスト
`ipv4.routes.external`                 | string  | -                  | no   | no   | NIC へのルートとアップリンクネットワークでの公開に使用する IPv4 静的ルートのカンマ区切りリスト
`ipv6.routes.external`                 | string  | -                  | no   | no   | NIC へのルートとアップリンクネットワークでの公開に使用する IPv6 静的ルートのカンマ区切りリスト
`boot.priority`                        | integer | -                  | no   | no   | VM のブート優先度 (高いほうが先にブート)
`security.acls`                        | string  | -                  | no   | no   | 適用するネットワーク ACL のカンマ区切りリスト
`security.acls.default.egress.action`  | string  | `reject`           | no   | no   | どの ACL ルールにもマッチしない外向きのトラフィックに使うアクション
`security.acls.default.egress.logged`  | bool    | `false`            | no   | no   | どの ACL ルールにもマッチしない外向きのトラフィックをログ出力するかどうか
`security.acls.default.ingress.action` | string  | `reject`           | no   | no   | どの ACL ルールにもマッチしない内向きのトラフィックに使うアクション
`security.acls.default.ingress.logged` | bool    | `false`            | no   | no   | どの ACL ルールにもマッチしない内向きのトラフィックをログ出力するかどうか

SR-IOV ハードウェアアクセラレーション:

`acceleration=sriov` を使用するためには互換性のある SR-IOV switchdev が使用できる物理 NIC が LXD ホスト内に存在する必要があります。
LXD は、物理 NIC (PF) が switchdev モードに設定されて OVN の統合 OVN ブリッジに接続されており、1 つ以上の仮想ファンクション (VF) がアクティブであることを想定しています。

これを実現するための前提となるセットアップの行程は以下の通りです。

PF と VF のセットアップ:

PF 上(以下の例では `0000:09:00.0` の PCI アドレスで `enp9s0f0np0` という名前) の VF をアクティベートしアンバインドします。
次に `switchdev` モードと PF 上の `hw-tc-offload` を有効にします。
最後に VF をリバインドします。

```
echo 4 > /sys/bus/pci/devices/0000:09:00.0/sriov_numvfs
for i in $(lspci -nnn | grep "Virtual Function" | cut -d' ' -f1); do echo 0000:$i > /sys/bus/pci/drivers/mlx5_core/unbind; done
devlink dev eswitch set pci/0000:09:00.0 mode switchdev
ethtool -K enp9s0f0np0 hw-tc-offload on
for i in $(lspci -nnn | grep "Virtual Function" | cut -d' ' -f1); do echo 0000:$i > /sys/bus/pci/drivers/mlx5_core/bind; done
```

OVS のセットアップ:

ハードウェアオフロードを有効にし、 PF NIC を統合ブリッジ (通常は `br-int` という名前) に追加します。

```
ovs-vsctl set open_vswitch . other_config:hw-offload=true
systemctl restart openvswitch-switch
ovs-vsctl add-port br-int enp9s0f0np0
ip link set enp9s0f0np0 up
```

## `nic`: `physical`

サポートされるインスタンスタイプ: コンテナ, VM

この NIC の指定に使えるプロパティ: `nictype`

物理デバイスそのものをパススルー。対象のデバイスはホストからは消失し、インスタンス内に出現します。

デバイス設定プロパティは以下の通りです。

キー               | 型      | デフォルト値       | 必須 | 説明
:--                | :--     | :--                | :--  | :--
`parent`           | string  | -                  | yes  | ホストデバイスの名前
`name`             | string  | カーネルが割り当て | no   | インスタンス内部でのインタフェース名
`mtu`              | integer | 親の MTU           | no   | 新しいインタフェースの MTU
`hwaddr`           | string  | ランダムに割り当て | no   | 新しいインタフェースの MAC アドレス
`vlan`             | integer | -                  | no   | アタッチ先の VLAN ID
`gvrp`             | bool    | `false`            | no   | GARP VLAN Registration Protocol を使って VLAN を登録する
`maas.subnet.ipv4` | string  | -                  | no   | インスタンスを登録する MAAS IPv4 サブネット
`maas.subnet.ipv6` | string  | -                  | no   | インスタンスを登録する MAAS IPv6 サブネット
`boot.priority`    | integer | -                  | no   | VM のブート優先度 (高いほうが先にブート)

## `nic`: `ipvlan`

サポートされるインスタンスタイプ: コンテナ

この NIC の指定に使えるプロパティ: `nictype`

既存のネットワークデバイスを元に新しいネットワークデバイスをセットアップしますが、異なる IP アドレスを用います。

LXD は現状 L2 と L3S モードで IPVLAN をサポートします。

このモードではゲートウェイは LXD により自動的に設定されますが、インスタンスが起動する前に
`ipv4.address` と `ipv6.address` の設定の 1 つあるいは両方を使うことにより IP アドレスを手動で指定する必要があります。

DNS に関しては、ネームサーバは自動的には設定されないので、インスタンス内部で設定する必要があります。

`ipvlan` の `nictype` を使用するには以下の `sysctl` の設定が必要です。

IPv4 アドレスを使用する場合

```
net.ipv4.conf.<parent>.forwarding=1
```

IPv6 アドレスを使用する場合

```
net.ipv6.conf.<parent>.forwarding=1
net.ipv6.conf.<parent>.proxy_ndp=1
```

デバイス設定プロパティは以下の通りです。

キー              | 型      | デフォルト値             | 必須 | 説明
:--               | :--     | :--                      | :--  | :--
`parent`          | string  | -                        | yes  | ホストデバイスの名前
`name`            | string  | カーネルが割り当て       | no   | インスタンス内部でのインタフェース名
`mtu`             | integer | 親の MTU                 | no   | 新しいインタフェースの MTU
`mode`            | string  | `l3s`                    | no   | IPVLAN のモード (`l2` か `l3s` のいずれか）
`hwaddr`          | string  | ランダムに割り当て       | no   | 新しいインタフェースの MAC アドレス
`ipv4.address`    | string  | -                        | no   | インスタンスに追加する IPv4 静的アドレスのカンマ区切りリスト。 `l2` モードでは CIDR 形式か単一アドレス形式で指定可能（単一アドレスの場合はサブネットは /24）
`ipv4.gateway`    | string  | `auto`                   | no   | `l3s` モードではデフォルト IPv4 ゲートウェイを自動的に追加するかどうか (`auto` か `none` を指定可能)。 `l2` モードではゲートウェイの IPv4 アドレスを指定。
`ipv4.host_table` | integer | -                        | no   | （メインのルーティングテーブルに加えて） IPv4 の静的ルートを追加する先のルーティングテーブル ID
`ipv6.address`    | string  | -                        | no   | インスタンスに追加する IPv6 静的アドレスのカンマ区切りリスト。 `l2` モードでは CIDR 形式か単一アドレス形式で指定可能（単一アドレスの場合はサブネットは /64）
`ipv6.gateway`    | string  | `auto` (`l3s`), - (`l2`) | no   | `l3s` モードではデフォルト IPv6 ゲートウェイを自動的に追加するかどうか (`auto` か `none` を指定可能)。 `l2` モードではゲートウェイの IPv6 アドレスを指定。
`ipv6.host_table` | integer | -                        | no   | （メインのルーティングテーブルに加えて） IPv6 の静的ルートを追加する先のルーティングテーブル ID
`vlan`            | integer | -                        | no   | アタッチ先の VLAN ID
`gvrp`            | bool    | `false`                  | no   | GARP VLAN Registration Protocol を使って VLAN を登録する

## `nic`: `p2p`

サポートされるインスタンスタイプ: コンテナ, VM

この NIC の指定に使えるプロパティ: `nictype`

仮想デバイスペアを作成し、片方はインスタンス内に配置し、もう片方はホストに残します。

デバイス設定プロパティは以下の通りです。

キー             | 型      | デフォルト値       | 必須 | 説明
:--              | :--     | :--                | :--  | :--
`name`           | string  | カーネルが割り当て | no   | インスタンス内部でのインタフェース名
`mtu`            | integer | カーネルが割り当て | no   | 新しいインタフェースの MTU
`hwaddr`         | string  | ランダムに割り当て | no   | 新しいインタフェースの MAC アドレス
`host_name`      | string  | ランダムに割り当て | no   | ホスト内でのインタフェースの名前
`limits.ingress` | string  | -                  | no   | 入力トラフィックの I/O 制限値（さまざまな単位が使用可能、 {ref}`instances-limit-units` 参照）
`limits.egress`  | string  | -                  | no   | 出力トラフィックの I/O 制限値（さまざまな単位が使用可能、 {ref}`instances-limit-units` 参照）
`limits.max`     | string  | -                  | no   | `limits.ingress` と `limits.egress` の両方を同じ値に変更する
`ipv4.routes`    | string  | -                  | no   | ホスト上で NIC に追加する IPv4 静的ルートのカンマ区切りリスト
`ipv6.routes`    | string  | -                  | no   | ホスト上で NIC に追加する IPv6 静的ルートのカンマ区切りリスト
`boot.priority`  | integer | -                  | no   | VM のブート優先度 (高いほうが先にブート)

## `nic`: `routed`

サポートされるインスタンスタイプ: コンテナ, VM

この NIC の指定に使えるプロパティ: `nictype`

この NIC タイプは運用上は IPVLAN に似ていて、ブリッジを作成することなくホストの MAC アドレスを共用してインスタンスが外部ネットワークに参加できるようにします。

しかしカーネルに IPVLAN サポートを必要としないこととホストとインスタンスが互いに通信できることが IPVLAN とは異なります。

さらにホスト上の `netfilter` のルールを尊重し、ホストのルーティングテーブルを使ってパケットをルーティングしますのでホストが複数のネットワークに接続している場合に役立ちます。

IP アドレスは `ipv4.address` と `ipv6.address` の設定のいずれかあるいは両方を使って、インスタンスが起動する前に手動で指定する必要があります。

コンテナでは仮想イーサネットデバイスペアを使用し、VM では TAP デバイスを使用します。そしてホスト側に下記のリンクローカルゲートウェイ IP アドレスを設定し、それらをインスタンス内のデフォルトゲートウェイに設定します。

    169.254.0.1
    fe80::1

コンテナではこれらはインスタンスの NIC インタフェースのデフォルトゲートウェイに自動的に設定されます。
しかし VM では IP アドレスとデフォルトゲートウェイは手動か cloud-init のような仕組みを使って設定する必要があります。

またお使いのコンテナイメージがインタフェースに対して DHCP を使うように設定されている場合、上記の自動的に追加される設定は削除される可能性が高く、その後手動か cloud-init のような仕組みを使って設定する必要があることにもご注意ください。

次にインスタンスの IP アドレス全てをインスタンスの `veth` インタフェースに向ける静的ルートをホスト上に設定します。

この NIC は `parent` のネットワークインタフェースのセットがあってもなくても利用できます。

`parent` ネットワークインタフェースのセットがある場合、インスタンスの IP の ARP/NDP のプロキシエントリーが親のインタフェースに追加され、インスタンスが親のインタフェースのネットワークにレイヤ 2 で参加できるようにします。

DNS に関してはネームサーバは自動的には設定されないので、インスタンス内で設定する必要があります。

次の `sysctl` の設定が必要です。

IPv4 アドレスを使用する場合は

```
net.ipv4.conf.<parent>.forwarding=1
```

IPv6 アドレスを使用する場合は

```
net.ipv6.conf.all.forwarding=1
net.ipv6.conf.<parent>.forwarding=1
net.ipv6.conf.all.proxy_ndp=1
net.ipv6.conf.<parent>.proxy_ndp=1
```

それぞれの NIC デバイスに複数の IP アドレスを追加できます。しかし複数の `routed` NIC インターフェースを使うほうが望ましいかもしれません。
その場合はデフォルトゲ－トウェイの衝突を避けるため、後続のインターフェースで `ipv4.gateway` と `ipv6.gateway` の値を `none` に設定するべきです。
さらにこれらの後続のインタフェースには `ipv4.host_address` と `ipv6.host_address` を用いて異なるホスト側のアドレスを設定することが有用かもしれません。

デバイス設定プロパティ

キー                  | 型      | デフォルト値       | 必須 | 説明
:--                   | :--     | :--                | :--  | :--
`parent`              | string  | -                  | no   | インスタンスが参加するホストデバイス名
`name`                | string  | カーネルが割り当て | no   | インスタンス内でのインタフェース名
`host_name`           | string  | ランダムに割り当て | no   | ホスト内でのインターフェース名
`mtu`                 | integer | 親の MTU           | no   | 新しいインタフェースの MTU
`hwaddr`              | string  | ランダムに割り当て | no   | 新しいインタフェースの MAC アドレス
`limits.ingress`      | string  | -                  | no   | 内向きトラフィックに対する bit/s での I/O 制限（さまざまな単位をサポート、 {ref}`instances-limit-units` 参照）
`limits.egress`       | string  | -                  | no   | 外向きトラフィックに対する bit/s での I/O 制限（さまざまな単位をサポート、 {ref}`instances-limit-units` 参照）
`limits.max`          | string  | -                  | no   | `limits.ingress` と `limits.egress` の両方を指定するのと同じ
`ipv4.routes`         | string  | -                  | no   | ホスト上で NIC に追加する IPv4 静的ルートのカンマ区切りリスト（L2 ARP/NDP プロキシを除く）
`ipv4.address`        | string  | -                  | no   | インスタンスに追加する IPv4 静的アドレスのカンマ区切りリスト
`ipv4.gateway`        | string  | `auto`             | no   | 自動的に IPv4 のデフォルトゲートウェイを追加するかどうか（ `auto` か `none` を指定可能）
`ipv4.host_address`   | string  | `169.254.0.1`      | no   | ホスト側の veth インターフェースに追加する IPv4 アドレス
`ipv4.host_table`     | integer | -                  | no   | （メインのルーティングテーブルに加えて） IPv4 の静的ルートを追加する先のルーティングテーブル ID
`ipv4.neighbor_probe` | bool    | `true`             | no   | IP アドレスが利用可能か知るために親のネットワークを調べるかどうか
`ipv6.address`        | string  | -                  | no   | インスタンスに追加する IPv6 静的アドレスのカンマ区切りリスト
`ipv6.routes`         | string  | -                  | no   | ホスト上で NIC に追加する IPv6 静的ルートのカンマ区切りリスト（L2 ARP/NDP プロキシを除く）
`ipv6.gateway`        | string  | `auto`             | no   | 自動的に IPv6 のデフォルトゲートウェイを追加するかどうか（ `auto` か `none` を指定可能）
`ipv6.host_address`   | string  | `fe80::1`          | no   | ホスト側の veth インターフェースに追加する IPv6 アドレス
`ipv6.host_table`     | integer | -                  | no   | （メインのルーティングテーブルに加えて） IPv6 の静的ルートを追加する先のルーティングテーブル ID
`ipv6.neighbor_probe` | bool    | `true`             | no   | IP アドレスが利用可能か知るために親のネットワークを調べるかどうか
`vlan`                | integer | -                  | no   | アタッチ先の VLAN ID
`gvrp`                | bool    | `false`            | no   | GARP VLAN Registration Protocol を使って VLAN を登録する

## `bridge`、`macvlan`、`ipvlan` を使った物理ネットワークへの接続

`bridged`、`macvlan`、`ipvlan` インターフェースタイプのいずれも、既存の物理ネットワークへ接続できます。

`macvlan` は、物理 NIC を効率的に分岐できます。つまり、物理 NIC からインスタンスで使える第 2 のインターフェースを取得できます。`macvlan` を使うことで、ブリッジデバイスと `veth` ペアの作成を減らせますし、通常はブリッジよりも良いパフォーマンスが得られます。

`macvlan` の欠点は、`macvlan` は外部との間で通信はできますが、自身の親デバイスとは通信できないことです。つまりインスタンスとホストが通信する必要がある場合は `macvlan` は使えません。

そのような場合は、 `bridge` デバイスを選ぶのが良いでしょう。`macvlan` では使えない MAC フィルタリングと I/O 制限も使えます。

`ipvlan` は `macvlan` と同様ですが、フォークされたデバイスが静的に割り当てられた IP アドレスを持ち、ネットワーク上の親の MAC アドレスを受け継ぐ点が異なります。

## SR-IOV

`sriov` インターフェースタイプで、SR-IOV が有効になったネットワークデバイスを使えます。このデバイスは、複数の仮想ファンクション（Virtual Functions: VFs）をネットワークデバイスの単一の物理ファンクション（Physical Function: PF）に関連付けます。
PF は標準の PCIe ファンクションです。一方、VFs は非常に軽量な PCIe ファンクションで、データの移動に最適化されています。
VFs は PF のプロパティを変更できないように、制限された設定機能のみを持っています。
VFs は通常の PCIe デバイスとしてシステム上に現れるので、通常の物理デバイスと同様にインスタンスに与えることができます。
`sriov` インターフェースタイプは、システム上の SR-IOV が有効になったネットワークデバイス名が、`parent` プロパティに設定されることを想定しています。
すると LXD は、システム上で使用可能な VFs があるかどうかをチェックします。デフォルトでは、LXD は検索で最初に見つかった使われていない VF を割り当てます。
有効になった VF が存在しないか、現時点で有効な VFs がすべて使われている場合は、サポートされている VF 数の最大値まで有効化し、最初の使用可能な VF をつかいます。
もしすべての使用可能な VF が使われているか、カーネルもしくはカードが VF 数を増加させられない場合は、LXD はエラーを返します。

`sriov` ネットワークデバイスは次のように作成します:

```
lxc config device add <instance> <device-name> nic nictype=sriov parent=<sriov-enabled-device>
```

特定の未使用な VF を使うように LXD に指示するには、`host_name` プロパティを追加し、有効な VF 名を設定します。

## MAAS を使った統合管理

もし、LXD ホストが接続されている物理ネットワークを MAAS を使って管理している場合で、インスタンスを直接 MAAS が管理するネットワークに接続したい場合は、MAAS とやりとりをしてインスタンスをトラッキングするように LXD を設定できます。

そのためには、デーモンに対して、`maas.api.url` と `maas.api.key` を設定しなければなりません。
そして、`maas.subnet.ipv4` と `maas.subnet.ipv6` の両方またはどちらかを、インスタンスもしくはプロファイルの `nic` エントリーに設定します。

これで、LXD はすべてのインスタンスを MAAS に登録し、適切な DHCP リースと DNS レコードがインスタンスに与えられます。

`ipv4.address` もしくは `ipv6.address` を NIC に設定した場合は、MAAS 上でも静的な割り当てとして登録されます。
