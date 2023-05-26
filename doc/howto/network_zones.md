---
discourse: 12033,13128
---

(network-zones)=
# ネットワークゾーンを設定するには

```{note}
ネットワークゾーンは {ref}`network-ovn` と  {ref}`network-bridge` で利用できます。
```

```{youtube} https://www.youtube.com/watch?v=2MqpJOogNVQ
```

ネットワークゾーンは LXD のネットワークの DNS レコードを保持するのに使用します。

ネットワークゾーンを使うと全てのインスタンスの有効な正引きと逆引きのレコードを自動的に維持できます。
多くのネットワークにまたがる複数のインスタンスからなる LXD クラスタを運用する際に有用です。

各インスタンスに DNS レコードを持つとインスタンス上のネットワークサービスにアクセスするのがより簡単になります。
また例えば外部への SMTP サービスをホストする際にも重要です。
インスタンスに正しい正引きと逆引きの DNS エントリがないと、送信されたメールが潜在的なスパムと判定されてしまうかもしれません。

各ネットワークは異なるゾーンに関連します。

- 正引き DNS レコード - カンマ区切りの複数のゾーン (プロジェクトごとに最大1つ)
- IPv4 逆引き DNS レコード - 単一のゾーン
- IPv6 逆引き DNS レコード - 単一のゾーン

LXD は全てのインスタンス、ネットワークゲートウェイ、ダウンストリーム (下流)
のネットワークポートの全てに対して正引きと逆引きのレコードを自動で管理し、
オペレータのプロダクションの DNS サーバへのゾーン転送のためのこれらのゾーンを提供します。

## プロジェクトビュー

プロジェクトには  `features.networks.zones` 機能があります。デフォルトでは無効です。
これは新しいネットワークゾーンがどのプロジェクト内に作成されるかを制御します。
この機能を有効にすると新しいゾーンはプロジェクト内に作成されますが、無効の場合はデフォルトプロジェクト内に作成されます。

これにより、複数のプロジェクトがデフォルトプロジェクト(すなわち`features.networks=false`と設定されたプロジェクト)内のネットワークを共有できるようになり、共有されたネットワークに対してプロジェクト指向の(プロジェクト内のインスタンスのアドレスのみを含むような)「ビュー」を提供するプロジェクト固有のDNSゾーンを持てるようになります。

## 生成されるレコード

### 正引きレコード

例えば、あなたのネットワークで `lxd.example.net` の正引き DNS レコードのゾーンを設定した場合、
以下の DNS 名を解決するレコードを生成します。

- ネットワーク内の全てのインスタンスに対して: `<instance_name>.lxd.example.net`
- ネットワークゲートウェイに対して: `<network_name>.gw.lxd.example.net`
- ダウンストリームネットワークポートに対して (ダウンストリーム OVN ネットワークを持つアップリンクのネットワーク上に設定されれうネットワークゾーンに対して): `<project_name>-<downstream_network_name>.uplink.lxd.example.net`
- ゾーンに手動で追加されたレコード

ゾーン設定に対して生成されたレコードは `dig` コマンドで確認できます。
これは`core.dns_address`が`<DNS_server_IP>:<DNS_server_PORT>`に設定されていることを前提としています。（その設定オプションを設定すると、バックエンドはすぐにそのアドレスでサービスを開始します。）

特定のゾーンに対して`dig`リクエストが許可されるようにするためには、そのゾーンの`peers.NAME.address`設定オプションを設定する必要があります。`NAME`はランダムなもので構いません。値は、`dig`が呼び出し元のIPアドレスと一致しなければなりません。同じランダムな`NAME`の`peers.NAME.key`は未設定のままにしておく必要があります。

例: `lxc network zone set lxd.example.net peers.whatever.address=192.0.2.1`

```{note}
`dig`が呼び出し元の同じマシンのアドレスであるだけでは十分ではありません。それは、`lxd`内のDNSサーバーが正確なリモートアドレスと考えるものと文字列で一致する必要があります。`dig`は`0.0.0.0`にバインドするため、必要なアドレスはおそらく、あなたが`core.dns_address`に提供したものと同じです。
```

例えば、`dig @<DNS_server_IP> -p <DNS_server_PORT> axfr lxd.example.net`と実行すると以下のような出力がでるかもしれません。

```{terminal}
:input: dig @192.0.2.200 -p 1053 axfr lxd.example.net
lxd.example.net.                        3600 IN SOA  lxd.example.net. ns1.lxd.example.net. 1669736788 120 60 86400 30
lxd.example.net.                        300  IN NS   ns1.lxd.example.net.
lxdtest.gw.lxd.example.net.             300  IN A    192.0.2.1
lxdtest.gw.lxd.example.net.             300  IN AAAA fd42:4131:a53c:7211::1
default-ovntest.uplink.lxd.example.net. 300  IN A    192.0.2.20
default-ovntest.uplink.lxd.example.net. 300  IN AAAA fd42:4131:a53c:7211:216:3eff:fe4e:b794
c1.lxd.example.net.                     300  IN AAAA fd42:4131:a53c:7211:216:3eff:fe19:6ede
c1.lxd.example.net.                     300  IN A    192.0.2.125
manualtest.lxd.example.net.             300  IN A    8.8.8.8
lxd.example.net.                        3600 IN SOA  lxd.example.net. ns1.lxd.example.net. 1669736788 120 60 86400 30
```

### 逆引きレコード

`192.0.2.0/24` を使用するネットワークに `2.0.192.in-addr.arpa` の IPv4 逆引き DNS レコードのゾーンを設定すると、正引きゾーンの1つを経由してネットワークを参照する全てのプロジェクトからのアドレスの逆引き `PTR` DNSレコードを生成します。

例えば `dig @<DNS_server_IP> -p <DNS_server_PORT> axfr 2.0.192.in-addr.arpa` を実行すると以下のような出力が得られるかもしれません。

```{terminal}
:input: dig @192.0.2.200 -p 1053 axfr 2.0.192.in-addr.arpa

2.0.192.in-addr.arpa.                  3600 IN SOA  2.0.192.in-addr.arpa. ns1.2.0.192.in-addr.arpa. 1669736828 120 60 86400 30
2.0.192.in-addr.arpa.                  300  IN NS   ns1.2.0.192.in-addr.arpa.
1.2.0.192.in-addr.arpa.                300  IN PTR  lxdtest.gw.lxd.example.net.
20.2.0.192.in-addr.arpa.               300  IN PTR  default-ovntest.uplink.lxd.example.net.
125.2.0.192.in-addr.arpa.              300  IN PTR  c1.lxd.example.net.
2.0.192.in-addr.arpa.                  3600 IN SOA  2.0.192.in-addr.arpa. ns1.2.0.192.in-addr.arpa. 1669736828 120 60 86400 30
```

## 組み込みの DNS サーバを有効にする

ネットワークゾーンを使用するには、組み込みの DNS サーバを有効にする必要があります。

そのためには、 LXD サーバのローカルアドレスに `core.dns_address` 設定オプション({ref}`server-options-core`参照)を設定してください。
既存のDNSとの衝突を避けるためポート53を使用しないことをお勧めします。
これは DNS サーバがリッスンするアドレスです。
LXD クラスタの場合、アドレスは各クラスタメンバーによって異なるかもしれないことに注意してください。

```{note}
組み込みの DNS サーバは AXFR 経由でのゾーン転送のみをサポートしており、
DNS レコードへの直接の問い合わせはできません。
つまりこの機能は外部の DNS サーバ (`bind9`, `nsd`, ...) の使用を前提としています。
外部の DNS サーバが LXD からの全体のゾーンを転送し、有効期限を過ぎたら更新し、
DNS 問い合わせに対する管理権限を持つ応答 (authoritative answers) を提供します。

ゾーン転送の認証はゾーン毎に設定され、各ゾーンでピアごとに IP アドレスと TSIG キーを設定して、
TSIG キーベースの認証を行います。
```

## ネットワークゾーンの作成と設定

ネットワークゾーンの作成には以下のコマンドを使用します。

```bash
lxc network zone create <network_zone> [configuration_options...]
```

以下の例は正引き DNS レコードのゾーン、IPv4 逆引き DNS レコードのゾーン、IPv6 逆引き DNS レコードのゾーンを作成する方法を示しています。

```bash
lxc network zone create lxd.example.net
lxc network zone create 2.0.192.in-addr.arpa
lxc network zone create 1.0.0.0.1.0.0.0.8.b.d.0.1.0.0.2.ip6.arpa
```

```{note}
ゾーン名は複数のプロジェクトをまたいでグローバルにユニークでなければなりません。
そのため、別のプロジェクト内の既存のゾーンのせいでゾーンの作成がエラーになることがありえます。
```

ネットワークを作成するときに設定オプションを指定できますし、後から以下のコマンドで設定もできます。

```bash
lxc network zone set <network_zone> <key>=<value>
```

YAML 形式でネットワークゾーンを編集するには以下のコマンドを使用します。

```bash
lxc network zone edit <network_zone>
```

### 設定オプション

ネットワークゾーンで利用可能な設定オプションは下記のとおりです。

キー                 | 型         | 必須 | デフォルト値 | 説明
:--                  | :--        | :--  | -            | :--
`peers.NAME.address` | string     | no   | -            | DNS サーバの IP アドレス
`peers.NAME.key`     | string     | no   | -            | サーバの TSIG キー
`dns.nameservers`    | string set | no   | -            | (NS レコード用の) DNS サーバの FQDN のカンマ区切りリスト
`network.nat`        | bool       | no   | `true`       | NAT されたサブネットのレコードを生成するかどうか
`user.*`             | *          | no   | -            | ユーザー提供の自由形式のキー・バリューペア

```{note}
`tsig-keygen`を使用してTSIGキーを生成するとき、キー名は`<zone_name>_<peer_name>.`というフォーマットに従わなければなりません。たとえば、ゾーン名が`lxd.example.net`でピア名が`bind9`の場合、キー名は`lxd.example.net_bind9.`でなければなりません。この形式に従わない場合、ゾーン転送が失敗する可能性があります。
```

## ネットワークにネットワークゾーンを追加する

ネットワークにゾーンを追加するにはネットワーク設定内に対応する設定オプションを設定します。

- 正引き DNS レコードには: `dns.zone.forward`
- IPv4 逆引き DNS レコードには: `dns.zone.reverse.ipv4`
- IPv6 逆引き DNS レコードには: `dns.zone.reverse.ipv6`

例えば

```bash
lxc network set <network_name> dns.zone.forward="lxd.example.net"
```

ゾーンはプロジェクトに属し、プロジェクトの `networks` 機能に紐づきます。
プロジェクトの `restricted.networks.zones` 設定キーを使ってプロジェクトを指定のドメインとサブドメインに制限できます。

## カスタムレコードを追加する

ネットワークゾーンは、全てのインスタンス、ネットワークゲートウェイ、ダウンストリームネットワークポートに対して
正引きと逆引きレコードを自動的に生成します。

そのためには `lxc network zone record` コマンドを使用します。

### レコードを作成する

レコードを作成するには以下のコマンドを使用します。

```bash
lxc network zone record create <network_zone> <record_name>
```

このコマンドはエントリ無しの空のレコードを作成しネットワークゾーンに追加します。

#### レコードのプロパティ

レコードは以下のプロパティを持ちます。

プロパティ    | 型         | 必須 | 説明
:--           | :--        | :--  | :--
`name`        | string     | yes  | レコードのユニークな名前
`description` | string     | no   | レコードの説明
`entries`     | entry list | no   | DNS エントリのリスト
`config`      | string set | no   | キー／バリュー形式の設定オプション (`user.*` カスタムキーのみサポート)

### エントリを追加または削除する

レコードにエントリを追加するには以下のコマンドを使います。

```bash
lxc network zone record entry add <network_zone> <record_name> <type> <value> [--ttl <TTL>]
```

このコマンドはレコードに指定した型と値を持つ DNS エントリを追加します。

例えば、デュアルスタックのウェブサーバを作成するには以下のような 2 つのエントリを持つレコードを追加します。

```bash
lxc network zone record entry add <network_zone> <record_name> A 1.2.3.4
lxc network zone record entry add <network_zone> <record_name> AAAA 1234::1234
```

エントリにカスタムの time-to-live (秒で指定) を設定するには `--ttl` フラグが使えます。
指定しない場合、デフォルトの 300 秒になります。

(`lxc network zone record edit` でレコード全体を編集するのを除いて) エントリを編集は出来ませんが、以下のコマンドでエントリを削除できます。

```bash
lxc network zone record entry remove <network_zone> <record_name> <type> <value>
```
