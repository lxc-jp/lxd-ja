# ネットワーク ACL の設定 <!-- Network ACL configuration -->

ネットワークアクセス制御リスト (ACL) はある種の Instance NIC デバイスに適用されるトラフィクルールを定義します。
これは同じネットワークに接続された異なるインスタンス間のネットワークアクセスと他のネットワークとのアクセスを制御する機能を提供します。
<!--
Network Access Control Lists (ACLs) define traffic rules that can then be applied to certain types of Instance NIC devices.
This provides the ability to control network access between different instances connected to the same network and
control access to and from other networks.
-->

ネットワーク ACL は希望の NIC に直接適用することも出来ますし、希望のネットワークに ACL を割り当てることでネットワークに接続する全ての NIC に適用することも出来ます。
<!--
Network ACLs can either be applied directly to the desired NICs or can be applied to all NICs connected to a
network by assigning the ACL to the desired network.
-->

特別な ACL を(明示的にあるいはネットワークから暗黙的に)適用した Instance NIC は他のルールから送信元あるいは送信先として参照される論理的なグループを形成します。
これにより IP のリストを維持管理したり追加のサブネットを作成することなくインスタンスのグループに対してルールを定義できます。
<!--
The Instance NICs that have a particular ACL applied (either explicitly or implicitly from the network) make up a
logical group that can be referenced from other rules as a source or destination. This makes it possible to define
rules for groups of instances without needing to maintain IP lists or create additional subnets.
-->

1 つ以上の ACL が NIC に (明示的またはネットワークから暗黙的に) ひとたび適用されると NIC にデフォルトの拒否ルールが追加され
適用された ACL のどのルールにもマッチしないトラフィックは拒否されます。
<!--
Once one or more ACLs are applied to a NIC (either explicitly or implicitly from the network) then a default reject
rule is added to the NIC, so if traffic doesn't match one of the rules in the applied ACLs then it is rejected.
-->

この挙動はネットワークと NIC レベルの `security.acls.default.ingress.action` と `security.acls.default.egress.action` 設定で変更できます。
NIC レベルの設定はネットワークレベルの設定を上書きします。
<!--
This behaviour can be modified by using the network and NIC level `security.acls.default.ingress.action` and
`security.acls.default.egress.action` settings. The NIC level settings will override the network level settings.
-->

ルールはインスタンス NIC に対しての特定の向き (ingress か egress) に対して定義します。
ingress のルールは NIC に向かうトラフィックに適用し、 egress のルールは NIC から出るトラフィックに適用します。
<!--
Rules are defined for a particular direction (ingress or egress) in relation to the Instance NIC.
Ingress rules apply to traffic going towards the NIC, and egress rules apply to traffic leaving the NIC.
-->

ルールはリスト形式で指定しますが、リスト内のルールの順番は重要ではなくフィルタリングには影響しません。
[ルールの順番と優先度](#rule-ordering-and-priorities) を参照してください。
<!--
Rules are provided as lists, however the order of the rules in the list is not important and does not affect
filtering. See [Rule ordering and priorities](#rule-ordering-and-priorities).
-->

有効なネットワーク ACL の名前は以下のルールに従う必要があります。
<!--
Valid Network ACL names must:
-->

- 1 文字から 63 文字の間である <!-- Be between 1 and 63 characters long -->
- ASCII の文字、数字、ハイフンからのみなる <!-- Be made up exclusively of letters, numbers and dashes from the ASCII table -->
- 数字やハイフンから始まらない <!-- Not start with a digit or a dash -->
- ハイフンで終わらない <!-- Not end with a dash -->

## プロパティー <!-- Properties -->
ACL のプロパティーには次のものがあります。
<!--
The following are ACL properties:
-->


Property         | Type       | Required | Description
:--              | :--        | :--      | :--
name             | string     | yes      | プロジェクト内でユニークなネットワーク ACL の名前 <!-- Unique name of Network ACL in Project -->
description      | string     | no       | ネットワーク ACL の説明 <!-- Description of Network ACL -->
ingress          | rule list  | no       | ingress のトラフィックルールのリスト <!-- Ingress traffic rules -->
egress           | rule list  | no       | egress のトラフィックルールのリスト <!-- Egress traffic rules -->
config           | string set | no       | 設定のキー・バリューペア (`user.*` カスタムキーのみサポート) <!-- Config key/value pairs (Only `user.*` custom keys supported) -->

ACL ルールには次のプロパティーがあります。
<!--
ACL rules have the following properties:
-->

Property          | Type       | Required | Description
:--               | :--        | :--      | :--
action            | string     | yes      | マッチしたトラフィックに適用するアクション(`allow`, `reject` または `drop`) <!-- Action to take for matching traffic (`allow`, `reject` or `drop`) -->
state             | string     | yes      | ルールの状態(`enabled`, `disabled` または `logged`) <!-- State of rule (`enabled`, `disabled` or `logged`) -->
description       | string     | no       | ルールの説明 <!-- Description of rule -->
source            | string     | no       | CIDR か IP の範囲、送信元の ACL の名前、あるいは(ingress ルールに対しての) ソースサブジェクト名セレクターのカンマ区切りリスト、または any の場合は空を指定 <!-- Comma separated list of CIDR or IP ranges, source subject name selectors (for ingress rules), or empty for any -->
destination       | string     | no       | CIDR か IP の範囲、送信先の ACL の名前、あるいは(egress ルールに対しての) デスティネーションサブジェクト名セレクターのカンマ区切りリスト、または any の場合は空を指定 <!-- Comma separated list of CIDR or IP ranges, destination subject name selectors (for egress rules), or empty for any -->
protocol          | string     | no       | マッチ対象のプロトコル(`icmp4`, `icmp6`, `tcp`, `udp`)、または any の場合は空を指定 <!-- Protocol to match (`icmp4`, `icmp6`, `tcp`, `udp`) or empty for any -->
source\_port      | string     | no       | protocol が `udp` か `tcp` の場合はポートかポートの範囲(開始-終了で両端含む)のカンマ区切りリスト、または any の場合は空を指定 <!-- If Protocol is `udp` or `tcp`, then comma separated list of ports or port ranges (start-end inclusive), or empty for any -->
destination\_port | string     | no       | protocol が `udp` か `tcp` の場合はポートかポートの範囲(開始-終了で両端含む)のカンマ区切りリスト、または any の場合は空を指定 <!-- If Protocol is `udp` or `tcp`, then comma separated list of ports or port ranges (start-end inclusive), or empty for any -->
icmp\_type        | string     | no       | protocol が `icmp4` か `icmp6` の場合は ICMP の Type 番号、または any の場合は空を指定 <!-- If Protocol is `icmp4` or `icmp6`, then ICMP Type number, or empty for any -->
icmp\_code        | string     | no       | protocol が `icmp4` か `icmp6` の場合は ICMP の Code 番号、または any の場合は空を指定 <!-- If Protocol is `icmp4` or `icmp6`, then ICMP Code number, or empty for any -->

## ルールの順序と優先度 <!-- Rule ordering and priorities -->

ルールは明示的に順序を指定できません。しかし、 LXD はルールを `action` プロパティーに基づいて次のように順序付けます。
<!--
Rules cannot be explicitly ordered. However LXD will order the rules based on the `action` property as follows:
-->

 - `drop`
 - `reject`
 - `allow`
 - 上記の全てにマッチしなかったトラフィックへの自動のデフォルトのアクション(デフォルトは `reject`) <!-- Automatic default action for any unmatched traffic (defaults to `reject`). -->

これは 1 つの NIC への複数のルールが結合されたルールの順序を指定することなしに適用できることを意味します。
ACL 内のどれか一つのルールがマッチされたらすぐにアクションが実行され、他のルールは考慮されません。
<!--
This means that multiple ACLs can be applied to a NIC without having to specify the combined rule ordering.
As soon as one of the rules in the ACLs matches then that action is taken and no other rules are considered.
-->

デフォルトの拒否アクションはネットワークと NIC レベルの `security.acls.default.ingress.action` と `security.acls.default.egress.action` 設定で変更できます。
NIC レベルの設定はネットワークレベルの設定を上書きします。
<!--
The default reject action can be modified by using the network and NIC level `security.acls.default.ingress.action`
and `security.acls.default.egress.action` settings. The NIC level settings will override the network level settings.
-->

# サブジェクト名セレクター <!-- Subject name selectors -->

サブジェクト名セレクターは ingress ルールの `source` フィールドと egress ルールの `destination` フィールドで使用可能です。
<!--
Subject name selectors can be used in the `source` field for ingress rules and in the `destination` field for
egress rules.
-->

（直接あるいは NIC が接続するネットワークに割り当てられた ACL 経由で） 特定の ACL を割り当てられた Instance NIC
は論理的なポートグループを形成し、他の ACL ルールから `<ACL_name>` 形式で ACL サブジェクト名として参照することが出来ます。
<!--
Instance NICs that are assigned a particular ACL (either directly or via the ACLs assigned to the network it is
connected to) make up a logical port group named after the ACL that can then be referenced as an ACL subject name
in other ACL rules using the format `<ACL_name>`.
-->

例 `source=foo`
<!--
E.g. `source=foo`
-->

ネットワークが [ネットワークピア](network-peers.md) をサポートする場合、ピア接続間のトラフィックを
`@<network_name>/<peer_name>` という形式のネットワークサブジェクトセレクターで参照できます。
<!--
If the network supports [network peers](network-peers.md) then you can also reference traffic to/from the peer
connection by way of a network subject selector in the format `@<network_name>/<peer_name>`.
-->

例 `source=@ovn1/mypeer`
<!--
E.g. `source=@ovn1/mypeer`
-->

ネットワークサブジェクトセレクターを使用する際は、 ACL 適用先のネットワークは指定されたピア接続を持っていなければなりません。
持っていない場合 ACL は適用されません。
<!--
When using a network subject selector, the network having the ACL applied to it must have the specified peer
connection or the ACL will refuse to be applied to it.
-->

`@internal` と `@external` という特別なネットワークサブジェクトセレクターもあります。
これらはそれぞれネットワークローカルのトラフィックと外部のトラフィックを表します。
<!--
There are also two special network subject selectors called `@internal` and `@external` which represent network
local and external traffic respectively.
-->

例 `source=@internal`
<!--
E.g. `source=@internal`
-->

## ブリッジの制限 <!-- Bridge limitations -->

<!--
Unlike OVN ACLs, `bridge` ACLs are applied *only* on the boundary between the bridge and the LXD host.
This means they can only be used to apply network policy for traffic going to/from external networks, and cannot be
used for intra-bridge firewalling (i.e for firewalling traffic between instances connected to the same bridge).
-->
OVN ACL とは違い、 `bridge` ACL はブリッジと LXD ホストの間の境界*のみ*に適用されます。これは外部へと外部からのトラフィックにネットワークポリシーを適用するために使うことしかできず、ブリッジ間のファイアウォール（例：同じブリッジに繋がれたインスタンス間のトラフィックに対するファイアウォール）には使えません。

<!--
Additionally `bridge` ACLs do not support using subject name selectors.
-->
さらに `bridge` ACL はサブジェクト名セレクターの使用をサポートしていません。

<!--
When using the `iptables` firewall driver, you cannot use IP range subjects (e.g. `192.168.1.1-192.168.1.10`).
-->
`iptables` ファイアウォールドライバーを使う際は、 IP レンジサブジェクト（例：`192.168.1.1-192.168.1.10`）は使用できません。

<!--
Baseline network service rules are added before ACL rules (in their respective INPUT/OUTPUT chains), because we
cannot differentiate between INPUT/OUTPUT and FORWARD traffic once we have jumped into the ACL chain. Because of
this ACL rules cannot be used to block baseline service rules.
-->
ベースラインのネットワークサービスルールが（対応する INPUT/OUTPUT チェイン内の） ACL ルールの前に適用されます。これは一旦 ACL チェインに入ってしまうと INPUT/OUTPUT と FORWARD トラフィックを区別できないからです。このため ACL ルールはベースラインのサービスルールをブロックするのには使えません。
