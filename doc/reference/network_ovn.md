---
discourse: 11033
---

(network-ovn)=
# OVN ネットワーク

<!-- Include start OVN intro -->
{abbr}`OVN (Open Virtual Network)`は仮想ネットワーク抽象化をサポートするソフトウェアで定義されたネットワークシステムです。
あなた自身のプライベートクラウドを構築するのに使用できます。
詳細は[`www.ovn.org`](https://www.ovn.org/)をご参照ください。
<!-- Include end OVN intro -->

`ovn`ネットワークタイプはOVN{abbr}`SDN (software-defined networking)`を使って論理的なネットワークの作成を可能にします。
この種のネットワークは複数の個別のネットワーク内で同じ論理ネットワークのサブネットを使うような検証環境やマルチテナントの環境で便利です。

LXDのOVNネットワークはより広いネットワークへの外向きのアクセスを可能にするため既存の管理された{ref}`network-bridge`や{ref}`network-physical`に接続できます。
デフォルトでは、OVN論理ネットワークからの全ての接続はアップリンクのネットワークによって割り当てられたIPにNATされます。

OVNネットワークをセットアップする基本的な手順については{ref}`network-ovn-setup`をご参照ください。

% Include content from [network_bridge.md](network_bridge.md)
```{include} network_bridge.md
    :start-after: <!-- Include start MAC identifier note -->
    :end-before: <!-- Include end MAC identifier note -->
```

(network-ovn-options)=
## 設定オプション

`ovn`ネットワークタイプでは現在以下の設定キーネームスペースがサポートされています。

- `bridge` (L2インタフェースの設定)
- `dns` (DNSサーバと名前解決の設定)
- `ipv4` (L3 IPv4設定)
- `ipv6` (L3 IPv6設定)
- `security` (ネットワークACL設定)
- `user` (key/valueの自由形式のユーザメタデータ)

```{note}
{{note_ip_addresses_CIDR}}
```

`ovn` ネットワークタイプには以下の設定オプションがあります。

キー                                   | 型      | 条件               | デフォルト         | 説明
:--                                    | :--     | :--                | :--                | :--
`network`                              | string  | -                  | -                  | 外部ネットワークへのアクセスに使うアップリンクのネットワーク
`bridge.hwaddr`                        | string  | -                  | -                  | ブリッジのMACアドレス
`bridge.mtu`                           | integer | -                  | `1442`             | ブリッジのMTU(デフォルトではホストからホストへのGeneveトンネルを許可します)
`dns.domain`                           | string  | -                  | `lxd`              | DHCPのクライアントに広告しDNSの名前解決に使用するドメイン
`dns.search`                           | string  | -                  | -                  | 完全なドメインサーチのカンマ区切りリスト(デフォルトは`dns.domain`の値)
`dns.zone.forward`                     | string  | -                  | -                  | 正引きDNSレコード用のDNSゾーン名のカンマ区切りリスト
`dns.zone.reverse.ipv4`                | string  | -                  | -                  | IPv4逆引きDNSレコード用のDNSゾーン名
`dns.zone.reverse.ipv6`                | string  | -                  | -                  | IPv6逆引きDNSレコード用のDNSゾーン名
`ipv4.address`                         | string  | 標準モード         | `auto`(作成時のみ) | ブリッジのIPv4アドレス(CIDR形式)。IPv4をオフにするには`none`、新しいランダムな未使用のサブネットを生成するには`auto`を指定。
`ipv4.dhcp`                            | bool    | IPv4アドレス       | `true`             | DHCPを使ってアドレスを割り当てるかどうか
`ipv4.l3only`                          | bool    | IPv4 address       | `false`            | layer 3 only モード を有効にするかどうか
`ipv4.nat`                             | bool    | IPv4アドレス       | `false`            | NATするかどうか(`ipv4.address`が未設定の場合デフォルト値は`true`でランダムな`ipv4.address`が生成されます)
`ipv4.nat.address`                     | string  | IPv4アドレス       | -                  | ネットワークからの外向きトラフィックに使用されるソースアドレス(アップリンクに`ovn.ingress_mode=routed`が必要)
`ipv6.address`                         | string  | 標準モード         | `auto`(作成時のみ) | ブリッジのIPv6アドレス(CIDR形式)。IPv6をオフにするには`none`、新しいランダムな未使用のサブネットを生成するには`auto`を指定。
`ipv6.dhcp`                            | bool    | IPv6アドレス       | `true`             | DHCP上に追加のネットワーク設定を提供するかどうか
`ipv6.dhcp.stateful`                   | bool    | IPv6 DHCP          | `false`            | DHCPを使ってアドレスを割り当てるかどうか
`ipv6.l3only`                          | bool    | IPv6 DHCP stateful | `false`            | layer 3 only モード を有効にするかどうか
`ipv6.nat`                             | bool    | IPv6アドレス       | `false`            | NATするかどうか(`ipv6.address`が未設定の場合デフォルト値は`true`でランダムな`ipv6.address`が生成されます)
`ipv6.nat.address`                     | string  | IPv6アドレス       | -                  | ネットワークからの外向きトラフィックに使用されるソースアドレス(アップリンクに`ovn.ingress_mode=routed`が必要)
`security.acls`                        | string  | -                  | -                  | このネットワークに接続するNICに適用するネットワークACLのカンマ区切りリスト
`security.acls.default.egress.action`  | string  | `security.acls`    | `reject`           | どのACLルールにもマッチしない外向きトラフィックに使うアクション
`security.acls.default.egress.logged`  | bool    | `security.acls`    | `false`            | どのACLルールにもマッチしない外向きトラフィックをログ出力するかどうか
`security.acls.default.ingress.action` | string  | `security.acls`    | `reject`           | どのACLルールにもマッチしない内向きトラフィックに使うアクション
`security.acls.default.ingress.logged` | bool    | `security.acls`    | `false`            | どのACLルールにもマッチしない内向きトラフィックをログ出力するかどうか
`user.*`                               | string  | -                  | -                  | ユーザ指定の自由形式のキー／バリューペア

(network-ovn-features)=
## サポートされている機能

`ovn`ネットワークタイプでは以下の機能がサポートされています。

- {ref}`network-acls`
- {ref}`network-forwards`
- {ref}`network-zones`
- {ref}`network-ovn-peers`
- {ref}`network-load-balancers`

```{toctree}
:maxdepth: 1
:hidden:

OVNのセットアップ </howto/network_ovn_setup>
ルーティング関係を作成 </howto/network_ovn_peers>
ネットワークロードバランサーを設定 </howto/network_load_balancers>
```
