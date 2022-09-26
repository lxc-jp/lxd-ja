(storage-ceph)=
# Ceph RBD - `ceph`

```{youtube} https://youtube.com/watch?v=kVLGbvRU98A
```

<!-- Include start Ceph intro -->
[Ceph](https://ceph.io/) はオープンソースのストレージプラットフォームで、データを {abbr}`RADOS (Reliable Autonomic Distributed Object Store)` に基づいたストレージクラスタ内に保管します。
非常にスケーラブルで、単一障害点がない分散システムであり非常に信頼性が高いです。

Ceph はブロックストレージ用とファイルシステム用に異なるコンポーネントを提供します。
<!-- Include end Ceph intro -->

Ceph {abbr}`RBD (RADOS Block Device)` はデータとワークロードを Ceph クラスタに分散する Ceph のブロックストレージコンポーネントです。
これは薄いプロビジョニングを使用し、リソースをオーバーコミットできることを意味します。

## 用語

<!-- Include start Ceph terminology -->
Ceph は保管するデータに *オブジェクト* という用語を使用します。
データを保存と管理する責任を持つデーモンは *Ceph {abbr}`OSD (Object Storage Daemon)`* です。
Ceph のストレージは *プール* に分割されます。これはオブジェクトを保管する論理的なパーティションです。
これらは *データプール*, *ストレージプール*, *OSD プール* とも呼ばれます。
<!-- Include end Ceph terminology -->

Ceph ブロックデバイスは *RBD イメージ* とも呼ばれ、これらの RBD イメージの *スナップショット* と *クローン* を作成できます。

## LXD の `ceph` ドライバ

```{note}
Ceph RBD ドライバを使用するには `ceph` と指定する必要があります。
これは少し誤解を招く恐れがあります。 Ceph の全ての機能ではなく Ceph RBD (ブロックストレージ) の機能しか使わないからです。
コンテントタイプ `filesystem` (イメージ、コンテナとカスタムファイルシステムボリューム) のストレージボリュームには `ceph` ドライバは Ceph RDB イメージをその上にファイルシステムがある状態で使用します ([`block.filesystem`](storage-ceph-vol-config) 参照)。

別の方法として、コンテントタイプ `filesystem` でストレージボリュームを作成するのに {ref}`CephFS <storage-cephfs>` を使用することもできます。
```

<!-- Include start Ceph driver cluster -->
他のストレージドライバとは異なり、このドライバはストレージシステムをセットアップはせず、既に Ceph クラスタをインストール済みであると想定します。
<!-- Include end Ceph driver cluster -->

<!-- Include start Ceph driver remote -->
このドライバはリモートのストレージを提供するという意味でも他のドライバとは異なる振る舞いをします。
結果として、内部ネットワークに依存し、ストレージへのアクセスはローカルのストレージより少し遅くなるかもしれません。
一方で、リモートのストレージを使うことはクラスタ構成では大きな利点があります。これはストレージプールを同期する必要なしに、全てのクラスタメンバが同じ内容を持つ同じストレージプールにアクセスできるからです。
<!-- Include end Ceph driver remote -->

LXD 内の `ceph` ドライバはイメージ、スナップショットに RBD イメージを使用し、インスタンスとスナップショットを作成するのにクローンを使用します。

<!-- Include start Ceph driver control -->
LXD は OSD ストレージプールに対して完全制御できることを想定します。
このため、 LXD OSD ストレージプール内に LXD が所有しないファイルシステムエンティティは LXD が消してしまうかもしれないので決して置くべきではありません。
<!-- Include end Ceph driver control -->

Ceph RBD 内で copy-on-write が動作する方法のため、親の RBD イメージは全ての子がいなくなるまで削除できません。
結果として LXD は削除されたがまだ参照されているオブジェクトを自動的にリネームします。
そのようなオブジェクトは全ての参照がいなくなりオブジェクトが安全に削除できるようになるまで `zombie_` 接頭辞をつけて維持されます。

### 制限

`ceph` ドライバには以下の制限があります。

インスタンス間でのカスタムボリュームの共有
: {ref}`コンテントタイプ <storage-content-types>` `filesystem` のカスタムストレージボリュームは異なるクラスタメンバの複数のインスタンス間で通常は共有できます。
  しかし、 Ceph RBD ドライバは RBD イメージ上にファイルシステムを置くことでコンテントタイプ `filesystem` のボリュームを「シミュレート」しているため、カスタムストレージボリュームは一度に1つのインスタンスにしか割り当てできません。
  コンテントタイプ `filesystem` のカスタムボリュームを共有する必要がある場合は代わりに {ref}`storage-cephfs` ドライバを使用してください。

複数インストールされた LXD 間で OSD ストレージプールの共有
: 複数インストールされた LXD 間で同じ OSD ストレージプールを共有することはサポートされていません。

タイプ "erasure" の OSD プールの使用
: タイプ "erasure" の OSD プールを使用するには事前に OSD プールを作成する必要があります。
  さらにタイプ "replicated" の別の OSD プールを作成する必要もあります。これはメタデータを保管するのに使用されます。
  これは Ceph RBD が `omap` をサポートしないために必要となります。
  どのプールが "erasure coded" であるかを指定するために [`ceph.osd.data_pool_name`](storage-ceph-pool-config) 設定オプションをイレージャーコーディングされたプールの名前に設定し [`source`](storage-ceph-pool-config) 設定オプションをリプリケートされたプールの名前に設定します。

## 設定オプション

`ceph` ドライバを使うストレージプールとこれらのプール内のストレージボリュームには以下の設定オプションが利用できます。

(storage-ceph-pool-config)=
### ストレージプール設定

キー                      | 型     | デフォルト値 | 説明
:--                       | :---   | :------      | :----------
`ceph.cluster_name`       | string | `ceph`       | 新しいストレージプールを作成する Ceph クラスタの名前
`ceph.osd.data_pool_name` | string | -            | OSD data pool の名前
`ceph.osd.pg_num`         | string | `32`         | OSD ストレージプール用の placement グループの数
`ceph.osd.pool_name`      | string | プールの名前 | OSD ストレージプールの名前
`ceph.rbd.clone_copy`     | bool   | `true`       | フルのデータセットコピーではなく RBD のライトウェイトクローンを使うかどうか
`ceph.rbd.du`             | bool   | `true`       | 停止したインスタンスのディスク使用データを取得するのに RBD `du` を使用するかどうか
`ceph.rbd.features`       | string | `layering`   | ボリュームで有効にする RBD の機能のカンマ区切りリスト
`ceph.user.name`          | string | `admin`      | ストレージプールとボリュームの作成に使用する Ceph ユーザー
`source`                  | string | -            | 使用する既存の OSD ストレージプール
`volatile.pool.pristine`  | string | `true`       | プールが作成時に空かどうか

{{volume_configuration}}

(storage-ceph-vol-config)=
### ストレージボリューム設定

キー                  | 型     | 条件                   | デフォルト値                                 | 説明
:--                   | :---   | :--------              | :------                                      | :----------
`block.filesystem`    | string | ブロックベースドライバ | `volume.block.filesystem` と同じ             | {{block_filesystem}}
`block.mount_options` | string | ブロックベースドライバ | `volume.block.mount_options` と同じ          | ブロックデバイスのマウントオプション
`security.shifted`    | bool   | カスタムボリューム     | `volume.security.shifted` と同じか `false`   | {{enable_ID_shifting}}
`security.unmapped`   | bool   | カスタムボリューム     | `volume.security.unmapped` と同じか `false`  | ボリュームの ID マッピングを無効にする
`size`                | string | 適切なドライバ         | `volume.size` と同じ                         | ストレージボリュームのサイズ/クォータ
`snapshots.expiry`    | string | カスタムボリューム     | `volume.snapshots.expiry` と同じ             | {{snapshot_expiry_format}}
`snapshots.pattern`   | string | カスタムボリューム     | `volume.snapshots.pattern` と同じか `snap%d` | {{snapshot_pattern_format}}
`snapshots.schedule`  | string | カスタムボリューム     | `volume.snapshots.schedule` と同じ           | {{snapshot_schedule_format}}
