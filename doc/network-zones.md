# ネットワークゾーンの設定 <!-- Network Zones configuration -->
ネットワークゾーンは LXD のネットワークの DNS レコードを保持するのに使用します。
<!--
Network zones are used to hold DNS records for LXD networks.
-->

各ネットワークは下記の最大3つに関連します。
<!--
Each network can be related to up to 3 zones for:
-->

 - 正引き DNS レコード <!-- Forward DNS records -->
 - IPv4 逆引き DNS レコード <!-- IPv4 reverse DNS records -->
 - IPv6 逆引き DNS レコード <!-- IPv6 reverse DNS records -->

これらはネットワーク設定の `dns.zone.forward`, `dns.zone.reverse.ipv4`,
`dns.zone.reverse.ipv6` で制御します。
LXD は全てのインスタンス、ネットワークゲートウェイ、ダウンストリーム (下流)
のネットワークポートの全てに対して正引きと逆引きのレコードを自動で管理します。
<!--
This is controlled through `dns.zone.forward`, `dns.zone.reverse.ipv4`
and `dns.zone.reverse.ipv6` in network configuration. LXD will then be
automatically managing forward and reverse records for all instances,
network gateways and downstream network ports.
-->

組み込みの DNS サーバーを有効にするには、サーバー設定内の `core.dns_address`
を設定する必要があります。
<!--
To enable the built-in DNS server, `core.dns_address` must be set in the
server configuration.
-->

組み込みの DNS サーバーは AXFR 経由でのゾーン転送のみをサポートしており、
DNS レコードへの直接の問い合わせはできません。
つまりこの機能は外部の DNS サーバー (bind9, nsd, ...) の使用を前提としています。
外部の DNS サーバーが LXD からの全体のゾーンを転送し、有効期限を過ぎたら更新し、
DNS 問い合わせに対する管理権限を持つ応答 (authoritative answers) を提供します。
<!--
The built-in DNS server only supports zone transfers through AXFR, it
cannot be directly queried for DNS records. This means that this feature
expects the use of an external DNS server (bind9, nsd, ...) which will
transfer the entire zone from LXD, refresh it upon expiry and provide
authoritative answers to DNS requests.
-->

ゾーン転送の認証はゾーン毎に設定され、各ゾーンでピアごとに IP アドレスと TSIG キーを設定して、
TSIG キーベースの認証を行います。
<!--
Authentication for zone transfer is configured on a per-zone basis with
peers defined in zone configuration and a combination of IP address
matching and TSIG key based authentication.
-->

ゾーンはプロジェクトに属し、プロジェクトの `networks` 機能に紐づけられます。
<!--
Zones belong to projects and are tied to the `networks` features of projects.
-->

ゾーン名は複数のプロジェクトをまたいでグローバルにユニークでなければなりません。
そのため、別のプロジェクト内の既存のゾーンのせいでゾーンの作成がエラーになることがありえます。
<!--
Zone names must be globally unique, even across projects, so it's
possible to get a creation error due to a zone already existing in
another project.
-->

`restricted.networks.zones` プロジェクト設定キーによりプロジェクトを
特定のドメインとサブドメインに制限できます。
<!--
It is possible to restrict projects to specific domains and sub-domains
through the `restricted.networks.zones` project configuration key.
-->

## プロパティー <!-- Properties -->
ネットワークゾーンのプロパティーは以下の通りです。
<!--
The following are network zone properties:
-->

プロパティー <!-- Property --> | 型 <!-- Type --> | 必須 <!-- Required --> | デフォルト値 <!-- Default --> | 説明 <!-- Description -->
:--                 | :--        | :--      | -       | :--
peers.NAME.address  | string     | no       | -       | DNS サーバーの IP アドレス <!-- IP address of a DNS server -->
peers.NAME.key      | string     | no       | -       | サーバー用の TSIG キー <!-- TSIG key for the server -->
dns.nameservers     | string set | no       | -       | (NS レコード用の) DNS サーバーの FQDN のカンマ区切りリスト <!-- Comma separated list of DNS server FQDNs (for NS records) -->
network.nat         | bool       | no       | true    | NAT されたサブネットのレコードを生成するかどうか <!-- Whether to generate records for NAT-ed subnets -->

さらに、 `user.` キーの名前空間もユーザー提供の自由形式のキー・バリュー用にサポートされています。
<!--
Additionally the `user.` key namespace is also supported for user-provided free-form key/value.
-->
