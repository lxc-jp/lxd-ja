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

各ネットワークは下記の最大3つに関連します。

 - 正引き DNS レコード
 - IPv4 逆引き DNS レコード
 - IPv6 逆引き DNS レコード

LXD は全てのインスタンス、ネットワークゲートウェイ、ダウンストリーム (下流)
のネットワークポートの全てに対して正引きと逆引きのレコードを自動で管理し、
オペレータのプロダクションの DNS サーバへのゾーン転送のためのこれらのゾーンを提供します。

## 生成されるレコード

例えば、あなたのネットワークで `lxd.example.net` の正引き DNS レコードのゾーンを設定した場合、
以下の DNS 名を解決するレコードを生成します。

- ネットワーク内の全てのインスタンスに対して: `<instance_name>.lxd.example.net`
- ネットワークゲートウェイに対して: `<network_name>.gw.lxd.example.net`
- ダウンストリームネットワークポートに対して (ダウンストリーム OVN ネットワークを持つアップリンクのネットワーク上に設定されれうネットワークゾーンに対して): `<project_name>-<downstream_network_name>.uplink.lxd.example.net`

ゾーン設定に対して生成されたレコードは `dig` コマンドで確認できます。
例えば、 `dig @<DNS_server_IP> -p 1053 axfr lxd.example.net` と実行すると以下のように出力されます。

```bash
lxd.example.net.              3600  IN  SOA lxd.example.net. hostmaster.lxd.example.net. 1648118965 120 60 86400 30
default-my-ovn.uplink.lxd.example.net. 300 IN A 192.0.2.100
my-instance.lxd.example.net.  300   IN  A   192.0.2.76
my-uplink.gw.lxd.example.net. 300   IN  A   192.0.2.1
foo.lxd.example.net.          300   IN  A   8.8.8.8
lxd.example.net.              3600  IN  SOA lxd.example.net. hostmaster.lxd.example.net. 1648118965 120 60 86400 30
```

`192.0.2.0/24` を使用するネットワークに `2.0.192.in-addr.arpa` の IPv4 逆引き DNS レコードのゾーンを設定すると、例えば `192.0.2.100` に対する逆引き DNS レコードを生成します。

## 組み込みの DNS サーバを有効にする

ネットワークゾーンを使用するには、組み込みの DNS サーバを有効にする必要があります。

そのためには、 LXD サーバのローカルアドレスに `core.dns_address` 設定オプション ({ref}`server` 参照) を設定してください。

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