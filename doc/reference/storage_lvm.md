(storage-lvm)=
# LVM - `lvm`

```{youtube} https://www.youtube.com/watch?v=AqLl2eMZE6U
```

{abbr}`LVM (Logical Volume Manager)` はファイルシステムというよりストレージマネージメントフレームワークです。
これは物理ストレージデバイスを管理するのに使用され、複数のストレージボリュームを作成し、配下の物理ストレージデバイスを使用し仮想化できるようにします。

この過程で物理ストレージをオーバーコミットすることが可能で、全ての利用可能なストレージが同時に使用されるわけではないシナリオに対して柔軟性を提供できることに注意してください。

LVM を使用するにはマシン上に `lvm2` がインストールされていることを確認してください。

## 用語

LVM は複数の物理ストレージデバイスを組み合わせて *ボリュームグループ* にすることができます。
その後このボリュームグループから異なるタイプの *論理ボリューム* を割り当てることができます。

サポートされるボリュームタイプの 1 つに *thin pool* があります。これは許可された最大サイズの合計は利用可能な物理ストレージより大きいような薄くプロビジョンされたボリュームを作成することでリソースをオーバーコミットすることを可能にします。
別のタイプは *ボリュームスナップショット* でこれは論理ボリュームの特定の状態をキャプチャーします。

## LXD の `lvm` ドライバ

LXD の `lvm` ドライバはイメージに論理ボリュームを、インスタンスとスナップショットにボリュームスナップショットを使用します。

LXD はボリュームグループを完全制御できると想定しています。
このため、 LXD が所有しないファイルシステムエンティティは LXD が消してしまうかもしれないので、LVM ボリュームグループ内に置くべきではありません。
しかし、既存のボリュームグループを再利用する必要がある場合 (例えば、あなたの環境ではボリュームグループが 1 つしかない場合)、[`lvm.vg.force_reuse`](storage-lvm-pool-config) を設定することでこれは可能です。

デフォルトでは LVM ストレージプールは LVM thin pool を使用しその中に全ての LXD ストレージエンティティ (イメージ、インスタンス、カスタムボリューム) の論理ボリュームを作成します。
この挙動はプール作成時に [`lvm.use_thinpool`](storage-lvm-pool-config) を `false` に設定することで変更できます。
この場合、LXD はスナップショットでない全てのストレージエンティティに "通常の" 論理ボリュームを使用します。
これは深刻なパフォーマンスの低下とディスクの空き容量の低下を `lvm` ドライバに必然的にもたらすことに注意してください (スピードとストレージ使用量の両面で `dir` ドライバに近くなります)。
この理由は thin pool でない論理ボリュームがスナップショットのスナップショットをサポートしないため、ほとんどのストレージ操作が `rsync` の使用にフォールバックするためです。
さらに、 thin でないスナップショットは作成時に最大のサイズのストレージを予約しなければならないため、 thin スナップショットよりもはるかに大容量のストレージを使用するからです。
このため、このオプションはどうしても必要なユースケースの場合にのみ選択すべきです。

インスタンスの入れ替わりが激しい環境 (例えば、継続的インテグレーション) では、LXD の操作が遅くなるのを回避するため `/etc/lvm/lvm.conf` 内のバックアップの `retain_min` と `retain_days` 設定を調整すべきです。

## 設定オプション

`lvm` ドライバを使うストレージプールとこれらのプール内のストレージボリュームには以下の設定オプションが利用できます。

(storage-lvm-pool-config)=
## ストレージプール設定

キー                         | 型     | デフォルト値                                              | 説明
:--                          | :---   | :------                                                   | :----------
`lvm.thinpool_name`          | string | `LXDThinPool`                                             | ボリュームが作成される thin pool
`lvm.thinpool_metadata_size` | string | `0` (auto)                                                | thin pool メタデータボリュームのサイズ (デフォルトは LVM が適切なサイズを計算)
`lvm.use_thinpool`           | bool   | `true`                                                    | ストレージプールは論理ボリュームに thin pool を使うかどうか
`lvm.vg.force_reuse`         | bool   | `false`                                                   | 既存の空でないボリュームグループの使用を強制
`lvm.vg_name`                | string | プールの名前                                              | 作成するボリュームグループ名
`rsync.bwlimit`              | string | `0` (no limit)                                            | ストレージエンティティーの転送に`rsync`を使う場合、ソケットI/Oに設定する上限を指定
`rsync.compression`          | bool   | `true`                                                    | ストレージプールをマイグレートする際に圧縮を使用するかどうか
`size`                       | string | 自動 (空きディスクスペースの 20%, >= 5 GiB and <= 30 GiB) | ループベースのプールを作成する際のストレージプールのサイズ (バイト単位、接尾辞のサポートあり、増やすとストレージプールのサイズを拡大)
`source`                     | string | -                                                         | 既存のブロックデバイスかループファイルかLVMボリュームグループのパス
`source.wipe`                | bool   | `false`                                                   | ストレージプールを作成する前に`source`で指定されたブロックデバイスの中身を消去する

{{volume_configuration}}

(storage-lvm-vol-config)=
## ストレージボリューム設定

キー                  | 型     | 条件               | デフォルト値                                 | 説明
:--                   | :---   | :--------          | :------                                      | :----------
`block.filesystem`    | string |                    | `volume.block.filesystem` と同じ             | {{block_filesystem}}
`block.mount_options` | string |                    | `volume.block.mount_options` と同じ          | block-backedなファイルシステムボリュームのマウントオプション
`lvm.stripes`         | string |                    | `volume.lvm.stripes` と同じ                  | 新しいボリューム (あるいは thin pool ボリューム) に使用するストライプ数
`lvm.stripes.size`    | string |                    | `volume.lvm.stripes.size` と同じ             | 使用するストライプのサイズ (最低 4096 バイトで 512 バイトの倍数を指定)
`security.shifted`    | bool   | カスタムボリューム | `volume.security.shifted` と同じか `false`   | {{enable_ID_shifting}}
`security.unmapped`   | bool   | カスタムボリューム | `volume.security.unmapped` と同じか `false`  | ボリュームへの ID マッピングを無効にする
`size`                | string |                    | `volume.size` と同じ                         | ストレージボリュームのサイズ/クォータ
`snapshots.expiry`    | string | カスタムボリューム | `volume.snapshots.expiry` と同じ             | {{snapshot_expiry_format}}
`snapshots.pattern`   | string | カスタムボリューム | `volume.snapshots.pattern` と同じか `snap%d` | {{snapshot_pattern_format}} [^*]
`snapshots.schedule`  | string | カスタムボリューム | `volume.snapshots.schedule` と同じ           | {{snapshot_schedule_format}}

[^*]: {{snapshot_pattern_detail}}

### ストレージバケット設定

ローカルのストレージプールドライバでストレージバケットを有効にし、 S3 プロトコル経由でアプリケーションがバケットにアクセスできるようにするには[`core.storage_buckets_address`](server-options-core)サーバー設定を調整する必要があります。

キー   | 型     | 条件           | デフォルト値         | 説明
:--    | :---   | :--------      | :------              | :----------
`size` | string | 適切なドライバ | `volume.size` と同じ | ストレージバケットのサイズ/クォータ
