# ネットワーク設定
<!-- Network configuration -->

LXDは以下のネットワークタイプをサポートします。
<!--
LXD supports the following network types:
-->

 - [bridge](#network-bridge): インスタンスを接続する L2 ブリッジを作成（ローカルの DHCP と DNS を提供可能）。これがデフォルトです。 <!-- Creates an L2 bridge for connecting instances to (can provide local DHCP and DNS). This is the default. -->
 - [macvlan](#network-macvlan): インスタンスを親の macvlan インターフェースに接続する際に使用するプリセットの設定を提供。 <!-- Provides preset configuration to use when connecting instances to a parent macvlan interface. -->
 - [sriov](#network-sriov): インスタンスを親の SR-IOV インターフェースに接続する際に使用するプリセットの設定を提供。 <!-- Provides preset configuration to use when connecting instances to a parent SR-IOV interface. -->
 - [ovn](#network-ovn): OVN SDN (software defined network) システムを使って論理的なネットワークを作成。 <!-- Creates a logical network using the OVN software defined networking system. -->
 - [physical](#network-physical): OVN ネットワークを親のインターフェースに接続する際に使用するプリセットの設定を提供。 <!-- Provides preset configuration to use when connecting OVN networks to a parent interface. -->

希望するタイプは以下のように `--type` 引数で指定できます。
<!--
The desired type can be specified using the `-\-type` argument, e.g.
-->

```bash
lxc network create <name> --type=bridge [options...]
```

`--type` 引数が指定されない場合は、デフォルトの `bridge` が使用されます。
<!--
If no `-\-type` argument is specified, the default type of `bridge` is used.
-->

設定キーは現状では全てのネットワークタイプでサポートされている以下のネームスペースによって名前空間が分けられています。
<!--
The configuration keys are namespaced with the following namespaces currently supported for all network types:
-->

 - `maas` (MAAS ネットワーク識別) <!-- (MAAS network identification) -->
 - `user` (ユーザーのメタデータに対する自由形式の key/value) <!-- (free form key/value for user metadata) -->

## <a name="network-bridge"></a> ネットワーク: ブリッジ <!-- network: bridge -->

LXD でのネットワークの設定タイプの 1 つとして、 LXD はネットワークブリッジの作成と管理をサポートしています。
LXD のブリッジは下層のネイティブな Linux のブリッジと Open vSwitch を利用できます。
<!--
As one of the possible network configuration types under LXD, LXD supports creating and managing network bridges.
LXD bridges can leverage underlying native Linux bridges and Open vSwitch.
-->

LXD のブリッジの作成と管理は `lxc network` コマンドで行えます。
LXD で作成されたブリッジはデフォルトでは "managed" です。
というのは LXD はさらにローカルの `dnsmasq` DHCP サーバをセットアップし、希望すれば (これがデフォルトです) ブリッジに対して NAT も行います。
<!--
Creation and management of LXD bridges is performed via the `lxc network` command.
A bridge created by LXD is by default "managed" which means that LXD also will additionally set up a local `dnsmasq`
DHCP server and if desired also perform NAT for the bridge (this is the default.)
-->

ブリッジが LXD に管理されているときは、 `bridge` ネームスペースを使って設定値を変更できます。
<!--
When a bridge is managed by LXD, configuration values under the `bridge` namespace can be used to configure it.
-->

さらに、 LXD は既存の Linux ブリッジを利用することも出来ます。
この場合、ブリッジは `lxc network` で作成する必要はなく、インスタンスかプロファイルのデバイス設定で下記のように単に参照できます。
<!--
Additionally, LXD can utilize a pre-existing Linux bridge. In this case, the bridge does not need to be created via
`lxc network` and can simply be referenced in an instance or profile device configuration as follows:
-->

```
devices:
  eth0:
     name: eth0
     nictype: bridged
     parent: br0
     type: nic
```

ネットワークフォワード: <!--Network forwards: -->

ブリッジのネットワークサポートは [network forwards](network-forwards.md#network-bridge) 参照。
<!--
Bridge networks support [network forwards](network-forwards.md#network-bridge).
-->

ネットワークの設定プロパティー: <!-- Network configuration properties: -->

LXD のネットワークの設定項目の完全なリストは以下の通りです。
<!--
A complete list of configuration settings for LXD networks can be found below.
-->

ブリッジネットワークでは以下の設定キーネームスペースが現状サポートされています。
<!--
The following configuration key namespaces are currently supported for bridge networks:
-->

 - `bridge` (L2 インタフェースの設定) <!-- (L2 interface configuration) -->
 - `fan` (Ubuntu FAN overlay に特有な設定) <!-- (configuration specific to the Ubuntu FAN overlay) -->
 - `tunnel` (ホスト間のトンネリングの設定) <!-- (cross-host tunneling configuration) -->
 - `ipv4` (L3 IPv4 設定) <!-- (L3 IPv4 configuration) -->
 - `ipv6` (L3 IPv6 設定) <!-- (L3 IPv6 configuration) -->
 - `dns` (DNS サーバと名前解決の設定) <!-- (DNS server and resolution configuration) -->
 - `raw` (raw の設定のファイルの内容) <!-- (raw configuration file content) -->

IP アドレスとサブネットは CIDR 形式 (`1.1.1.1/24` や `fd80:1234::1/64`) で指定することを想定しています。
<!--
It is expected that IP addresses and subnets are given using CIDR notation (`1.1.1.1/24` or `fd80:1234::1/64`).
-->

例外としてトンネルのローカルとリモートのアドレスは単なるアドレス (`1.1.1.1` や `fd80:1234::1`) を指定します。
<!--
The exception being tunnel local and remote addresses which are just plain addresses (`1.1.1.1` or `fd80:1234::1`).
-->

キー <!-- Key -->                    | 型 <!-- Type --> | 条件 <!-- Condition -->           | デフォルト <!-- Default -->                                        | 説明 <!-- Description -->
:--                                  | :--       | :--                                      | :--                                                                | :--
bgp.peers.NAME.address               | string    | bgp server                               | -                                                                  | ピアのアドレス (IPv4 か IPv6) <!-- Peer address (IPv4 or IPv6) -->
bgp.peers.NAME.asn                   | integer   | bgp server                               | -                                                                  | ピアの AS 番号 <!-- Peer AS number -->
bgp.peers.NAME.password              | string    | bgp server                               | - (パスワード無し) <!-- (no password) -->                          | ピアのセッションパスワード（省略可能） <!-- Peer session password (optional) -->
bgp.ipv4.nexthop                     | string    | bgp server                               | ローカルアドレス <!-- local address -->                            | 広告されたプリフィクスの next-hop をオーバーライド <!-- Override the next-hop for advertised prefixes -->
bgp.ipv6.nexthop                     | string    | bgp server                               | ローカルアドレス <!-- local address -->                            | 広告されたプリフィクスの next-hop をオーバーライド <!-- Override the next-hop for advertised prefixes -->
bridge.driver                        | string    | -                                        | native                                                             | ブリッジのドライバ ("native" か "openvswitch") <!-- Bridge driver ("native" or "openvswitch") -->
bridge.external\_interfaces          | string    | -                                        | -                                                                  | ブリッジに含める未設定のネットワークインタフェースのカンマ区切りリスト <!-- Comma separate list of unconfigured network interfaces to include in the bridge -->
bridge.hwaddr                        | string    | -                                        | -                                                                  | ブリッジの MAC アドレス <!-- MAC address for the bridge -->
bridge.mode                          | string    | -                                        | standard                                                           | ブリッジの稼働モード ("standard" か "fan") <!-- Bridge operation mode ("standard" or "fan") -->
bridge.mtu                           | integer   | -                                        | 1500                                                               | ブリッジの MTU (tunnel か fan かでデフォルト値は変わります) <!-- Bridge MTU (default varies if tunnel or fan setup) -->
dns.domain                           | string    | -                                        | lxd                                                                | DHCP のクライアントに広告し DNS の名前解決に使用するドメイン <!-- Domain to advertise to DHCP clients and use for DNS resolution -->
dns.mode                             | string    | -                                        | managed                                                            | DNS の登録モード ("none" は DNS レコード無し、 "managed" は LXD が静的レコードを生成、 "dynamic" はクライアントがレコードを生成) <!-- DNS registration mode ("none" for no DNS record, "managed" for LXD generated static records or "dynamic" for client generated records) -->
dns.search                           | string    | -                                        | -                                                                  | 完全なドメインサーチのカンマ区切りリスト（デフォルトは `dns.domain` の値） <!-- Full comma separated domain search list, defaulting to `dns.domain` value -->
dns.zone.forward                     | string    | -                                        | managed                                                            | 正引き DNS レコード用の DNS ゾーン名 <!-- DNS zone name for forward DNS records -->
dns.zone.reverse.ipv4                | string    | -                                        | managed                                                            | IPv4 逆引き DNS レコード用の DNS ゾーン名 <!-- DNS zone name for IPv4 reverse DNS records -->
dns.zone.reverse.ipv6                | string    | -                                        | managed                                                            | IPv6 逆引き DNS レコード用の DNS ゾーン名 <!-- DNS zone name for IPv6 reverse DNS records -->
fan.overlay\_subnet                  | string    | ファンモード <!-- fan mode -->           | 240.0.0.0/8                                                        | FAN の overlay として使用するサブネット (CIDR 形式) <!-- Subnet to use as the overlay for the FAN (CIDR notation) -->
fan.type                             | string    | ファンモード <!-- fan mode -->           | vxlan                                                              | FAN のトンネル・タイプ ("vxlan" か "ipip") <!-- The tunneling type for the FAN ("vxlan" or "ipip") -->
fan.underlay\_subnet                 | string    | ファンモード <!-- fan mode -->           | 自動（作成時のみ） <!-- auto (on create only) -->                  | FAN の underlay として使用するサブネット (CIDR 形式)。デフォルトのゲートウェイサブネットを使うには "auto" を指定。 <!-- Subnet to use as the underlay for the FAN (CIDR notation). Use "auto" to use default gateway subnet-->
ipv4.address                         | string    | 標準モード <!-- standard mode -->        | 自動（作成時のみ） <!-- auto (on create only) -->                  | ブリッジの IPv4 アドレス (CIDR 形式)。 IPv4 をオフにするには "none" 、新しいランダムな未使用のサブネットを生成するには "auto" を指定。 <!-- IPv4 address for the bridge (CIDR notation). Use "none" to turn off IPv4 or "auto" to generate a new random unused subnet -->
ipv4.dhcp                            | boolean   | ipv4 アドレス <!-- address -->           | true                                                               | DHCP を使ってアドレスを割り当てるかどうか <!-- Whether to allocate addresses using DHCP -->
ipv4.dhcp.expiry                     | string    | ipv4 dhcp                                | 1h                                                                 | DHCP リースの有効期限 <!-- When to expire DHCP leases -->
ipv4.dhcp.gateway                    | string    | ipv4 dhcp                                | ipv4.address                                                       | サブネットのゲートウェイのアドレス <!-- Address of the gateway for the subnet -->
ipv4.dhcp.ranges                     | string    | ipv4 dhcp                                | 全てのアドレス <!-- all addresses -->                              | DHCP に使用する IPv4 の範囲 (開始-終了 形式) のカンマ区切りリスト <!-- Comma separated list of IP ranges to use for DHCP (FIRST-LAST format) -->
ipv4.firewall                        | boolean   | ipv4 アドレス <!-- address -->           | true                                                               | このネットワークに対するファイアウォールのフィルタリングルールを生成するかどうか <!-- Whether to generate filtering firewall rules for this network -->
ipv4.nat.address                     | string    | ipv4 アドレス <!-- address -->           | -                                                                  | ブリッジからの送信時に使うソースアドレス <!-- The source address used for outbound traffic from the bridge -->
ipv4.nat                             | boolean   | ipv4 アドレス <!-- address -->           | false                                                              | NAT にするかどうか（通常のブリッジではデフォルト値は true で ipv4.address が生成され、fan ブリッジでは常にデフォルト値は true になります） <!-- Whether to NAT (defaults to true for regular bridges where ipv4.address is generated and always defaults to true for fan bridges) -->
ipv4.nat.order                       | string    | ipv4 アドレス <!-- address -->           | before                                                             | 必要な NAT のルールを既存のルールの前に追加するか後に追加するか <!-- Whether to add the required NAT rules before or after any pre-existing rules -->
ipv4.ovn.ranges                      | string    | -                                        | -                                                                  | 子供の OVN ネットワークルーターに使用する IPv4 アドレスの範囲（開始-終了 形式）のカンマ区切りリスト <!-- Comma separate list of IPv4 ranges to use for child OVN network routers (FIRST-LAST format) -->
ipv4.routes                          | string    | ipv4 アドレス <!-- address -->           | -                                                                  | ブリッジへルーティングする追加の IPv4 CIDR サブネットのカンマ区切りリスト <!-- Comma separated list of additional IPv4 CIDR subnets to route to the bridge -->
ipv4.routing                         | boolean   | ipv4 アドレス <!-- address -->           | true                                                               | ブリッジの内外にトラフィックをルーティングするかどうか <!-- Whether to route traffic in and out of the bridge -->
ipv6.address                         | string    | 標準モード <!-- standard mode -->        | 自動（作成時のみ） <!-- auto (on create only) -->                  | ブリッジの IPv6 アドレス (CIDR 形式)。 IPv6 をオフにするには "none" 、新しいランダムな未使用のサブネットを生成するには "auto" を指定。 <!-- IPv6 address for the bridge (CIDR notation). Use "none" to turn off IPv6 or "auto" to generate a new random unused subnet -->
ipv6.dhcp                            | boolean   | ipv6 アドレス <!-- address -->           | true                                                               | DHCP 上で追加のネットワーク設定を提供するかどうか <!-- Whether to provide additional network configuration over DHCP -->
ipv6.dhcp.expiry                     | string    | ipv6 dhcp                                | 1h                                                                 | DHCP リースの有効期限 <!-- When to expire DHCP leases -->
ipv6.dhcp.ranges                     | string    | ipv6 ステートフル <!-- stateful --> dhcp | 全てのアドレス <!-- all addresses -->                              | DHCP に使用する IPv6 の範囲 (開始-終了 形式) のカンマ区切りリスト <!-- Comma separated list of IPv6 ranges to use for DHCP (FIRST-LAST format) -->
ipv6.dhcp.stateful                   | boolean   | ipv6 dhcp                                | false                                                              | DHCP を使ってアドレスを割り当てるかどうか <!-- Whether to allocate addresses using DHCP -->
ipv6.firewall                        | boolean   | ipv6 アドレス <!-- address -->           | true                                                               | このネットワークに対するファイアウォールのフィルタリングルールを生成するかどうか <!-- Whether to generate filtering firewall rules for this network -->
ipv6.nat.address                     | string    | ipv6 アドレス <!-- address -->           | -                                                                  | ブリッジからの送信時に使うソースアドレス <!-- The source address used for outbound traffic from the bridge -->
ipv6.nat                             | boolean   | ipv6 アドレス <!-- address -->           | false                                                              | NAT にするかどうか (未設定の場合はデフォルト値は true になりランダムな ipv6.address が生成されます) <!-- Whether to NAT (will default to true if unset and a random ipv6.address is generated) -->
ipv6.nat.order                       | string    | ipv6 アドレス <!-- address -->           | before                                                             | 必要な NAT のルールを既存のルールの前に追加するか後に追加するか <!-- Whether to add the required NAT rules before or after any pre-existing rules -->
ipv6.ovn.ranges                      | string    | -                                        | -                                                                  | 子供の OVN ネットワークルーターに使用する IPv6 アドレスの範囲（開始-終了 形式) のカンマ区切りリスト <!-- Comma separate list of IPv6 ranges to use for child OVN network routers (FIRST-LAST format) -->
ipv6.routes                          | string    | ipv6 アドレス <!-- address -->           | -                                                                  | ブリッジへルーティングする追加の IPv4 CIDR サブネットのカンマ区切りリスト <!-- Comma separated list of additional IPv6 CIDR subnets to route to the bridge -->
ipv6.routing                         | boolean   | ipv6 アドレス <!-- address -->           | true                                                               | ブリッジの内外にトラフィックをルーティングするかどうか <!-- Whether to route traffic in and out of the bridge -->
maas.subnet.ipv4                     | string    | ipv4 アドレス <!-- address -->           | -                                                                  | インスタンスを登録する MAAS IPv4 サブネット (NIC で `network` プロパティを使う場合に有効) <!-- MAAS IPv4 subnet to register instances in (when using `network` property on nic) -->
maas.subnet.ipv6                     | string    | ipv6 アドレス <!-- address -->           | -                                                                  | インスタンスを登録する MAAS IPv6 サブネット (NIC で `network` プロパティを使う場合に有効) <!-- MAAS IPv6 subnet to register instances in (when using `network` property on nic) -->
raw.dnsmasq                          | string    | -                                        | -                                                                  | 設定に追加する dnsmasq の設定ファイル <!-- Additional dnsmasq configuration to append to the configuration file-->
tunnel.NAME.group                    | string    | vxlan                                    | 239.0.0.1                                                          | vxlan のマルチキャスト設定 (local と remote が未設定の場合に使われます) <!-- Multicast address for vxlan (used if local and remote aren't set) -->
tunnel.NAME.id                       | integer   | vxlan                                    | 0                                                                  | vxlan トンネルに使用するトンネル ID <!-- Specific tunnel ID to use for the vxlan tunnel -->
tunnel.NAME.interface                | string    | vxlan                                    | -                                                                  | トンネルに使用するホスト・インタフェース <!-- Specific host interface to use for the tunnel -->
tunnel.NAME.local                    | string    | gre か <!-- or --> vxlan                 | -                                                                  | トンネルに使用するローカルアドレス (マルチキャスト vxlan の場合は不要) <!-- Local address for the tunnel (not necessary for multicast vxlan) -->
tunnel.NAME.port                     | integer   | vxlan                                    | 0                                                                  | vxlan トンネルに使用するポート <!-- Specific port to use for the vxlan tunnel -->
tunnel.NAME.protocol                 | string    | 標準モード <!-- standard mode -->        | -                                                                  | トンネリングのプロトコル ("vxlan" か "gre") <!-- Tunneling protocol ("vxlan" or "gre") -->
tunnel.NAME.remote                   | string    | gre か <!-- or --> vxlan                 | -                                                                  | トンネルに使用するリモートアドレス (マルチキャスト vxlan の場合は不要) <!-- Remote address for the tunnel (not necessary for multicast vxlan) -->
tunnel.NAME.ttl                      | integer   | vxlan                                    | 1                                                                  | マルチキャストルーティングトポロジーに使用する固有の TTL <!-- Specific TTL to use for multicast routing topologies -->
security.acls                        | string    | -                                        | -                                                                  | このネットワークに接続されたNICに適用するカンマ区切りのネットワークACL（[ブリッジの制限](network-acls.md#_4)参照） <!-- Comma separated list of Network ACLs to apply to NICs connected to this network (see [Limitations](network-acls.md#bridge-limitations)) -->
security.acls.default.ingress.action | string    | security.acls                            | reject                                                             | どの ACL ルールにもマッチしない ingress トラフィックに使うアクション <!-- Action to use for ingress traffic that doesn't match any ACL rule -->
security.acls.default.egress.action  | string    | security.acls                            | reject                                                             | どの ACL ルールにもマッチしない egress トラフィックに使うアクション <!-- Action to use for egress traffic that doesn't match any ACL rule -->
security.acls.default.ingress.logged | boolean   | security.acls                            | false                                                              | どの ACL ルールにもマッチしない ingress トラフィックをログ出力するかどうか <!-- Whether to log ingress traffic that doesn't match any ACL rule -->
security.acls.default.egress.logged  | boolean   | security.acls                            | false                                                              | どの ACL ルールにもマッチしない egress トラフィックをログ出力するかどうか <!-- Whether to log egress traffic that doesn't match any ACL rule -->

これらのキーは lxc コマンドで以下のように設定できます。
<!--
Those keys can be set using the lxc tool with:
-->

```bash
lxc network set <network> <key> <value>
```

### systemd-resolved との統合 <!-- Integration with systemd-resolved -->
LXD が動いているシステムが DNS のルックアップに systemd-resolved を使用している場合、 LXD が名前解決できるドメインを systemd-resolved に指定することができます。
これには systemd-resolved にどのブリッジ、ネームサーバーのアドレス、そして DNS ドメインかを伝える必要があります。
<!--
If the system running LXD uses systemd-resolved to perform DNS
lookups, it's possible to notify resolved of the domain(s) that
LXD is able to resolve.  This requires telling resolved the
specific bridge(s), nameserver address(es), and dns domain(s).
-->

例えば、 LXD が `lxdbr0` インターフェースを使用している場合、 `lxc network get lxdbr0 ipv4.address` コマンドで IPv4 アドレス（IPv4 アドレスの代わりに IPv6 アドレスを使うこともできますし、 IPv4 アドレスと IPv6 アドレスの両方を使うこともできます）と `lxc network get lxdbr0 dns.domain` （ドメインが設定されていない場合は上記の表に示されているデフォルト値の `lxd` が使用されます）でドメインを取得します。
そして systemd-resolved に以下のように指定します。
<!--
For example, if LXD is using the `lxdbr0` interface, get the
ipv4 address with `lxc network get lxdbr0 ipv4.address` command
(the ipv6 can be used instead or in addition), and the domain
with `lxc network get lxdbr0 dns.domain` (if unset, the domain
is `lxd` as shown in the table above).  Then notify resolved:
-->

```
systemd-resolve --interface lxdbr0 --set-domain '~lxd' --set-dns n.n.n.n
```

上記の `lxdbr0` は実際のブリッジの名前に、 `n.n.n.n` はネームサーバーの実際の（サブネットマスクを除いた） アドレスに置き換えて実行してください。
<!--
Replace `lxdbr0` with the actual bridge name, and `n.n.n.` with
the actual address of the nameserver (without the subnet netmask).
-->

さらに `lxd` はドメイン名に置き換えてください。
ドメイン名の前の `~` が重要ですので注意してください。
`~` はこのドメインだけをルックアップするためにこのネームサーバーを使うように systemd-resolved に指定します。
実際のドメイン名が何であるかにかかわらず `~` を前につけるべきです。
また、 `~` という文字はシェルが展開するかもしれないので、クォートに囲んでエスケープする必要があるかもしれません。
<!--
Also replace `lxd` with the domain name.  Note the `~` before the
domain name is important; it tells resolved to use this
nameserver to look up only this domain; no matter what your
actual domain name is, you should prefix it with `~`.  Also,
since the shell may expand the `~` character, you may need to
include it in quotes.
-->

systemd のより新しいリリースでは `systemd-resolve` コマンドは deprecated になっていますが、（これを書いている時点では）後方互換性のためまだ提供されています。
systemd-resolved に伝えるための新しい方法は `resolvectl` コマンドを使うことです。
これは以下の 2 ステップで実行します。
<!--
In newer releases of systemd, the `systemd-resolve` command has been
deprecated, however it is still provided for backwards compatibility
(as of this writing).  The newer method to notify resolved is using
the `resolvectl` command, which would be done in two steps:
-->

```
resolvectl dns lxdbr0 n.n.n.n
resolvectl domain lxdbr0 '~lxd'
```

この systemd-resolved の設定はブリッジが存在する間のみ存続します。
ですので、リブートと LXD が再起動するたびにこのコマンドを繰り返し実行する必要があります（これを自動化するには下記を参照してください）。
また、これはブリッジの `dns.mode` が `none` でないときにしか機能しないことに注意してください。
<!--
This resolved configuration will persist as long as the bridge
exists, so you must repeat this command each reboot and after
LXD is restarted (see below on how to automate this).

Also note this only works if the bridge `dns.mode` is not `none`.
-->

`dns.domain` の使用に依存する場合 DNS の名前解決ができるように resolved の DNSSEC を無効にする必要があるかもしれないことに注意してください。
これは `resolved.conf` の `DNSSEC` オプションで設定できます。
<!--
Note that depending on the `dns.domain` used, you may need to disable
DNSSEC in resolved to allow for DNS resolution. This can be done through
the `DNSSEC` option in `resolved.conf`.
-->

LXD が `lxdbr0` インターフェースを作成した場合、 `systemd-resolved` の DNS 設定をシステム起動時に適用するのを自動化するには
以下のような設定を含む systemd の unit ファイル `/etc/systemd/system/lxd-dns-lxdbr0.service` を作成する必要があります。
<!--
To automate the `systemd-resolved` DNS configuration when LXD creates the `lxdbr0` interface so that it is applied
on system start you need to create a systemd unit file `/etc/systemd/system/lxd-dns-lxdbr0.service` containing:
-->

```
[Unit]
Description=LXD per-link DNS configuration for lxdbr0
BindsTo=sys-subsystem-net-devices-lxdbr0.device
After=sys-subsystem-net-devices-lxdbr0.device

[Service]
Type=oneshot
ExecStart=/usr/bin/resolvectl dns lxdbr0 n.n.n.n
ExecStart=/usr/bin/resolvectl domain lxdbr0 '~lxd'

[Install]
WantedBy=sys-subsystem-net-devices-lxdbr0.device
```

`n.n.n.n` を `lxdbr0` ブリッジの IP アドレスで必ず置き換えてください。
<!--
Be sure to replace `n.n.n.n` in that file with the IP of the `lxdbr0` bridge.
-->

自動起動を有効にし、起動するには以下のようにします。
<!--
Then enable and start it using:
-->

```
sudo systemctl daemon-reload
sudo systemctl enable --now lxd-dns-lxdbr0
```

`lxdbr0` インタフェースが既に存在する（例: LXD が実行中である場合など）場合、以下のようにサービスが起動済みかを確認できます。
<!--
If the `lxdbr0` interface already exists (i.e LXD is running), then you can check that the new service has started:
-->

```
sudo systemctl status lxd-dns-lxdbr0.service
● lxd-dns-lxdbr0.service - LXD per-link DNS configuration for lxdbr0
     Loaded: loaded (/etc/systemd/system/lxd-dns-lxdbr0.service; enabled; vendor preset: enabled)
     Active: inactive (dead) since Mon 2021-06-14 17:03:12 BST; 1min 2s ago
    Process: 9433 ExecStart=/usr/bin/resolvectl dns lxdbr0 n.n.n.n (code=exited, status=0/SUCCESS)
    Process: 9434 ExecStart=/usr/bin/resolvectl domain lxdbr0 ~lxd (code=exited, status=0/SUCCESS)
   Main PID: 9434 (code=exited, status=0/SUCCESS)
```

次に設定が適用されているかを以下のように確認します。
<!--
You can then check it has applied the settings using:
-->

```
sudo resolvectl status lxdbr0
Link 6 (lxdbr0)
      Current Scopes: DNS
DefaultRoute setting: no
       LLMNR setting: yes
MulticastDNS setting: no
  DNSOverTLS setting: no
      DNSSEC setting: no
    DNSSEC supported: no
  Current DNS Server: n.n.n.n
         DNS Servers: n.n.n.n
          DNS Domain: ~lxd
```

### IPv6 プリフィクスサイズ <!-- IPv6 prefix size -->
最適な動作には 64 のプリフィクスサイズが望ましいです。
より大きなサブネット（ 64 より小さいプリフィクス）も正しく動作するでしょうが、SLAAC環境下では有用ではないことが多いでしょう。
<!--
For optimal operation, a prefix size of 64 is preferred.
Larger subnets (prefix smaller than 64) should work properly too but
aren't typically that useful for SLAAC.
-->

IPv6 アドレスの割り当てにステートフル DHCPv6 を使用している場合は、より小さなサブネットも理論的には利用可能ですが、 dnsmasq にきちんとサポートされておらず問題が起きるかもしれません。
これらの 1 つをどうしても使わなければならない場合、静的割り当てか別のスタンドアロンの RA デーモンを使用可能です。
<!--
Smaller subnets while in theory possible when using stateful DHCPv6 for
IPv6 allocation aren't properly supported by dnsmasq and may be the
source of issue. If you must use one of those, static allocation or
another standalone RA daemon be used.
-->

### Firewalld で DHCP, DNS を許可する <!-- Allow DHCP, DNS with Firewalld -->

firewalld を使用しているホストで LXD が実行する DHCP と DNS サーバーにインスタンスがアクセスできるようにするには、ホストのブリッジインターフェースを firewalld の `trusted` ゾーンに追加する必要があります。
<!--
In order to allow instances to access the DHCP and DNS server that LXD runs on the host when using firewalld
you need to add the host's bridge interface to the `trusted` zone in firewalld.
-->

（リブート後も設定が残るように）恒久的にこれを行うには以下のコマンドを実行してください。
<!--
To do this permanently (so that it persists after a reboot) run the following command:
-->

```
firewall-cmd --zone=trusted --change-interface=<LXD network name> --permanent
```

例えばブリッジネットワークが `lxdbr0` という名前の場合、以下のコマンドを実行します。
<!--
E.g. for a bridged network called `lxdbr0` run the command:
-->

```
firewall-cmd --zone=trusted --change-interface=lxdbr0 --permanent
```

これにより LXD 自身のファイアーウォールのルールが有効になります。
<!--
This will then allow LXD's own firewall rules to take effect.
-->


### Firewalld に LXD の iptables ルールを制御させるには <!-- How to let Firewalld control the LXD's iptables rules -->

firewalld と LXD を一緒に使う場合、 iptables のルールがオーバーラップするかもしれません。例えば firewalld が LXD デーモンより後に起動すると firewalld が LXD の iptables ルールを削除し、 LXD コンテナーが外向きのインターネットアクセスが全くできなくなるかもしれません。
これを修正する 1 つの方法は LXD の iptables ルールを firewalld に移譲し、 LXD の iptables ルールは無効にすることです。
<!--
When using firewalld and LXD together, iptables rules can overlaps. For example, firewalld could erase LXD iptables rules if it is started after LXD daemon, then LXD container will not be able to do any oubound internet access.
One way to fix it is to delegate to firewalld the LXD's iptables rules and to disable the LXD ones.
-->

最初のステップは [Firewalld で DHCP, DNS を許可する](#allow-dhcp-dns-with-firewalld) ことです。
<!--
First step is to [allow DNS and DHCP](#allow-dhcp-dns-with-firewalld).
-->

次に LXD に iptables ルールを設定しないように（firewalld が設定するので）伝えます。
<!--
Then to tell to LXD totally stop to set iptables rules (because firewalld will do it):
-->
```
lxc network set lxdbr0 ipv4.nat false
lxc network set lxdbr0 ipv6.nat false
lxc network set lxdbr0 ipv6.firewall false
lxc network set lxdbr0 ipv4.firewall false
```

最後に firewalld のルールを LXD の利用ケースに応じて有効にします（この例では、ブリッジインターフェースが `lxdbr0` で付与されている IP の範囲が `10.0.0.0/24` だとしています）。
<!--
Finally, to enable iptables firewalld's rules for LXD usecase (in this example, we suppose the bridge interface is `lxdbr0` and the associated IP range is `10.0.0.0/24`:
-->
```
firewall-cmd --permanent --direct --add-rule ipv4 filter INPUT 0 -i lxdbr0 -s 10.0.0.0/24 -m comment --comment "generated by firewalld for LXD" -j ACCEPT
firewall-cmd --permanent --direct --add-rule ipv4 filter OUTPUT 0 -o lxdbr0 -d 10.0.0.0/24 -m comment --comment "generated by firewalld for LXD" -j ACCEPT
firewall-cmd --permanent --direct --add-rule ipv4 filter FORWARD 0 -i lxdbr0 -s 10.0.0.0/24 -m comment --comment "generated by firewalld for LXD" -j ACCEPT
firewall-cmd --permanent --direct --add-rule ipv4 nat POSTROUTING 0 -s 10.0.0.0/24 ! -d 10.0.0.0/24 -m comment --comment "generated by firewalld for LXD" -j MASQUERADE
firewall-cmd --reload
```

firewalld にルールが設定されたかを確認するには以下のようにします。
<!--
To check the rules are taken into account by firewalld:
-->
```
firewall-cmd --direct --get-all-rules 
```

警告：上記の手順はフールプルーフなアプローチではなく、不注意にセキュリティリスクをもたらすことにつながる可能性があります。
<!--
Warning: what is exposed above is not a fool-proof approach and may end up inadvertently introducing a security risk.
-->

## <a name="network-macvlan"></a> ネットワーク: macvlan <!-- network: macvlan -->

macvlan ネットワークタイプではインスタンスを macvlan NIC を使って親のインターフェースに接続する際に使用するプリセットを指定可能です。
これによりインスタンスの NIC 自体は下層の詳しい設定を一切知ることなく、接続する `network` を単に指定するだけで設定できます。
<!--
The macvlan network type allows one to specify presets to use when connecting instances to a parent interface
using macvlan NICs. This allows the instance NIC itself to simply specify the `network` it is connecting to without
knowing any of the underlying configuration details.
-->

ネットワーク設定プロパティー:
<!--
Network configuration properties:
-->

キー <!-- Key -->        | 型 <!-- Type --> | 条件 <!-- Condition --> | デフォルト <!-- Default --> | 説明 <!-- Description -->
:--                             | :--       | :--                            | :--                  | :--
maas.subnet.ipv4                | string    | ipv4 アドレス <!-- address --> | -                    | インスタンスを登録する MAAS IPv4 サブネット（nic の `network` プロパティを使用する場合） <!-- MAAS IPv4 subnet to register instances in (when using `network` property on nic) -->
maas.subnet.ipv6                | string    | ipv6 アドレス <!-- address --> | -                    | インスタンスを登録する MAAS IPv6 サブネット（nic の `network` プロパティを使用する場合） <!-- MAAS IPv6 subnet to register instances in (when using `network` property on nic) -->
mtu                             | integer   | -                              | -                    | 作成するインターフェースの MTU <!-- The MTU of the new interface -->
parent                          | string    | -                              | -                    | macvlan NIC を作成する親のインターフェース <!-- Parent interface to create macvlan NICs on -->
vlan                            | integer   | -                              | -                    | アタッチする先の VLAN ID <!-- The VLAN ID to attach to -->
gvrp                            | boolean   | -                              | false                | GARP VLAN Registration Protocol を使って VLAN を登録する <!-- Register VLAN using GARP VLAN Registration Protocol -->

## <a name="network-sriov"></a> ネットワーク: sriov <!-- network: sriov -->

sriov ネットワークタイプではインスタンスを sriov NIC を使って親のインターフェースに接続する際に使用するプリセットを指定可能です。
これによりインスタンスの NIC 自体は下層の詳しい設定を一切知ることなく、接続する `network` を単に指定するだけで設定できます。
<!--
The sriov network type allows one to specify presets to use when connecting instances to a parent interface
using sriov NICs. This allows the instance NIC itself to simply specify the `network` it is connecting to without
knowing any of the underlying configuration details.
-->

ネットワーク設定プロパティー:
<!--
Network configuration properties:
-->

キー <!-- Key -->        | 型 <!-- Type --> | 条件 <!-- Condition -->  | デフォルト <!-- Default --> | 説明 <!-- Description -->
:--                             | :--       | :--                            | :--                   | :--
maas.subnet.ipv4                | string    | ipv4 アドレス <!-- address --> | -                     | インスタンスを登録する MAAS IPv4 サブネット（nic の `network` プロパティを使用する場合） <!-- MAAS IPv4 subnet to register instances in (when using `network` property on nic) -->
maas.subnet.ipv6                | string    | ipv6 アドレス <!-- address --> | -                     | インスタンスを登録する MAAS IPv6 サブネット（nic の `network` プロパティを使用する場合） <!-- MAAS IPv6 subnet to register instances in (when using `network` property on nic) -->
mtu                             | integer   | -                              | -                     | 作成するインターフェースの MTU <!-- The MTU of the new interface -->
parent                          | string    | -                              | -                     | sriov NIC を作成する親のインターフェース <!-- Parent interface to create sriov NICs on -->
vlan                            | integer   | -                              | -                     | アタッチする先の VLAN ID <!-- The VLAN ID to attach to -->

## <a name="network-ovn"></a> ネットワーク: ovn <!-- network: ovn -->

ovn ネットワークタイプは OVN SDN を使って論理的なネットワークの作成を可能にします。
これは複数の個別のネットワーク内で同じ論理ネットワークのサブネットを使うような検証環境やマルチテナントの環境で便利です。
<!--
The ovn network type allows the creation of logical networks using the OVN SDN. This can be useful for labs and
multi-tenant environments where the same logical subnets are used in multiple discrete networks.
-->

LXD の OVN ネットワークはより広いネットワークへの外向きのアクセスを可能にするため既存の管理された LXD のブリッジネットワークに接続できます。
OVN 論理ネットワークからの全ての接続は親のネットワークによって割り当てられた動的 IP に NAT されます。
<!--
A LXD OVN network can be connected to an existing managed LXD bridge network in order for it to gain outbound
access to the wider network. All connections from the OVN logical networks are NATed to a dynamic IP allocated by
the parent network.
-->

### スタンドアロンの LXD での OVN の設定 <!-- Standalone LXD OVN setup -->

これは外向きの通信のために親のネットワーク lxdbr0 に接続されたスタンドアロンの OVN ネットワークを作成する手順です。
<!--
This will create a standalone OVN network that is connected to the parent network lxdbr0 for outbound connectivity.
-->

OVN のツールをインストールし、ローカルノードで OVN の統合ブリッジを設定します。
<!--
Install the OVN tools and configure the OVN integration bridge on the local node:
-->

```
sudo apt install ovn-host ovn-central
sudo ovs-vsctl set open_vswitch . \
  external_ids:ovn-remote=unix:/var/run/ovn/ovnsb_db.sock \
  external_ids:ovn-encap-type=geneve \
  external_ids:ovn-encap-ip=127.0.0.1
```

以下を使用して OVN ネットワークとインスタンスを作成します。
<!--
Create an OVN network and an instance using it:
-->

```
lxc network set lxdbr0 ipv4.dhcp.ranges=... ipv4.ovn.ranges=... # OVN ゲートウェイに IP のレンジを割り当て
lxc network create ovntest --type=ovn network=lxdbr0
lxc init images:ubuntu/20.04 c1
lxc config device override c1 eth0 network=ovntest
lxc start c1
lxc ls
+------+---------+---------------------+----------------------------------------------+-----------+-----------+
| NAME |  STATE  |        IPV4         |                     IPV6                     |   TYPE    | SNAPSHOTS |
+------+---------+---------------------+----------------------------------------------+-----------+-----------+
| c1   | RUNNING | 10.254.118.2 (eth0) | fd42:887:cff3:5089:216:3eff:fef0:549f (eth0) | CONTAINER | 0         |
+------+---------+---------------------+----------------------------------------------+-----------+-----------+
```

<!--
```
lxc network set lxdbr0 ipv4.dhcp.ranges=... ipv4.ovn.ranges=... # Allocate IP range for OVN gateways.
lxc network create ovntest -\-type=ovn network=lxdbr0
lxc init images:ubuntu/20.04 c1
lxc config device override c1 eth0 network=ovntest
lxc start c1
lxc ls
+\-\-\-\-\-\-+\-\-\-\-\-\-\-\-\-+\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-+\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-+\-\-\-\-\-\-\-\-\-\-\-+\-\-\-\-\-\-\-\-\-\-\-+
| NAME |  STATE  |        IPV4         |                     IPV6                     |   TYPE    | SNAPSHOTS |
+\-\-\-\-\-\-+\-\-\-\-\-\-\-\-\-+\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-+\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-+\-\-\-\-\-\-\-\-\-\-\-+\-\-\-\-\-\-\-\-\-\-\-+
| c1   | RUNNING | 10.254.118.2 (eth0) | fd42:887:cff3:5089:216:3eff:fef0:549f (eth0) | CONTAINER | 0         |
+\-\-\-\-\-\-+\-\-\-\-\-\-\-\-\-+\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-+\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-+\-\-\-\-\-\-\-\-\-\-\-+\-\-\-\-\-\-\-\-\-\-\-+
```
-->

ネットワークフォワード: <!--Network forwards: -->

OVN のネットワークサポートは [ネットワークフォワード](network-forwards.md) 参照。
<!--
OVN networks support [network forwards](network-forwards.md).
-->

ネットワークピア: <!-- Network peers: -->

OVN のネットワークピアサポートは [ネットワークピア](network-peers.md) 参照。
<!--
OVN networks support [network peers](network-peers.md).
-->

ネットワークの設定プロパティー: <!-- Network configuration properties: -->

キー <!-- Key -->                    | 型 <!-- Type --> | 条件 <!-- Condition -->    | デフォルト <!-- Default -->                                | 説明 <!-- Description -->
:--                                  | :--       | :--                               | :--                                                        | :--
bridge.hwaddr                        | string    | -                                 | -                                                          | ブリッジの MAC アドレス <!-- MAC address for the bridge -->
bridge.mtu                           | integer   | -                                 | 1442                                                       | ブリッジの MTU (デフォルトではホストからホストへの geneve トンネルを許可します) <!-- Bridge MTU (default allows host to host geneve tunnels) -->
dns.domain                           | string    | -                                 | lxd                                                        | DHCP のクライアントに広告し DNS の名前解決に使用するドメイン <!-- Domain to advertise to DHCP clients and use for DNS resolution -->
dns.search                           | string    | -                                 | -                                                          | 完全なドメインサーチのカンマ区切りリスト（デフォルトは `dns.domain` の値） <!-- Full comma separated domain search list, defaulting to `dns.domain` value -->
dns.zone.forward                     | string    | -                                 | -                                                          | 正引き DNS レコード用の DNS ゾーン名 <!-- DNS zone name for forward DNS records -->
dns.zone.reverse.ipv4                | string    | -                                 | -                                                          | IPv4 逆引き DNS レコード用の DNS ゾーン名 <!-- DNS zone name for IPv4 reverse DNS records -->
dns.zone.reverse.ipv6                | string    | -                                 | -                                                          | IPv6 逆引き DNS レコード用の DNS ゾーン名 <!-- DNS zone name for IPv6 reverse DNS records -->
ipv4.address                         | string    | 標準モード <!-- standard mode --> | 自動（作成時のみ） <!-- auto (on create only) -->          | ブリッジの IPv4 アドレス (CIDR 形式)。 IPv4 をオフにするには "none" 、新しいランダムな未使用のサブネットを生成するには "auto" を指定。 <!-- IPv4 address for the bridge (CIDR notation). Use "none" to turn off IPv4 or "auto" to generate a new random unused subnet -->
ipv4.dhcp                            | boolean   | ipv4 アドレス <!-- address -->    | true                                                       | DHCP を使ってアドレスを割り当てるかどうか <!-- Whether to allocate addresses using DHCP -->
ipv4.nat                             | boolean   | ipv4 アドレス <!-- address -->    | false                                                      | NAT するかどうか（ipv4.address が未設定の場合デフォルト値は true でランダムな ipv4.address が生成されます） <!-- Whether to NAT (will default to true if unset and a random ipv4.address is generated) -->
ipv4.nat.address                     | string    | ipv4 アドレス <!-- address -->    | -                                                          | ネットワークからの外向きトラフィックに使用されるソースアドレス (アップリンクに `ovn.ingress_mode=routed` が必要) <!-- The source address used for outbound traffic from the network (requires uplink `ovn.ingress_mode=routed`) -->
ipv6.address                         | string    | 標準モード <!-- standard mode --> | 自動（作成時のみ） <!-- auto (on create only) -->          | ブリッジの IPv6 アドレス (CIDR 形式)。 IPv6 をオフにするには "none" 、新しいランダムな未使用のサブネットを生成するには "auto" を指定。 <!-- IPv6 address for the bridge (CIDR notation). Use "none" to turn off IPv6 or "auto" to generate a new random unused subnet -->
ipv6.nat.address                     | string    | ipv6 アドレス <!-- address -->    | -                                                          | ネットワークからの外向きトラフィックに使用されるソースアドレス (アップリンクに `ovn.ingress_mode=routed` が必要) <!-- The source address used for outbound traffic from the network (requires uplink `ovn.ingress_mode=routed`) -->
ipv6.dhcp                            | boolean   | ipv6 アドレス <!-- address -->    | true                                                       | DHCP 上に追加のネットワーク設定を提供するかどうか <!-- Whether to provide additional network configuration over DHCP -->
ipv6.dhcp.stateful                   | boolean   | ipv6 dhcp                         | false                                                      | DHCP を使ってアドレスを割り当てるかどうか <!-- Whether to allocate addresses using DHCP -->
ipv6.nat                             | boolean   | ipv6 アドレス <!-- address -->    | false                                                      | NAT するかどうか（ipv6.address が未設定の場合デフォルト値は true でランダムな ipv6.address が生成されます） <!-- Whether to NAT (will default to true if unset and a random ipv6.address is generated) -->
network                              | string    | -                                 | -                                                          | 外部ネットワークへの外向きのアクセスに使うアップリンクのネットワーク <!-- Uplink network to use for external network access -->
security.acls                        | string    | -                                 | -                                                          | このネットワークに接続する NIC に適用するネットワーク ACL のカンマ区切りリスト <!-- Comma separated list of Network ACLs to apply to NICs connected to this network -->
security.acls.default.ingress.action | string    | security.acls                     | reject                                                     | どの ACL ルールにもマッチしない ingress トラフィックに使うアクション <!-- Action to use for ingress traffic that doesn't match any ACL rule -->
security.acls.default.egress.action  | string    | security.acls                     | reject                                                     | どの ACL ルールにもマッチしない egress トラフィックに使うアクション <!-- Action to use for egress traffic that doesn't match any ACL rule -->
security.acls.default.ingress.logged | boolean   | security.acls                     | false                                                      | どの ACL ルールにもマッチしない ingress トラフィックをログ出力するかどうか <!-- Whether to log ingress traffic that doesn't match any ACL rule -->
security.acls.default.egress.logged  | boolean   | security.acls                     | false                                                      | どの ACL ルールにもマッチしない egress トラフィックをログ出力するかどうか <!-- Whether to log egress traffic that doesn't match any ACL rule -->

## <a name="network-physical"></a> ネットワーク: physical <!-- network: physical -->

physical ネットワークは OVN ネットワークを親インターフェースに接続する際に使用するプリセットの設定を提供します。
<!--
The physical network type allows one to specify presets to use when connecting OVN networks to a parent interface.
-->

ネットワーク設定プロパティー:
<!--
Network configuration properties:
-->


キー <!-- Key -->               | 型 <!-- Type --> | 条件 <!-- Condition -->    | デフォルト <!-- Default -->               | 説明 <!-- Description -->
:--                             | :--       | :--                               | :--                                       | :--
bgp.peers.NAME.address          | string    | bgp server                        | -                                         | `ovn` ダウンストリームネットワークで使用するピアアドレス (IPv4 か IPv6) <!-- Peer address (IPv4 or IPv6) for use by `ovn` downstream networks -->
bgp.peers.NAME.asn              | integer   | bgp server                        | -                                         | `ovn` ダウンストリームネットワークで使用する AS 番号 <!-- Peer AS number for use by `ovn` downstream networks -->
bgp.peers.NAME.password         | string    | bgp server                        | - (パスワード無し) <!-- (no password) --> | `ovn` ダウンストリームネットワークで使用するピアのセッションパスワード（省略可能） <!-- Peer session password (optional) for use by `ovn` downstream networks -->
maas.subnet.ipv4                | string    | ipv4 アドレス <!-- address -->    | -                                         | インスタンスを登録する MAAS IPv4 サブネット (NIC で `network` プロパティを使う場合に有効) <!-- MAAS IPv4 subnet to register instances in (when using `network` property on nic) -->
maas.subnet.ipv6                | string    | ipv6 アドレス <!-- address -->    | -                                         | インスタンスを登録する MAAS IPv6 サブネット (NIC で `network` プロパティを使う場合に有効) <!-- MAAS IPv6 subnet to register instances in (when using `network` property on nic) -->
mtu                             | integer   | -                                 | -                                         | 作成するインターフェースの MTU <!-- The MTU of the new interface -->
parent                          | string    | -                                 | -                                         | sriov NIC を作成する親のインターフェース <!-- Parent interface to create sriov NICs on -->
vlan                            | integer   | -                                 | -                                         | アタッチする先の VLAN ID <!-- The VLAN ID to attach to -->
gvrp                            | boolean   | -                                 | false                                     | GARP VLAN Registration Protocol を使って VLAN を登録する <!-- Register VLAN using GARP VLAN Registration Protocol -->
ipv4.gateway                    | string    | 標準モード <!-- standard mode --> | -                                         | ゲートウェイとネットワークの IPv4 アドレス（CIDR表記） <!-- IPv4 address for the gateway and network (CIDR notation) -->
ipv4.ovn.ranges                 | string    | -                                 | -                                         | 子供の OVN ネットワークルーターに使用する IPv4 アドレスの範囲（開始-終了 形式) のカンマ区切りリスト <!-- Comma separate list of IPv4 ranges to use for child OVN network routers (FIRST-LAST format) -->
ipv4.routes                     | string    | ipv4 アドレス <!-- address -->    | -                                         | 子供の OVN ネットワークの ipv4.routes.external 設定で利用可能な追加の IPv4 CIDR サブネットのカンマ区切りリスト <!-- Comma separated list of additional IPv4 CIDR subnets that can be used with child OVN networks ipv4.routes.external setting -->
ipv4.routes.anycast             | boolean   | ipv4 アドレス <!-- address -->    | false                                     | 複数のネットワーク／NICで同時にオーバーラップするルートが使われることを許可するかどうか <!-- Allow the overlapping routes to be used on multiple networks/NIC at the same time. -->
ipv6.gateway                    | string    | 標準モード <!-- standard mode --> | -                                         | ゲートウェイとネットワークの IPv6 アドレス（CIDR表記） <!-- IPv6 address for the gateway and network (CIDR notation) -->
ipv6.ovn.ranges                 | string    | -                                 | -                                         | 子供の OVN ネットワークルーターに使用する IPv6 アドレスの範囲（開始-終了 形式) のカンマ区切りリスト <!-- Comma separate list of IPv6 ranges to use for child OVN network routers (FIRST-LAST format) -->
ipv6.routes                     | string    | ipv6 アドレス <!-- address -->    | -                                         | 子供の OVN ネットワークの ipv6.routes.external 設定で利用可能な追加の IPv6 CIDR サブネットのカンマ区切りリスト <!-- Comma separated list of additional IPv6 CIDR subnets that can be used with child OVN networks ipv6.routes.external setting -->
ipv6.routes.anycast             | boolean   | ipv6 アドレス <!-- address -->    | false                                     | 複数のネットワーク／NICで同時にオーバーラップするルートが使われることを許可するかどうか <!-- Allow the overlapping routes to be used on multiple networks/NIC at the same time. -->
dns.nameservers                 | string    | 標準モード <!-- standard mode --> | -                                         | 物理ネットワークの DNS サーバー IP のリスト <!-- List of DNS server IPs on physical network -->
ovn.ingress\_mode               | string    | 標準モード <!-- standard mode --> | l2proxy                                   | OVN NIC の外部 IP アドレスがアップリンクネットワークで広告される方法を設定します。 `l2proxy` (proxy ARP/NDP) か `routed` です。 <!-- Sets the method that OVN NIC external IPs will be advertised on uplink network. Either `l2proxy` (proxy ARP/NDP) or `routed`. -->

## BGP の統合 <!-- BGP integration -->
LXD は BGP サーバーとして機能でき、アップストリームの BGP ルーターとセッションを確立し LXD が使用しているアドレスとサブネットを広告できます。
<!--
LXD can act as a BGP server, effectively allowing to establish sessions with upstream BGP routers and announce the addresses and subnets that it's using.
-->

これにより LXD サーバーやクラスターが内部／外部のアドレス空間を直接使い、正しいホストにルーティングされた特定のサブネットやアドレスをターゲットインスタンスにフォワードできます。
<!--
This can be used to allow a LXD server or cluster to directly use internal/external address space, getting the specific subnets or addresses routed to the correct host for it to forward onto the target instance.
-->

このためには `core.bgp_address`, `core.bgp_asn` と `core.bgp_routerid` が設定されている必要があります。
これらが設定されると LXD は BGP セッションのリッスンを開始します。
<!--
For this to work, `core.bgp_address`, `core.bgp_asn` and `core.bgp_routerid` must be set.
Once those are set, LXD will start listening for BGP sessions.
-->

ピアは `bridged` と `physical` で管理されたネットワーク上に定義できます。さらに `bridged` の場合は next-hop をオーバーライドするためにサーバー毎の設定キーの組が利用できます。それらが指定されない場合は next-hop はデフォルトとして BGP セッションに使用されるアドレスになります。
<!--
Peers can be defined on both `bridged` and `physical` managed networks. Additionally in the `bridged` case, a set of per-server configuration keys are also available to override the next-hop. When those aren't specified, the next-hop defaults to the address used for the BGP session.
-->

`physical` ネットワークの場合はアップリンクのネットワークが利用可能なサブネットのリストと BGP 設定を持つような `ovn` ネットワークに使用されます。
親のネットワークが一旦設定されると、子の OVN ネットワークは BGP で広告された外部のサブネットとアドレスを受け取り next-hop は問題のネットワークの OVN ルーターアドレスに設定されます。
<!--
The `physical` network case is used for `ovn` networks where the uplink network is the one holding the list of allowed subnets and the BGP configuration. Once that parent network is configured, children OVN networks will get their external subnets and addresses announced over BGP with the next-hop set to the OVN router address for the network in question.
-->

現在公開されるアドレスとネットワークは以下のとおりです。
<!--
The addresses and networks currently being advertised are:
-->
 - `nat` プロパティーが `true` に設定されない場合はネットワークの `ipv4.address` か `ipv6.address` <!-- Network `ipv4.address` or `ipv6.address` subnets when the matching `nat` property isn't set to `true` -->
 - `nat` プロパティーが設定される場合はネットワークの `ipv4.address` と `ipv6.address` <!-- Network `ipv4.nat.address` and `ipv6.nat.address` when those are set -->
 - `ipv4.routes.external` か `ipv6.routes.external` 経由で定義されるインスタンスの NIC ルート <!-- Instance NIC routes defined through `ipv4.routes.external` or `ipv6.routes.external` -->

現時点では、特定のピアに特定のルートやアドレスのみを公開する方法はありません。代わりにアップストリームのルーターでプリフィクスをフィルターすることを現状ではお勧めします。
<!--
At this time, there isn't a way to only announce some specific routes/addresses to particular peers. Instead it's currently recommended to filter prefixes on the upstream routers.
-->
