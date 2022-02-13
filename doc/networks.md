# ネットワーク設定

 - {ref}`ブリッジ <network-bridge>`: インスタンスを接続する L2 ブリッジを作成（ローカルの DHCP と DNS を提供可能）。これがデフォルトです。
 - {ref}`macvlan <network-macvlan>`: インスタンスを親の macvlan インターフェースに接続する際に使用するプリセットの設定を提供。
 - {ref}`sriov <network-sriov>`: インスタンスを親の SR-IOV インターフェースに接続する際に使用するプリセットの設定を提供。
 - {ref}`ovn <network-ovn>`: OVN SDN (software defined network) システムを使って論理的なネットワークを作成。
 - {ref}`物理 <network-physical>`: OVN ネットワークを親のインターフェースに接続する際に使用するプリセットの設定を提供。

希望するタイプは以下のように `--type` 引数で指定できます。

```bash
lxc network create <name> --type=bridge [options...]
```

`--type` 引数が指定されない場合は、デフォルトの `bridge` が使用されます。

設定キーは現状では全てのネットワークタイプでサポートされている以下のネームスペースによって名前空間が分けられています。

 - `maas` (MAAS ネットワーク識別)
 - `user` (ユーザーのメタデータに対する自由形式の key/value)

(network-bridge)=
## ネットワーク: ブリッジ

LXD でのネットワークの設定タイプの 1 つとして、 LXD はネットワークブリッジの作成と管理をサポートしています。
LXD のブリッジは下層のネイティブな Linux のブリッジと Open vSwitch を利用できます。

LXD のブリッジの作成と管理は `lxc network` コマンドで行えます。
LXD で作成されたブリッジはデフォルトでは "managed" です。
というのは LXD はさらにローカルの `dnsmasq` DHCP サーバをセットアップし、希望すれば (これがデフォルトです) ブリッジに対して NAT も行います。

ブリッジが LXD に管理されているときは、 `bridge` ネームスペースを使って設定値を変更できます。

さらに、 LXD は既存の Linux ブリッジを利用することも出来ます。
この場合、ブリッジは `lxc network` で作成する必要はなく、インスタンスかプロファイルのデバイス設定で下記のように単に参照できます。

```
devices:
  eth0:
     name: eth0
     nictype: bridged
     parent: br0
     type: nic
```

ネットワークフォワード:

ブリッジのネットワークサポートは [network forwards](network-forwards.md#network-bridge) 参照。

ネットワークの設定プロパティ:

LXD のネットワークの設定項目の完全なリストは以下の通りです。

ブリッジネットワークでは以下の設定キーネームスペースが現状サポートされています。

 - `bridge` (L2 インタフェースの設定)
 - `fan` (Ubuntu FAN overlay に特有な設定)
 - `tunnel` (ホスト間のトンネリングの設定)
 - `ipv4` (L3 IPv4 設定)
 - `ipv6` (L3 IPv6 設定)
 - `dns` (DNS サーバと名前解決の設定)
 - `raw` (raw の設定のファイルの内容)

IP アドレスとサブネットは CIDR 形式 (`1.1.1.1/24` や `fd80:1234::1/64`) で指定することを想定しています。

例外としてトンネルのローカルとリモートのアドレスは単なるアドレス (`1.1.1.1` や `fd80:1234::1`) を指定します。

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
ipv4.nat.address                     | string    | ipv4 アドレス           | -                   | ブリッジからの送信時に使うソースアドレス
ipv4.nat                             | boolean   | ipv4 アドレス           | false               | NAT にするかどうか（通常のブリッジではデフォルト値は true で ipv4.address が生成され、fan ブリッジでは常にデフォルト値は true になります）
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
ipv6.nat.address                     | string    | ipv6 アドレス           | -                   | ブリッジからの送信時に使うソースアドレス
ipv6.nat                             | boolean   | ipv6 アドレス           | false               | NAT にするかどうか (未設定の場合はデフォルト値は true になりランダムな ipv6.address が生成されます)
ipv6.nat.order                       | string    | ipv6 アドレス           | before              | 必要な NAT のルールを既存のルールの前に追加するか後に追加するか
ipv6.ovn.ranges                      | string    | -                       | -                   | 子供の OVN ネットワークルーターに使用する IPv6 アドレスの範囲（開始-終了 形式) のカンマ区切りリスト
ipv6.routes                          | string    | ipv6 アドレス           | -                   | ブリッジへルーティングする追加の IPv4 CIDR サブネットのカンマ区切りリスト
ipv6.routing                         | boolean   | ipv6 アドレス           | true                | ブリッジの内外にトラフィックをルーティングするかどうか
maas.subnet.ipv4                     | string    | ipv4 アドレス           | -                   | インスタンスを登録する MAAS IPv4 サブネット (NIC で `network` プロパティを使う場合に有効)
maas.subnet.ipv6                     | string    | ipv6 アドレス           | -                   | インスタンスを登録する MAAS IPv6 サブネット (NIC で `network` プロパティを使う場合に有効)
raw.dnsmasq                          | string    | -                       | -                   | 設定に追加する dnsmasq の設定ファイル
tunnel.NAME.group                    | string    | vxlan                   | 239.0.0.1           | vxlan のマルチキャスト設定 (local と remote が未設定の場合に使われます)
tunnel.NAME.id                       | integer   | vxlan                   | 0                   | vxlan トンネルに使用するトンネル ID
tunnel.NAME.interface                | string    | vxlan                   | -                   | トンネルに使用するホスト・インタフェース
tunnel.NAME.local                    | string    | gre か vxlan            | -                   | トンネルに使用するローカルアドレス (マルチキャスト vxlan の場合は不要)
tunnel.NAME.port                     | integer   | vxlan                   | 0                   | vxlan トンネルに使用するポート
tunnel.NAME.protocol                 | string    | 標準モード              | -                   | トンネリングのプロトコル ("vxlan" か "gre")
tunnel.NAME.remote                   | string    | gre か vxlan            | -                   | トンネルに使用するリモートアドレス (マルチキャスト vxlan の場合は不要)
tunnel.NAME.ttl                      | integer   | vxlan                   | 1                   | マルチキャストルーティングトポロジーに使用する固有の TTL
security.acls                        | string    | -                       | -                   | このネットワークに接続されたNICに適用するカンマ区切りのネットワークACL（[ブリッジの制限](network-acls.md#_4)参照）
security.acls.default.ingress.action | string    | security.acls           | reject              | どの ACL ルールにもマッチしない ingress トラフィックに使うアクション
security.acls.default.egress.action  | string    | security.acls           | reject              | どの ACL ルールにもマッチしない egress トラフィックに使うアクション
security.acls.default.ingress.logged | boolean   | security.acls           | false               | どの ACL ルールにもマッチしない ingress トラフィックをログ出力するかどうか
security.acls.default.egress.logged  | boolean   | security.acls           | false               | どの ACL ルールにもマッチしない egress トラフィックをログ出力するかどうか

これらのキーは lxc コマンドで以下のように設定できます。

```bash
lxc network set <network> <key> <value>
```

### systemd-resolved との統合
LXD が動いているシステムが DNS のルックアップに systemd-resolved を使用している場合、 LXD が名前解決できるドメインを systemd-resolved に指定することができます。
これには systemd-resolved にどのブリッジ、ネームサーバのアドレス、そして DNS ドメインかを伝える必要があります。

例えば、 LXD が `lxdbr0` インターフェースを使用している場合、 `lxc network get lxdbr0 ipv4.address` コマンドで IPv4 アドレス（IPv4 アドレスの代わりに IPv6 アドレスを使うこともできますし、 IPv4 アドレスと IPv6 アドレスの両方を使うこともできます）と `lxc network get lxdbr0 dns.domain` （ドメインが設定されていない場合は上記の表に示されているデフォルト値の `lxd` が使用されます）でドメインを取得します。
そして systemd-resolved に以下のように指定します。

```
systemd-resolve --interface lxdbr0 --set-domain '~lxd' --set-dns n.n.n.n
```

上記の `lxdbr0` は実際のブリッジの名前に、 `n.n.n.n` はネームサーバの実際の（サブネットマスクを除いた） アドレスに置き換えて実行してください。

さらに `lxd` はドメイン名に置き換えてください。
ドメイン名の前の `~` が重要ですので注意してください。
`~` はこのドメインだけをルックアップするためにこのネームサーバを使うように systemd-resolved に指定します。
実際のドメイン名が何であるかにかかわらず `~` を前につけるべきです。
また、 `~` という文字はシェルが展開するかもしれないので、クォートに囲んでエスケープする必要があるかもしれません。

systemd のより新しいリリースでは `systemd-resolve` コマンドは deprecated になっていますが、（これを書いている時点では）後方互換性のためまだ提供されています。
systemd-resolved に伝えるための新しい方法は `resolvectl` コマンドを使うことです。
これは以下の 2 ステップで実行します。

```
resolvectl dns lxdbr0 n.n.n.n
resolvectl domain lxdbr0 '~lxd'
```

この systemd-resolved の設定はブリッジが存在する間のみ存続します。
ですので、リブートと LXD が再起動するたびにこのコマンドを繰り返し実行する必要があります（これを自動化するには下記を参照してください）。
また、これはブリッジの `dns.mode` が `none` でないときにしか機能しないことに注意してください。

`dns.domain` の使用に依存する場合 DNS の名前解決ができるように resolved の DNSSEC を無効にする必要があるかもしれないことに注意してください。
これは `resolved.conf` の `DNSSEC` オプションで設定できます。

LXD が `lxdbr0` インターフェースを作成した場合、 `systemd-resolved` の DNS 設定をシステム起動時に適用するのを自動化するには
以下のような設定を含む systemd の unit ファイル `/etc/systemd/system/lxd-dns-lxdbr0.service` を作成する必要があります。

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

自動起動を有効にし、起動するには以下のようにします。

```
sudo systemctl daemon-reload
sudo systemctl enable --now lxd-dns-lxdbr0
```

`lxdbr0` インタフェースが既に存在する（例: LXD が実行中である場合など）場合、以下のようにサービスが起動済みかを確認できます。

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

### IPv6 プリフィクスサイズ
最適な動作には 64 のプリフィクスサイズが望ましいです。
より大きなサブネット（ 64 より小さいプリフィクス）も正しく動作するでしょうが、SLAAC環境下では有用ではないことが多いでしょう。

IPv6 アドレスの割り当てにステートフル DHCPv6 を使用している場合は、より小さなサブネットも理論的には利用可能ですが、 dnsmasq にきちんとサポートされておらず問題が起きるかもしれません。
これらの 1 つをどうしても使わなければならない場合、静的割り当てか別のスタンドアロンの RA デーモンを使用可能です。

### Firewalld で DHCP, DNS を許可する

firewalld を使用しているホストで LXD が実行する DHCP と DNS サーバにインスタンスがアクセスできるようにするには、ホストのブリッジインターフェースを firewalld の `trusted` ゾーンに追加する必要があります。

（リブート後も設定が残るように）恒久的にこれを行うには以下のコマンドを実行してください。

```
firewall-cmd --zone=trusted --change-interface=<LXD network name> --permanent
```

例えばブリッジネットワークが `lxdbr0` という名前の場合、以下のコマンドを実行します。

```
firewall-cmd --zone=trusted --change-interface=lxdbr0 --permanent
```

これにより LXD 自身のファイアーウォールのルールが有効になります。


### Firewalld に LXD の iptables ルールを制御させるには

firewalld と LXD を一緒に使う場合、 iptables のルールがオーバーラップするかもしれません。例えば firewalld が LXD デーモンより後に起動すると firewalld が LXD の iptables ルールを削除し、 LXD コンテナが外向きのインターネットアクセスが全くできなくなるかもしれません。
これを修正する 1 つの方法は LXD の iptables ルールを firewalld に移譲し、 LXD の iptables ルールは無効にすることです。

最初のステップは [Firewalld で DHCP, DNS を許可する](#allow-dhcp-dns-with-firewalld) ことです。

次に LXD に iptables ルールを設定しないように（firewalld が設定するので）伝えます。
```
lxc network set lxdbr0 ipv4.nat false
lxc network set lxdbr0 ipv6.nat false
lxc network set lxdbr0 ipv6.firewall false
lxc network set lxdbr0 ipv4.firewall false
```

最後に firewalld のルールを LXD の利用ケースに応じて有効にします（この例では、ブリッジインターフェースが `lxdbr0` で付与されている IP の範囲が `10.0.0.0/24` だとしています）。
```
firewall-cmd --permanent --direct --add-rule ipv4 filter INPUT 0 -i lxdbr0 -s 10.0.0.0/24 -m comment --comment "generated by firewalld for LXD" -j ACCEPT
firewall-cmd --permanent --direct --add-rule ipv4 filter OUTPUT 0 -o lxdbr0 -d 10.0.0.0/24 -m comment --comment "generated by firewalld for LXD" -j ACCEPT
firewall-cmd --permanent --direct --add-rule ipv4 filter FORWARD 0 -i lxdbr0 -s 10.0.0.0/24 -m comment --comment "generated by firewalld for LXD" -j ACCEPT
firewall-cmd --permanent --direct --add-rule ipv4 nat POSTROUTING 0 -s 10.0.0.0/24 ! -d 10.0.0.0/24 -m comment --comment "generated by firewalld for LXD" -j MASQUERADE
firewall-cmd --reload
```

firewalld にルールが設定されたかを確認するには以下のようにします。
```
firewall-cmd --direct --get-all-rules
```

警告：上記の手順はフールプルーフなアプローチではなく、不注意にセキュリティリスクをもたらすことにつながる可能性があります。

(network-macvlan)=
## ネットワーク: macvlan

macvlan ネットワークタイプではインスタンスを macvlan NIC を使って親のインターフェースに接続する際に使用するプリセットを指定可能です。
これによりインスタンスの NIC 自体は下層の詳しい設定を一切知ることなく、接続する `network` を単に指定するだけで設定できます。

ネットワーク設定プロパティ:

キー                            | 型        | 条件          | デフォルト           | 説明
:--                             | :--       | :--           | :--                  | :--
maas.subnet.ipv4                | string    | ipv4 アドレス | -                    | インスタンスを登録する MAAS IPv4 サブネット（nic の `network` プロパティを使用する場合）
maas.subnet.ipv6                | string    | ipv6 アドレス | -                    | インスタンスを登録する MAAS IPv6 サブネット（nic の `network` プロパティを使用する場合）
mtu                             | integer   | -             | -                    | 作成するインターフェースの MTU
parent                          | string    | -             | -                    | macvlan NIC を作成する親のインターフェース
vlan                            | integer   | -             | -                    | アタッチする先の VLAN ID
gvrp                            | boolean   | -             | false                | GARP VLAN Registration Protocol を使って VLAN を登録する

(network-sriov)=
## ネットワーク: sriov

sriov ネットワークタイプではインスタンスを sriov NIC を使って親のインターフェースに接続する際に使用するプリセットを指定可能です。
これによりインスタンスの NIC 自体は下層の詳しい設定を一切知ることなく、接続する `network` を単に指定するだけで設定できます。

ネットワーク設定プロパティ:

キー                            | 型        | 条件          | デフォルト            | 説明
:--                             | :--       | :--           | :--                   | :--
maas.subnet.ipv4                | string    | ipv4 アドレス | -                     | インスタンスを登録する MAAS IPv4 サブネット（nic の `network` プロパティを使用する場合）
maas.subnet.ipv6                | string    | ipv6 アドレス | -                     | インスタンスを登録する MAAS IPv6 サブネット（nic の `network` プロパティを使用する場合）
mtu                             | integer   | -             | -                     | 作成するインターフェースの MTU
parent                          | string    | -             | -                     | sriov NIC を作成する親のインターフェース
vlan                            | integer   | -             | -                     | アタッチする先の VLAN ID

(network-ovn)=
## ネットワーク: ovn

ovn ネットワークタイプは OVN SDN を使って論理的なネットワークの作成を可能にします。
これは複数の個別のネットワーク内で同じ論理ネットワークのサブネットを使うような検証環境やマルチテナントの環境で便利です。

LXD の OVN ネットワークはより広いネットワークへの外向きのアクセスを可能にするため既存の管理された LXD のブリッジネットワークに接続できます。
OVN 論理ネットワークからの全ての接続は親のネットワークによって割り当てられた動的 IP に NAT されます。

### スタンドアロンの LXD での OVN の設定

これは外向きの通信のために親のネットワーク lxdbr0 に接続されたスタンドアロンの OVN ネットワークを作成する手順です。

OVN のツールをインストールし、ローカルノードで OVN の統合ブリッジを設定します。

```
sudo apt install ovn-host ovn-central
sudo ovs-vsctl set open_vswitch . \
  external_ids:ovn-remote=unix:/var/run/ovn/ovnsb_db.sock \
  external_ids:ovn-encap-type=geneve \
  external_ids:ovn-encap-ip=127.0.0.1
```

以下を使用して OVN ネットワークとインスタンスを作成します。

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


ネットワークフォワード:

OVN のネットワークサポートは [ネットワークフォワード](network-forwards.md) 参照。

ネットワークピア:

OVN のネットワークピアサポートは [ネットワークピア](network-peers.md) 参照。

ネットワークの設定プロパティ:

キー                                 | 型        | 条件             | デフォルト                  | 説明
:--                                  | :--       | :--              | :--                         | :--
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
ipv6.nat.address                     | string    | ipv6 アドレス    | -                           | ネットワークからの外向きトラフィックに使用されるソースアドレス (アップリンクに `ovn.ingress_mode=routed` が必要)
ipv6.dhcp                            | boolean   | ipv6 アドレス    | true                        | DHCP 上に追加のネットワーク設定を提供するかどうか
ipv6.dhcp.stateful                   | boolean   | ipv6 dhcp        | false                       | DHCP を使ってアドレスを割り当てるかどうか
ipv6.nat                             | boolean   | ipv6 アドレス    | false                       | NAT するかどうか（ipv6.address が未設定の場合デフォルト値は true でランダムな ipv6.address が生成されます）
network                              | string    | -                | -                           | 外部ネットワークへの外向きのアクセスに使うアップリンクのネットワーク
security.acls                        | string    | -                | -                           | このネットワークに接続する NIC に適用するネットワーク ACL のカンマ区切りリスト
security.acls.default.ingress.action | string    | security.acls    | reject                      | どの ACL ルールにもマッチしない ingress トラフィックに使うアクション
security.acls.default.egress.action  | string    | security.acls    | reject                      | どの ACL ルールにもマッチしない egress トラフィックに使うアクション
security.acls.default.ingress.logged | boolean   | security.acls    | false                       | どの ACL ルールにもマッチしない ingress トラフィックをログ出力するかどうか
security.acls.default.egress.logged  | boolean   | security.acls    | false                       | どの ACL ルールにもマッチしない egress トラフィックをログ出力するかどうか

(network-physical)=
## ネットワーク: 物理

物理ネットワークは OVN ネットワークを親インターフェースに接続する際に使用するプリセットの設定を提供します。

ネットワーク設定プロパティ:


キー                            | 型        | 条件             | デフォルト                                | 説明
:--                             | :--       | :--              | :--                                       | :--
bgp.peers.NAME.address          | string    | bgp server       | -                                         | `ovn` ダウンストリームネットワークで使用するピアアドレス (IPv4 か IPv6)
bgp.peers.NAME.asn              | integer   | bgp server       | -                                         | `ovn` ダウンストリームネットワークで使用する AS 番号
bgp.peers.NAME.password         | string    | bgp server       | - (パスワード無し)                        | `ovn` ダウンストリームネットワークで使用するピアのセッションパスワード（省略可能）
maas.subnet.ipv4                | string    | ipv4 アドレス    | -                                         | インスタンスを登録する MAAS IPv4 サブネット (NIC で `network` プロパティを使う場合に有効)
maas.subnet.ipv6                | string    | ipv6 アドレス    | -                                         | インスタンスを登録する MAAS IPv6 サブネット (NIC で `network` プロパティを使う場合に有効)
mtu                             | integer   | -                | -                                         | 作成するインターフェースの MTU
parent                          | string    | -                | -                                         | sriov NIC を作成する親のインターフェース
vlan                            | integer   | -                | -                                         | アタッチする先の VLAN ID
gvrp                            | boolean   | -                | false                                     | GARP VLAN Registration Protocol を使って VLAN を登録する
ipv4.gateway                    | string    | 標準モード       | -                                         | ゲートウェイとネットワークの IPv4 アドレス（CIDR表記）
ipv4.ovn.ranges                 | string    | -                | -                                         | 子供の OVN ネットワークルーターに使用する IPv4 アドレスの範囲（開始-終了 形式) のカンマ区切りリスト
ipv4.routes                     | string    | ipv4 アドレス    | -                                         | 子供の OVN ネットワークの ipv4.routes.external 設定で利用可能な追加の IPv4 CIDR サブネットのカンマ区切りリスト
ipv4.routes.anycast             | boolean   | ipv4 アドレス    | false                                     | 複数のネットワーク／NICで同時にオーバーラップするルートが使われることを許可するかどうか
ipv6.gateway                    | string    | 標準モード       | -                                         | ゲートウェイとネットワークの IPv6 アドレス（CIDR表記）
ipv6.ovn.ranges                 | string    | -                | -                                         | 子供の OVN ネットワークルーターに使用する IPv6 アドレスの範囲（開始-終了 形式) のカンマ区切りリスト
ipv6.routes                     | string    | ipv6 アドレス    | -                                         | 子供の OVN ネットワークの ipv6.routes.external 設定で利用可能な追加の IPv6 CIDR サブネットのカンマ区切りリスト
ipv6.routes.anycast             | boolean   | ipv6 アドレス    | false                                     | 複数のネットワーク／NICで同時にオーバーラップするルートが使われることを許可するかどうか
dns.nameservers                 | string    | 標準モード       | -                                         | 物理ネットワークの DNS サーバ IP のリスト
ovn.ingress\_mode               | string    | 標準モード       | l2proxy                                   | OVN NIC の外部 IP アドレスがアップリンクネットワークで広告される方法を設定します。 `l2proxy` (proxy ARP/NDP) か `routed` です。

## BGP の統合
LXD は BGP サーバとして機能でき、アップストリームの BGP ルーターとセッションを確立し LXD が使用しているアドレスとサブネットを広告できます。

これにより LXD サーバやクラスターが内部／外部のアドレス空間を直接使い、正しいホストにルーティングされた特定のサブネットやアドレスをターゲットインスタンスにフォワードできます。

このためには `core.bgp_address`, `core.bgp_asn` と `core.bgp_routerid` が設定されている必要があります。
これらが設定されると LXD は BGP セッションのリッスンを開始します。

ピアは `bridged` と `physical` で管理されたネットワーク上に定義できます。さらに `bridged` の場合は next-hop をオーバーライドするためにサーバ毎の設定キーの組が利用できます。それらが指定されない場合は next-hop はデフォルトとして BGP セッションに使用されるアドレスになります。

`physical` ネットワークの場合はアップリンクのネットワークが利用可能なサブネットのリストと BGP 設定を持つような `ovn` ネットワークに使用されます。
親のネットワークが一旦設定されると、子の OVN ネットワークは BGP で広告された外部のサブネットとアドレスを受け取り next-hop は問題のネットワークの OVN ルーターアドレスに設定されます。

現在公開されるアドレスとネットワークは以下のとおりです。
 - `nat` プロパティが `true` に設定されない場合はネットワークの `ipv4.address` か `ipv6.address`
 - `nat` プロパティが設定される場合はネットワークの `ipv4.address` と `ipv6.address`
 - `ipv4.routes.external` か `ipv6.routes.external` 経由で定義されるインスタンスの NIC ルート

現時点では、特定のピアに特定のルートやアドレスのみを公開する方法はありません。代わりにアップストリームのルーターでプリフィクスをフィルターすることを現状ではお勧めします。
