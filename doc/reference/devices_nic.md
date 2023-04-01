(devices-nic)=
# タイプ: `nic`

```{note}
`nic`デバイスタイプはコンテナとVMの両方でサポートされます。

ほとんどのNICはコンテナとVMの両方でホットプラグをサポートします。
例外は各NICタイプの説明を参照してください。
```

ネットワークデバイス(*ネットワークインタフェースコントローラー*や*NIC*とも呼びます)はネットワークへの接続を提供します。
LXDはさまざまな異なるタイプのネットワークデバイス(*NICタイプ*)をサポートします。

## `nictype` 対 `network`

インスタンスにネットワークデバイスを追加する際には、追加したいデバイスのタイプを選択するのに2つの方法があります。`nictype`プロパティを指定するか`network`プロパティを使うかです。

これらの2つのデバイスオプションは相互排他であり、デバイスを作成時にどちらか1つのみ指定可能です。
しかし、`network`オプションを指定する際には、`nictype`オプションはネットワークタイプから自動的に導出されることに注意してください。

`nictype`
: `nictype`デバイスオプションを使用する際は、LXDに管理されていないネットワークインタフェースを指定できます。
  このため、LXD がネットワークインタフェースを使用するために必要な全ての情報を指定する必要があります。

  この方法を使用する際は、`nictype`オプションはデバイス作成時に指定する必要があり、作成後は変更できません。

`network`
: `network`デバイスオプションを使用する際は、NICは既存の{ref}`管理されたネットワーク <managed-networks>`にリンクされます。
  この場合、LXDはネットワークについて必要な情報を全て持っているので、デバイス追加時にはネットワーク名を指定するだけでよいです。

  この方法を使用する際は、`nictype`オプションはLXDが自動的に導出します。
  値は読み取り専用で変更できません。

  ネットワークから継承される他のデバイスオプションはNIC固有のデバイスオプションの「管理」カラムで「yes」と記載されています。
  `network`の方法を使う場合、NICのこれらのオプションを直接カスタマイズはできません。

詳細な情報は{ref}`networks`を参照してください。

## 利用可能なNIC

次のNICは`nictype`か`network`オプションを使って追加できます。

- [`bridged`](nic-bridged): ホスト上に存在する既存のブリッジを使い、ホストのブリッジをインスタンスに接続する仮想デバイスペアを作成します。
- [`macvlan`](nic-macvlan): 既存のネットワークデバイスをベースにMACアドレスが異なる新しいネットワークデバイスを作成します。
- [`sriov`](nic-sriov): SR-IOVが有効な物理ネットワークデバイスの仮想ファンクション(virtual function)をインスタンスにパススルーします。
- [`physical`](nic-physical): ホストの物理デバイスをインスタンスにパススルーします。
  対象のデバイスはホスト上では見えなくなり、インスタンス内に出現します。

次のNICは`network`オプションでのみ追加できます。

- [`ovn`](nic-ovn): 既存のOVNネットワークを使用し、インスタンスが接続する仮想デバイスペアを作成します。

次のNICは`nictype`オプションでのみ追加できます。

- [`ipvlan`](nic-ipvlan): 既存のネットワークデバイスをベースにMACアドレスは同じですがIPアドレスが異なる新しいネットワークデバイスを作成します。
- [`p2p`](nic-p2p): 仮想デバイスペアを作成し、片方をインスタンス内に置き、残りの片方をホスト上に残します。
- [`routed`](nic-routed): 仮想デバイスペアを作成し、ホストからインスタンスに繋いで静的ルートとプロキシARP/NDPエントリーを作成します。これにより指定された親インタフェースのネットワークにインスタンスが参加できるようになります。

利用可能なデバイスオプションはNICタイプによって異なり、以下のセクションの表に一覧表示されます。

(nic-bridged)=
### `nictype`: `bridged`

```{note}
このNICタイプは`nictype`オプションか`network`オプションで選択できます。
```

`bridged` NICはホストの既存のブリッジを使用し、ホストのブリッジをインスタンスに接続するための仮想デバイスのペアを作成します。

#### デバイスオプション

`bridged` タイプのNICデバイスには以下のデバイスオプションがあります。

キー                      | 型      | デフォルト値       | 管理 | 説明
:--                       | :--     | :--                | :--  | :--
`boot.priority`           | integer | -                  | no   | VMのブート優先度(高いほうが先にブート)
`host_name`               | string  | ランダムに割り当て | no   | ホスト内でのインタフェースの名前
`hwaddr`                  | string  | ランダムに割り当て | no   | 新しいインタフェースのMACアドレス
`ipv4.address`            | string  | -                  | no   | DHCPでインスタンスに割り当てるIPv4アドレス(`security.ipv4_filtering`設定時に全てのIPv4トラフィックを制限するには`none`と設定可能)
`ipv4.routes`             | string  | -                  | no   | ホスト上でNICに追加するIPv4静的ルートのカンマ区切りリスト
`ipv4.routes.external`    | string  | -                  | no   | NICにルーティングしアップリンクのネットワーク(BGP)で公開するIPv4静的ルートのカンマ区切りリスト
`ipv6.address`            | string  | -                  | no   | DHCPでインスタンスに割り当てるIPv6アドレス(`security.ipv6_filtering`設定時に全てのIPv6トラフィックを制限するには`none`と設定可能)
`ipv6.routes`             | string  | -                  | no   | ホスト上でNICに追加するIPv6静的ルートのカンマ区切りリスト
`ipv6.routes.external`    | string  | -                  | no   | NICにルーティングしアップリンクのネットワーク(BGP)で公開するIPv6静的ルートのカンマ区切りリスト
`limits.egress`           | string  | -                  | no   | 外向きトラフィックのI/O制限値(さまざまな単位が使用可能、{ref}`instances-limit-units`参照)
`limits.ingress`          | string  | -                  | no   | 内向きトラフィックのI/O制限値(さまざまな単位が使用可能、{ref}`instances-limit-units`参照)
`limits.max`              | string  | -                  | no   | 内向きと外向きの両方のトラフィックI/O制限値(`limits.ingress`と`limits.egress`の両方を設定するのと同じ)
`maas.subnet.ipv4`        | string  | -                  | yes  | インスタンスを登録するMAAS IPv4サブネット
`maas.subnet.ipv6`        | string  | -                  | yes  | インスタンスを登録するMAAS IPv6サブネット
`mtu`                     | integer | 親の MTU           | yes  | 新しいインタフェースのMTU
`name`                    | string  | カーネルが割り当て | no   | インスタンス内でのインタフェースの名前
`network`                 | string  | -                  | no   | (`nictype`を直接設定する代わりに)デバイスをリンクする先の管理されたネットワーク
`parent`                  | string  | -                  | yes  | ホストデバイスの名前(`nictype`を直接設定する場合は必須)
`queue.tx.length`         | integer | -                  | no   | NICの送信キューの長さ
`security.ipv4_filtering` | bool    | `false`            | no   | インスタンスが他のインスタンスのIPv4アドレスになりすますのを防ぐ(これを設定すると`mac_filtering`も有効になります)
`security.ipv6_filtering` | bool    | `false`            | no   | インスタンスが他のインスタンスのIPv6アドレスになりすますのを防ぐ(これを設定すると`mac_filtering`も有効になります)
`security.mac_filtering`  | bool    | `false`            | no   | インスタンスが他のインスタンスのMACアドレスになりすますのを防ぐ
`security.port_isolation` | bool    | `false`            | no   | NICがポート隔離を有効にしたネットワーク内の他のNICと通信するのを防ぐ
`vlan`                    | integer | -                  | no   | タグなしのトラフィックに使用するVLAN ID(デフォルトのVLANからポートを削除するには`none`を指定)
`vlan.tagged`             | integer | -                  | no   | タグありのトラフィックに参加するVLAN IDまたはVLANの範囲のカンマ区切りリスト

(nic-macvlan)=
### `nictype`: `macvlan`

```{note}
このNICタイプは`nictype`オプションか`network`オプションで選択できます。
```

`macvlan` NICは既存のNICをベースにしますが、MACアドレスが異なる新しいネットワークデバイスをセットアップします。

#### デバイスオプション

`macvlan`タイプのNICデバイスには以下のデバイスオプションがあります。

キー               | 型      | デフォルト値       | 管理 | 説明
:--                | :--     | :--                | :--  | :--
`boot.priority`    | integer | -                  | no   | VMのブート優先度(高いほうが先にブート)
`gvrp`             | bool    | `false`            | no   | GARP VLAN Registration Protocolを使ってVLANを登録する
`hwaddr`           | string  | ランダムに割り当て | no   | 新しいインタフェースのMACアドレス
`maas.subnet.ipv4` | string  | -                  | yes  | インスタンスを登録するMAAS IPv4サブネット
`maas.subnet.ipv6` | string  | -                  | yes  | インスタンスを登録するMAAS IPv6サブネット
`mtu`              | integer | 親の MTU           | yes  | 新しいインタフェースのMTU
`name`             | string  | カーネルが割り当て | no   | インスタンス内部でのインタフェース名
`network`          | string  | -                  | no   | (`nictype`を直接設定する代わりに)デバイスをリンクする先の管理されたネットワーク
`parent`           | string  | -                  | yes  | ホストデバイスの名前(`nictype`を直接設定する場合は必須)
`vlan`             | integer | -                  | no   | アタッチ先のVLAN ID

(nic-sriov)=
### `nictype`: `sriov`

```{note}
このNICタイプは`nictype`オプションか`network`オプションで選択できます。
```

`sriov` NICはSR-IOVを有効にした物理ネットワークデバイスの仮想ファンクションをインスタンスにパススルーします。

SR-IOVを有効にしたネットワークデバイスは一組の仮想ファンクション(VF)をネットワークデバイスの単一の物理ファンクション(PF)に関連付けます。
PFは標準的なPCIe関数です。
一方、VFはデータの移動に最適化された非常に軽量なPCIe関数です。
PFのプロパティを変えるのを防ぐため、VFの構成機能は限定されています。

VFはシステムには通常のPCIeデバイスのように見えますので、通常の物理デバイスと全く同じようにインスタンスにパススルーできます。

VFの割り当て
: `sriov`インタフェースタイプは`parent`プロパティを通してシステム上のSR-IOVを有効にしたネットワークデバイスの名前を渡されることを想定しています。
  するとLXDはシステム上の任意の利用可能なVFをチェックします。

  デフォルトでは、LXDは見つけた最初の未使用なVFを割り当てます。
  有効になっているものが1つもないか、有効なVFが全て使用中の場合、サポートされているVFの数を最大に上げて最初の未使用なVFを使用します。
  全ての利用可能なVFが使用中か、カーネルまたはカードがVFの数の増加をサポートしない場合は、LXDはエラーを返します。

  ```{note}
  LXDに特定のVFを使わせたい場合、`sriov` NICの代わりに`physical` NICを使用し、`parent`オプションをVF名に設定してください。
  ```

#### デバイスオプション

`sriov`タイプのNICデバイスには以下のデバイスオプションがあります。

キー                     | 型      | デフォルト値       | 管理 | 説明
:--                      | :--     | :--                | :--  | :--
`boot.priority`          | integer | -                  | no   | VMのブート優先度(高いほうが先にブート)
`hwaddr`                 | string  | ランダムに割り当て | no   | 新しいインタフェースのMACアドレス
`maas.subnet.ipv4`       | string  | -                  | yes  | インスタンスを登録するMAAS IPv4サブネット
`maas.subnet.ipv6`       | string  | -                  | yes  | インスタンスを登録するMAAS IPv6サブネット
`mtu`                    | integer | カーネルが割り当て | yes  | 新しいインタフェースのMTU
`name`                   | string  | カーネルが割り当て | no   | インスタンス内部でのインタフェース名
`network`                | string  | -                  | no   |  (`nictype`を直接設定する代わりに)デバイスをリンクする先の管理されたネットワーク
`parent`                 | string  | -                  | yes  | ホストデバイスの名前(`nictype`を直接設定する場合は必須)
`security.mac_filtering` | bool    | `false`            | no   | インスタンスが他のインスタンスのMACアドレスになりすますのを防ぐ
`vlan`                   | integer | -                  | no   | アタッチ先のVLAN ID

(nic-ovn)=
### `nictype`: `ovn`

```{note}
- このNICタイプは`network`オプションでのみ選択できます。
- このNICタイプはコンテナでのみホットプラグをサポートし、VM ではサポートしません。
```

`ovn` NICは既存のOVNネットワークを使用し、それにインスタンスが接続する仮想デバイスペアを作成します。

(devices-nic-hw-acceleration)=
SR-IOVハードウェアアクセラレーション
: `acceleration=sriov`を使用するには、LXDホスト内のEthernetスイッチデバイスのドライバモデル(`switchdev`)をサポートする互換性のあるSR-IOV物理NICを持っている必要があります。
  LXDは物理NIC(PF)が`switchdev`モードに設定され、OVN統合OVSブリッジに接続され、1つ以上の仮想ファンクション(VF)がアクティブになっていることを前提とします。

  これを実現するには、基本的な前提条件となる以下のセットアップ手順に従ってください。

   1. PFとVFをセットアップする

      1. PF上でいくつかのVFをアクティベートし(以下の例では`enp9s0f0np0`とし、PCIアドレスは`0000:09:00.0`とします)、アンバインドします。
      1. `switchdev`モードとPF上の`hw-tc-offload`を有効にします。
      1. VFをリバインドします。

      ```
      echo 4 > /sys/bus/pci/devices/0000:09:00.0/sriov_numvfs
      for i in $(lspci -nnn | grep "Virtual Function" | cut -d' ' -f1); do echo 0000:$i > /sys/bus/pci/drivers/mlx5_core/unbind; done
      devlink dev eswitch set pci/0000:09:00.0 mode switchdev
      ethtool -K enp9s0f0np0 hw-tc-offload on
      for i in $(lspci -nnn | grep "Virtual Function" | cut -d' ' -f1); do echo 0000:$i > /sys/bus/pci/drivers/mlx5_core/bind; done
      ```

   1. ハードウェアオフロードを有効にし、統合ブリッジ(通常`br-int`と呼ばれます)に PF NICを追加してOVSをセットアップします。

      ```
      ovs-vsctl set open_vswitch . other_config:hw-offload=true
      systemctl restart openvswitch-switch
      ovs-vsctl add-port br-int enp9s0f0np0
      ip link set enp9s0f0np0 up
      ```

#### デバイスオプション

`sriov` タイプのNICデバイスには以下のデバイスオプションがあります。

キー                                   | 型      | デフォルト値       | 管理 | 説明
:--                                    | :--     | :--                | :--  | :--
`acceleration`                         | string  | `none`             | no   | ハードウェアオフローディングを有効にする(`none`か`sriov`、{ref}`devices-nic-hw-acceleration`参照)
`boot.priority`                        | integer | -                  | no   | VMのブート優先度(高いほうが先にブート)
`host_name`                            | string  | ランダムに割り当て | no   | ホスト内部でのインタフェース名
`hwaddr`                               | string  | ランダムに割り当て | no   | 新しいインタフェースのMACアドレス
`ipv4.address`                         | string  | -                  | no   | DHCPでインスタンスに割り当てるIPv4アドレス
`ipv4.routes`                          | string  | -                  | no   | NICへルーティングするIPv4静的ルートのカンマ区切りリスト
`ipv4.routes.external`                 | string  | -                  | no   | NICへのルーティングとアップリンクネットワークでの公開に使用するIPv4静的ルートのカンマ区切りリスト
`ipv6.address`                         | string  | -                  | no   | DHCPでインスタンスに割り当てるIPv6アドレス
`ipv6.routes`                          | string  | -                  | no   | NICへルーティングするIPv6静的ルートのカンマ区切りリスト
`ipv6.routes.external`                 | string  | -                  | no   | NICへのルーティングとアップリンクネットワークでの公開に使用するIPv6静的ルートのカンマ区切りリスト
`name`                                 | string  | カーネルが割り当て | no   | インスタンス内部でのインタフェース名
`network`                              | string  | -                  | yes  | デバイスの接続先の管理されたネットワーク(必須)
`security.acls`                        | string  | -                  | no   | 適用するネットワークACLのカンマ区切りリスト
`security.acls.default.egress.action`  | string  | `reject`           | no   | どのACLルールにもマッチしない外向きトラフィックに使うアクション
`security.acls.default.egress.logged`  | bool    | `false`            | no   | どのACLルールにもマッチしない外向きトラフィックをログ出力するかどうか
`security.acls.default.ingress.action` | string  | `reject`           | no   | どのACLルールにもマッチしない内向きトラフィックに使うアクション
`security.acls.default.ingress.logged` | bool    | `false`            | no   | どのACLルールにもマッチしない内向きトラフィックをログ出力するかどうか

(nic-physical)=
### `nictype`: `physical`

```{note}
- このNICタイプは`nictype`オプションまたは`network`オプションで選択できます。
- それぞれの親デバイスに対して`physical` NIC は1つだけ持つことができます。
```

`physical` NICはホストからパススルーされるそのままの物理デバイスを提供します。
対象のデバイスはホストから消失し、インスタンス内に出現します(これは各ターゲットデバイスに`physical` NICは1つだけ持つことができることを意味します)。

#### デバイスオプション

`physical`タイプのNICデバイスには以下のデバイスオプションがあります。

キー               | 型      | デフォルト値       | 管理    | 説明
:--                | :--     | :--                | :--     | :--
`boot.priority`    | integer | -                  | no      | VMのブート優先度(高いほうが先にブート)
`gvrp`             | bool    | `false`            | no      | GARP VLAN Registration Protocolを使ってVLANを登録する
`hwaddr`           | string  | ランダムに割り当て | no      | 新しいインタフェースのMACアドレス
`maas.subnet.ipv4` | string  | -                  | no      | インスタンスを登録するMAAS IPv4サブネット
`maas.subnet.ipv6` | string  | -                  | no      | インスタンスを登録するMAAS IPv6サブネット
`mtu`              | integer | 親の MTU           | no      | 新しいインタフェースのMTU
`name`             | string  | カーネルが割り当て | no      | インスタンス内部でのインタフェース名
`network`          | string  | -                  | no      | デバイスのリンク先(`nictype`を直接指定する代わりに)の管理ネットワーク
`parent`           | string  | -                  | yes     | ホストデバイスの名前(必須)
`vlan`             | integer | -                  | no      | アタッチ先のVLAN ID

(nic-ipvlan)=
### `nictype`: `ipvlan`

```{note}
- このNICタイプはコンテナのみで利用でき、仮想マシンでは利用できません。
- このNICタイプは`nictype`オプションでのみ選択できます。
- このNICタイプはホットプラグをサポートしません。
```

`ipvlan` NICは既存のネットワークデバイスを元に、同じMACアドレスですがIPアドレスは異なるような新しいネットワークデバイスをセットアップします。

LXDは現状L2とL3SモードでIPVLANをサポートします。
このモードでは、ゲートウェイはLXDにより自動的に設定されますが、コンテナが起動する前に`ipv4.address`と`ipv6.address`の設定の1つあるいは両方を使うことによりIPアドレスを手動で指定する必要があります。

DNS
: ネームサーバは自動的には設定されないので、コンテナ内部で設定する必要があります。
  このためには、以下の`sysctl`の設定をしてください。

   - IPv4アドレスを使用する場合

     ```
     net.ipv4.conf.<parent>.forwarding=1
     ```

   - IPv6アドレスを使用する場合

     ```
     net.ipv6.conf.<parent>.forwarding=1
     net.ipv6.conf.<parent>.proxy_ndp=1
     ```

#### デバイスオプション

`ipvlan`タイプのNICデバイスには以下のデバイスオプションがあります。

キー              | 型      | デフォルト値             | 説明
:--               | :--     | :--                      | :--
`gvrp`            | bool    | `false`                  | GARP VLAN Registration Protocolを使ってVLANを登録する
`hwaddr`          | string  | ランダムに割り当て       | 新しいインタフェースのMACアドレス
`ipv4.address`    | string  | -                        | インスタンスに追加するIPv4静的アドレスのカンマ区切りリスト(`l2`モードでは、CIDR形式か`/24`のサブネットの単一アドレスで指定可能)
`ipv4.gateway`    | string  | `auto` (`l3s`), - (`l2`) | `l3s`モードでは、デフォルトIPv4ゲートウェイを自動的に追加するかどうか(`auto`か`none`を指定可能)。`l2`モードでは、ゲートウェイのIPv4アドレス。
`ipv4.host_table` | integer | -                        | (メインのルーティングテーブルに加えて)IPv4の静的ルートを追加する先のカスタムポリシー・ルーティングテーブルID
`ipv6.address`    | string  | -                        | インスタンスに追加するIPv6静的アドレスのカンマ区切りリスト(`l2`モードでは、CIDR 形式か`/64`のサブネットの単一アドレスで指定可能)
`ipv6.gateway`    | string  | `auto` (`l3s`), - (`l2`) | `l3s`モードでは、デフォルトIPv6ゲートウェイを自動的に追加するかどうか(`auto`か`none`を指定可能)。`l2`モードで、はゲートウェイのIPv6アドレス。
`ipv6.host_table` | integer | -                        | (メインのルーティングテーブルに加えて)IPv6の静的ルートを追加する先のカスタムポリシー・ルーティングテーブルID
`mode`            | string  | `l3s`                    | IPVLANのモード(`l2`か`l3s`のいずれか)
`mtu`             | integer | 親の MTU                 | 新しいインタフェースのMTU
`name`            | string  | カーネルが割り当て       | インスタンス内部でのインタフェース名
`queue.tx.length` | integer | -                        | NICの送信キューの長さ
`parent`          | string  | -                        | ホストデバイスの名前(必須)
`vlan`            | integer | -                        | アタッチ先のVLAN ID

(nic-p2p)=
### `nictype`: `p2p`

```{note}
このNICタイプは`nictype`オプションでのみ選択できます。
```

`p2p` NICは仮想デバイスペアを作成し、片方はインスタンス内に配置し、もう片方はホストに残します。

#### デバイスオプション

`p2p`タイプのNICデバイスには以下のデバイスオプションがあります。

キー             | 型      | デフォルト値       | 説明
:--              | :--     | :--                | :--
`boot.priority`  | integer | -                  | VMのブート優先度 (高いほうが先にブート)
`host_name`      | string  | ランダムに割り当て | ホスト内でのインタフェースの名前
`hwaddr`         | string  | ランダムに割り当て | 新しいインタフェースのMACアドレス
`ipv4.routes`    | string  | -                  | ホスト上でNICに追加するIPv4静的ルートのカンマ区切りリスト
`ipv6.routes`    | string  | -                  | ホスト上でNICに追加するIPv6静的ルートのカンマ区切りリスト
`limits.egress`  | string  | -                  | 外向きトラフィックのI/O制限値(さまざまな単位が使用可能、{ref}`instances-limit-units`参照)
`limits.ingress` | string  | -                  | 内向きトラフィックのI/O制限値(さまざまな単位が使用可能、{ref}`instances-limit-units`参照)
`limits.max`     | string  | -                  | 内向きと外向きの両方のトラフィックI/O制限値(`limits.ingress`と`limits.egress`の両方を設定するのと同じ)
`mtu`            | integer | カーネルが割り当て | 新しいインタフェースのMTU
`name`           | string  | カーネルが割り当て | インスタンス内部でのインタフェース名

(nic-routed)=
### `nictype`: `routed`

```{note}
このNICタイプは`nictype`オプションでのみ選択できます。
```

`routed` NICタイプはホストをインスタンスに接続する仮想デバイスペアを作成し、インスタンスが指定された親インタフェースのネットワークに参加できるように、静的ルートとプロキシARP/NDPエントリをセットアップします。
コンテナでは仮想イーサネットデバイスペアを使用し、VMではTAPデバイスを使用します。

このNICタイプは運用上はIPVLANに似ていて、ブリッジを設定することなくホストのMACアドレスを共用して、インスタンスが外部ネットワークに参加できるようにします。
しかし、カーネルにIPVLANサポートを必要としないことと、ホストとインスタンスが互いに通信できることが`ipvlan`とは異なります。

このNICタイプは`netfilter`のルールを尊重し、ホストのルーティングテーブルを使ってパケットをルーティングしますので、ホストが複数のネットワークに接続している場合に役立ちます。

IP アドレス、ゲートウェイ、ルーティング
: インスタンスが起動する前にIPアドレスを(`ipv4.address`と`ipv6.address`の設定のいずれかあるいは両方を使って)手動で指定する必要があります。

  コンテナでは、NICはホスト上に下記のリンクローカルゲートウェイIPアドレスを設定し、それらをコンテナのNICインタフェースのデフォルトゲートウェイに設定します。

    169.254.0.1
    fe80::1

  VMでは、ゲートウェイは手動か`cloud-init`のような仕組みを使って設定する必要があります。

  ```{note}
  お使いのコンテナイメージがインタフェースに対してDHCPを使うように設定されている場合、上記の自動的に追加される設定は削除される可能性が高いです。
  この場合、IPアドレスとゲートウェイを手動か`cloud-init`のような仕組みを使って設定する必要があります。
  ```

  このNICタイプはインスタンスのIPアドレス全てをインスタンスの`veth`インタフェースに向ける静的ルートをホスト上に設定します。

複数のIPアドレス
: それぞれのNICデバイスに複数のIPアドレスを追加できます。

  しかし、代わりに複数の`routed` NICインタフェースを使うほうが望ましいかもしれません。
  この場合、`ipv4.gateway`と`ipv6.gateway`の値を`none`に設定し、後続のインタフェースがデフォルトゲートウェイの衝突を避けるようにします。
  さらに、これらの後続のインタフェースに`ipv4.host_address`と`ipv6.host_address`を使って異なるホスト側のアドレスを指定することを検討してください。

親のインタフェース
: このNICは`parent`のネットワークインタフェースのセットがあってもなくても利用できます。

  `parent`ネットワークインタフェースのセットがある場合、インスタンスのIPのプロキシARP/NDPエントリが親のインタフェースに追加され、インスタンスが親のインタフェースのネットワークにレイヤ2で参加できるようにします。

DNS
: ネームサーバは自動的には設定されないので、インスタンス内で設定する必要があります。
  このためには、以下の`sysctl`の設定をしてください。

   - IPv4アドレスを使用する場合

     ```
     net.ipv4.conf.<parent>.forwarding=1
     ```

   - IPv6アドレスを使用する場合

     ```
     net.ipv6.conf.all.forwarding=1
     net.ipv6.conf.<parent>.forwarding=1
     net.ipv6.conf.all.proxy_ndp=1
     net.ipv6.conf.<parent>.proxy_ndp=1
     ```

#### デバイスオプション

`routed`タイプのNICデバイスには以下のデバイスオプションがあります。

キー                  | 型      | デフォルト値       | 説明
:--                   | :--     | :--                | :--
`gvrp`                | bool    | `false`            | GARP VLAN Registration Protocolを使ってVLANを登録する
`host_name`           | string  | ランダムに割り当て | ホスト内でのインタフェース名
`hwaddr`              | string  | ランダムに割り当て | 新しいインタフェースのMACアドレス
`ipv4.address`        | string  | -                  | インスタンスに追加するIPv4静的アドレスのカンマ区切りリスト
`ipv4.gateway`        | string  | `auto`             | 自動的にIPv4デフォルトゲートウェイを追加するかどうか(`auto`か`none`を指定可能)
`ipv4.host_address`   | string  | `169.254.0.1`      | ホスト側のvethインタフェースに追加するIPv4アドレス
`ipv4.host_table`     | integer | -                  | (メインのルーティングテーブルに加えて)IPv4の静的ルートを追加する先のカスタムポリシー・ルーティングテーブルID
`ipv4.neighbor_probe` | bool    | `true`             | IPアドレスが利用可能か知るために親のネットワークを調べるかどうか
`ipv4.routes`         | string  | -                  | ホスト上でNICに追加するIPv4静的ルートのカンマ区切りリスト(L2 ARP/NDPプロキシを除く)
`ipv6.address`        | string  | -                  | インスタンスに追加するIPv6静的アドレスのカンマ区切りリスト
`ipv6.gateway`        | string  | `auto`             | 自動的にIPv6のデフォルトゲートウェイを追加するかどうか(`auto`か`none`を指定可能)
`ipv6.host_address`   | string  | `fe80::1`          | ホスト側のvethインタフェースに追加するIPv6アドレス
`ipv6.host_table`     | integer | -                  | (メインのルーティングテーブルに加えて)IPv6の静的ルートを追加する先のカスタムポリシー・ルーティングテーブルID
`ipv6.neighbor_probe` | bool    | `true`             | IPアドレスが利用可能か知るために親のネットワークを調べるかどうか
`ipv6.routes`         | string  | -                  | ホスト上でNICに追加するIPv6静的ルートのカンマ区切りリスト(L2 ARP/NDPプロキシを除く)
`limits.ingress`      | string  | -                  | 内向きトラフィックに対するbit/sでのI/O制限値(さまざまな単位をサポート、{ref}`instances-limit-units`参照)
`limits.egress`       | string  | -                  | 外向きトラフィックに対するbit/sでのI/O制限値(さまざまな単位をサポート、{ref}`instances-limit-units`参照)
`limits.max`          | string  | -                  | 内向きと外向き両方のトラフィックのI/O 制限値(`limits.ingress`と`limits.egress`の両方を設定するのと同じ)
`mtu`                 | integer | 親の MTU           | 新しいインタフェースのMTU
`name`                | string  | カーネルが割り当て | インスタンス内でのインタフェース名
`parent`              | string  | -                  | インスタンスが参加するホストデバイス名
`queue.tx.length`     | integer | -                  | NICの送信キューの長さ
`vlan`                | integer | -                  | アタッチ先の VLAN ID

## `bridge`、`macvlan`、`ipvlan`を使った物理ネットワークへの接続

`bridged`、`macvlan`、`ipvlan`インタフェースタイプのいずれも、既存の物理ネットワークへ接続するために使用できます。

`macvlan`は、物理NICを効率的に分岐できます。つまり、物理NICからインスタンスで使える第2のインタフェースを取得できます。
この方法はブリッジデバイスと仮想イーサネットデバイスペアの作成を不要にしますし、通常はブリッジよりも良いパフォーマンスが得られます。

`macvlan`の欠点は、`macvlan`はインスタンス自身と外部との間で通信はできますが、親デバイスとは通信できないことです。
つまりインスタンスとホストが通信する必要がある場合は`macvlan`は使えません。

そのような場合は、`bridge`デバイスを選ぶのが良いでしょう。
`macvlan`では使えないMACフィルタリングとI/O制限も使えます。

`ipvlan`は`macvlan`と同様ですが、フォークされたデバイスが静的に割り当てられたIPアドレスを持ち、ネットワーク上の親のMACアドレスを受け継ぐ点が異なります。

## MAASを使った統合管理

もし、LXDホストが接続されている物理ネットワークをMAASを使って管理している場合で、インスタンスを直接MAASが管理するネットワークにアタッチしたい場合は、MAASとやりとりをしてインスタンスをトラッキングするようにLXDを設定できます。

そのためには、デーモンに対して、`maas.api.url`と`maas.api.key`を設定しなければなりません。そして、`maas.subnet.ipv4`と`maas.subnet.ipv6`の両方またはどちらかを、インスタンスもしくはプロファイルの`nic`エントリーに設定します。

これで、LXDは全てのインスタンスをMAASに登録し、適切なDHCPリースとDNSレコードをインスタンスに与えます。

`ipv4.address`もしくは`ipv6.address`キーをNIC に設定した場合は、MAAS上で静的な割り当てとして登録されます。
