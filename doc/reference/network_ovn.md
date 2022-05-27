---
discourse: 11033
---

(network-ovn)=
# OVN ネットワーク

<!-- Include start OVN intro -->
{abbr}`OVN (Open Virtual Network)` は仮想ネットワーク抽象化をサポートするソフトウェアで定義されたネットワークシステムです。
あなた自身のプライベートクラウドを構築するのに使用できます。
詳細は https://www.ovn.org/ をご参照ください。
<!-- Include end OVN intro -->

`ovn` ネットワークタイプは OVN {abbr}`SDN (software-defined networking)` を使って論理的なネットワークの作成を可能にします。
この種のネットワークは複数の個別のネットワーク内で同じ論理ネットワークのサブネットを使うような検証環境やマルチテナントの環境で便利です。

LXD の OVN ネットワークはより広いネットワークへの外向きのアクセスを可能にするため既存の管理された {ref}`network-bridge` や {ref}`network-physical` に接続できます。
デフォルトでは、 OVN 論理ネットワークからの全ての接続はアップリンクのネットワークによって割り当てられた IP に NAT されます。

OVN ネットワークをセットアップする基本的な手順については {ref}`network-ovn-setup` をご参照ください。

(network-ovn-options)=
## 設定オプション

`ovn` ネットワークタイプでは現在以下の設定キーネームスペースがサポートされています。

 - `bridge` (L2 インタフェースの設定)
 - `dns` (DNS サーバと名前解決の設定)
 - `ipv4` (L3 IPv4 設定)
 - `ipv6` (L3 IPv6 設定)
 - `security` (ネットワーク ACL 設定)
 - `user` (key/value の自由形式のユーザメタデータ)

```{note}
{{note_ip_addresses_CIDR}}
```

`ovn` ネットワークタイプには以下の設定オプションがあります。

キー                                 | 型        | 条件             | デフォルト                  | 説明
:--                                  | :--       | :--              | :--                         | :--
network                              | string    | -                | -                           | 外部ネットワークへの外向きのアクセスに使うアップリンクのネットワーク
bridge.hwaddr                        | string    | -                | -                           | ブリッジの MAC アドレス
bridge.mtu                           | integer   | -                | 1442                        | ブリッジの MTU (デフォルトではホストからホストへの geneve トンネルを許可します)
dns.domain                           | string    | -                | lxd                         | DHCP のクライアントに広告し DNS の名前解決に使用するドメイン
dns.search                           | string    | -                | -                           | 完全なドメインサーチのカンマ区切りリスト（デフォルトは `dns.domain` の値）
dns.zone.forward                     | string    | -                | -                           | 正引き DNS レコード用の DNS ゾーン名
dns.zone.reverse.ipv4                | string    | -                | -                           | IPv4 逆引き DNS レコード用の DNS ゾーン名
dns.zone.reverse.ipv6                | string    | -                | -                           | IPv6 逆引き DNS レコード用の DNS ゾーン名
ipv4.address                         | string    | 標準モード       | 自動（作成時のみ）          | ブリッジの IPv4 アドレス (CIDR 形式)。 IPv4 をオフにするには "none" 、新しいランダムな未使用のサブネットを生成するには "auto" を指定。
ipv4.dhcp                            | boolean   | ipv4 アドレス    | true                        | DHCP を使ってアドレスを割り当てるかどうか
ipv4.nat                             | boolean   | ipv4 アドレス    | false                       | NAT するかどうか（ipv4.address が未設定の場合デフォルト値は true でランダムな ipv4.address が生成されます）
ipv4.nat.address                     | string    | ipv4 アドレス    | -                           | ネットワークからの外向きトラフィックに使用されるソースアドレス (アップリンクに `ovn.ingress_mode=routed` が必要)
ipv6.address                         | string    | 標準モード       | 自動（作成時のみ）          | ブリッジの IPv6 アドレス (CIDR 形式)。 IPv6 をオフにするには "none" 、新しいランダムな未使用のサブネットを生成するには "auto" を指定。
ipv6.dhcp                            | boolean   | ipv6 アドレス    | true                        | DHCP 上に追加のネットワーク設定を提供するかどうか
ipv6.dhcp.stateful                   | boolean   | ipv6 dhcp        | false                       | DHCP を使ってアドレスを割り当てるかどうか
ipv6.nat                             | boolean   | ipv6 アドレス    | false                       | NAT するかどうか（ipv6.address が未設定の場合デフォルト値は true でランダムな ipv6.address が生成されます）
ipv6.nat.address                     | string    | ipv6 アドレス    | -                           | ネットワークからの外向きトラフィックに使用されるソースアドレス (アップリンクに `ovn.ingress_mode=routed` が必要)
security.acls                        | string    | -                | -                           | このネットワークに接続する NIC に適用するネットワーク ACL のカンマ区切りリスト
security.acls.default.ingress.action | string    | security.acls    | reject                      | どの ACL ルールにもマッチしない ingress トラフィックに使うアクション
security.acls.default.egress.action  | string    | security.acls    | reject                      | どの ACL ルールにもマッチしない egress トラフィックに使うアクション
security.acls.default.ingress.logged | boolean   | security.acls    | false                       | どの ACL ルールにもマッチしない ingress トラフィックをログ出力するかどうか
security.acls.default.egress.logged  | boolean   | security.acls    | false                       | どの ACL ルールにもマッチしない egress トラフィックをログ出力するかどうか
user.*                               | string    | -                | -                           | ユーザ指定の自由形式のキー／バリューペア

(network-ovn-features)=
## サポートされている機能

`ovn` ネットワークタイプでは以下の機能がサポートされています。

- {ref}`network-acls`
- {ref}`network-forwards`
- {ref}`network-zones`
- {ref}`network-ovn-peers`

```{toctree}
:maxdepth: 1
:hidden:

OVN のセットアップ </howto/network_ovn_setup>
ルーティング関係を作成 </howto/network_ovn_peers>
```
