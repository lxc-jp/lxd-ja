# ネットワーク設定 <!-- Network configuration -->
LXD はネットワークの作成と管理をサポートしています。サポートしている
ブリッジの設定オプションは以下の一覧の通りです。
<!--
LXD supports creating and managing bridges, below is a list of the
configuration options supported for those bridges.
-->

この機能は API 拡張 "network" の一部だったことに注意してください。
<!--
Note that this feature was introduced as part of API extension "network".
-->

次のネームスペースの key/value 設定が現在サポートされています。
<!--
The key/value configuration is namespaced with the following namespaces
currently supported:
-->

 - `bridge` (L2 インタフェースの設定) <!-- (L2 interface configuration) -->
 - `fan` (Ubuntu FAN overlay に特有な設定) <!-- (configuration specific to the Ubuntu FAN overlay) -->
 - `tunnel` (ホスト間のトンネリングの設定) <!-- (cross-host tunneling configuration) -->
 - `ipv4` (L3 IPv4 設定) <!-- (L3 IPv4 configuration) -->
 - `ipv6` (L3 IPv6 設定) <!-- (L3 IPv6 configuration) -->
 - `dns` (DNS サーバと名前解決の設定) <!-- (DNS server and resolution configuration) -->
 - `raw` (raw の設定のファイルの内容) <!-- (raw configuration file content) -->
 - `user` (ユーザのメタデータに対する自由形式の key/value) <!-- (free form key/value for user metadata) -->

IP アドレスとサブネットは CIDR 形式 (`1.1.1.1/24` や `fd80:1234::1/64`) で指定することを想定しています。例外としてトンネルのローカルとリモートのアドレスは単なるアドレス (`1.1.1.1` や `fd80:1234::1`) を指定します。
<!--
It is expected that IP addresses and subnets are given using CIDR notation (`1.1.1.1/24` or `fd80:1234::1/64`).
The exception being tunnel local and remote addresses which are just plain addresses (`1.1.1.1` or `fd80:1234::1`).
-->

キー <!-- Key -->                            | 型 <!-- Type -->      | 条件 <!-- Condition -->             | デフォルト <!-- Default -->                   | 説明 <!-- Description -->
:--                             | :--       | :--                   | :--                       | :--
bridge.driver                   | string    | -                     | native                    | ブリッジのドライバ ("native" か "openvswitch") <!-- Bridge driver ("native" or "openvswitch") -->
bridge.external\_interfaces     | string    | -                     | -                         | ブリッジに含める未設定のネットワークインタフェースのカンマ区切りリスト <!-- Comma separate list of unconfigured network interfaces to include in the bridge -->
bridge.hwaddr                   | string    | -                     | -                         | ブリッジの MAC アドレス <!-- MAC address for the bridge -->
bridge.mode                     | string    | -                     | standard                  | ブリッジの稼働モード ("standard" か "fan") <!-- Bridge operation mode ("standard" or "fan") -->
bridge.mtu                      | integer   | -                     | 1500                      | ブリッジの MTU (tunnel か fan かでデフォルト値は変わります) <!-- Bridge MTU (default varies if tunnel or fan setup) -->
dns.domain                      | string    | -                     | lxd                       | DHCP のクライアントに広告し DNS の名前解決に使用するドメイン <!-- Domain to advertise to DHCP clients and use for DNS resolution -->
dns.mode                        | string    | -                     | managed                   | DNS の登録モード ("none" は DNS レコード無し、 "managed" は LXD が静的レコードを生成、 "dynamic" はクライアントがレコードを生成) <!-- DNS registration mode ("none" for no DNS record, "managed" for LXD generated static records or "dynamic" for client generated records) -->
fan.overlay\_subnet             | string    | fan mode              | 240.0.0.0/8               | FAN の overlay として使用するサブネット (CIDR 形式) <!-- Subnet to use as the overlay for the FAN (CIDR notation) -->
fan.type                        | string    | fan mode              | vxlan                     | FAN のトンネル・タイプ ("vxlan" か "ipip") <!-- The tunneling type for the FAN ("vxlan" or "ipip") -->
fan.underlay\_subnet            | string    | fan mode              | デフォルトゲートウェイのサブネット <!-- default gateway subnet -->    | FAN の underlay として使用するサブネット (CIDR 形式) <!-- Subnet to use as the underlay for the FAN (CIDR notation) -->
ipv4.address                    | string    | standard mode         | ランダムな未使用のサブネット <!-- random unused subnet -->      | ブリッジの IPv4 アドレス (CIDR 形式)。 IPv4 をオフにするには "none" 、新しいアドレスを生成するには "auto" を指定。 <!-- IPv4 address for the bridge (CIDR notation). Use "none" to turn off IPv4 or "auto" to generate a new one -->
ipv4.dhcp                       | boolean   | ipv4 address          | true                      | DHCP を使ってアドレスを割り当てるかどうか <!-- Whether to allocate addresses using DHCP -->
ipv4.dhcp.expiry                | string    | ipv4 dhcp             | 1h                        | DHCP リースの有効期限 <!-- When to expire DHCP leases -->
ipv4.dhcp.gateway               | string    | ipv4 dhcp             | ipv4.address              | サブネットのゲートウェイのアドレス <!-- Address of the gateway for the subnet -->
ipv4.dhcp.ranges                | string    | ipv4 dhcp             | 全てのアドレス <!-- all addresses -->             | DHCP に使用する IPv4 の範囲 (開始-終了 形式) のカンマ区切りリスト <!-- Comma separated list of IP ranges to use for DHCP (FIRST-LAST format) -->
ipv4.firewall                   | boolean   | ipv4 address          | true                      | このネットワークに対するファイアウォールのフィルタリングルールを生成するかどうか <!-- Whether to generate filtering firewall rules for this network -->
ipv4.nat                        | boolean   | ipv4 address          | false                     | NAT にするかどうか (未設定の場合はデフォルト値は true になりランダムな ipv4.address が生成されます) <!-- Whether to NAT (will default to true if unset and a random ipv4.address is generated) -->
ipv4.nat.order                  | string    | ipv4 address          | before                    | 必要な NAT のルールを既存のルールの前に追加するか後に追加するか <!-- Whether to add the required NAT rules before or after any pre-existing rules -->
ipv4.routes                     | string    | ipv4 address          | -                         | ブリッジへルーティングする追加の IPv4 CIDR サブネットのカンマ区切りリスト <!-- Comma separated list of additional IPv4 CIDR subnets to route to the bridge -->
ipv4.routing                    | boolean   | ipv4 address          | true                      | ブリッジの内外にトラフィックをルーティングするかどうか <!-- Whether to route traffic in and out of the bridge -->
ipv6.address                    | string    | standard mode         | ランダムな未使用のサブネット <!-- random unused subnet -->      | ブリッジの IPv6 アドレス (CIDR 形式)。 IPv6 をオフにするには "none" 、新しいアドレスを生成するには "auto" を指定。 <!-- IPv6 address for the bridge (CIDR notation). Use "none" to turn off IPv6 or "auto" to generate a new one -->
ipv6.dhcp                       | boolean   | ipv6 address          | true                      | DHCP 上で追加のネットワーク設定を提供するかどうか <!-- Whether to provide additional network configuration over DHCP -->
ipv6.dhcp.expiry                | string    | ipv6 dhcp             | 1h                        | DHCP リースの有効期限 <!-- When to expire DHCP leases -->
ipv6.dhcp.ranges                | string    | ipv6 stateful dhcp    | 全てのアドレス <!-- all addresses -->             | DHCP に使用する IPv6 の範囲 (開始-終了 形式) のカンマ区切りリスト <!-- Comma separated list of IPv6 ranges to use for DHCP (FIRST-LAST format) -->
ipv6.dhcp.stateful              | boolean   | ipv6 dhcp             | false                     | DHCP を使ってアドレスを割り当てるかどうか <!-- Whether to allocate addresses using DHCP -->
ipv6.firewall                   | boolean   | ipv6 address          | true                      | このネットワークに対するファイアウォールのフィルタリングルールを生成するかどうか <!-- Whether to generate filtering firewall rules for this network -->
ipv6.nat                        | boolean   | ipv6 address          | false                     | NAT にするかどうか (未設定の場合はデフォルト値は true になりランダムな ipv6.address が生成されます) <!-- Whether to NAT (will default to true if unset and a random ipv6.address is generated) -->
ipv6.nat.order                  | string    | ipv6 address          | before                    | 必要な NAT のルールを既存のルールの前に追加するか後に追加するか <!-- Whether to add the required NAT rules before or after any pre-existing rules -->
ipv6.routes                     | string    | ipv6 address          | -                         | ブリッジへルーティングする追加の IPv4 CIDR サブネットのカンマ区切りリスト <!-- Comma separated list of additional IPv6 CIDR subnets to route to the bridge -->
ipv6.routing                    | boolean   | ipv6 address          | true                      | ブリッジの内外にトラフィックをルーティングするかどうか <!-- Whether to route traffic in and out of the bridge -->
raw.dnsmasq                     | string    | -                     | -                         | 設定に追加する dnsmasq の設定 <!-- Additional dnsmasq configuration to append to the configuration -->
tunnel.NAME.group               | string    | vxlan                 | 239.0.0.1                 | vxlan のマルチキャスト設定 (local と remote が未設定の場合に使われます) <!-- Multicast address for vxlan (used if local and remote aren't set) -->
tunnel.NAME.id                  | integer   | vxlan                 | 0                         | vxlan トンネルに使用するトンネル ID <!-- Specific tunnel ID to use for the vxlan tunnel -->
tunnel.NAME.interface           | string    | vxlan                 | -                         | トンネルに使用するホスト・インタフェース <!-- Specific host interface to use for the tunnel -->
tunnel.NAME.local               | string    | gre or vxlan          | -                         | トンネルに使用するローカルアドレス (マルチキャスト vxlan の場合は不要) <!-- Local address for the tunnel (not necessary for multicast vxlan) -->
tunnel.NAME.port                | integer   | vxlan                 | 0                         | vxlan トンネルに使用するポート <!-- Specific port to use for the vxlan tunnel -->
tunnel.NAME.protocol            | string    | standard mode         | -                         | トンネリングのプロトコル ("vxlan" か "gre") <!-- Tunneling protocol ("vxlan" or "gre") -->
tunnel.NAME.remote              | string    | gre or vxlan          | -                         | トンネルに使用するリモートアドレス (マルチキャスト vxlan の場合は不要) <!-- Remote address for the tunnel (not necessary for multicast vxlan) -->
tunnel.NAME.ttl                 | integer   | vxlan                 | 1                         | マルチキャストルーティングトポロジーに使用する固有の TTL <!-- Specific TTL to use for multicast routing topologies -->


これらのキーは lxc コマンドで以下のように設定できます。
<!--
Those keys can be set using the lxc tool with:
-->

```bash
lxc network set <network> <key> <value>
```
