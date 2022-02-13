# ネットワーク ACL の設定

ネットワークアクセス制御リスト (ACL) はある種の Instance NIC デバイスに適用されるトラフィクルールを定義します。
これは同じネットワークに接続された異なるインスタンス間のネットワークアクセスと他のネットワークとのアクセスを制御する機能を提供します。

ネットワーク ACL は希望の NIC に直接適用することも出来ますし、希望のネットワークに ACL を割り当てることでネットワークに接続する全ての NIC に適用することも出来ます。

特別な ACL を(明示的にあるいはネットワークから暗黙的に)適用した Instance NIC は他のルールから送信元あるいは送信先として参照される論理的なグループを形成します。
これにより IP のリストを維持管理したり追加のサブネットを作成することなくインスタンスのグループに対してルールを定義できます。

1 つ以上の ACL が NIC に (明示的またはネットワークから暗黙的に) ひとたび適用されると NIC にデフォルトの拒否ルールが追加され
適用された ACL のどのルールにもマッチしないトラフィックは拒否されます。

この挙動はネットワークと NIC レベルの `security.acls.default.ingress.action` と `security.acls.default.egress.action` 設定で変更できます。
NIC レベルの設定はネットワークレベルの設定を上書きします。

ルールはインスタンス NIC に対しての特定の向き (ingress か egress) に対して定義します。
ingress のルールは NIC に向かうトラフィックに適用し、 egress のルールは NIC から出るトラフィックに適用します。

ルールはリスト形式で指定しますが、リスト内のルールの順番は重要ではなくフィルタリングには影響しません。
[ルールの順番と優先度](#rule-ordering-and-priorities) を参照してください。

有効なネットワーク ACL の名前は以下のルールに従う必要があります。

- 1 文字から 63 文字の間である
- ASCII の文字、数字、ハイフンからのみなる
- 数字やハイフンから始まらない
- ハイフンで終わらない

## プロパティ
ACL のプロパティには次のものがあります。


Property         | Type       | Required | Description
:--              | :--        | :--      | :--
name             | string     | yes      | プロジェクト内でユニークなネットワーク ACL の名前
description      | string     | no       | ネットワーク ACL の説明
ingress          | rule list  | no       | ingress のトラフィックルールのリスト
egress           | rule list  | no       | egress のトラフィックルールのリスト
config           | string set | no       | 設定のキー・バリューペア (`user.*` カスタムキーのみサポート)

ACL ルールには次のプロパティがあります。

Property          | Type       | Required | Description
:--               | :--        | :--      | :--
action            | string     | yes      | マッチしたトラフィックに適用するアクション(`allow`, `reject` または `drop`)
state             | string     | yes      | ルールの状態(`enabled`, `disabled` または `logged`)
description       | string     | no       | ルールの説明
source            | string     | no       | CIDR か IP の範囲、送信元の ACL の名前、あるいは(ingress ルールに対しての) ソースサブジェクト名セレクターのカンマ区切りリスト、または any の場合は空を指定
destination       | string     | no       | CIDR か IP の範囲、送信先の ACL の名前、あるいは(egress ルールに対しての) デスティネーションサブジェクト名セレクターのカンマ区切りリスト、または any の場合は空を指定
protocol          | string     | no       | マッチ対象のプロトコル(`icmp4`, `icmp6`, `tcp`, `udp`)、または any の場合は空を指定
source\_port      | string     | no       | protocol が `udp` か `tcp` の場合はポートかポートの範囲(開始-終了で両端含む)のカンマ区切りリスト、または any の場合は空を指定
destination\_port | string     | no       | protocol が `udp` か `tcp` の場合はポートかポートの範囲(開始-終了で両端含む)のカンマ区切りリスト、または any の場合は空を指定
icmp\_type        | string     | no       | protocol が `icmp4` か `icmp6` の場合は ICMP の Type 番号、または any の場合は空を指定
icmp\_code        | string     | no       | protocol が `icmp4` か `icmp6` の場合は ICMP の Code 番号、または any の場合は空を指定

## ルールの順序と優先度

ルールは明示的に順序を指定できません。しかし、 LXD はルールを `action` プロパティに基づいて次のように順序付けます。

 - `drop`
 - `reject`
 - `allow`
 - 上記の全てにマッチしなかったトラフィックへの自動のデフォルトのアクション(デフォルトは `reject`)

これは 1 つの NIC への複数のルールが結合されたルールの順序を指定することなしに適用できることを意味します。
ACL 内のどれか一つのルールがマッチされたらすぐにアクションが実行され、他のルールは考慮されません。

デフォルトの拒否アクションはネットワークと NIC レベルの `security.acls.default.ingress.action` と `security.acls.default.egress.action` 設定で変更できます。
NIC レベルの設定はネットワークレベルの設定を上書きします。

## サブジェクト名セレクター

サブジェクト名セレクターは ingress ルールの `source` フィールドと egress ルールの `destination` フィールドで使用可能です。

（直接あるいは NIC が接続するネットワークに割り当てられた ACL 経由で） 特定の ACL を割り当てられた Instance NIC
は論理的なポートグループを形成し、他の ACL ルールから `<ACL_name>` 形式で ACL サブジェクト名として参照することが出来ます。

例 `source=foo`

ネットワークが [ネットワークピア](network-peers.md) をサポートする場合、ピア接続間のトラフィックを
`@<network_name>/<peer_name>` という形式のネットワークサブジェクトセレクターで参照できます。

例 `source=@ovn1/mypeer`

ネットワークサブジェクトセレクターを使用する際は、 ACL 適用先のネットワークは指定されたピア接続を持っていなければなりません。
持っていない場合 ACL は適用されません。

`@internal` と `@external` という特別なネットワークサブジェクトセレクターもあります。
これらはそれぞれネットワークローカルのトラフィックと外部のトラフィックを表します。

例 `source=@internal`

## ブリッジの制限

OVN ACL とは違い、 `bridge` ACL はブリッジと LXD ホストの間の境界*のみ*に適用されます。これは外部へと外部からのトラフィックにネットワークポリシーを適用するために使うことしかできず、ブリッジ間のファイアウォール（例：同じブリッジに繋がれたインスタンス間のトラフィックに対するファイアウォール）には使えません。

さらに `bridge` ACL はサブジェクト名セレクターの使用をサポートしていません。

`iptables` ファイアウォールドライバーを使う際は、 IP レンジサブジェクト（例：`192.168.1.1-192.168.1.10`）は使用できません。

ベースラインのネットワークサービスルールが（対応する INPUT/OUTPUT チェイン内の） ACL ルールの前に適用されます。これは一旦 ACL チェインに入ってしまうと INPUT/OUTPUT と FORWARD トラフィックを区別できないからです。このため ACL ルールはベースラインのサービスルールをブロックするのには使えません。
