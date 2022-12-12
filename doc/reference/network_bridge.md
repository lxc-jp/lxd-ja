---
discourse: 7322
---

(network-bridge)=
# ブリッジネットワーク

LXDでのネットワークの設定タイプの1つとして、LXDはネットワークブリッジの作成と管理をサポートしています。
<!-- Include start bridge intro -->
ネットワークブリッジはインスタンスNICが接続できる仮想的なL2イーサネットスイッチを作成し、インスタンスが他のインスタンスやホストと通信できるようにします。
LXDのブリッジは下層のネイティブなLinuxのブリッジとOpen vSwitchを利用できます。
<!-- Include end bridge intro -->

`bridge`ネットワークはそれを利用する複数のインスタンスを接続するL2ブリッジを作成しそれらのインスタンスを単一のL2ネットワークセグメントにします。
LXD で作成されたブリッジは"managed"です。
つまり、ブリッジインタフェース自体を作成するのに加えて、LXDさらにDHCP、IPv6ルート広告とDNSサービスを提供するローカルの`dnsmasq`プロセスをセットアップします。
デフォルトではブリッジに対してNATも行います。

LXDブリッジネットワークでファイアウォールを設定するための手順については{ref}`network-bridge-firewall`を参照してください。

<!-- Include start MAC identifier note -->

```{note}
静的なDHCP割当はMACアドレスをDHCP識別子として使用するクライアントに依存します。
この方法はインスタンスをコピーする際に衝突するリースを回避し、静的に割り当てられたリースが正しく動くようにします。
```

<!-- Include end MAC identifier note -->

## IPv6プリフィクスサイズ

ブリッジネットワークでIPv6を使用する場合、64のプリフィクスサイズを使用するべきです。

より大きなサブネット(つまり64より小さいプリフィクスを使用する)も正常に動くはずですが、通常それらは{abbr}`SLAAC (Stateless Address Auto-configuration)`には役立ちません。

より小さなサブネットも(IPv6の割当にはステートフルDHCPv6を使用する場合) 理論上は可能ですが、`dnsmasq`に適切にサポートされていないので問題が起きるかもしれません。より小さなサブネットを作らなければならない場合は、静的割当を使うか別のルータ広告デーモンを使用してください。

(network-bridge-options)=
## 設定オプション

`bridge`ネットワークタイプでは現在以下の設定キーネームスペースがサポートされています。

- `bgp` (BGPピア設定)
- `bridge` (L2インタフェースの設定)
- `dns` (DNSサーバと名前解決の設定)
- `fan` (Ubuntu FAN overlayに特有な設定)
- `ipv4` (L3 IPv4設定)
- `ipv6` (L3 IPv6設定)
- `maas` (MAASネットワーク識別)
- `security` (ネットワークACL設定)
- `raw` (rawの設定のファイルの内容)
- `tunnel` (ホスト間のトンネリングの設定)
- `user` (key/valueの自由形式のユーザメタデータ)

```{note}
{{note_ip_addresses_CIDR}}
```

`bridge`ネットワークタイプには以下の設定オプションがあります。

キー                                   | 型      | 条件                  | デフォルト         | 説明
:--                                    | :--     | :--                   | :--                | :--
`bgp.peers.NAME.address`               | string  | BGPサーバ             | -                  | ピアのアドレス(IPv4かIPv6)
`bgp.peers.NAME.asn`                   | integer | BGPサーバ             | -                  | ピアのAS番号
`bgp.peers.NAME.password`              | string  | BGPサーバ             | - (パスワード無し) | ピアのセッションパスワード(省略可能)
`bgp.peers.NAME.holdtime`              | integer | BGPサーバ             | `180`              | ピアセッションホールドタイム(秒で指定、省略可能)
`bgp.ipv4.nexthop`                     | string  | BGPサーバ             | ローカルアドレス   | 広告されたプリフィクスのnext-hopをオーバーライド
`bgp.ipv6.nexthop`                     | string  | BGPサーバ             | ローカルアドレス   | 広告されたプリフィクスのnext-hopをオーバーライド
`bridge.driver`                        | string  | -                     | `native`           | ブリッジのドライバ(`native`か`openvswitch`)
`bridge.external_interfaces`           | string  | -                     | -                  | ブリッジに含める未設定のネットワークインタフェースのカンマ区切りリスト
`bridge.hwaddr`                        | string  | -                     | -                  | ブリッジのMACアドレス
`bridge.mode`                          | string  | -                     | `standard`         | ブリッジの稼働モード(`standard`か`fan`)
`bridge.mtu`                           | integer | -                     | `1500`             | ブリッジのMTU(tunnelかfanかでデフォルト値は変わります)
`dns.domain`                           | string  | -                     | `lxd`              | DHCPのクライアントに広告しDNSの名前解決に使用するドメイン
`dns.mode`                             | string  | -                     | `managed`          | DNSの登録モード(`none`はDNSレコード無し、`managed`はLXDが静的レコードを生成、`dynamic`はクライアントがレコードを生成)
`dns.search`                           | string  | -                     | -                  | 完全なドメインサーチのカンマ区切りリスト(デフォルトは`dns.domain`の値)
`dns.zone.forward`                     | string  | -                     | `managed`          | 正引きDNSレコード用のDNSゾーン名のカンマ区切りリスト
`dns.zone.reverse.ipv4`                | string  | -                     | `managed`          | IPv4逆引きDNSレコード用のDNSゾーン名
`dns.zone.reverse.ipv6`                | string  | -                     | `managed`          | IPv6逆引きDNSレコード用のDNSゾーン名
`fan.overlay_subnet`                   | string  | ファンモード          | `240.0.0.0/8`      | FANのoverlayとして使用するサブネット(CIDR形式)
`fan.type`                             | string  | ファンモード          | `vxlan`            | FANのトンネル・タイプ(`vxlan`か`ipip`)
`fan.underlay_subnet`                  | string  | ファンモード          | `auto`(作成時のみ) | FANのunderlayとして使用するサブネット(CIDR形式)。デフォルトのゲートウェイサブネットを使うには`auto`を指定。
`ipv4.address`                         | string  | 標準モード            | `auto`(作成時のみ) | ブリッジのIPv4アドレス(CIDR形式)。IPv4をオフにするには`none`、新しいランダムな未使用のサブネットを生成するには`auto`を指定。
`ipv4.dhcp`                            | bool    | IPv4アドレス          | `true`             | DHCPを使ってアドレスを割り当てるかどうか
`ipv4.dhcp.expiry`                     | string  | IPv4 DHCP             | `1h`               | DHCPリースの有効期限
`ipv4.dhcp.gateway`                    | string  | IPv4 DHCP             | IPv4アドレス       | サブネットのゲートウェイのアドレス
`ipv4.dhcp.ranges`                     | string  | IPv4 DHCP             | 全てのアドレス     | DHCPに使用するIPv4の範囲(開始-終了の形式)のカンマ区切りリスト
`ipv4.firewall`                        | bool    | IPv4アドレス          | `true`             | このネットワークに対するファイアウォールのフィルタリングルールを生成するかどうか
`ipv4.nat`                             | bool    | IPv4アドレス          | `false`            | NATにするかどうか(通常のブリッジではデフォルト値は`true`で`ipv4.address`が生成され、fanブリッジでは常にデフォルト値は`true`になります)
`ipv4.nat.address`                     | string  | IPv4アドレス          | -                  | ブリッジからの送信時に使うソースアドレス
`ipv4.nat.order`                       | string  | IPv4アドレス          | `before`           | 必要なNATのルールを既存のルールの前に追加するか後に追加するか
`ipv4.ovn.ranges`                      | string  | -                     | -                  | 子供のOVNネットワークルーターに使用するIPv4アドレスの範囲(開始-終了の形式)のカンマ区切りリスト
`ipv4.routes`                          | string  | IPv4アドレス          | -                  | ブリッジへルーティングする追加のIPv4 CIDRサブネットのカンマ区切りリスト
`ipv4.routing`                         | bool    | IPv4アドレス          | `true`             | ブリッジの内外にトラフィックをルーティングするかどうか
`ipv6.address`                         | string  | 標準モード            | `auto`(作成時のみ) | ブリッジのIPv6アドレス(CIDR形式)。IPv6をオフにするには`none`、新しいランダムな未使用のサブネットを生成するには`auto`を指定。
`ipv6.dhcp`                            | bool    | IPv6アドレス          | `true`             | DHCP上で追加のネットワーク設定を提供するかどうか
`ipv6.dhcp.expiry`                     | string  | IPv6 DHCP             | `1h`               | DHCPリースの有効期限
`ipv6.dhcp.ranges`                     | string  | IPv6ステートフル DHCP | 全てのアドレス     | DHCPに使用するIPv6の範囲(開始-終了の形式)のカンマ区切りリスト
`ipv6.dhcp.stateful`                   | bool    | IPv6 DHCP             | `false`            | DHCP を使ってアドレスを割り当てるかどうか
`ipv6.firewall`                        | bool    | IPv6アドレス          | `true`             | このネットワークに対するファイアウォールのフィルタリングルールを生成するかどうか
`ipv6.nat`                             | bool    | IPv6アドレス          | `false`            | NATにするかどうか(未設定の場合はデフォルト値は`true`になりランダムな`ipv6.address`が生成されます)
`ipv6.nat.address`                     | string  | IPv6アドレス          | -                  | ブリッジからの送信時に使うソースアドレス
`ipv6.nat.order`                       | string  | IPv6アドレス          | `before`           | 必要なNATのルールを既存のルールの前に追加するか後に追加するか
`ipv6.ovn.ranges`                      | string  | -                     | -                  | 子供のOVNネットワークルーターに使用するIPv6アドレスの範囲(開始-終了の形式) のカンマ区切りリスト
`ipv6.routes`                          | string  | IPv6アドレス          | -                  | ブリッジへルーティングする追加のIPv4 CIDRサブネットのカンマ区切りリスト
`ipv6.routing`                         | bool    | IPv6アドレス          | `true`             | ブリッジの内外にトラフィックをルーティングするかどうか
`maas.subnet.ipv4`                     | string  | IPv4アドレス          | -                  | インスタンスを登録するMAAS IPv4サブネット(NICで`network`プロパティを使う場合に有効)
`maas.subnet.ipv6`                     | string  | IPv6アドレス          | -                  | インスタンスを登録するMAAS IPv6サブネット(NICで`network`プロパティを使う場合に有効)
`raw.dnsmasq`                          | string  | -                     | -                  | 設定に追加する`dnsmasq`の設定ファイル
`security.acls`                        | string  | -                     | -                  | このネットワークに接続されたNICに適用するカンマ区切りのネットワークACL({ref}`network-acls-bridge-limitations`参照)
`security.acls.default.egress.action`  | string  | `security.acls`       | `reject`           | どのACLルールにもマッチしない外向きトラフィックに使うアクション
`security.acls.default.egress.logged`  | bool    | `security.acls`       | `false`            | どのACLルールにもマッチしない外向きトラフィックをログ出力するかどうか
`security.acls.default.ingress.action` | string  | `security.acls`       | `reject`           | どのACLルールにもマッチしない内向きトラフィックに使うアクション
`security.acls.default.ingress.logged` | bool    | `security.acls`       | `false`            | どのACLルールにもマッチしない内向きトラフィックをログ出力するかどうか
`tunnel.NAME.group`                    | string  | `vxlan`               | `239.0.0.1`        | `vxlan`のマルチキャスト設定(localとremoteが未設定の場合に使われます)
`tunnel.NAME.id`                       | integer | `vxlan`               | `0`                | `vxlan`トンネルに使用するトンネルID
`tunnel.NAME.interface`                | string  | `vxlan`               | -                  | トンネルに使用するホスト・インタフェース
`tunnel.NAME.local`                    | string  | `gre`か`vxlan`        | -                  | トンネルに使用するローカルアドレス(マルチキャストの場合は不要)
`tunnel.NAME.port`                     | integer | `vxlan`               | `0`                | `vxlan`トンネルに使用するポート
`tunnel.NAME.protocol`                 | string  | 標準モード            | -                  | トンネリングのプロトコル(`vxlan`か`gre`)
`tunnel.NAME.remote`                   | string  | `gre`か`vxlan`        | -                  | トンネルに使用するリモートアドレス(マルチキャストの場合は不要)
`tunnel.NAME.ttl`                      | integer | `vxlan`               | `1`                | マルチキャストルーティングトポロジーに使用する固有の TTL
`user.*`                               | string  | -                     | -                  | ユーザ指定の自由形式のキー／バリューペア

(network-bridge-features)=
## サポートされている機能

`bridge`ネットワークタイプでは以下の機能がサポートされています。

- {ref}`network-acls`
- {ref}`network-forwards`
- {ref}`network-zones`
- {ref}`network-bgp`
- [`systemd-resolved`と統合するには](network-bridge-resolved)

```{toctree}
:maxdepth: 1
:hidden:

resolvedとの統合 </howto/network_bridge_resolved>
ファイアウォールの設定 </howto/network_bridge_firewalld>
```
