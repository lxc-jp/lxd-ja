(storage-cephfs)=
# CephFS - `cephfs`

% Include content from [storage_ceph.md](storage_ceph.md)
```{include} storage_ceph.md
    :start-after: <!-- Include start Ceph intro -->
    :end-before: <!-- Include end Ceph intro -->
```

{abbr}`CephFS (Ceph File System)` はロバストでフル機能の POSIX 互換の分散ファイルシステムを提供する Ceph のファイルシステムコンポーネントです。
内部的には ファイルを Ceph オブジェクトにマップし、ファイルのメタデータ (例えば、ファイルの所有権、ディレクトリパス、アクセス権限) を別のデータプールに保管します。

## 用語

% Include content from [storage_ceph.md](storage_ceph.md)
```{include} storage_ceph.md
    :start-after: <!-- Include start Ceph terminology -->
    :end-before: <!-- Include end Ceph terminology -->
```

*CephFS ファイルシステム* は 2 つの OSD ストレージプールから構成され、ひとつは実際のデータ、もうひとつはファイルメタデータに使用されます。

## LXD の `cephfs` ドライバ

```{note}
`cephfs` ドライバはコンテントタイプ `filesystem` のカスタムストレージボリュームにのみ使用できます。

他のストレージボリュームには {ref}`storage-ceph` ドライバを使用してください。
そのドライバはコンテントタイプ `filesystem` のカスタムストレージボリュームにも使用できますが、 Ceph RBD イメージを使って実装しています。
```

% Include content from [storage_ceph.md](storage_ceph.md)
```{include} storage_ceph.md
    :start-after: <!-- Include start Ceph driver cluster -->
    :end-before: <!-- Include end Ceph driver cluster -->
```

使用したい CephFS ファイルシステムは事前に作成する必要があり {ref}`source <storage-cephfs-pool-config>` オプションで指定する必要があります。

% Include content from [storage_ceph.md](storage_ceph.md)
```{include} storage_ceph.md
    :start-after: <!-- Include start Ceph driver remote -->
    :end-before: <!-- Include end Ceph driver remote -->
```

% Include content from [storage_ceph.md](storage_ceph.md)
```{include} storage_ceph.md
    :start-after: <!-- Include start Ceph driver control -->
    :end-before: <!-- Include end Ceph driver control -->
```

LXD の `cephfs` ドライバはサーバ側でスナップショットが有効な場合はスナップショットをサポートします。

## 設定オプション

`cephfs` ドライバを使うストレージプールとこれらのプール内のストレージボリュームには以下の設定オプションが利用できます。

(storage-cephfs-pool-config)=
## ストレージプール設定
キー                   | 型     | デフォルト値 | 説明
:--                    | :---   | :------      | :----------
cephfs.cluster\_name   | string | ceph         | CephFS ファイルシステムを含む Ceph クラスタの名前
cephfs.fscache         | bool   | false        | カーネルの fscache と cachefilesd を使用するか
cephfs.path            | string | /            | CephFS をマウントするベースのパス
cephfs.user.name       | string | admin        | 使用する Ceph のユーザー
source                 | string | -            | 使用する既存の CephFS ファイルシステムかファイルシステムパス
volatile.pool.pristine | string | true         | 作成時に CephFS ファイルシステムが空だったか

{{volume_configuration}}

## ストレージボリューム設定
キー               | 型     | 条件               | デフォルト値                             | 説明
:--                | :---   | :--------          | :------                                  | :----------
security.shifted   | bool   | custom volume      | volume.security.shifted と同じか false   | {{enable_ID_shifting}}
security.unmapped  | bool   | custom volume      | volume.security.unmapped と同じか false  | ボリュームの ID マッピングを無効にする
size               | string | appropriate driver | volume.size と同じ                       | ストレージボリュームのサイズ/クォータ
snapshots.expiry   | string | custom volume      | volume.snapshots.expiry と同じ           | {{snapshot_expiry_format}}
snapshots.pattern  | string | custom volume      | volume.snapshots.pattern と同じか snap%d | {{snapshot_pattern_format}}
snapshots.schedule | string | custom volume      | volume.snapshots.schedule と同じ         | {{snapshot_schedule_format}}
