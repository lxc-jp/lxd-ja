(expp-security)=
# セキュリティについて

```{youtube} https://www.youtube.com/watch?v=cOOzKdYHkus
```

% Include content from [../../README.md](../../README.md)
```{include} ../../README.md
    :start-after: <!-- Include start security -->
    :end-before: <!-- Include end security -->
```

詳細な情報は以下のセクションを参照してください。

セキュリティ上の問題を発見した場合、その問題の報告方法については [LXDのセキュリティポリシー](https://github.com/lxc-jp/lxd-ja/blob/master/SECURITY.md) (原文: [LXD security policy](https://github.com/lxc/lxd/blob/master/SECURITY.md))を参照してください。 <!-- wokeignore:rule=master -->

## サポートされているバージョン

サポートされていないバージョンのLXDは実運用環境では絶対に使用しないでください。

% Include content from [../../SECURITY.md](../../SECURITY.md)
```{include} ../../SECURITY.md
    :start-after: <!-- Include start supported versions -->
    :end-before: <!-- Include end supported versions -->
```

(security-daemon-access)=
## LXDデーモンへのアクセス

LXDはUnixソケットを介してローカルにアクセスできるデーモンで、設定されていれば{abbr}`TLS(Transport Layer Security)`ソケットを介してリモートにアクセスすることもできます。
ソケットにアクセスできる人は、ホストデバイスやファイルシステムをアタッチしたり、すべてのインスタンスのセキュリティ機能をいじったりするなど、LXDを完全に制御することができます。

したがって、デーモンへのアクセスを信頼できるユーザに制限するようにしてください。

### LXD デーモンへのローカルアクセス

LXDデーモンはrootで動作し、ローカル通信用のUnixソケットを提供します。
LXD のアクセス制御は、グループメンバーシップに基づいて行われます。
root ユーザーと `lxd` グループのすべてのメンバーがローカルデーモンと対話できます。

````{important}
% Include content from [../../README.md](../../README.md)
```{include} ../../README.md
    :start-after: <!-- Include start security note -->
    :end-before: <!-- Include end security note -->
```
````

(security_remote_access)=
### リモート API へのアクセス

デフォルトでは、デーモンへのアクセスはローカルでのみ可能です。
`core.https_address`という設定オプション({doc}`server`参照)を設定することで、同じAPIを{abbr}`TLS (Transport Layer Security)`ソケットでネットワーク上に公開することができます。
リモートクライアントは、LXDに接続して、公開用にマークされたイメージにアクセスできます。

リモートクライアントがAPIにアクセスできるように、信頼できるクライアントとして認証する方法がいくつかあります。
詳細は{doc}`authentication`を参照してください。

本番環境では、`core.https_address`に、(ホスト上の任意のアドレスではなく)サーバーが利用可能な単一のアドレスを設定する必要があります。
さらに、許可されたホスト/サブネットからのみLXDポートへのアクセスを許可するファイアウォールルールを設定する必要があります。

## コンテナのセキュリティ

LXDコンテナはセキュリティのために幅広い機能を使うことができます。

デフォルトでは、コンテナは *非特権* (*unprivileged*) であり、ユーザーネームスペース内で動作することを意味し、コンテナ内のユーザーの能力を、コンテナが所有するデバイスに対する制限された権限を持つホスト上の通常のユーザーに制限します。

コンテナ間のデータ共有が必要ない場合は、`security.idmap.isolated`({ref}`instance-options-security`参照)を有効にすることで、各コンテナに対して重複しないUID/GIDマップを使用し、他のコンテナに対する潜在的な{abbr}`DoS(サービス拒否)`攻撃を防ぐことができます。

LXDはまた、*特権* (*privileged*) コンテナを実行することができます。
しかし、これは(訳注:コンテナ内だけで)安全にroot権限を使えるわけではなく、そのようなコンテナの中でルートアクセスを持つユーザは、閉じ込められた状態から逃れる方法を見つけるだけでなく、ホストをDoSすることができてしまう点に注意してください。

コンテナのセキュリティと私たちが使っているカーネルの機能についてのより詳細な情報は
[LXCセキュリティページ](https://linuxcontainers.org/ja/lxc/security/)にあります。

### コンテナ名の漏洩

デフォルトの設定ではシステム上の全ての cgroup と、さらに転じて、全ての実行中のコンテナを一覧表示することが簡単にできてしまいます。

コンテナを開始する前に `/sys/kernel/slab` と `/proc/sched_debug` へのアクセスをブロックすることでコンテナ名の漏洩を防げます。
このためには以下のコマンドを実行してください。

    chmod 400 /proc/sched_debug
    chmod 700 /sys/kernel/slab/

## ネットワークセキュリティ

ネットワークインタフェースは必ず安全に設定してください。
どのような点を考慮すべきかは、使用するネットワークモードによって異なります。

### ブリッジ型NICのセキュリティ

LXDのデフォルトのネットワークモードは、各インスタンスが接続する「管理された」プライベートネットワークのブリッジを提供することです。
このモードでは、ホスト上に`lxdbr0`というインタフェースがあり、それがインスタンスのブリッジとして機能します。

ホストは、管理されたブリッジごとに`dnsmasq`のインスタンスを実行し、IPアドレスの割り当てと、権威DNSおよび再帰DNSサービスの提供を担当します。

DHCPv4を使用しているインスタンスには、IPv4アドレスが割り当てられ、インスタンス名のDNSレコードが作成されます。
これにより、インスタンスがDHCPリクエストに偽のホスト名情報を提供して、DNSレコードを偽装することができなくなります。

`dnsmasq`サービスは、IPv6のルータ広告機能も提供します。
つまり、インスタンスはSLAACを使って自分のIPv6アドレスを自動設定するので、`dnsmasq`による割り当ては行われません。
しかし、DHCPv4を使用しているインスタンスは、SLAAC IPv6アドレスに相当するAAAAのDNSレコードも取得します。
これは、インスタンスがIPv6アドレスを生成する際に、IPv6プライバシー拡張を使用していないことを前提としています。

このデフォルト構成では、DNS名を偽装することはできませんが、インスタンスはイーサネットブリッジに接続されており、希望するレイヤー2トラフィックを送信することができます。これは、信頼されていないインスタンスがブリッジ上でMACまたはIPの偽装を効果的に行うことができることを意味します。

デフォルトの設定では、ブリッジに接続されたインスタンスがブリッジに(潜在的に悪意のある)IPv6ルータ広告を送信することで、LXDホストのIPv6ルーティングテーブルを修正することも可能です。
これは、`lxdbr0`インターフェイスが`/proc/sys/net/ipv6/conf/lxdbr0/accept_ra`を`2`に設定して作成されているためで、`forwarding`が有効であるにもかかわらず、LXDホストがルーター広告を受け入れることを意味しています(詳細は[`/proc/sys/net/ipv4/*` Variables](https://www.kernel.org/doc/Documentation/networking/ip-sysctl.txt)を参照してください)。

しかし、LXDはいくつかのブリッジ型{abbr}`NIC(Network interface controller)`セキュリティ機能を提供しており、インスタンスがネットワーク上に送信することを許可されるトラフィックの種類を制御するために使用することができます。
これらのNIC設定は、インスタンスが使用しているプロファイルに追加する必要がありますが、以下のように個々のインスタンスに追加することもできます。

ブリッジ型NICには、以下のようなセキュリティ機能があります。

キー                      | タイプ | デフォルト | 必須 | 説明
:--                       | :--    | :--        | :--  | :--
`security.mac_filtering`  | bool   | `false`    | no   | インスタンスが他のインスタンスの MAC アドレスを詐称することを防ぐ。
`security.ipv4_filtering` | bool   | `false`    | no   | インスタンスが他のインスタンスの IPv4 アドレスになりすますことを防ぎます(`mac_filtering` を有効にします)。
`security.ipv6_filtering` | bool   | `false`    | no   | インスタンスが他のインスタンスの IPv6 アドレスになりすますことを防ぎます(`mac_filtering` を有効にします)。

プロファイルで設定されたデフォルトのブリッジ型NICの設定は、インスタンスごとに以下の方法で上書きすることができます。

```
lxc config device override <instance> <NIC> security.mac_filtering=true
```

これらの機能を併用することで、ブリッジに接続されているインスタンスがMACアドレスやIPアドレスを詐称することを防ぐことができます。
これらのオプションは、ホスト上で利用可能なものに応じて、`xtables`(`iptables`、`ip6tables`、`ebtables`)または`nftables`を使用して実装されます。

これらのオプションは、ネストされたコンテナが異なるMACアドレスを持つ親ネットワークを使用すること(ブリッジされたNICや`macvlan` NICを使用すること)を効果的に防止することができるのは注目に値します。

IPフィルタリング機能は、スプーフィングされたIPを含むARPおよびNDPアドバタイジングをブロックし、スプーフィングされたソースアドレスを含むすべてのパケットをブロックします。

`security.ipv4_filtering`または`security.ipv6_filtering`が有効で、(`ipvX.address=none`またはブリッジでDHCPサービスが有効になっていないため)インスタンスにIPアドレスが割り当てられない場合、そのプロトコルのすべてのIPトラフィックがインスタンスからブロックされます。

`security.ipv6_filtering` が有効な場合、IPv6 のルータ広告がインスタンスからブロックされます。

`security.ipv4_filtering`または`security.ipv6_filtering`が有効な場合、ARP、IPv4またはIPv6ではないイーサネットフレームはすべてドロップされます。
これにより、スタックされたVLAN `Q-in-Q` (802.1ad) フレームがIPフィルタリングをバイパスすることを防ぎます。

### ルート化されたNICのセキュリティ

「ルーテッド」と呼ばれる別のネットワークモードがあります。
このモードでは、コンテナとホストの間に仮想イーサネットデバイペアを提供します。
このネットワークモードでは、LXDホストがルータとして機能し、コンテナのIP宛のトラフィックをコンテナの`veth`インターフェイスに誘導するスタティックルートがホストに追加されます。

デフォルトでは、コンテナからのルータ広告がLXDホスト上のIPv6ルーティングテーブルを変更するのを防ぐために、ホスト上に作成された`veth`インタフェースは、その`accept_ra`設定が無効になっています。
それに加えて、コンテナが持っていることをホストが知らないIPに対するソースアドレスの偽装を防ぐために、ホスト上の`rp_filter`が`1`に設定されています。
