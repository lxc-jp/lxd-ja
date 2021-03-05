# ネットワーク ACL の設定 <!-- Network ACL configuration -->

ネットワークアクセス制御リスト (ACL) はある種の Instance NIC デバイスに適用されるトラフィクルールを定義します。
これは同じネットワークに接続された異なるインスタンス間のネットワークアクセスと外部のネットワークとのアクセスを制御する機能を提供します。
<!--
Network Access Control Lists (ACLs) define traffic rules that can then be applied to certain types of Instance NIC devices.
This provides the ability to control network access between different instances connected to the same network and
control access to and from the external network.
-->

ネットワーク ACL は希望の NIC に直接適用することも出来ますし、希望のネットワークに ACL の適用することでネットワークに接続する全ての NIC に適用することも出来ます。
<!--
Network ACLs can either be applied directly to the desired NICs or can be applied to all NICs connected to a
network by assigning applying the ACL to the desired network.
-->

特別な ACL を(明示的にあるいはネットワークから暗黙的に)適用された Instance NIC は他のルールから送信元あるいは送信先として参照される論理的なグループを形成します。
これにより IP のリストを維持管理したり追加のサブネットを作成することなくインスタンスのグループに対してルールを定義できます。
<!--
The Instance NICs that have a particular ACL applied (either explicitly or implicitly from the network) make up a
logical group that can be referenced from other rules as a source or destination. This makes it possible to define
rules for groups of instances without needing to maintain IP lists or create additional subnets.
-->

ネットワーク ACL には暗黙のデフォルトルール(`default.action` が定義されない限り `reject` がデフォルト)があるため、トラフィックが ACL に定義されたルールのいずれにもマッチしない場合は drop されます。
<!--
Network ACLs come with an implicit default rule (that defaults to `reject` unless `default.action` is set), so if
traffic doesn't match one of the defined rules in an ACL then all other traffic is dropped.
-->

ルールは Instance NIC に対しての特定の向き(ingress か egress)について定義されます。
ingress のルールは NIC に向かうトラフィックに適用され、 egress のルールは NIC から出るトラフィックに適用されます。
<!--
Rules are defined on for a particular direction (ingress or egress) in relation to the Instance NIC.
Ingress rules apply to traffic going towards the NIC, and egress rules apply to traffic leave the NIC.
-->

ルールはリストとして定義され、リスト内でのルールの順番は重要ではなく、フィルタリングには影響しません。
<!--
Rules are provided as lists, however the order of the rules in the list is not important and does not affect filtering.
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
config           | string set | no       | (`user.*` カスタムキー以外の) 設定のキー・バリューペアのセット(下記参照) <!-- Config key/value pairs (in addition to `user.*` custom keys, see below) -->

config のプロパティー
<!--
Config properties:
-->

Property         | Type       | Required | Description
:--              | :--        | :--      | :--
default.action   | string     | no       | デフォルトルールに到達したトラフィックに適用するアクション(デフォルトは `reject`) <!-- What action to take for traffic hitting the default rule (default `reject`) -->
default.logged   | boolean    | no       | デフォルトルールに到達したトラフィックをログ出力するかどうか(デフォルトは `false`) <!-- Whether or not to log traffic hitting the default rule (default `false`) -->

ACL ルールには次のプロパティーがあります。
<!--
ACL rules have the following properties:
-->

Property          | Type       | Required | Description
:--               | :--        | :--      | :--
action            | string     | yes      | マッチしたトラフィックに適用するアクション(`allow`, `reject` または `drop`) <!-- Action to take for matching traffic (`allow`, `reject` or `drop`) -->
state             | string     | yes      | ルールの状態(`enabled`, `disabled` または `logged`) <!-- State of rule (`enabled`, `disabled` or `logged`) -->
description       | string     | no       | ルールの説明 <!-- Description of rule -->
source            | string     | no       | CIDR か IP の範囲、送信元の ACL の名前、あるいは(ingress ルールに対しての) #external/#internal のカンマ区切りリスト、または any の場合は空を指定 <!-- Comma separated list of CIDR or IP ranges, source ACL names or #external/#internal (for ingress rules), or empty for any -->
destination       | string     | no       | CIDR か IP の範囲、送信先の ACL の名前、あるいは(egress ルールに対しての) #external/#internal のカンマ区切りリスト、または any の場合は空を指定 <!-- Comma separated list of CIDR or IP ranges, destination ACL names or #external/#internal (for egress rules), or empty for any -->
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
 - 上記の全てにマッチしなかったトラフィックへのデフォルトルールのアクション(`default.action` が未指定の場合のデフォルトは `reject`) <!-- Automatic default rule action for any unmatched traffic (defaults to `reject` if `default.action` not specified). -->

 これは 1 つの NIC への複数のルールが結合されたルールの順序を指定することなしに適用できることを意味します。
 ACL 内のどれか一つのルールがマッチされたらすぐにアクションが実行され、他のルールは考慮されません。
<!--
 This means that multiple ACLs can be applied to a NIC without having to specify the combined rule ordering.
 As soon as one of the rules in the ACLs matches then that action is taken and no other rules are considered.
-->

## ポートグループセレクター <!-- Port group selectors -->

特定の ACL を割り当てられた Instance NIC は論理的なポートグループを形成し、他の ACL ルールから名前で参照することが出来ます。
<!--
The Instance NICs that are assigned a particular ACL make up a logical port group that can then be referenced by
name in other ACL rules.
-->

また `#internal` と `#external` という 2 つの特殊なセレクターがあり、これらはネットワークのそれぞれローカルと外部のトラフィックを表します。
<!--
There are also two special selectors called `#internal` and `#external` which represent network local and external
traffic respectively.
-->

ポートグループセレクターは ingress ルールの `source` フィールドと egress ルールの `destination` フィールドで使用可能です。
<!--
Port group selectors can be used in the `source` field for ingress rules and in the `destination` field for egress rules.
-->
