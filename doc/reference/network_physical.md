(network-physical)=
# 物理ネットワーク

<!-- Include start physical intro -->
物理 (`physical`) ネットワークタイプは既存のネットワークに接続します。これはネットワークインタフェースまたはブリッジになることができ、OVN のためのアップリンクネットワークとしての役目を果たします。
<!-- Include end physical intro -->

このネットワークタイプは OVN ネットワークを親インターフェースに接続する際に使用するプリセットの設定を提供したり、インスタンスが物理インタフェースを NIC として使用できるようにします。この場合、インスタンス NIC は接続先の設定詳細を知ること無く単に `network` オプションを設定できるようにします。

(network-physical-options)=
## 設定オプション

物理ネットワークでは現在以下の設定キーネームスペースがサポートされています。

 - `bgp` (BGP ピア設定)
 - `dns` (DNS サーバと名前解決の設定)
 - `ipv4` (L3 IPv4 設定)
 - `ipv6` (L3 IPv6 設定)
 - `maas` (MAAS ネットワーク識別)
 - `ovn` (OVN 設定)
 - `user` (key/value の自由形式のユーザメタデータ)

```{note}
{{note_ip_addresses_CIDR}}
```

物理ネットワークタイプには以下の設定オプションがあります。

キー                            | 型        | 条件             | デフォルト                                | 説明
:--                             | :--       | :--              | :--                                       | :--
gvrp                            | boolean   | -                | false                                     | GARP VLAN Registration Protocol を使って VLAN を登録する
mtu                             | integer   | -                | -                                         | 作成するインターフェースの MTU
parent                          | string    | -                | -                                         | ネットワークで使う既存のインターフェース
vlan                            | integer   | -                | -                                         | アタッチする先の VLAN ID
bgp.peers.NAME.address          | string    | bgp server       | -                                         | `ovn` ダウンストリームネットワークで使用するピアアドレス (IPv4 か IPv6)
bgp.peers.NAME.asn              | integer   | bgp server       | -                                         | `ovn` ダウンストリームネットワークで使用する AS 番号
bgp.peers.NAME.password         | string    | bgp server       | - (パスワード無し)                        | `ovn` ダウンストリームネットワークで使用するピアのセッションパスワード（省略可能）
dns.nameservers                 | string    | 標準モード       | -                                         | 物理 (`physical`) ネットワークの DNS サーバ IP のリスト
ipv4.gateway                    | string    | 標準モード       | -                                         | ゲートウェイとネットワークの IPv4 アドレス（CIDR表記）
ipv4.ovn.ranges                 | string    | -                | -                                         | 子供の OVN ネットワークルーターに使用する IPv4 アドレスの範囲（開始-終了 形式) のカンマ区切りリスト
ipv4.routes                     | string    | ipv4 アドレス    | -                                         | 子供の OVN ネットワークの ipv4.routes.external 設定で利用可能な追加の IPv4 CIDR サブネットのカンマ区切りリスト
ipv4.routes.anycast             | boolean   | ipv4 アドレス    | false                                     | 複数のネットワーク／NICで同時にオーバーラップするルートが使われることを許可するかどうか
ipv6.gateway                    | string    | 標準モード       | -                                         | ゲートウェイとネットワークの IPv6 アドレス（CIDR表記）
ipv6.ovn.ranges                 | string    | -                | -                                         | 子供の OVN ネットワークルーターに使用する IPv6 アドレスの範囲（開始-終了 形式) のカンマ区切りリスト
ipv6.routes                     | string    | ipv6 アドレス    | -                                         | 子供の OVN ネットワークの ipv6.routes.external 設定で利用可能な追加の IPv6 CIDR サブネットのカンマ区切りリスト
ipv6.routes.anycast             | boolean   | ipv6 アドレス    | false                                     | 複数のネットワーク／NICで同時にオーバーラップするルートが使われることを許可するかどうか
maas.subnet.ipv4                | string    | ipv4 アドレス    | -                                         | インスタンスを登録する MAAS IPv4 サブネット (NIC で `network` プロパティを使う場合に有効)
maas.subnet.ipv6                | string    | ipv6 アドレス    | -                                         | インスタンスを登録する MAAS IPv6 サブネット (NIC で `network` プロパティを使う場合に有効)
ovn.ingress\_mode               | string    | 標準モード       | l2proxy                                   | OVN NIC の外部 IP アドレスがアップリンクネットワークで広告される方法を設定します。 `l2proxy` (proxy ARP/NDP) か `routed` です。
user.*                          | string    | -                | -                                         | ユーザ指定の自由形式のキー／バリューペア

(network-physical-features)=
## サポートされている機能

物理ネットワークタイプでは以下の機能がサポートされています。

- {ref}`network-bgp`
