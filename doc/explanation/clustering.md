---
discourse: 15728
---

(exp-clustering)=
# クラスタリングについて

```{youtube} https://www.youtube.com/watch?v=nrOR6yaO_MY
```

全体のワークロードを複数のサーバに分散するため、 LXD はクラスタリングモードで動かせます。
このシナリオでは、クラスタメンバーとそのインスタンスの設定を保持する同じ分散データベースを任意の台数の LXD サーバで共有します。
LXD クラスタは `lxc` クライアントまたは REST API を使って管理できます。

この機能は [`clustering`](../api-extensions.md#clustering) API 拡張の一部として導入され、 LXD 3.0 以降で利用可能です。

```{tip}
ベーシックなLXDクラスタを素早くセットアップしたい場合、[MicroCloud](https://discuss.linuxcontainers.org/t/introducing-microcloud/15871)をチェックしてみてください。
```

(clustering-members)=
## クラスタメンバー

LXD クラスタは 1 台のブートストラップサーバと少なくともさらに 2 台のクラスタメンバーから構成されます。
クラスタは状態を [分散データベース](../database.md) に保管します。これは Raft アルゴリズムを使用して複製される [Dqlite](https://dqlite.io/) データベースです。

2 台のメンバーだけでもクラスタを作成することは出来なくはないですが、少なくとも 3 台のクラスタメンバーを強く推奨します。
このセットアップでは、クラスタは少なくとも 1 台のメンバーの消失に耐えることができ、分散状態の過半数を確立できます。

クラスタを作成する際、 Dqlite データベースは 3 番目のメンバーがクラスタにジョインするまではブートストラップサーバ上でのみ稼働します。
そして 2 番目と 3 番目のサーバはデータベースの複製を受信します。

詳細は {ref}`cluster-form` を参照してください。

(clustering-member-roles)=
### メンバーロール

3 台のメンバーのクラスタでは、全てのメンバーがクラスタの状態を保管する分散データベースを複製します。
クラスタのメンバーがさらに増えると、一部のメンバーだけがデータベースを複製します。
残りのメンバーはデータベースへアクセスしますが、複製はしません。

任意の時点で、選出されたリーダーが 1 つ存在し、他のメンバーの健康状態をモニターします。

データベースを複製する各メンバーは *voter* か *stand-by* のロールを持ちます。
クラスタリーダーがオフラインになると voter の 1 つが新しいリーダーに選出されます。
voter のメンバーがオフラインになると stand-by メンバーが自動的に voter に昇格します。
データベース (そしてクラスタ) は voter の過半数がオンラインである限り利用可能です。

以下のロールが LXD クラスタメンバーに割り当て可能です。
自動のロールは LXD 自身によって割り当てられユーザによる変更は出来ません。

| ロール                  | 自動     | 説明 |
| :---                  | :--------     | :---------- |
| `database`            | yes           | 分散データベースの voter メンバー |
| `database-leader`     | yes           | 分散データベースの現在のリーダー |
| `database-standby`    | yes           | 分散データベースの stand-by (voter ではない) メンバー |
| `event-hub`           | no            | 内部 LXD イベントへの交換ポイント (hub) (最低 2 つは必要) |
| `ovn-chassis`         | no            | OVN ネットワークのアップリンクゲートウェイの候補 |

voter メンバーのデフォルトの数 ([`cluster.max_voters`](server)) は 3 です。
stand-by メンバーのデフォルトの数 ([`cluster.max_standby`](server)) は 2 です。
この設定では、クラスタを稼働したまま一度に最大で 1 つの voter メンバーの電源を切ることができます。

詳細は {ref}`cluster-manage` を参照してください。

(clustering-offline-members)=
#### オフラインメンバーと障害耐性

クラスタメンバーがダウンして設定されたオフラインの閾値を超えると、ステータスはオフラインと記録されます。
この場合、このメンバーに対する操作はできなくなり、全てのメンバーの状態変更を必要とする操作もできなくなります。

オフラインのメンバーがオンラインに戻るとすぐに操作が再びできるようになります。

オフラインになったメンバーがリーダーそのものだった場合、他のメンバーは新しいリーダーを選出します。

サーバを再びオンラインに復旧できないあるいはしたくない場合、[クラスタからメンバーを削除](cluster-manage-delete-members) できます。

応答しないメンバーをオフラインと判断する秒数は [`cluster.offline_threshold`](server) 設定で調整できます。
デフォルト値は 20 秒です。
最小値は 10 秒です。

詳細は {ref}`cluster-recover` を参照してください。

#### failure domain

オフラインになったメンバーにロールを割り当てる際に、どのクラスタメンバーを優先するかを指示するために failure domain を使用できます。
例えば、現在データベースロールを持つクラスタメンバーがシャットダウンした場合、 LXD はデータベースロールを同じ failure domain 内の別のクラスタメンバーがあればそれに割り当てようとします。

クラスタメンバーの failure domain を更新するには、 `lxc cluster edit <member>` コマンドを使って `failure_domain` プロパティを `default` から他の文字列に変更します。

(clustering-member-config)=
### メンバー設定

LXD クラスタメンバーは一般的に同一のシステムと想定されています。
それはクラスタにジョインする全ての LXD サーバはブートストラップサーバとストレージプールとネットワークについて同一の設定を持つ必要があるということです。

少し異なるディスクの順序やネットワークインタフェースの名前付けのようなことに対応するため、ストレージとネットワークに関連してメンバー固有のいくつかの設定が例外的に用意されています。

クラスタ内にそのような設定が存在する場合、追加するサーバにはそれらの設定に対する値を提供する必要があります。
たいていの場合、これはインタラクティブな `lxd init` コマンドで実行され、ユーザにストレージやネットワークに関連する設定の値の入力を求めます。

通常これらの設定には以下のものが含まれます。

- ストレージプールのソースデバイスとサイズ
- ZFS プール、 LVM thin pool、または LVM ボリュームグループの名前
- ブリッジネットワークの外部インタフェースと BGP の next-hop
- 管理された `physical` または `macvlan` ネットワークの親のネットワークデバイス名

詳細は {ref}`cluster-config-storage` と {ref}`cluster-config-networks` を参照してください。

事前に質問を調べたい (スクリプトでの自動化に有用) 場合、 `/1.0/cluster` API エンドポイントをクエリしてください。
これは `lxc query /1.0/cluster` あるいは他の API クライアントを使って実行できます。

## イメージ

デフォルトでは、 LXD はデータベースメンバーと同じ数のクラスタメンバーにイメージを複製します。
通常これはクラスタ内で最大 3 つのコピーを持つことを意味します。

障害耐性とイメージがローカルで利用できる確率を改善するためこの数を増やすことができます。
そのためには、 [`cluster.images_minimal_replica`](server) 設定を変更してください。
すべてのクラスタメンバーにイメージをコピーするには `-1` という特別な値を使用できます。

(cluster-groups)=
## クラスタグループ

LXD のクラスタではクラスタグループにメンバーを追加できます。
これらのクラスタグループは、全ての利用可能なメンバーのサブセットに属するクラスタメンバー上で、インスタンスを起動するのに使用できます。
例えば、GPU を持つ全てのメンバーからなるクラスタメンバーを作って、GPU が必要な全てのインスタンスをこのクラスタグループ上で起動できます。

デフォルトでは、全てのクラスタメンバーは `default` グループに属します。

詳細は {ref}`howto-cluster-groups` と {ref}`cluster-target-instance` を参照してください。

(clustering-instance-placement)=
## インスタンスの自動配置

クラスタのセットアップでは各インスタンスはクラスタメンバーの 1 つの上で稼働します。
インスタンスを起動する際、特定のクラスタメンバー、クラスタグループをターゲットにするか、あるいは LXD に自動的にどれかのクラスタメンバーに割り当てさせることもできます。

デフォルトでは、自動的な割り当てはインスタンス数が一番少ないクラスタメンバーを選択します。
複数のメンバーが同じインスタンス数の場合は、それらの 1 つがランダムで選ばれます。

しかし、この挙動を [`scheduler.instance`](cluster-member-config) 設定で制御することもできます。

- クラスタメンバーの `scheduler.instance` が `all` に設定されると、以下の条件でこのクラスタメンバーが選ばれます。

   - インスタンスが `--target` を指定せずに作成され、かつクラスタメンバーのインスタンス数が最小である。
   - インスタンスがこのクラスタメンバー上で稼働するようにターゲットされた。
   - インスタンスがこのクラスタメンバーが所属するクラスタグループのメンバー上で稼働するようにターゲットされ、かつクラスタメンバーがそのクラスタグループの他のメンバーと比べてインスタンス数が最小である。

- クラスタメンバーの `scheduler.instance` が `manual` に設定されると、以下の条件でこのクラスタメンバーが選ばれます。

   - インスタンスがこのクラスタメンバー上で稼働するようにターゲットされた。

- クラスタメンバーの `scheduler.instance` が `group` に設定されると、以下の条件でこのクラスタメンバーが選ばれます。

   - インスタンスがこのクラスタメンバー上で稼働するようにターゲットされた。
   - インスタンスがこのクラスタメンバーが所属するクラスタグループのメンバー上で稼働するようにターゲットされ、かつクラスタメンバーがそのクラスタグループの他のメンバーと比べてインスタンス数が最小である。

(clustering-instance-placement-scriptlet)=
### インスタンス配置スクリプトレット

LXDでは埋め込まれたスクリプト(スクリプトレット)を使って自動的なインスタンス配置を制御するカスタムロジックを使用できます。
この方法は、組み込みのインスタンス配置機能よりも柔軟性が高いです。

インスタンス配置スクリプトレットは[Starlark言語](https://github.com/bazelbuild/starlark) (Pythonのサブセット)で記述する必要があります。
スクリプトレットは、LXDがインスタンスをどこに配置するかを知る必要があるたびに呼び出されます。
スクリプトレットは、配置されるインスタンスに関する情報と、インスタンスをホストできる候補のクラスタメンバーに関する情報を受け取ります。
スクリプトレットからクラスタメンバー候補の状態と利用可能なハードウェアリソースについての情報を要求することもできます。

インスタンス配置スクリプトレットは`instance_placement`関数を以下のシグネチャで実装する必要があります。

   `instance_placement(request, candidate_members)`:

- `request`は、[`scriptlet.InstancePlacement`](https://pkg.go.dev/github.com/lxc/lxd/shared/api/scriptlet/#InstancePlacement) の展開された表現を含むオブジェクトです。このリクエストには、`project`および`reason`フィールドが含まれています。`reason`は、`new`、`evacuation`、または`relocation`のいずれかです。
- `candidate_members`は、[`api.ClusterMember`](https://pkg.go.dev/github.com/lxc/lxd/shared/api#ClusterMember) エントリを表すクラスタメンバーオブジェクトの`list`です。

例:

```python
def instance_placement(request, candidate_members):
    # 情報ログ出力の例。これはLXDのログに出力されます。
    log_info("instance placement started: ", request)

    # インスタンスのリクエストに基づいてロジックを適用する例。
    if request.name == "foo":
        # エラーログ出力の例。これはLXDのログに出力されます。
        log_error("Invalid name supplied: ", request.name)

        fail("Invalid name") # エラーで終了してインスタンス配置を拒否します。

    # 提供された第1候補のサーバにインスタンスを配置する。
    set_target(candidate_members[0].server_name)

    return # インスタンス配置を進めるために空を返す。
```

スクリプトレットはLXDに適用するためには`instances.placement.scriptlet`グローバル設定に設定する必要があります。

例えばスクリプトレットが`instance_placement.star`というファイルに保存されている場合、LXDには以下のように適用できます。

    cat instance_placement.star | lxc config set instances.placement.scriptlet=-

LXDに現在適用されているスクリプトレットを見るには`lxc config get instances.placement.scriptlet`コマンドを使用してください。

スクリプトレットでは(Starlarkで提供される関数に加えて)以下の関数が利用できます。

- `log_info(*messages)`: `info`レベルでLXDのログにログエントリを追加します。`messages`は1つ以上のメッセージの引数です。
- `log_warn(*messages)`: `warn`レベルでLXDのログにログエントリを追加します。`messages`は1つ以上のメッセージの引数です。
- `log_error(*messages)`: `error`レベルでLXDのログにログエントリを追加します。`messages`は1つ以上のメッセージの引数です。
- `set_cluster_member_target(member_name)`: インスタンスが作成されるべきクラスタメンバーを設定します。`member_name`はインスタンスが作成されるべきクラスタメンバーの名前です。この関数が呼ばれなければ、LXDは組み込みのインスタンス配置ロジックを使用します。
- `get_cluster_member_state(member_name)`: クラスタメンバーの状態を取得します。[`api.ClusterMemberState`](https://pkg.go.dev/github.com/lxc/lxd/shared/api#ClusterMemberState)の形式でクラスタメンバーの状態を含むオブジェクトを返します。`member_name`は状態を取得する対象のクラスタメンバーの名前です。
- `get_cluster_member_resources(member_name)`: クラスタメンバーのリソースについての情報を取得します。[`api.Resources`](https://pkg.go.dev/github.com/lxc/lxd/shared/api#Resources)の形式でリソースについての情報を含むオブジェクトを返します。`member_name`はリソース情報を取得する対象のクラスタメンバーの名前です。
- `get_instance_resources()`: インスタンスが必要とするリソースについての情報を取得します。 [`scriptlet.InstanceResources`](https://pkg.go.dev/github.com/lxc/lxd/shared/api/scriptlet/#InstanceResources)の形式でリソース情報を含むオブジェクトを返します。

```{note}
オブジェクト内のフィールド名は対応するGoの型のJSONフィールド名と同じです。
```
