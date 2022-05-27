---
discourse: 7322
---

(network-bridge)=
# ブリッジネットワーク

LXD でのネットワークの設定タイプの 1 つとして、 LXD はネットワークブリッジの作成と管理をサポートしています。
<!-- Include start bridge intro -->
ネットワークブリッジはインスタンス NIC が接続できる仮想的な L2 イーサネットスイッチを作成し、インスタンスが他のインスタンスやホストと通信できるようにします。
LXD のブリッジは下層のネイティブな Linux のブリッジと Open vSwitch を利用できます。
<!-- Include end bridge intro -->

`bridge` ネットワークはそれを利用する複数のインスタンスを接続する L2 ブリッジを作成しそれらのインスタンスを単一の L2 ネットワークセグメントにします。
LXD で作成されたブリッジは "managed" です。
つまり、ブリッジインタフェース自体を作成するのに加えて、LXD さらに DHCP、 IPv6 ルート広告と DNS サービスを提供するローカルの `dnsmasq` プロセスをセットアップします。
デフォルトではブリッジに対して NAT も行います。

LXD ブリッジネットワークでファイアウォールを設定するための手順については {ref}`network-bridge-firewall` を参照してください。

## IPv6 プリフィクスサイズ

ブリッジネットワークで IPv6 を使用する場合、 64 のプリフィクスサイズを使用するべきです。

より大きなサブネット (つまり 64 より小さいプリフィクスを使用する) も正常に動くはずですが、通常それらは {abbr}`SLAAC (Stateless Address Auto-configuration)` には役立ちません。

より小さなサブネットも (IPv6 の割当にはステートフル DHCPv6 を使用する場合) 理論上は可能ですが、 `dnsmasq` に適切にサポートされていないので問題が起きるかもしれません。より小さなサブネットを作らなければならない場合は、静的割当を使うか別のルータ広告デーモンを使用してください。

(network-bridge-options)=
## 設定オプション

ブリッジネットワークでは現在以下の設定キーネームスペースがサポートされています。

 - `bgp` (BGP ピア設定)
 - `bridge` (L2 インタフェースの設定)
 - `dns` (DNS サーバと名前解決の設定)
 - `fan` (Ubuntu FAN overlay に特有な設定)
 - `ipv4` (L3 IPv4 設定)
 - `ipv6` (L3 IPv6 設定)
 - `maas` (MAAS ネットワーク識別)
 - `security` (ネットワーク ACL 設定)
 - `raw` (raw の設定のファイルの内容)
 - `tunnel` (ホスト間のトンネリングの設定)
 - `user` (key/value の自由形式のユーザメタデータ)

```{note}
{{note_ip_addresses_CIDR}}
```

ブリッジネットワークタイプには以下の設定オプションがあります。

キー                                 | 型        | 条件                    | デフォルト          | 説明
:--                                  | :--       | :--                     | :--                 | :--
bgp.peers.NAME.address               | string    | bgp server              | -                   | ピアのアドレス (IPv4 か IPv6)
bgp.peers.NAME.asn                   | integer   | bgp server              | -                   | ピアの AS 番号
bgp.peers.NAME.password              | string    | bgp server              | - (パスワード無し)  | ピアのセッションパスワード（省略可能）
bgp.ipv4.nexthop                     | string    | bgp server              | ローカルアドレス    | 広告されたプリフィクスの next-hop をオーバーライド
bgp.ipv6.nexthop                     | string    | bgp server              | ローカルアドレス    | 広告されたプリフィクスの next-hop をオーバーライド
bridge.driver                        | string    | -                       | native              | ブリッジのドライバ ("native" か "openvswitch")
bridge.external\_interfaces          | string    | -                       | -                   | ブリッジに含める未設定のネットワークインタフェースのカンマ区切りリスト
bridge.hwaddr                        | string    | -                       | -                   | ブリッジの MAC アドレス
bridge.mode                          | string    | -                       | standard            | ブリッジの稼働モード ("standard" か "fan")
bridge.mtu                           | integer   | -                       | 1500                | ブリッジの MTU (tunnel か fan かでデフォルト値は変わります)
dns.domain                           | string    | -                       | lxd                 | DHCP のクライアントに広告し DNS の名前解決に使用するドメイン
dns.mode                             | string    | -                       | managed             | DNS の登録モード ("none" は DNS レコード無し、 "managed" は LXD が静的レコードを生成、 "dynamic" はクライアントがレコードを生成)
dns.search                           | string    | -                       | -                   | 完全なドメインサーチのカンマ区切りリスト（デフォルトは `dns.domain` の値）
dns.zone.forward                     | string    | -                       | managed             | 正引き DNS レコード用の DNS ゾーン名
dns.zone.reverse.ipv4                | string    | -                       | managed             | IPv4 逆引き DNS レコード用の DNS ゾーン名
dns.zone.reverse.ipv6                | string    | -                       | managed             | IPv6 逆引き DNS レコード用の DNS ゾーン名
fan.overlay\_subnet                  | string    | ファンモード            | 240.0.0.0/8         | FAN の overlay として使用するサブネット (CIDR 形式)
fan.type                             | string    | ファンモード            | vxlan               | FAN のトンネル・タイプ ("vxlan" か "ipip")
fan.underlay\_subnet                 | string    | ファンモード            | 自動（作成時のみ）  | FAN の underlay として使用するサブネット (CIDR 形式)。デフォルトのゲートウェイサブネットを使うには "auto" を指定。
ipv4.address                         | string    | 標準モード              | 自動（作成時のみ）  | ブリッジの IPv4 アドレス (CIDR 形式)。 IPv4 をオフにするには "none" 、新しいランダムな未使用のサブネットを生成するには "auto" を指定。
ipv4.dhcp                            | boolean   | ipv4 アドレス           | true                | DHCP を使ってアドレスを割り当てるかどうか
ipv4.dhcp.expiry                     | string    | ipv4 dhcp               | 1h                  | DHCP リースの有効期限
ipv4.dhcp.gateway                    | string    | ipv4 dhcp               | ipv4.address        | サブネットのゲートウェイのアドレス
ipv4.dhcp.ranges                     | string    | ipv4 dhcp               | 全てのアドレス      | DHCP に使用する IPv4 の範囲 (開始-終了 形式) のカンマ区切りリスト
ipv4.firewall                        | boolean   | ipv4 アドレス           | true                | このネットワークに対するファイアウォールのフィルタリングルールを生成するかどうか
ipv4.nat                             | boolean   | ipv4 アドレス           | false               | NAT にするかどうか（通常のブリッジではデフォルト値は true で ipv4.address が生成され、fan ブリッジでは常にデフォルト値は true になります）
ipv4.nat.address                     | string    | ipv4 アドレス           | -                   | ブリッジからの送信時に使うソースアドレス
ipv4.nat.order                       | string    | ipv4 アドレス           | before              | 必要な NAT のルールを既存のルールの前に追加するか後に追加するか
ipv4.ovn.ranges                      | string    | -                       | -                   | 子供の OVN ネットワークルーターに使用する IPv4 アドレスの範囲（開始-終了 形式）のカンマ区切りリスト
ipv4.routes                          | string    | ipv4 アドレス           | -                   | ブリッジへルーティングする追加の IPv4 CIDR サブネットのカンマ区切りリスト
ipv4.routing                         | boolean   | ipv4 アドレス           | true                | ブリッジの内外にトラフィックをルーティングするかどうか
ipv6.address                         | string    | 標準モード              | 自動（作成時のみ）  | ブリッジの IPv6 アドレス (CIDR 形式)。 IPv6 をオフにするには "none" 、新しいランダムな未使用のサブネットを生成するには "auto" を指定。
ipv6.dhcp                            | boolean   | ipv6 アドレス           | true                | DHCP 上で追加のネットワーク設定を提供するかどうか
ipv6.dhcp.expiry                     | string    | ipv6 dhcp               | 1h                  | DHCP リースの有効期限
ipv6.dhcp.ranges                     | string    | ipv6 ステートフル dhcp  | 全てのアドレス      | DHCP に使用する IPv6 の範囲 (開始-終了 形式) のカンマ区切りリスト
ipv6.dhcp.stateful                   | boolean   | ipv6 dhcp               | false               | DHCP を使ってアドレスを割り当てるかどうか
ipv6.firewall                        | boolean   | ipv6 アドレス           | true                | このネットワークに対するファイアウォールのフィルタリングルールを生成するかどうか
ipv6.nat                             | boolean   | ipv6 アドレス           | false               | NAT にするかどうか (未設定の場合はデフォルト値は true になりランダムな ipv6.address が生成されます)
ipv6.nat.address                     | string    | ipv6 アドレス           | -                   | ブリッジからの送信時に使うソースアドレス
ipv6.nat.order                       | string    | ipv6 アドレス           | before              | 必要な NAT のルールを既存のルールの前に追加するか後に追加するか
ipv6.ovn.ranges                      | string    | -                       | -                   | 子供の OVN ネットワークルーターに使用する IPv6 アドレスの範囲（開始-終了 形式) のカンマ区切りリスト
ipv6.routes                          | string    | ipv6 アドレス           | -                   | ブリッジへルーティングする追加の IPv4 CIDR サブネットのカンマ区切りリスト
ipv6.routing                         | boolean   | ipv6 アドレス           | true                | ブリッジの内外にトラフィックをルーティングするかどうか
maas.subnet.ipv4                     | string    | ipv4 アドレス           | -                   | インスタンスを登録する MAAS IPv4 サブネット (NIC で `network` プロパティを使う場合に有効)
maas.subnet.ipv6                     | string    | ipv6 アドレス           | -                   | インスタンスを登録する MAAS IPv6 サブネット (NIC で `network` プロパティを使う場合に有効)
raw.dnsmasq                          | string    | -                       | -                   | 設定に追加する dnsmasq の設定ファイル
security.acls                        | string    | -                       | -                   | このネットワークに接続されたNICに適用するカンマ区切りのネットワークACL（{ref}`network-acls-bridge-limitations`参照）
security.acls.default.ingress.action | string    | security.acls           | reject              | どの ACL ルールにもマッチしない ingress トラフィックに使うアクション
security.acls.default.egress.action  | string    | security.acls           | reject              | どの ACL ルールにもマッチしない egress トラフィックに使うアクション
security.acls.default.ingress.logged | boolean   | security.acls           | false               | どの ACL ルールにもマッチしない ingress トラフィックをログ出力するかどうか
security.acls.default.egress.logged  | boolean   | security.acls           | false               | どの ACL ルールにもマッチしない egress トラフィックをログ出力するかどうか
tunnel.NAME.group                    | string    | vxlan                   | 239.0.0.1           | vxlan のマルチキャスト設定 (local と remote が未設定の場合に使われます)
tunnel.NAME.id                       | integer   | vxlan                   | 0                   | vxlan トンネルに使用するトンネル ID
tunnel.NAME.interface                | string    | vxlan                   | -                   | トンネルに使用するホスト・インタフェース
tunnel.NAME.local                    | string    | gre か vxlan            | -                   | トンネルに使用するローカルアドレス (マルチキャスト vxlan の場合は不要)
tunnel.NAME.port                     | integer   | vxlan                   | 0                   | vxlan トンネルに使用するポート
tunnel.NAME.protocol                 | string    | 標準モード              | -                   | トンネリングのプロトコル ("vxlan" か "gre")
tunnel.NAME.remote                   | string    | gre か vxlan            | -                   | トンネルに使用するリモートアドレス (マルチキャスト vxlan の場合は不要)
tunnel.NAME.ttl                      | integer   | vxlan                   | 1                   | マルチキャストルーティングトポロジーに使用する固有の TTL
user.*                               | string    | -                       | -                   | ユーザ指定の自由形式のキー／バリューペア

(network-bridge-features)=
## サポートされている機能

ブリッジネットワークタイプでは以下の機能がサポートされています。

- {ref}`network-acls`
- {ref}`network-forwards`
- {ref}`network-zones`
- {ref}`network-bgp`
- {ref}`network-bridge-resolved`


```{toctree}
:maxdepth: 1
:hidden:

systemd-resolved との統合 </howto/network_bridge_resolved>
ファイアウォールの設定 </howto/network_bridge_firewalld>
```
