# ネットワークフォワード設定 <!-- Network Forward configuration -->

ネットワークフォワードは外部 IP アドレス（あるいは外部 IP アドレスの特定のポート）をフォワード設定が属するネットワーク内の内部 IP アドレス（あるいは内部 IP アドレスの特定のポート）にフォワードする機能です。
<!--
Network forwards allow an external IP address (or specific ports on it) to be forwarded to an internal IP address
(or specific ports on it) in the network that the forward belongs to.
-->

各フォワード設定には単一の外部リッスンアドレスが必要で、オプショナルでデフォルトのターゲットアドレス（これによりポート指定にマッチしないトラフィックをそのアドレスにフォワードします）と複数のポート指定（これによりリッスンアドレス上の特定のポートをデフォルトのターゲットアドレスとは異なるアドレス上の指定のポートにフォワードします）を指定できます。
<!--
Each forward requires a single external listen address, combined with an optional default target address
(which causes any traffic not matched by a port specification to be forwarded to it) and an optional set of port
specifications (that allow specific port(s) on the listen address to be forwarded to specific port(s) on a target
address that is different than the default target address).
-->

全てのターゲットアドレスはフォワード設定が関連付けられるネットワークと同じサブネット内でなければなりません。
<!--
All target addresses must be within the same subnet as the network that the forward is associated to.
-->

デフォルトのターゲットアドレスはフォワードの `config` セットの `target_address` フィールドで指定されます。
<!--
The default target address is specified in the forward's `config` set using the `target_address` field.
-->

指定可能なリッスンアドレスはフォワード設定が関連付けられる [ネットワーク種別](#network-types) によって異なります。
<!--
The listen addresses allowed vary depending on which [network type](#network-types) the forward is associated to.
-->

## プロパティー <!-- Properties -->
ネットワークフォワードのプロパティーには以下のものがあります。
<!--
The following are network forward properties:
-->

プロパティー <!-- Property --> | 型 <!-- Type --> | 必須 <!-- Required --> | 説明 <!-- Description -->
:--              | :--        | :--      | :--
listen\_address  | string     | yes      | リッスンする IP アドレス <!-- IP address to listen on -->
description      | string     | no       | ネットワークフォワードの説明 <!-- Description of Network Forward -->
config           | string set | no       | 設定のキーバリューペア（`target_address` と `user.*` カスタムキーのみサポート） <!-- Config key/value pairs (Only `target_address` and `user.*` custom keys supported) -->
ports            | port list  | no       | ネットワークフォワードのポートリスト <!-- Network forward port list -->

ネットワークフォワードポートのプロパティーには以下のものがあります。
<!--
Network forward ports have the following properties:
-->

プロパティー <!-- Property --> | 型 <!-- Type --> | 必須 <!-- Required --> | 説明 <!-- Description -->
:--               | :--        | :--      | :--
protocol          | string     | yes      | ポートのプロトコル (`tcp` or `udp`) <!-- Protocol for port (`tcp` or `udp`) -->
listen\_port      | string     | yes      | リッスンするポート (例 `80,90-100`) <!-- Listen port(s) (e.g. `80,90-100`) -->
target\_address   | string     | yes      | フォワード先の IP アドレス <!-- IP address to forward to -->
target\_port      | string     | no       | ターゲットのポート (例 `70,80-90` or `90`)、 空の場合は `listen_port` と同じ <!-- Target port(s) (e.g. `70,80-90` or `90`), same as `listen_port` if empty -->
description       | string     | no       | ポートの説明 <!-- Description of port(s) -->

## <a name="network-types"></a> ネットワーク種別 <!-- Network types -->

以下のネットワーク種別がフォワードをサポートします。詳細は各ネットワーク種別の項を参照してください。
<!--
The following network types support forwards. See each network type section for more details.
-->

 - [bridge](#network-bridge)
 - [ovn](#network-ovn)


### <a name="network-bridge"></a> ブリッジ <!-- network: bridge -->

衝突しない任意のリッスンアドレスを指定可能です。
<!--
Any non-conflicting listen address is allowed.
-->

使用するリッスンアドレスは他のネットワークで使用中のサブネットとオーバーラップはできません。
<!--
The listen address used cannot overlap with a subnet that is in use with another network.
-->

### <a name="network-ovn"></a> ovn <!-- network: ovn -->

使用可能なリッスンアドレスはアップリンクのネットワークの `ipv{n}.routes` 設定と（設定されていれば）プロジェクトの `restricted.networks.subnets` 設定で定義されているアドレスです。
<!--
The allowed listen addresses are those that are defined in the uplink network's `ipv{n}.routes` settings, and the
project's `restricted.networks.subnets` setting (if set).
-->

使用するリッスンアドレスは他のネットワークで使用中のサブネットとオーバーラップはできません。
<!--
The listen address used cannot overlap with a subnet that is in use with another network.
-->
