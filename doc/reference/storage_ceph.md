(storage-ceph)=
# Ceph

- イメージとして RBD イメージを使い、インスタンスやスナップショットを作成するためにスナップショットやクローンを実行します
- RBD でコピーオンライトが動作するため、すべての子がなくなるまでは、親のファイルシステムは削除できません。
  その結果、LXD は削除されたにもかかわらずまだ参照されているオブジェクトに、自動的に `zombie_` というプレフィックスを付与します。
  そして、参照されなくなるまでそれを保持します。そして安全に削除します
- LXD は OSD ストレージプールを完全にコントロールできると想定していることに注意してください。
  LXD OSD ストレージプール内に、LXD が所有しないファイルシステムエンティティを維持し続けないことをおすすめします。
  LXD がそれらを削除する可能性があるからです
- 複数の LXD インスタンス間で、同じストレージプールを共有することはサポートしないことに注意してください。
  `lxd import` を使って既存インスタンスをバックアップする目的のときのみ、OSD ストレージプールを複数の LXD インスタンスで共有できます。
  このような場合には、`ceph.osd.force_reuse` プロパティを true に設定する必要があります。
  設定しない場合、LXD は他の LXD インスタンスが OSD ストレージプールを使っていることを検出した場合には、OSD ストレージプールの再利用を拒否します
- LXD が使う Ceph クラスタを設定するときは、OSD ストレージプールを保持するために使うストレージエンティティ用のファイルシステムとして `xfs` の使用をおすすめします。
  ストレージエンティティ用のファイルシステムとして ext4 を使用することは、Ceph の開発元では推奨していません。
  LXD と関係ない予期しない不規則な障害が発生するかもしれません
- "erasure" タイプの Ceph osd プールを使うためには事前に作成した osd pool とメタデータを保管するための "replicated" タイプの別の osd pool が必要です。
  これは RBD と CephFS が omap をサポートしないために必要となります。
  そのプールが "earasure coded" かを指定するにはリプリケートされたプールに
  `ceph.osd.data_pool_name=<erasure-coded-pool-name>` と
  `source=<replicated-pool-name>` を使用する必要があります。

## ストレージプール設定
キー                      | 型     | デフォルト値 | 説明
:--                       | :---   | :------      | :----------
ceph.cluster\_name        | string | ceph         | 新しいストレージプールを作成する Ceph クラスタの名前
ceph.osd.data\_pool\_name | string | -            | osd data pool の名前
ceph.osd.force\_reuse     | bool   | false        | 別の LXD インスタンスで既に使用されている osd ストレージプールの使用を強制するか
ceph.osd.pg\_num          | string | 32           | osd ストレージプール用の placement グループの数
ceph.osd.pool\_name       | string | プールの名前 | osd ストレージプールの名前
ceph.rbd.clone\_copy      | bool   | true         | フルのデータセットコピーではなく RBD のライトウェイトクローンを使うかどうか
ceph.rbd.du               | bool   | true         | 停止したインスタンスのディスク使用データを取得するのに rbd du を使用するかどうか
ceph.rbd.features         | string | layering     | ボリュームで有効にする RBD の機能のカンマ区切りリスト
ceph.user.name            | string | admin        | ストレージプールとボリュームの作成に使用する Ceph ユーザー
volatile.pool.pristine    | string | true         | プールが作成時に空かどうか

## ストレージボリューム設定
キー                 | 型     | 条件               | デフォルト値                       | 説明
:--                  | :---   | :--------          | :------                            | :----------
block.filesystem     | string | block based driver | volume.block.filesystem と同じ     | ストレージボリュームのファイルシステム
block.mount\_options | string | block based driver | volume.block.mount\_options と同じ | ブロックデバイスのマウントオプション
security.shifted     | bool   | custom volume      | false                              | id シフトオーバーレイを有効にする（複数の独立したインスタンスによるアタッチを許可する）
security.unmapped    | bool   | custom volume      | false                              | ボリュームへの id マッピングを無効にする
size                 | string | appropriate driver | volume.size と同じ                 | ストレージボリュームのサイズ
snapshots.expiry     | string | custom volume      | -                                  | スナップショットがいつ削除されるかを制御（`1M 2H 3d 4w 5m 6y` のような設定形式を想定）
snapshots.pattern    | string | custom volume      | snap%d                             | スナップショット名を表す Pongo2 テンプレート文字列（スケジュールされたスナップショットと名前指定なしのスナップショットに使用）
snapshots.schedule   | string | custom volume      | -                                  | Cron の書式 (`<minute> <hour> <dom> <month> <dow>`)、またはスケジュールアイリアスのカンマ区切りリスト `<@hourly> <@daily> <@midnight> <@weekly> <@monthly> <@annually> <@yearly>`


## Ceph ストレージプールを作成するには以下のコマンドが使用できます

- Ceph クラスタ "ceph" 内に "pool1" という OSD ストレージプールを作成する

```bash
lxc storage create pool1 ceph
```

- Ceph クラスタ "my-cluster" 内に "pool1" という OSD ストレージプールを作成する

```bash
lxc storage create pool1 ceph ceph.cluster_name=my-cluster
```

- ディスク上の名前を "my-osd" で "pool1" という名前の OSD ストレージプールを作成する

```bash
lxc storage create pool1 ceph ceph.osd.pool_name=my-osd
```

- 既存の OSD ストレージプール "my-already-existing-osd" を使用する

```bash
lxc storage create pool1 ceph source=my-already-existing-osd
```

- 既存の osd イレージャーコードされたプール "ecpool" と osd リプリケートされたプール "rpl-pool" を使用する

```bash
lxc storage create pool1 ceph source=rpl-pool ceph.osd.data_pool_name=ecpool
```
