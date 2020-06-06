# ネットワーク設定
<!-- Network configuration -->

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
 - `user` (ユーザーのメタデータに対する自由形式の key/value) <!-- (free form key/value for user metadata) -->

## ブリッジ <!-- Bridges -->

LXD でのネットワークの設定タイプの 1 つとして、 LXD はネットワークブリッジの作成と管理をサポートしています。
LXD のブリッジは下層のネイティブな Linux のブリッジと Open vSwitch を利用できます。
<!--
As one of the possible network configuration types under LXD,
LXD supports creating and managing network bridges. LXD bridges 
can leverage underlying native Linux bridges and Open vSwitch. 
-->

LXD のブリッジの作成と管理は `lxc network` コマンドで行えます。
LXD で作成されたブリッジはデフォルトでは "managed" です。
というのは LXD はさらにローカルの `dnsmasq` DHCP サーバをセットアップし、希望すれば (これがデフォルトです) ブリッジに対して NAT も行います。
<!--
Creation and management of LXD bridges is performed via the `lxc network`
command. A bridge created by LXD is by default "managed" which 
means that LXD also will additionally set up a local `dnsmasq` 
DHCP server and if desired also perform NAT for the bridge (this 
is the default.)
-->

ブリッジが LXD に管理されているときは、 `bridge` ネームスペースを使って設定値を変更できます。
<!--
When a bridge is managed by LXD, configuration values
under the `bridge` namespace can be used to configure it.
-->

さらに、 LXD は既存の Linux ブリッジを利用することも出来ます。
この場合、ブリッジは `lxc network` で作成する必要はなく、インスタンスかプロファイルのデバイス設定で下記のように単に参照できます。
<!--
Additionally, LXD can utilize a pre-existing Linux
bridge. In this case, the bridge does not need to be created via
`lxc network` and can simply be referenced in an instance or
profile device configuration as follows:
-->

```
devices:
  eth0:
     name: eth0
     nictype: bridged
     parent: br0
     type: nic
```

## 設定項目 <!-- Configuration Settings -->

LXD のネットワークの設定項目の完全なリストは以下の通りです。
<!--
A complete list of configuration settings for LXD networks can
be found below.
-->

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
dns.search                      | string    | -                     | -                         | 完全なドメインサーチのカンマ区切りリスト（デフォルトは dns.domain ） <!-- Full comma eparate domain search list, defaulting to dns.domain -->
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
ipv4.nat.address                | string    | ipv4 address          | -                         | ブリッジからの送信時に使うソースアドレス <!-- The source address used for outbound traffic from the bridge -->
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
ipv6.nat.address                | string    | ipv6 address          | -                         | ブリッジからの送信時に使うソースアドレス <!-- The source address used for outbound traffic from the bridge -->
ipv6.routes                     | string    | ipv6 address          | -                         | ブリッジへルーティングする追加の IPv4 CIDR サブネットのカンマ区切りリスト <!-- Comma separated list of additional IPv6 CIDR subnets to route to the bridge -->
ipv6.routing                    | boolean   | ipv6 address          | true                      | ブリッジの内外にトラフィックをルーティングするかどうか <!-- Whether to route traffic in and out of the bridge -->
maas.subnet.ipv4                | string    | ipv4 address          | -                         | インスタンスを登録する MAAS IPv4 サブネット (NIC で `network` プロパティを使う場合に有効) <!-- MAAS IPv4 subnet to register instances in (when using `network` property on nic) -->
maas.subnet.ipv6                | string    | ipv6 address          | -                         | インスタンスを登録する MAAS IPv6 サブネット (NIC で `network` プロパティを使う場合に有効) <!-- MAAS IPv6 subnet to register instances in (when using `network` property on nic) -->
raw.dnsmasq                     | string    | -                     | -                         | 設定に追加する dnsmasq の設定ファイル <!-- Additional dnsmasq configuration to append to the configuration file-->
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

## systemd-resolved との統合 <!-- Integration with systemd-resolved -->

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
systemd-resolve --interface lxdbr0 --set-domain '~lxd' --set-dns 1.2.3.4
```

上記の `lxdbr0` は実際のブリッジの名前に、 `1.2.3.4` はネームサーバーの実際の（サブネットマスクを除いた） アドレスに置き換えて実行してください。
<!--
Replace `lxdbr0` with the actual bridge name, and `1.2.3.4` with
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
resolvectl dns lxdbr0 1.2.3.4
resolvectl domain lxdbr0 '~lxd'
```

この systemd-resolved の設定はブリッジが存在する間のみ存続します。
ですので、リブートと LXD が再起動するたびにこのコマンドを繰り返し実行する必要があります。
また、これはブリッジの `dns.mode` が `none` でないときにしか機能しないことに注意してください。
<!--
This resolved configuration will persist as long as the bridge
exists, so you must repeat this command each reboot and after
LXD is restarted.  Also note this only works if the bridge
`dns.mode` is not `none`.
-->

## IPv6 プリフィクスサイズ <!-- IPv6 prefix size -->
最適な動作には 64 のプリフィクスサイズが望ましいです。
より大きなサブネット（ 64 より小さいプリフィクス）も動作はするはずですが、通常は SLAAC に対するほど有用ではありません。
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

## Firewalld で DHCP, DNS を許可する <!-- Allow DHCP, DNS with Firewalld -->

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


## Firewalld に LXD の iptables ルールを制御させるには <!-- How to let Firewalld control the LXD's iptables rules -->

firewalld と LXD を一緒に使う場合、 iptables のルールがオーバーラップするかもしれません。例えば firewalld が LXD デーモンより後に起動すると firewalld が LXD の iptables ルールを削除し、 LXD コンテナーが外向きのインターネットアクセスが全くできなくなるかもしれません。
これを修正する 1 つの方法は LXD の iptables ルールを firewalld に移譲し、 LXD の iptables ルールは無効にすることです。
<!--
When using firewalld and LXD together, iptables rules can overlaps. For example, firewalld could erase LXD iptables rules if it is started after LXD daemon, then LXD container will not be able to do any oubound internet access.
On way to fix it is to delegate to firewalld the LXD's iptables rules and to disable the LXD ones.
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
