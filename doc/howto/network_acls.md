---
discourse: 13223
---

(network-acls)=
# ネットワーク ACL を設定するには

```{note}
ネットワーク ACL は {ref}`OVN NIC タイプ <nic-ovn>`、{ref}`network-ovn` と {ref}`network-bridge` (いくつか制限あり、{ref}`network-acls-bridge-limitations` 参照) で利用できます。
```

```{youtube} https://www.youtube.com/watch?v=mu34G0cX6Io
```

ネットワーク {abbr}`ACL (Access Control Lists; アクセス制御リスト)` は同じネットワークに接続された異なるインスタンス間のネットワークアクセスや、他のネットワークとのアクセスを制御するトラフィクルールを定義します。

ネットワーク ACL は インスタンスの {abbr}`NIC (Network Interface Controller; ネットワークインタフェースコントローラ)` やネットワークに直接適用できます。
ネットワークに適用するときは、ネットワークに接続された全ての NIC に ACL が適用されます。

特定の ACL を(明示的にあるいはネットワークから暗黙的に)適用したインスタンス NIC は論理的なグループを形成し、他のルールから送信元あるいは送信先として参照できます。
より詳細な情報は {ref}`network-acls-groups` を参照してください。

## ACL を作成する

ACL を作成するには以下のコマンドを使用します。

```bash
lxc network acl create <ACL_name> [configuration_options...]
```

このコマンドはルール無しの ACL を作成します。
次のステップとして ACL に {ref}`ルールを追加します <network-acls-rules>`。

有効なネットワーク ACL の名前は以下のルールに従う必要があります。

- 名前は 1 文字から 63 文字の間である
- 名前は ASCII の文字、数字、ハイフンからのみなる
- 名前は数字やハイフンから始まらない
- 名前はハイフンで終わらない

### ACL のプロパティ

ACL のプロパティには次のものがあります。

プロパティ    | 型         | 必須 | 説明
:--           | :--        | :--  | :--
`name`        | string     | yes  | プロジェクト内でユニークなネットワーク ACL の名前
`description` | string     | no   | ネットワーク ACL の説明
`ingress`     | rule list  | no   | 内向きのトラフィックルールのリスト
`egress`      | rule list  | no   | 外向きのトラフィックルールのリスト
`config`      | string set | no   | キー・バリューペア形式での設定オプション (`user.*` カスタムキーのみサポート)


(network-acls-rules)=
## ルールの追加と削除

それぞれの ACL はルールの 2 つのリストを含みます。

- *イングレス (ingress)* ルールは NIC に向かう内向きのトラフィックに適用されます。
- *イーグレス (egress)* ルールは NIC から出ていく外向きのトラフィックに適用されます。

ACL にルールを追加するには、以下のコマンドを使います。 `<direction>` には `ingress` か `egress` のどちらかを指定します。

```bash
lxc network acl rule add <ACL_name> <direction> [properties...]
```

このコマンドは指定した方向 (direction) に対応するリストにルールを追加します。

({ref}`ACL 全体を編集する <network-acls-edit>` 場合を除き) ルールを編集することはできませんが、以下のコマンドでルールを削除はできます。

```bash
lxc network acl rule remove <ACL_name> <direction> [properties...]
```

ユニークにルールを特定するのに必要な全てのプロパティを指定するか、またはマッチした全てのルールを削除するためコマンドに `--force` を追加する必要があります。

### ルールの順番と優先度

ルールはリストとして提供されます。
しかしリスト内のルールの順番は重要ではなくフィルタリングには影響しません。

LXD は以下のように `action` プロパティに基づいてルールの順番を自動的に決めます。

- `drop`
- `reject`
- `allow`
- 上記の全てにマッチしなかったトラフィックに対する自動のデフォルトアクション (デフォルトでは `reject`、{ref}`network-acls-defaults` 参照)。

これは NIC に複数の ACL を適用する際、組み合わせたルールの順番を指定する必要がないことを意味します。
ACL 内のあるルールがマッチすれば、そのルールが採用され、他のルールは考慮されません。

### ルールのプロパティ

ACL ルールには次のプロパティがあります。

プロパティ         | 型     | 必須 | 説明
:--                | :--    | :--  | :--
`action`           | string | yes  | マッチしたトラフィックに適用するアクション(`allow`, `reject` または `drop`)
`state`            | string | yes  | ルールの状態(`enabled`, `disabled` または `logged`)、未設定の場合のデフォルト値は `enabled`
`description`      | string | no   | ルールの説明
`source`           | string | no   | CIDR か IP の範囲、送信元の ACL の名前、あるいは(ingress ルールに対しての) ソースサブジェクト名セレクターのカンマ区切りリスト、または any の場合は空を指定
`destination`      | string | no   | CIDR か IP の範囲、送信先の ACL の名前、あるいは(egress ルールに対しての) デスティネーションサブジェクト名セレクターのカンマ区切りリスト、または any の場合は空を指定
`protocol`         | string | no   | マッチ対象のプロトコル(`icmp4`, `icmp6`, `tcp`, `udp`)、または any の場合は空を指定
`source_port`      | string | no   | protocol が `udp` か `tcp` の場合はポートかポートの範囲(開始-終了で両端含む)のカンマ区切りリスト、または any の場合は空を指定
`destination_port` | string | no   | protocol が `udp` か `tcp` の場合はポートかポートの範囲(開始-終了で両端含む)のカンマ区切りリスト、または any の場合は空を指定
`icmp_type`        | string | no   | protocol が `icmp4` か `icmp6` の場合は ICMP の Type 番号、または any の場合は空を指定
`icmp_code`        | string | no   | protocol が `icmp4` か `icmp6` の場合は ICMP の Code 番号、または any の場合は空を指定

(network-acls-selectors)=
### ルール内でセレクタを使う

```{note}
この機能は {ref}`OVN NIC タイプ <nic-ovn>` と {ref}`network-ovn` でのみサポートされます。
```

(ingress ルールの) `source` フィールドと (egress ルールの) `destination` フィールドは CIDR や IP の範囲の代わりにセレクタの使用をサポートします。

この機能を使えば、 IP のリストを管理したり追加のサブネットを作ることなしに、 ACL グループかネットワークセレクタを使ってインスタンスのグループに対するルールを定義できます。

(network-acls-groups)=
#### ACL グループ

(明示的にあるいはネットワーク経由で暗黙的に) 特定の ACL を適用されたインスタンス NIC は論理的なポートグループを形成します。

そのような ACL グループは *サブジェクト名セレクタ* と呼ばれ、他の ACL グループ内で ACL 名を用いて参照できます。

例えば `foo` という名前の ACL がある場合、この ACL が適用されたインスタンス NIC のグループを `source=foo` で参照できます。

#### ネットワークセレクタ

*ネットワークサブジェクトセレクタ* を用いて、ネットワーク上の外向きと内向きのトラフィックにルールを定義できます。

`@internal` と `@external` という 2 つの特別なネットワークサブジェクトセレクタがあります。
これらはそれぞれネットワークのローカルと外向きのトラフィックを示します。
例:

```bash
source=@internal
```

ネットワークが [ネットワークピア](network_ovn_peers.md) をサポートする場合、ピア接続間のトラフィックを
`@<network_name>/<peer_name>` という形式のネットワークサブジェクトセレクタで参照できます。
例:

```bash
source=@ovn1/mypeer
```

ネットワークサブジェクトセレクターを使用する際は、 ACL 適用先のネットワークは指定されたピア接続を持っていなければなりません。
持っていない場合 ACL は適用できません。

### トラフィックのログ

一般的には ACL はインスタンスとネットワーク間のネットワークトラフィックを制御するためのものです。
しかし、特定のネットワークトラフィックをログ出力するためにルールを使うこともできます。
これはモニタリングや、ルールを実際に有効にする前にテストするのに役立ちます。

ログのためにルールを追加するには `state=logged` プロパティ付きでルールを作成してください。
ACL 内の全てのログのルールに対するログ出力は以下のコマンドで表示できます。

```bash
lxc network acl show-log <ACL_name>
```

(network-acls-edit)=
### ACL を編集する

ACL を編集するには以下のコマンドを使用します。

```bash
lxc network acl edit <ACL_name>
```

このコマンドは ACL を編集用に YAML 形式でオープンします。
ACL 設定とルールの両方を編集できます。

## ACL の適用

ACL の設定が終わったらネットワークかインスタンス NIC に適用する必要があります。

そのためにはネットワークか NIC の設定の `security.acls` リストに ACL を追加してください。
ネットワークの場合は、以下のコマンドを使います。

```bash
lxc network set <network_name> security.acls="<ACL_name>"
```

インスタンス NIC の場合は、以下のコマンドを使います。

```bash
lxc config device set <instance_name> <device_name> security.acls="<ACL_name>"
```

(network-acls-defaults)=
## デフォルトアクションの設定

1 つ以上の ACL が NIC に (明示的にあるいはネットワーク経由で暗黙的に) 適用されると、 NIC にデフォルトの reject ルールが追加されます。
このルールは適用された ACL 内のどのルールにもマッチしない全てのトラフィックを拒否 (reject) します。

この挙動はネットワークと NIC レベルの `security.acls.default.ingress.action` と `security.acls.default.egress.action` 設定で変更できます。
NIC レベルの設定はネットワークレベルの設定を上書きします。

例えば、ネットワークに接続された全てのインスタンスの内向きのトラフィックを許可 (`allow`) するには以下のコマンドを使用します。

```bash
lxc config device set <instance_name> <device_name> security.acls.default.ingress.action=allow
```

インスタンス NIC に同じデフォルトアクションを設定するには以下のコマンドを使用します。

```bash
lxc config device set <instance_name> <device_name> security.acls.default.ingress.action=allow
```

(network-acls-bridge-limitations)=
## ブリッジの制限

ブリッジネットワークにネットワーク ACL を使用する場合は以下の制限に気を付けてください。

- OVN ACL とは違い、ブリッジ ACL はブリッジと LXD ホストの間の境界のみに適用されます。これは外部へと外部からのトラフィックにネットワークポリシーを適用するために使うことしかできないことを意味します。ブリッジ間のファイアウォール、つまり同じブリッジに接続されたインスタンス間のトラフィックを制御するファイアウォールには使えません。
- {ref}`ACL グループとネットワークセレクタ <network-acls-selectors>` はサポートされません。
- `iptables` ファイアウォールドライバを使う際は、 IP レンジサブジェクト（例：`192.168.1.1-192.168.1.10`）は使用できません。
- ベースラインのネットワークサービスルールが（対応する INPUT/OUTPUT チェイン内の） ACL ルールの前に適用されます。これは一旦 ACL チェインに入ってしまうと INPUT/OUTPUT と FORWARD トラフィックを区別できないからです。このため ACL ルールはベースラインのサービスルールをブロックするのには使えません。
